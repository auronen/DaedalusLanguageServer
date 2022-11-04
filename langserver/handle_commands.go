package langserver

import (
	"fmt"

	dls "github.com/kirides/DaedalusLanguageServer"
	lsp "go.lsp.dev/protocol"
)

const (
	//DlsConstants string = "daedalus.dls-constants"
	//DlsDialogues string = "daedalus.dls-dialogues"
	DlsAll     string = "daedalus.dls-all"
	DlsAutorun string = "daedalus.dls-autorun"
)

func (h *LspHandler) handleCommand(req dls.RpcContext, params lsp.ExecuteCommandParams) error {
	h.logger.Infof("Command received: %s ", params.Command)
	if params.Command == DlsAll {
		h.logger.Infof("Generating all")
		h.generateAllCSV()
		h.logger.Infof("Done")
		return nil
	} else if params.Command == DlsAutorun {
		h.logger.Infof("Generating autorun scripts")
		h.generatezPECSV()
		h.logger.Infof("Done")
		return nil
	}
	return fmt.Errorf("no command found")
}
