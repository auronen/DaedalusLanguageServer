package langserver

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"go.lsp.dev/jsonrpc2"
	lsp "go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type LspConfig struct {
	FileEncoding    string `json:"fileEncoding"`
	SrcFileEncoding string `json:"srcFileEncoding"`
}

// LspHandler ...
type LspHandler struct {
	TextDocumentSync   *textDocumentSync
	bufferManager      *BufferManager
	parsedDocuments    *parseResultsManager
	initialDiagnostics map[string][]lsp.Diagnostic
	handlers           RpcMux
	config             LspConfig
	onceParseAll       sync.Once

	baseLspHandler
	initialized bool
}

var (
	// ErrWalkAbort should be returned if a walk function should abort early
	ErrWalkAbort = fmt.Errorf("OK")
)

// NewLspHandler ...
func NewLspHandler(conn jsonrpc2.Conn, logger Logger) *LspHandler {
	bufferManager := NewBufferManager()
	parsedDocuments := newParseResultsManager(logger)
	return &LspHandler{
		baseLspHandler: baseLspHandler{
			logger: logger,
			conn:   conn,
		},
		handlers: RpcMux{
			pathToType: map[string]Handler{},
		},
		initialized:     false,
		bufferManager:   bufferManager,
		parsedDocuments: parsedDocuments,
		TextDocumentSync: &textDocumentSync{
			baseLspHandler: baseLspHandler{
				logger: logger,
				conn:   conn,
			},
			bufferManager:   bufferManager,
			parsedDocuments: parsedDocuments,
		},
	}
}

func getTypeFieldsAsCompletionItems(docs *parseResultsManager, symbolName string, locals map[string]Symbol) ([]lsp.CompletionItem, error) {
	symName := strings.ToUpper(symbolName)
	sym, ok := locals[symName]
	if !ok {
		sym, ok = docs.LookupGlobalSymbol(symName, SymbolVariable|SymbolClass|SymbolInstance|SymbolPrototype)
	}

	if !ok {
		return []lsp.CompletionItem{}, nil
	}
	if clsSym, ok := sym.(ClassSymbol); ok {
		return fieldsToCompletionItems(clsSym.Fields), nil
	}
	if protoSym, ok := sym.(ProtoTypeOrInstanceSymbol); ok {
		return getTypeFieldsAsCompletionItems(docs, protoSym.Parent, locals)
	}
	if varSym, ok := sym.(VariableSymbol); ok {
		return getTypeFieldsAsCompletionItems(docs, varSym.Type, locals)
	}
	return []lsp.CompletionItem{}, nil
}

func (h *LspHandler) getTextDocumentCompletion(params *lsp.CompletionParams) ([]lsp.CompletionItem, error) {
	result := make([]lsp.CompletionItem, 0, 200)
	parsedDoc, err := h.parsedDocuments.Get(h.uriToFilename(params.TextDocument.URI))
	if err == nil {
		doc := h.bufferManager.GetBuffer(h.uriToFilename(params.TextDocument.URI))
		di := DefinitionIndex{Line: int(params.Position.Line), Column: int(params.Position.Character)}

		scci, err := getSignatureCompletions(params, h)
		if err != nil {
			h.logger.Debugf("signature completion error %v: ", err)
		}
		if len(scci) > 0 {
			return scci, nil
		}

		// dot-completion
		proto, _, err := doc.GetParentSymbolReference(params.Position)
		if err == nil && proto != "" {
			locals := parsedDoc.ScopedVariables(di)
			return getTypeFieldsAsCompletionItems(h.parsedDocuments, proto, locals)
		}

		// locally scoped variables ordered at the top
		localSortIdx := 0
		for _, fn := range parsedDoc.Functions {
			if fn.BodyDefinition.InBBox(di) {
				for _, p := range fn.Parameters {
					ci, err := completionItemFromSymbol(p)
					if err != nil {
						continue
					}
					localSortIdx++
					ci.Detail += " (parameter)"
					ci.SortText = fmt.Sprintf("%000d", localSortIdx)
					result = append(result, ci)
				}
				for _, p := range fn.LocalVariables {
					ci, err := completionItemFromSymbol(p)
					if err != nil {
						continue
					}
					localSortIdx++
					ci.Detail += " (local)"
					ci.SortText = fmt.Sprintf("%000d", localSortIdx)
					result = append(result, ci)
				}
				break
			}
		}
		for _, fn := range parsedDoc.Instances {
			if fn.BodyDefinition.InBBox(di) {
				return getTypeFieldsAsCompletionItems(h.parsedDocuments, fn.Parent, map[string]Symbol{})
			}
		}
		for _, fn := range parsedDoc.Prototypes {
			if fn.BodyDefinition.InBBox(di) {
				return getTypeFieldsAsCompletionItems(h.parsedDocuments, fn.Parent, map[string]Symbol{})
			}
		}
	}

	h.parsedDocuments.WalkGlobalSymbols(func(s Symbol) error {
		ci, err := completionItemFromSymbol(s)
		if err != nil {
			return nil
		}
		result = append(result, ci)
		return nil
	}, SymbolAll)

	return result, nil
}

