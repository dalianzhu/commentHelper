package main

import (
	"log"
	"testing"
)

// TestParseType ...
func TestParseType(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []*NeedCommentLine
	}{
		{
			name: "test1",
			args: args{
				text: `
const HandleFullErr = QueueErrType("handleFullError")
const handleWillFullErr = QueueErrType("handleWillFullError")
const HandleMsgErr = QueueErrType("handleMsgError")

var HandlerMsgErr1 string
var handlerMsgErr2 string
var HandlerMsgErr3 = func()

type Msg string
type msg HandlerMsg
type Msg2 HandlerMsg
type Msg3 *HandlerMsg
type Msg4 func()
`,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseType(tt.args.text)
			for _, v := range got {
				log.Printf("TestParseType got:%v, %v", v.Name, v.OriginLineNo)
			}
		})
	}
}
