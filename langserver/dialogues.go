package langserver

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Comment struct {
	text string
	line int
}

func newCommentFromToken(tok antlr.Token) Comment {
	return Comment{
		text: tok.GetText(),
		line: tok.GetLine(),
	}
}

type Dialogue struct {
	soundName  string
	text       string
	sourceFile string
	line       int
}

func newDialogue(sndName, txt, srcFile string, ln int) Dialogue {
	return Dialogue{
		soundName:  sndName,
		text:       txt,
		sourceFile: srcFile,
		line:       ln,
	}
}

type StringLiteral struct {
	text       string
	sourceFile string
	line       int
	context    string
}

func newStringLiteral(txt, source string, ln int, ctx string) StringLiteral {
	return StringLiteral{
		text:       txt,
		sourceFile: source,
		line:       ln,
		context:    ctx,
	}
}
