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

	fmt.Fprintf(os.Stderr, "-> duplicate string IDs {\n")
	for _, e := range dupl {
		fmt.Fprintf(os.Stderr, "\t[%s]: %s\n", e.stringID, e.stringContent)
	}
	fmt.Fprintf(os.Stderr, "}\n")

	var response lsp.ApplyWorkspaceEditResponse

	edits := make(map[uri.URI][]lsp.TextEdit)

	for _, w := range h.workspaces { // for every workspace
		for _, res := range w.parsedDocuments.parseResults { // for all parsed documents
			file := strings.TrimPrefix(res.Source, path.Join(path.Dir(w.path), path.Base(path.Dir(w.path))))
			// fmt.Fprintf(os.Stderr, "path: %s\n", uri.File(file))
			fmt.Fprintf(os.Stderr, "Document: %s num of strings: %d\n", path.Base(file), len(res.StringLocations))

			for i, entry := range ts { // for every entry in the translation files
				if positions, ok := res.StringLocations[entry.stringID]; ok {
					ts[i].substituted = true
					for _, pos := range positions {
						var cont string
						if pos.quotes {
							cont = strings.ReplaceAll(entry.stringContent, "\"", "'");
						} else {
							cont = entry.stringContent
						}
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
							NewText: cont,
						})
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
		//	failedFiles = append(failedFiles, strings.TrimPrefix(res.Source, path.Join(path.Dir(w.path), path.Base(path.Dir(w.path)))))
		// This reports nothing...
		//h.logger.Infof("Reason: %s", response.FailureReason)
	}
		h.logger.Infof("%s", h.debugPrintEdits(edits))

	numOfEdits := 0
	for ed := range edits {
		numOfEdits += len(ed)
	}

	h.logger.Infof("Number of files: %d Number of edits: %d", len(edits), numOfEdits)

	numberOfIDs := 0
	for _, entry := range ts {
		if !entry.substituted {
			numberOfIDs++
			// fmt.Fprint(os.Stderr, entry.stringID+"\n")
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













func (h *LspHandler) generateTranslationFiles() {

}





func (h *LspHandler) substsLogs() int {

	edits := make(map[uri.URI][]lsp.TextEdit)

	intermediate := make(map[string][]logPos)
	num := 0
	for _, w := range h.workspaces { // for every workspace
		for _, res := range w.parsedDocuments.parseResults { // for all parsed documents
			for key, values := range res.logs {
				intermediate[key] = append(intermediate[key], values...)
				num += 1
			}
		}
	}
	h.logger.Infof("%s", h.debugPrintEdits(edits))

	for _, log_entries := range intermediate {
		var constants string
		var uri_tmp uri.URI
		var fnc_line int
		for i, l := range log_entries {
			var line string
			if strings.HasPrefix(strings.ToLower(l.entryID), "use_") {
				line = l.entryID[4:] + "_" + fmt.Sprint(i+1)
			} else if strings.HasPrefix(strings.ToLower(l.entryID), "use") {
				line = l.entryID[3:] + "_" + fmt.Sprint(i+1)
			} else {
				line = l.entryID + "_" + fmt.Sprint(i+1)
			}
			edits[l.uri] = append(edits[l.uri], lsp.TextEdit{
				Range: lsp.Range{
					Start: lsp.Position{
						Line: uint32(l.line),
						Character: uint32(l.start),
					},
					End: lsp.Position{
						Line: uint32(l.line),
						Character: uint32(l.end),
					},
				},
				NewText: line,
			})
			// fmt.Fprintf(os.Stderr, "const string %s = %s;\n", line, l.content)
			constants += "const string " + line + " = " + l.content + ";\n"
			uri_tmp = l.uri
			fnc_line = l.function_line
		}

		edits[uri_tmp] = append(edits[uri_tmp], lsp.TextEdit{
			Range: lsp.Range{
				Start: lsp.Position{
					Line: uint32(fnc_line-1),
					Character: 0,
				},
				End: lsp.Position{
					Line: uint32(fnc_line-1),
					Character: 0,
				},
			},
			NewText: constants,
		})
	}



	var response lsp.ApplyWorkspaceEditResponse

	h.conn.Call(context.Background(),
		lsp.MethodWorkspaceApplyEdit,
		lsp.ApplyWorkspaceEditParams{
			Edit: lsp.WorkspaceEdit{
				Changes: edits,
			}}, &response)

	return num;
}
