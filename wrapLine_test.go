package main

import (
	"go/format"
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
