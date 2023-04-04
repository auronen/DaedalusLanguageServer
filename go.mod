module github.com/kirides/DaedalusLanguageServer

go 1.19

require (
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230329005541-ce241de812fb
	github.com/google/uuid v1.3.0
	github.com/segmentio/encoding v0.3.6
	go.lsp.dev/jsonrpc2 v0.10.0
	go.lsp.dev/uri v0.3.0
	go.uber.org/zap v1.24.0
	golang.org/x/exp v0.0.0-20230321023759-10a507213a29
	golang.org/x/text v0.8.0
)

require go.lsp.dev/pkg v0.0.0-20210717090340-384b27a52fb2 // indirect

require (
	github.com/integralist/go-findroot v0.0.0-20160518114804-ac90681525dc
	github.com/segmentio/asm v1.1.3 // indirect
	go.lsp.dev/protocol v0.12.0
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
)