func (h *LspHandler) lookUpScopedSymbol(documentURI, identifier string, position lsp.Position) Symbol {
	p, err := h.parsedDocuments.Get(documentURI)
	if err != nil {
		return nil
	}

	di := DefinitionIndex{Line: int(position.Line), Column: int(position.Character)}

	var sym Symbol
	p.WalkScopedVariables(di, func(s Symbol) bool {
		if strings.EqualFold(s.Name(), identifier) {
			sym = s
			return false
		}
		return true
	})
	return sym
}

func (h *LspHandler) lookUpSymbol(documentURI string, position lsp.Position) (Symbol, error) {
	doc := h.bufferManager.GetBuffer(documentURI)
	if doc == "" {
		return nil, fmt.Errorf("document %q not found", documentURI)
	}
	identifier := doc.GetWordRangeAtPosition(position)

	if v := h.lookUpScopedSymbol(documentURI, identifier, position); v != nil {
		return v, nil
	}

	symbol, found := h.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(identifier), SymbolAll)

	if !found {
		return nil, fmt.Errorf("identifier %q not found", identifier)
	}

	return symbol, nil
}

func (h *LspHandler) handleCommand(ctx context.Context, params *lsp.ExecuteCommandParams) error {
	if params.Command == "daedalus.dls-hello" {
		h.logger.Infof("Generating...")
		h.generateConstStringCSV()
		h.logger.Infof("Done")
		// This should return progress
		return nil
	} else if params.Command == "daedalus.dialogues" {
		h.logger.Infof("Generating dialogues")
		h.generateDialogueCSV()
		h.logger.Infof("Done")
	}
	h.logger.Infof("Hello - no")
	return fmt.Errorf("no command found")
}
func (h *LspHandler) handleSignatureInfo(ctx context.Context, params *lsp.TextDocumentPositionParams) (lsp.SignatureHelp, error) {
	fnCtx, err := getFunctionCallContext(h, params.TextDocument.URI, params.Position)
	if err != nil {
		return lsp.SignatureHelp{}, err
	}

	fn := fnCtx.Function

	var fnParams []lsp.ParameterInformation
	for _, p := range fn.Parameters {
		doc := findJavadocParam(fn.Documentation(), p.Name())
		var mdDoc interface{}
		if doc != "" {
			mdDoc = &lsp.MarkupContent{
				Kind:  lsp.Markdown,
				Value: fmt.Sprintf("**%s** - *%s*", p.Name(), cleanUpParamDesc(doc)),
			}
		}
		fnParams = append(fnParams, lsp.ParameterInformation{
			Label:         p.String(),
			Documentation: mdDoc,
		})
	}

	return lsp.SignatureHelp{
		Signatures: []lsp.SignatureInformation{
			{
				Documentation: &lsp.MarkupContent{
					Kind:  lsp.Markdown,
					Value: simpleJavadocMD(fn),
				},
				Label:      fn.String(),
				Parameters: fnParams,
			},
		},
		ActiveParameter: uint32(fnCtx.ParamIdx),
		ActiveSignature: 0,
	}, nil
}

func (h *LspHandler) handleGoToDefinition(ctx context.Context, params *lsp.TextDocumentPositionParams) (lsp.Location, error) {
	symbol, err := h.lookUpSymbol(h.uriToFilename(params.TextDocument.URI), params.Position)
	if err != nil {
		return lsp.Location{}, err
	}

	return lsp.Location{
		URI: uri.File(symbol.Source()),
		Range: lsp.Range{
			Start: lsp.Position{
				Character: uint32(symbol.Definition().Start.Column),
				Line:      uint32(symbol.Definition().Start.Line - 1),
			},
			End: lsp.Position{
				Character: uint32(symbol.Definition().Start.Column + len(symbol.Name())),
				Line:      uint32(symbol.Definition().Start.Line - 1),
			},
		}}, nil
}

func (h *LspHandler) handleTextDocumentCompletion(req RpcContext, data lsp.CompletionParams) error {
	items, err := h.getTextDocumentCompletion(&data)
	replyEither(req.Context(), req, items, err)
	return nil
}

