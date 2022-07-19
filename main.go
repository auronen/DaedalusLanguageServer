package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"

	"langsrv/langserver"

	"go.uber.org/zap"

	_ "net/http/pprof"

	"go.lsp.dev/jsonrpc2"
)

func BV(s []debug.BuildSetting, key string) string {
	for _, v := range s {
		if v.Key == key {
			if key == "vcs.revision" && len(v.Value) > 7 {
				return v.Value[:7]
			}
			return v.Value
		}
	}
	return ""
}

func logBuildInfo(log *zap.SugaredLogger) {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	BV := func(key string) string {
		for _, v := range bi.Settings {
			if v.Key == key {
				if key == "vcs.revision" && len(v.Value) > 7 {
					return v.Value[:7]
				}
				return v.Value
			}
		}
		return ""
	}

	log.Infof("Running %q built with %s at %s", BV("vcs.revision"), bi.GoVersion, BV("vcs.time"))
}

func main() {
	pprofPort := flag.Int("pprof", -1, "enables pprof on the specified port")
	logLevel := zap.LevelFlag("loglevel", zap.InfoLevel, "debug/info/warn/error")
	flag.Parse()

	logCfg := zap.NewDevelopmentConfig()
	logCfg.Level = zap.NewAtomicLevelAt(*logLevel)
	logger, err := logCfg.Build()

	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	log := logger.Sugar()

	logBuildInfo(log)

	pprofServer := &pprofServer{}
	defer pprofServer.Stop()

	if *pprofPort > 0 {
		pprofServer.ChangeAddr(fmt.Sprintf("127.0.0.1:%d", *pprofPort))
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()

	conn := connectLanguageServer(&rwc{os.Stdin, os.Stdout})
	lspHandler := langserver.NewLspHandler(conn, log)

	lspHandler.OnConfigChanged(func(config langserver.LspConfig) {
		level := strings.ToLower(config.LogLevel)
		switch {
		case strings.HasPrefix(level, "d"):
			logCfg.Level.SetLevel(zap.DebugLevel)
		case strings.HasPrefix(level, "w"):
			logCfg.Level.SetLevel(zap.WarnLevel)
		case strings.HasPrefix(level, "e"):
			logCfg.Level.SetLevel(zap.ErrorLevel)
		case strings.HasPrefix(level, "i"):
			logCfg.Level.SetLevel(zap.InfoLevel)
		default:
			logCfg.Level.SetLevel(zap.InfoLevel)
		}
	})

	lspHandler.OnConfigChanged(func(config langserver.LspConfig) {
		if *pprofPort <= 0 {
			// only when not set by args
			log.Infof("Updating pprof address to %s", config.PprofAddr)
			pprofServer.ChangeAddr(config.PprofAddr)
		}
	})

	conn.Go(ctx, func(ctx context.Context, reply jsonrpc2.Replier, req jsonrpc2.Request) error {
		err := lspHandler.Handle(ctx, reply, req)
		if err != nil {
			if errors.Is(err, langserver.ErrUnhandled) {
				log.Debugf("%v\n", err)
			} else {
				log.Errorf("%s: %v\n", req.Method(), err)
			}
		}
		return nil // unhandled
	})
	<-conn.Done()
}

type rwc struct {
	r io.ReadCloser
	w io.WriteCloser
}

func (rwc *rwc) Read(b []byte) (int, error)  { return rwc.r.Read(b) }
func (rwc *rwc) Write(b []byte) (int, error) { return rwc.w.Write(b) }
func (rwc *rwc) Close() error {
	rwc.r.Close()
	return rwc.w.Close()
}

func connectLanguageServer(rwc io.ReadWriteCloser) jsonrpc2.Conn {
	bufStream := jsonrpc2.NewStream(rwc)
	rootConn := jsonrpc2.NewConn(bufStream)
	return rootConn

}
