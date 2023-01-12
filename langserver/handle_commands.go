package langserver

import (
	"context"

	dls "github.com/kirides/DaedalusLanguageServer"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
	"go.lsp.dev/uri"
)

const (
	CommandSetupWorkspace      string = "daedalus.dls-setup-workspace"
	CommandTranslateAll        string = "daedalus.dls-all"
	CommandTranslateAutorun    string = "daedalus.dls-autorun"
	CommandTranslateSubstitute string = "daedalus.dls-edit"
)

func (h *LspHandler) handleWorkspaceExecuteCommand(req dls.RpcContext, params lsp.ExecuteCommandParams) error {
	if params.Command == CommandTranslateSubstitute {
		h.logger.Infof("Command arguments: %v", params.Arguments)
		h.logger.Infof("Collecting languages")

		languages, err := csvGetLanguages()
		if err != nil {
			h.logger.Errorf("Error geeting file names: %v", err)
		}

		languages = []string{"cs", "en", "de", "es"}

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

		for _, ws := range h.workspaces {
			ws.parseGameAndMenuForTranslation(h.config)
		}
		failedFiles := h.substituteTranslation(lang)
		h.logger.Infof("Substitution done")
		h.logger.Infof("Files failed to substitute (%d files)", len(failedFiles))
		for _, f := range failedFiles {
			h.logger.Infof("%s", uri.File(f))
		}
	}
	return req.Reply(req.Context(), nil, nil)
}
