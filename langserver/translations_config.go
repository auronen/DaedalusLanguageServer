package langserver

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/segmentio/encoding/json"
)

type translationConfiguration struct {
	FileMasks         []string `json:"fileMasks"`
	FunctionBlackList []string `json:"functionBlackList"`
	MemberVariables   []string `json:"memberVariables"`
}

func readConf(file string, globalConfig *translationConfiguration) {
	// Open our jsonFile
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &globalConfig)

	fmt.Fprintf(os.Stderr, "\nTranslation config: %v\n", prettyJSON(globalConfig))
}

func initTranslationConfig(ws LspWorkspace) translationConfiguration {
	var conf translationConfiguration
	targetDir := filepath.Join(ws.path, ".dls", "translations.json")
	readConf(targetDir, &conf)
	return conf
}
