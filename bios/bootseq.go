package bios

import (
	"fmt"
	"io/ioutil"
)

type BootSeq struct {
	Keys         map[string]string `json:"keys"`
	Contents     []*ContentRef     `json:"contents"`
	BootSequence []*OperationType  `json:"boot_sequence"`
}

func ReadBootSeq(filename string) (out *BootSeq, err error) {
	rawBootSeq, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading boot seq: %s", err)
	}

	if err := yamlUnmarshal(rawBootSeq, &out); err != nil {
		return nil, fmt.Errorf("parsing boot seq yaml: %s", err)
	}

	return
}

type ContentRef struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Hash string `json:"hash"`
}
