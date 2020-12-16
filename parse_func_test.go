package main

import (
	"log"
	"testing"
)

// TestParseFunc ...
func TestParseFunc(t *testing.T) {

	text := `package controller
import (
    "context"
)


func NewResendMsg(ctx context.Context, uuid string, timestamp int64, data interface{}) *ResenderMsg {
}

func newResendMsg(ctx context.Context, uuid string, timestamp int64, data interface{}) *ResenderMsg {
}

type ResenderMsg struct {
}

func (r *ResenderMsg) isSendCtx() context.Context {
}

func (r *ResenderMsg) IsSendCancel() {
}
`

	willComments := ParseFunc(text)
	for _, item := range willComments {
		log.Printf("line:%v, funcName:=%v\n", item.OriginLineNo, item.Name)
	}
}
