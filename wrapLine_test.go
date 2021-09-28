package main

import (
	"go/format"
	"io/ioutil"
	"log"
	"testing"
)

func Test_wrapLine(t *testing.T) {
	text := `
package main
const HandleFullErr = QueueErrType("handleFullError","handleFullError","handleFullError","handleFullError","handleFullError","handleFullError")
/*const HandleFullErr = QueueErrType("handleFullError","handleFullError","handleFullError","handleFullError","handleFullError","handleFullError")*/
// const HandleFullErr = QueueErrType("handleFullError","handleFullError","handleFullError","handleFullError","handleFullError","handleFullError")
const HandleFullErr1 = QueueErrType("handleFullError, handleFullError, handleFullError, handleFullError, handleFullError,handleFullError,handleFullError,handleFullError,handleFullError,handleFullError,handleFullError")`
	ret := wrapLine(text)
	fmtRet, err := format.Source([]byte(ret))
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("ret:%s", fmtRet)

}

func Test_wrapLine2(t *testing.T) {
	bts, err := ioutil.ReadFile("./wrapLineExample.txt")
	if err != nil {
		t.Fatal(err)
	}
	ret := wrapLine(string(bts))
	fmtRet, err := format.Source([]byte(ret))
	log.Printf("ret:%s", fmtRet)
}
