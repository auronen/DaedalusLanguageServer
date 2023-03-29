package langserver

import (
	"context"
	"fmt"
	"os"
	"strings"

	// "path"
	// "strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
	// og "go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

const (
	CommandSetupWorkspace      string = "daedalus.dls-setup-workspace"
	CommandTranslateAll        string = "daedalus.dls-all"
	CommandTranslateAutorun    string = "daedalus.dls-autorun"
	CommandTranslateSubstitute string = "daedalus.dls-edit"
	CommandTranslateUnresolved string = "daedalus.dls-unresolved"
	CommandTranslateLogs       string = "daedalus.dls-logs"
)

func (h *LspHandler) handleWorkspaceExecuteCommand(req dls.RpcContext, params lsp.ExecuteCommandParams) error {
	if params.Command == CommandTranslateSubstitute {
		h.logger.Infof("Command arguments: %v", params.Arguments)
		h.logger.Infof("Collecting languages")

		languages, err := csvGetLanguages()
		if err != nil {
			h.logger.Errorf("Error geeting file names: %v", err)
		}

		// languages = []string{"cs", "en", "tr", "zh_Hans"}

		buttons := make([]lsp.MessageActionItem, 0)
		for _, lang := range languages {
			h.logger.Infof("Language: %s", lang)
			buttons = append(buttons, lsp.MessageActionItem{
				Title: lang,
			})
		}
		req := lsp.ShowMessageRequestParams{
			Type:    lsp.Info,
			Message: "Pick a language",
			Actions: buttons,
		}

		var result lsp.MessageActionItem
		id, err := h.conn.Call(context.Background(), lsp.MethodWindowShowMessageRequest, req, &result)
		if err != nil {
			h.logger.Errorf("Error requesting message %q: %v", id, err)
		}
		h.logger.Debugf("Result: %#v", result)
		lang := result.Title

		h.logger.Infof("Substitution started")

		// for _, ws := range h.workspaces {
		// 	ws.parseGameAndMenuForTranslation(h.config)
		// }
		failedFiles := h.substituteTranslation(lang)
		h.logger.Infof("Substitution done")
		h.logger.Infof("Files failed to substitute (%d files)", len(failedFiles))
		for _, f := range failedFiles {
			h.logger.Infof("%s", uri.File(f))
		}
		return nil
	} else if params.Command == CommandTranslateAll {
		h.logger.Infof("Generating translation file.")

		// h.generateAllsimpleCSV()
		h.generateAllCSV()

		h.logger.Infof("DONE generating translation file. :herosmile:")
		return nil
	} else if params.Command == CommandTranslateAutorun {
		h.logger.Infof("Generating translation files")

		for _, w := range h.workspaces { // for every workspace
			for _, res := range w.parsedDocuments.parseResults { // for all parsed documents
				for key := range res.StringLocations {
					if strings.Contains(key, ".") {
						parts := strings.Split(key, ".")
						if len(parts) == 3 {
							fmt.Fprintf(os.Stderr, "\"%s\",\"%s\"\n", parts[1] + "." + parts[2], key)
						}
					}
				}
			}
		}

		h.logger.Infof("DONE")
	} else if params.Command == CommandTranslateUnresolved {
		h.logger.Infof("Generating unresolved strings.")
		// prints strings that need to be abstracted away in translation files
		num := 0
		for _, w := range h.workspaces { // for every workspace
			for _, res := range w.parsedDocuments.parseResults { // for all parsed documents
				for _, us := range res.UnresolvedString {
					// if strings.HasPrefix(us.content, "const") {
					// 	fmt.Fprintf(os.Stderr, "%s:%d,%s\n", res.Source, us.line, us.content)
					// 	num += 1
					// }

					fmt.Fprintf(os.Stderr, "%s:%d,%s,%s\n", res.Source, us.line, us.id, us.content)
					num += 1
				}
			}
		}
		h.logger.Infof("DONE %d strings", num)
	} else if params.Command == CommandTranslateLogs {
		h.logger.Infof("Generating and substituting logs")

		num := h.substsLogs()

		h.logger.Infof("DONE %d strings", num)
	}
	return req.Reply(req.Context(), nil, nil)
}
