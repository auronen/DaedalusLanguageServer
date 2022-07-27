package langserver

import (
	"fmt"

	dls "github.com/kirides/DaedalusLanguageServer"
	lsp "go.lsp.dev/protocol"
)

const (
	DlsConstants string = "daedalus.dls-constants"
)

func (h *LspHandler) handleCommand(req dls.RpcContext, params lsp.ExecuteCommandParams) error {
	if params.Command == DlsConstants {
		h.logger.Infof("Generating...")
		h.generateConstStringCSV()
		h.logger.Infof("Done")
		// This should return progress
		return nil
	} /* else if params.Command == "daedalus.dialogues" {
		h.logger.Infof("Generating dialogues")
		h.generateDialogueCSV()
		h.logger.Infof("Done")
		return nil
	} */
	return fmt.Errorf("no command found")
}
