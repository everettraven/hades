package utils

import (
	"errors"

	"github.com/everettraven/hades/resources"
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
)

// UnitTestHCLUtil - Structure to hold info about unit tests
type UnitTestHCLUtil struct {
	UnitTests []UnitTestUtil `hcl:"unittest,block"`
}

//Test - struct to hold test information
type Test struct {
	Title string               `hcl:"title,attr"`
	Cmds  []*resources.Command `hcl:"command,block"`
	Os    *resources.OS        `hcl:"os,block"`
}

//Host - struct to hold host information
type Host struct {
	IP   string `hcl:"ip"`
	Port string `hcl:"port,optional"`
	User string `hcl:"user,optional"`
}

//HostHCLUtil - struct to hold host info read from a hosts HCL file
type HostHCLUtil struct {
	Hosts []Host `hcl:"host,block"`
}

//NewHost - Function to create a new host object
func NewHost(ip string, port string) Host {
	host := new(Host)
	host.IP = ip
	host.Port = port
	return *host
}

// Parse - Function used to parse the HCL according to the type passed in
func Parse(filepath string, out interface{}) (interface{}, error) {
	parser := hclparse.NewParser()
	parseData, parseDiags := parser.ParseHCLFile(filepath)

	switch out.(type) {
	default:
		return nil, errors.New("Unexpected type passed in to the HCL Parse function")
	case Test:
		var testList Test

		if parseDiags.HasErrors() {
			return testList, errors.New("Parse Diags: " + parseDiags.Error())
		}

		decodeDiags := gohcl.DecodeBody(parseData.Body, nil, &testList)

		if decodeDiags.HasErrors() {
			return testList, errors.New("Decode Diags: " + decodeDiags.Error())
		}

		return testList, nil

	case UnitTestHCLUtil:
		var testList UnitTestHCLUtil
		if parseDiags.HasErrors() {
			return testList, errors.New("Parse Diags: " + parseDiags.Error())
		}

		decodeDiags := gohcl.DecodeBody(parseData.Body, nil, &testList)

		if decodeDiags.HasErrors() {
			return testList, errors.New("Decode Diags: " + decodeDiags.Error())
		}

		return testList, nil

	case HostHCLUtil:
		var testList HostHCLUtil
		if parseDiags.HasErrors() {
			return testList, errors.New("Parse Diags: " + parseDiags.Error())
		}

		decodeDiags := gohcl.DecodeBody(parseData.Body, nil, &testList)

		if decodeDiags.HasErrors() {
			return testList, errors.New("Decode Diags: " + decodeDiags.Error())
		}

		return testList, nil

	}

}
