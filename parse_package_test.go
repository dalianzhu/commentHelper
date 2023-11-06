package main

import (
	"log"
	"testing"
)

func TestParsePackage(t *testing.T) {
	text := `package controller
import (
    "context"
)

`

	willComments := ParsePackage(text)
	for _, item := range willComments {
		log.Printf("line:%v, funcName:=%v\n", item.OriginLineNo, item.Name)
	}
}
