package utils

import (
	"errors"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
)

// TestHCLUtil - Structure to hold info about unit tests
type TestHCLUtil struct {
	Tests []TestUtil `hcl:"test,block"`
}

// ParseUnitTests - Function used to parse the HCL of the unit tests
func ParseUnitTests(filepath string) (TestHCLUtil, error) {
	parser := hclparse.NewParser()
	parseData, parseDiags := parser.ParseHCLFile(filepath)

	var testList TestHCLUtil

	if parseDiags.HasErrors() {
		return testList, errors.New("Parse Diags: " + parseDiags.Error())
	}

	decodeDiags := gohcl.DecodeBody(parseData.Body, nil, &testList)

	if decodeDiags.HasErrors() {
		return testList, errors.New("Decode Diags: " + decodeDiags.Error())
	}

	return testList, nil
}
