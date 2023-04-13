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
	SymbolBlacklist   []string `json:"SymbolBlacklist"`
}

func readConf(file string, globalConfig *translationConfiguration) error {
	// Open our jsonFile
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if os.IsNotExist(err) {
		path, err := findRepoRoot();
		if err != nil {
			return err
		}
		// TODO: create translations.json file based on G1 or G2
		f, err := os.Create(filepath.Join(path, ".dls", "translations.json"))
		if err != nil {
			return err
		}
		defer f.Close()
	} else if err != nil {
		return err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &globalConfig)

	// fmt.Fprintf(os.Stderr, "\nTranslation config: %v\n", prettyJSON(globalConfig))
	return nil
}

func initTranslationConfig(ws *LspWorkspace) translationConfiguration {
	if ws == nil {
		return translationConfiguration{}
	}
	var conf translationConfiguration
	targetDir := filepath.Join(ws.path, ".dls", "translations.json")
	err := readConf(targetDir, &conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error reading the translation config: %v", err)
		return translationConfiguration{}
	}
	return conf
}

func initTranslationConfigWithPath(path string) translationConfiguration {
	var conf translationConfiguration
	targetDir := filepath.Join(path, ".dls", "translations.json")
	err := readConf(targetDir, &conf)

    fmt.Fprintf(os.Stderr, "tr config: %v", conf)
	if err != nil {
    fmt.Fprintf(os.Stderr, "There was an error reading the translation config: %v", err)
		return translationConfiguration{}
	}
	return conf
}
