package langserver

import (
	"fmt"
	"strings"
)

type FuncSignature struct {
	ReturnType string
	Parameters []string
}

func newFuncSignature(ret string, params ...string) FuncSignature {
	return FuncSignature{
		ReturnType: ret,
		Parameters: params,
	}
}

func parseFuncSigantureDirective(input string) (FuncSignature, error) {
	parts := strings.Split(input, ",")
	if len(parts) < 1 {
		return FuncSignature{}, fmt.Errorf("signature directive needs to contain at least one type")
	}
	if len(parts) > 1 { // not sure if this is really needed
		return newFuncSignature(parts[0], parts[1:]...), nil
	} else {
		return newFuncSignature(parts[0]), nil
	}
}

func getFuncSignatureString(input []string) (string, error) {
	directive := strings.Join(input, ",")
	sign, err := parseFuncSigantureDirective(directive)
	if err == nil {
		return sign.signatureString(), nil
	}
	return "", err
}

func (fs FuncSignature) signatureString() string {
	pars := fs.Parameters
	for i, p := range pars {
		pars[i] = "var " + p
	}
	return "func " + fs.ReturnType + " f (" + strings.Join(pars, ",") + ")"

}

func (fs FuncSignature) isEqual(funcSig FunctionSymbol) bool {
	if !strings.EqualFold(fs.ReturnType, funcSig.ReturnType) {
		return false
	}
	if len(fs.Parameters) != len(funcSig.Parameters) {
		return false
	}
	for index, p := range fs.Parameters {
		if !strings.EqualFold(p, funcSig.Parameters[index].Type) {
			return false
		}
	}
	return true
}
