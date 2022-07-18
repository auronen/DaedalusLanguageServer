package langserver

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
