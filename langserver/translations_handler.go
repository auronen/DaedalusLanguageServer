package langserver

import (
	"context"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/integralist/go-findroot/find"
	lsp "go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

// Returns absolute path of the git repo
func findRepoRoot() (string, error) {
	root, err := find.Repo()
	if err != nil {
		return "", err
	}
	return root.Path, nil
}

func uniqueTranslatedStrings(input []TranslatedString) ([]TranslatedString, []TranslatedString) {
	u := make([]TranslatedString, 0, len(input))
	duplicates := make([]TranslatedString, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val.stringID]; !ok {
			m[val.stringID] = true
			u = append(u, val)
		} else {
			duplicates = append(duplicates, val)
		}
	}
	sort.Slice(duplicates, func(i, j int) bool {
		return duplicates[i].stringID < duplicates[j].stringID
	})
	return u, duplicates
}

// test lsp server's ability to edit code
func (h *LspHandler) substituteTranslation(language string) (failedFiles []string) {

	h.logger.Infof("Creating edits")

	// Read all of the translation files for a selected language and parse them into TranslatedString slice
	// TODO: How to pass in arguments from commands?
	// TODO: Create encoding lookup table
	ts, err := csvReadLanguage(language)
	if err != nil {
		h.logger.Errorf("Error: %s", err)
	}
	h.logger.Infof("About to substitute %d strings from csv translation files.", len(ts))

	ts, dupl := uniqueTranslatedStrings(ts)

	h.conn.Notify(context.Background(),
		lsp.MethodWindowShowMessage,
		lsp.ShowMessageRequestParams{
			Message: fmt.Sprintf("Squashed %d duplicate string IDs in translation files (using the first one found).\nPlease remove the duplicate IDs from translation fifles.\nFor more information check the logs.", len(dupl)),
			Type:    lsp.MessageTypeInfo,
		},
	)

	fmt.Fprintf(os.Stderr, "duplicate string IDs {\n")
	for _, e := range dupl {
		fmt.Fprintf(os.Stderr, "\t[%s]: %s\n", e.stringID, e.stringContent)
	}
	fmt.Fprintf(os.Stderr, "}\n")

	var response lsp.ApplyWorkspaceEditResponse

	edits := make(map[uri.URI][]lsp.TextEdit)

	for _, w := range h.workspaces { // for every workspace
		for _, res := range w.parsedDocuments.parseResults { // for all parsed documents
			//edits := make(map[uri.URI][]lsp.TextEdit)
			file := strings.TrimPrefix(res.Source, path.Join(path.Dir(w.path), path.Base(path.Dir(w.path))))
			for i, entry := range ts { // for every entry in the translation files
				//h.logger.Infof("entries: %d", i)
				if positions, ok := res.StringLocations[entry.stringID]; ok {
					ts[i].substituted = true
					for _, pos := range positions {
						if pos.quotes { // if it had quotes
							edits[uri.File(file)] = append(edits[uri.File(file)], lsp.TextEdit{
								Range: lsp.Range{
									Start: lsp.Position{
										Line:      uint32(pos.line - 1),
										Character: uint32(pos.start),
									},
									End: lsp.Position{
										Line:      uint32(pos.line - 1),
										Character: uint32(pos.end),
									},
								},
								NewText: "\"" + entry.stringContent + "\"",
							})
						} else { // if it was a comment
							edits[uri.File(file)] = append(edits[uri.File(file)], lsp.TextEdit{
								Range: lsp.Range{
									Start: lsp.Position{
										Line:      uint32(pos.line - 1),
										Character: uint32(pos.start),
									},
									End: lsp.Position{
										Line:      uint32(pos.line - 1),
										Character: uint32(pos.end),
									},
								},
								NewText: "//" + strings.TrimSpace(entry.stringContent),
							})
						}
					}
				}
			}
			if len(edits) == 0 {
				continue
			}
			// this is probably not needed :thinking:

			sort.Slice(edits[uri.File(file)], func(i, j int) bool {
				return edits[uri.File(file)][i].Range.Start.Line < edits[uri.File(file)][j].Range.Start.Line
			})

		}
	}

	// send the request to the editor
	// Not applicable: one file at a time - I could not get the full workspace to be done in one request
	h.conn.Call(context.Background(),
		lsp.MethodWorkspaceApplyEdit,
		lsp.ApplyWorkspaceEditParams{
			Edit: lsp.WorkspaceEdit{
				Changes: edits,
			}}, &response)

	if !response.Applied {
		h.logger.Infof("Did not apply: ")
		h.logger.Infof("%s", h.debugPrintEdits(edits))
		//	failedFiles = append(failedFiles, strings.TrimPrefix(res.Source, path.Join(path.Dir(w.path), path.Base(path.Dir(w.path)))))
		// This reports nothing...
		//h.logger.Infof("Reason: %s", response.FailureReason)
	}

	numOfEdits := 0
	for ed := range edits {
		numOfEdits += len(ed)
	}

	h.logger.Infof("Number of edits: %d", numOfEdits)

	numberOfIDs := 0
	for _, entry := range ts {
		if !entry.substituted {
			numberOfIDs++
			fmt.Fprint(os.Stderr, entry.stringID+"\n")
		}
	}
	h.logger.Infof("Number of string IDs not substituted: %d", numberOfIDs)

	return failedFiles
}

func (h *LspHandler) debugPrintEdits(edits map[uri.URI][]lsp.TextEdit) string {
	sb := strings.Builder{}
	for key, content := range edits {
		sb.WriteString(key.Filename() + "\n")
		for _, edit := range content {
			sb.WriteString(fmt.Sprintf("%v\n", edit))
		}
	}
	return sb.String()
}