func (h *LspHandler) handleTextDocumentDefinition(req RpcContext, data lsp.TextDocumentPositionParams) error {
	found, err := h.handleGoToDefinition(req.Context(), &data)
	if err != nil {
		return req.Reply(req.Context(), nil, nil)
	}
	return req.Reply(req.Context(), found, nil)
}

func (h *LspHandler) handleTextDocumentHover(req RpcContext, data lsp.TextDocumentPositionParams) error {
	found, err := h.lookUpSymbol(h.uriToFilename(data.TextDocument.URI), data.Position)
	if err != nil {
		return req.Reply(req.Context(), nil, nil)
	}
	h.LogDebug("Found Symbol for Hover: %s\n", found.String())
	return req.Reply(req.Context(), lsp.Hover{
		Range: &lsp.Range{
			Start: data.Position,
			End:   data.Position,
		},
		Contents: lsp.MarkupContent{
			Kind:  lsp.Markdown,
			Value: strings.TrimSpace(simpleJavadocMD(found) + "\n\n```daedalus\n" + found.String() + "\n```"),
		},
	}, nil)
}
func (h *LspHandler) handleTextDocumentSignatureHelp(req RpcContext, data lsp.TextDocumentPositionParams) error {
	result, err := h.handleSignatureInfo(req.Context(), &data)
	if err != nil {
		return req.Reply(req.Context(), nil, nil)
	}
	return req.Reply(req.Context(), result, nil)
}

// TODO: handleExecuteCommand should control all commands (like format, check duplicate string literals and so on)
/* or other commands related to the translations:
preview language - checks the repo and changes all the strings to selected language temporarily, maybe use InlayHints?
generate language scripts - generates new script base in selected language
generate CSVs - generates language files to be translated
*/
func (h *LspHandler) handleExecuteCommand(req RpcContext, data lsp.ExecuteCommandParams) error {
	err := h.handleCommand(req.Context(), &data)
	if err != nil {
		return req.Reply(req.Context(), nil, nil)
	}
	// return results/progress or some info?
	return req.Reply(req.Context(), nil, nil)
}

func (h *LspHandler) onInitialized() {
	h.handlers.Register(lsp.MethodTextDocumentCompletion, MakeHandler(h.handleTextDocumentCompletion))
	h.handlers.Register(lsp.MethodTextDocumentDefinition, MakeHandler(h.handleTextDocumentDefinition))
	h.handlers.Register(lsp.MethodTextDocumentHover, MakeHandler(h.handleTextDocumentHover))
	h.handlers.Register(lsp.MethodTextDocumentSignatureHelp, MakeHandler(h.handleTextDocumentSignatureHelp))

	h.handlers.Register(lsp.MethodWorkspaceExecuteCommand, MakeHandler(h.handleExecuteCommand))
	// textDocument/didOpen/didSave/didChange
	h.handlers.Register(lsp.MethodTextDocumentDidOpen, MakeHandler(h.TextDocumentSync.handleTextDocumentDidOpen))
	h.handlers.Register(lsp.MethodTextDocumentDidChange, MakeHandler(h.TextDocumentSync.handleTextDocumentDidChange))
	h.handlers.Register(lsp.MethodTextDocumentDidSave, MakeHandler(h.TextDocumentSync.handleTextDocumentDidSave))
}

func prettyJSON(val interface{}) string {
	v, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(v)
}

