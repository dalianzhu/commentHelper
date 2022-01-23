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

func NewCommonPBFilter[T proto.Message](ttype string,
	vals []string, getters map[string]func(T) any) (*CommonPBFilter[T], error) {
}

func (c *CommonPBFilter[TPB]) DoFilter(data interfaces.BusinessData) (interfaces.BusinessData, error) {
}
`

	willComments := ParseFunc(text)
	for _, item := range willComments {
		log.Printf("line:%v, funcName:=%v\n", item.OriginLineNo, item.Name)
	}
}
