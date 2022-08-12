package langserver

import (
	"fmt"

	dls "github.com/kirides/DaedalusLanguageServer"
	lsp "go.lsp.dev/protocol"
)

const (
	//DlsConstants string = "daedalus.dls-constants"
	//DlsDialogues string = "daedalus.dls-dialogues"
	DlsAll string = "daedalus.dls-all"
)

func (h *LspHandler) handleCommand(req dls.RpcContext, params lsp.ExecuteCommandParams) error {
	if params.Command == DlsAll {
		h.logger.Infof("Generating all")
		h.generateAllCSV()
		h.logger.Infof("Done")
		return nil
	}
	return fmt.Errorf("no command found")
}
