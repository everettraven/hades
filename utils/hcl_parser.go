package utils

import (
	"errors"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
)

// UnitTestHCLUtil - Structure to hold info about unit tests
type UnitTestHCLUtil struct {
	UnitTests []UnitTestUtil `hcl:"unittest,block"`
}

// ParseUnitTests - Function used to parse the HCL of the unit tests
func ParseUnitTests(filepath string) (UnitTestHCLUtil, error) {
	parser := hclparse.NewParser()
	parseData, parseDiags := parser.ParseHCLFile(filepath)

	var testList UnitTestHCLUtil

	if parseDiags.HasErrors() {
		return testList, errors.New("Parse Diags: " + parseDiags.Error())
	}

	decodeDiags := gohcl.DecodeBody(parseData.Body, nil, &testList)

	if decodeDiags.HasErrors() {
		return testList, errors.New("Decode Diags: " + decodeDiags.Error())
	}

	return testList, nil
}
