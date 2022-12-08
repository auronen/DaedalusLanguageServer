package langserver

// TranslatedString represents string read from a csv file
//
// stringId      - string identifier
// stringContent - content of the string (the actual text)
type TranslatedString struct {
	stringID      string
	stringContent string
}

func newTranslatedString(stringID, stringContent string) TranslatedString {
	return TranslatedString{
		stringID:      stringID,
		stringContent: stringContent,
	}
}

// SymbolPosition holds the position of a string
//
// document - document URI
// line     - line
// start    - column where the string starts
// end      - column where the string ends
type SymbolPosition struct {
	document string
	line     int
	start    int
	end      int
	quotes   bool
}

func newSymbolPosition(doc string, ln, st, en int, quotes ...bool) SymbolPosition {
	par := true
	if len(quotes) > 0 {
		par = quotes[0]
	}
	return SymbolPosition{
		document: doc,
		line:     ln,
		start:    st,
		end:      en,
		quotes:   par,
	}
}
