package langserver

// TranslatedString represents string read from a csv file
//
// stringId      - string identifier
// stringContent - content of the string (the actual text)
type TranslatedString struct {
	stringID      string
	stringContent string
	substituted   bool
}

func newTranslatedString(stringID, stringContent string) TranslatedString {
	return TranslatedString{
		stringID:      stringID,
		stringContent: stringContent,
		substituted:   false,
	}
}

// SymbolPosition holds the position of a string
//
// document - document URI
// line     - line
// start    - column where the string starts
// end      - column where the string ends
// quotes   - is this string in quotes (string literal) or does it come after // (comments - SVMs and AI_Outputs)
type SymbolPosition struct {
	document            string
	content             string
	translation_comment string
	line                int
	start               int
	end                 int
	quotes              bool
}

func newSymbolPosition(doc, cont string, ln, st, en int, params ...interface{}) SymbolPosition {
	// Default values
	translation_comment := ""
	par := true
	if len(params) == 1 {
		if quotes, ok := params[0].(bool); ok {
			par = quotes
		} else if tr_com, ok := params[0].(string); ok {
			translation_comment = tr_com
		}
	}
	return SymbolPosition{
		document:            doc,
		content:             cont,
		line:                ln,
		start:               st,
		end:                 en,
		quotes:              par,
		translation_comment: translation_comment,
	}
}