// Deliver ...
func (h *LspHandler) Handle(ctx context.Context, reply jsonrpc2.Replier, r jsonrpc2.Request) error {
	h.LogDebug("Requested %q", r.Method())

	// if r.Params != nil {
	// 	var paramsMap map[string]interface{}
	// 	json.Unmarshal(*r.Params, &paramsMap)
	// 	fmt.Fprintf(os.Stderr, "Params: %+v\n", paramsMap)
	// }
	switch r.Method() {
	case lsp.MethodInitialize:
		if err := reply(ctx, lsp.InitializeResult{
			Capabilities: lsp.ServerCapabilities{
				CompletionProvider: &lsp.CompletionOptions{
					TriggerCharacters: []string{"."},
				},
				DefinitionProvider: true,
				HoverProvider:      true,
				SignatureHelpProvider: &lsp.SignatureHelpOptions{
					TriggerCharacters: []string{"(", ","},
				},
				TextDocumentSync: lsp.TextDocumentSyncOptions{
					Change:    lsp.TextDocumentSyncKindFull,
					OpenClose: true,
					Save: &lsp.SaveOptions{
						IncludeText: true,
					},
				},
				ExecuteCommandProvider: &lsp.ExecuteCommandOptions{
					Commands: []string{"daedalus.dls-hello", "daedalus.dialogues"},
				},
			},
		}, nil); err != nil {
			return fmt.Errorf("not initialized")
		}
		h.initialized = true
		h.onInitialized()
		return nil
	case lsp.MethodWorkspaceDidChangeConfiguration:
		var params struct {
			Settings struct {
				DaedalusLanguageServer LspConfig `json:"daedalusLanguageServer"`
			} `json:"settings"`
		}

		var configRaw map[string]interface{}
		_ = json.Unmarshal(r.Params(), &configRaw)
		h.LogDebug("%s (debug): %v", r.Method(), prettyJSON(configRaw))

		_ = json.Unmarshal(r.Params(), &params)
		h.config = params.Settings.DaedalusLanguageServer
		h.LogInfo("%s: %v", r.Method(), prettyJSON(configRaw))

		h.parsedDocuments.SetFileEncoding(h.config.FileEncoding)
		h.parsedDocuments.SetSrcEncoding(h.config.SrcFileEncoding)
		return nil
	case lsp.MethodInitialized:
		return nil
	}

	// DEFAULT / OTHERWISE

	if !h.initialized {
		reply(ctx, nil, jsonrpc2.Errorf(jsonrpc2.ServerNotInitialized, "Not initialized yet"))
		return fmt.Errorf("not initialized yet")
	}

	// Recover if something bad happens in the handlers...
	defer func() {
		err := recover()
		if err != nil {
			h.LogWarn("Recovered from panic at %s: %v\n", r.Method, err)
		}
	}()

	if r.Method() == lsp.MethodTextDocumentDidOpen {
		h.onceParseAll.Do(func() {
			h.LogInfo("Starting parsing of sources...")
			go func() {
				exe, _ := os.Executable()
				var resultsX []*ParseResult
				if f, err := findPath(filepath.Join(filepath.Dir(exe), "DaedalusBuiltins", "builtins.src")); err == nil {
					resultsX, err = h.parsedDocuments.ParseSource(f)
					if err != nil {
						h.LogError("Error parsing %q: %v", f, err)
						return
					}
				}

				if externalsSrc, err := findPath(filepath.Join("_externals", "externals.src")); err == nil {
					customBuiltins, err := h.parsedDocuments.ParseSource(externalsSrc)
					if err != nil {
						h.LogError("Error parsing %q: %v", externalsSrc, err)
					} else {
						resultsX = append(resultsX, customBuiltins...)
					}
				} else if externalsDaedalus, err := findPath(filepath.Join("_externals", "externals.d")); err == nil {
					parsed, err := h.parsedDocuments.ParseFile(externalsDaedalus)
					if err != nil {
						h.LogError("Error parsing %q: %v", externalsDaedalus, err)
					} else {
						resultsX = append(resultsX, parsed)
					}
				}

				for _, v := range []string{"Gothic.src", "Camera.src", "Menu.src", "Music.src", "ParticleFX.src", "SFX.src", "VisualFX.src"} {
					if full, err := findPath(v); err == nil {
						results, err := h.parsedDocuments.ParseSource(full)
						if err != nil {
							h.LogError("Error parsing %s: %v", full, err)
							return
						}
						resultsX = append(resultsX, results...)
					} else {
						h.LogDebug("Did not parse %q: %v", v, err)
					}
				}

				var diagnostics []lsp.Diagnostic
				tmpDiags := make(map[string][]lsp.Diagnostic)

				for _, p := range resultsX {
					if p.SyntaxErrors != nil && len(p.SyntaxErrors) > 0 {
						diagnostics = make([]lsp.Diagnostic, 0, len(p.SyntaxErrors))
						for _, se := range p.SyntaxErrors {
							diagnostics = append(diagnostics, se.Diagnostic())
						}
						tmpDiags[p.Source] = diagnostics
					}
				}
				h.initialDiagnostics = tmpDiags
			}()
		})
	}

	if h.initialDiagnostics != nil && len(h.initialDiagnostics) > 0 {
		h.LogInfo("Publishing initial diagnostics (%d).", len(h.initialDiagnostics))
		for k, v := range h.initialDiagnostics {
			h.LogInfo("> %s", k)
			h.conn.Notify(ctx, lsp.MethodTextDocumentPublishDiagnostics, lsp.PublishDiagnosticsParams{
				URI:         lsp.DocumentURI(uri.File(k)),
				Diagnostics: v,
			})
		}
		h.initialDiagnostics = map[string][]lsp.Diagnostic{}
	}

	handled, err := h.handlers.Handle(ctx, reply, r)
	if err != nil && handled {
		return err
	}
	return h.baseLspHandler.Handle(ctx, reply, r)
}
