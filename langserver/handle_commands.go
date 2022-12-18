package langserver

import (
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
		h.logger.Infof("Substitution started")

		for _, ws := range h.workspaces {
			ws.parseGameAndMenuForTranslation(h.config)
		}
		failedFiles := h.substituteTranslation()
		h.logger.Infof("Substitution done")
		h.logger.Infof("Files failed to substitute (%d files)", len(failedFiles))
		for _, f := range failedFiles {
			h.logger.Infof("%s", uri.File(f))
		}

	}
	return req.Reply(req.Context(), nil, nil)
}
