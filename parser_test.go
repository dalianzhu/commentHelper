package main

import (
	"testing"
)

// TestExtractText ...
func TestExtractText(t *testing.T) {

	ret := AddCommentToText(origin)
	ret = AddCommentToText(ret)
	// fmt.Println(ret)
	for i := range []byte(ret) {
		if ret[i] != []byte(checkRet)[i] {
			t.Fail()
			return
		}
	}
}

var origin = `package controller

import (
    "context"
    log "github.com/sirupsen/logrus"
    "reflect"
    "sync"
    "sync/atomic"
    "time"
)

type abc *string
// Abc ...
type Abc *string

// 一个滑动窗口的resender

var hello int
var Hello int

func NewResendMsg(ctx context.Context, uuid string, timestamp int64, data interface{}) *ResenderMsg {
	var CST = time.Now()
}


type ResenderMsg struct {
    Data         interface{}
    isSend       context.Context
}

func (r *ResenderMsg) isSendCtx() context.Context {
    return r.isSend
}

func (r *ResenderMsg) IsSendCancel() {
    r.isSendCancel()
}

func NewMsgResender(ctx context.Context, subId int, sender func(data *ResenderMsg) error) *MsgResender {
}

// MsgResender ...
type MsgResender struct {
    subId       int
}

func (m *MsgResender) UnSendCount() int64 {
    return m.unsendCount
}

// Wait ...
func (m *MsgResender) Wait() {
}

func (m *MsgResender) Put(key string, data *ResenderMsg) {
}

// Pop ...
func (m *MsgResender) Pop(f func(key string, data *ResenderMsg) bool) bool {
}

func (m *MsgResender) IsSend(key string) {

}

func (m *MsgResender) Maintenance() {
}

func NewCommonPBFilter[T proto.Message](ttype string,
	vals []string, getters map[string]func(T) any) (*CommonPBFilter[T], error) {
}

type CommonPBFilter[TPB proto.Message] struct {
	Getter func(TPB) any
	vals   map[any]bool
}

func (c *CommonPBFilter[TPB]) DoFilter(data interfaces.BusinessData) (interfaces.BusinessData, error) {
}
func (s *DefaultPBConverter[T, U]) Convert(
	msg interfaces.OriginMessage) (businessData interfaces.BusinessData, err error) {
}

type PbPtr[U any] interface {
	*U // U:protoStruct T: *protoStruct
	protoreflect.ProtoMessage
}`
var checkRet = `package controller

import (
    "context"
    log "github.com/sirupsen/logrus"
    "reflect"
    "sync"
    "sync/atomic"
    "time"
)

type abc *string
// Abc ...
type Abc *string

// 一个滑动窗口的resender

var hello int
// Hello ...
var Hello int

// NewResendMs ...
func NewResendMsg(ctx context.Context, uuid string, timestamp int64, data interface{}) *ResenderMsg {
// CST ...
	var CST = time.Now()
}


// ResenderMsg ...
type ResenderMsg struct {
    Data         interface{}
    isSend       context.Context
}

func (r *ResenderMsg) isSendCtx() context.Context {
    return r.isSend
}

// IsSendCancel ...
func (r *ResenderMsg) IsSendCancel() {
    r.isSendCancel()
}

// NewMsgResender ...
func NewMsgResender(ctx context.Context, subId int, sender func(data *ResenderMsg) error) *MsgResender {
}

// MsgResender ...
type MsgResender struct {
    subId       int
}

// UnSendCount ...
func (m *MsgResender) UnSendCount() int64 {
    return m.unsendCount
}

// Wait ...
func (m *MsgResender) Wait() {
}

// Put ...
func (m *MsgResender) Put(key string, data *ResenderMsg) {
}

// Pop ...
func (m *MsgResender) Pop(f func(key string, data *ResenderMsg) bool) bool {
}

// IsSend ...
func (m *MsgResender) IsSend(key string) {

}

// Maintenance ...
func (m *MsgResender) Maintenance() {
}

// NewCommonPBFilter ...
func NewCommonPBFilter[T proto.Message](ttype string,
	vals []string, getters map[string]func(T) any) (*CommonPBFilter[T], error) {
}

// CommonPBFilter ...
type CommonPBFilter[TPB proto.Message] struct {
	Getter func(TPB) any
	vals   map[any]bool
}

// DoFilter ...
func (c *CommonPBFilter[TPB]) DoFilter(data interfaces.BusinessData) (interfaces.BusinessData, error) {
}
// Convert ...
func (s *DefaultPBConverter[T, U]) Convert(
	msg interfaces.OriginMessage) (businessData interfaces.BusinessData, err error) {
}

// PbPtr ...
type PbPtr[U any] interface {
	*U // U:protoStruct T: *protoStruct
	protoreflect.ProtoMessage
}`
