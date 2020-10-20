package main

import (
	"log"
	"testing"
)

// TestExtractText ...
func TestExtractText(t *testing.T) {
	text := `package controller

import (
    "context"
    log "github.com/sirupsen/logrus"
    "reflect"
    "sync"
    "sync/atomic"
    "time"
)

type abc *string
type Abc *string

// 一个滑动窗口的resender

var hello int
var Hello int

func NewResendMsg(ctx context.Context, uuid string, timestamp int64, data interface{}) *ResenderMsg {
    r := &ResenderMsg{}
    r.Uuid = uuid
    r.Timestamp = timestamp
    r.Data = data
    r.isSend, r.isSendCancel = context.WithCancel(ctx)
    return r
}


type ResenderMsg struct {
    Data         interface{}
    Timestamp    int64
    SendTimes    int
    isSend       context.Context
    isSendCancel context.CancelFunc
    Uuid         string
}

func (r *ResenderMsg) isSendCtx() context.Context {
    return r.isSend
}

func (r *ResenderMsg) IsSendCancel() {
    r.isSendCancel()
}

func NewMsgResender(ctx context.Context, subId int, sender func(data *ResenderMsg) error) *MsgResender {
    s := new(MsgResender)
    s.ctx = ctx
    s.subId = subId
    s.queue = make(chan *ResenderMsg, 50)
    s.dataMap = sync.Map{}
    s.unsendCount = 0
    s.sender = sender
    return s
}

type MsgResender struct {
    subId       int
    dataMap     sync.Map
    unsendCount int64
    sender      func(data *ResenderMsg) error
    queue       chan *ResenderMsg
    ctx         context.Context
}

func (m *MsgResender) UnSendCount() int64 {
    return m.unsendCount
}

func (m *MsgResender) Wait() {
    for m.unsendCount != 0 {
        time.Sleep(time.Second)
    }
}

func (m *MsgResender) Put(key string, data *ResenderMsg) {
    data.Uuid = key
    atomic.AddInt64(&m.unsendCount, 1)
    m.dataMap.Store(key, data)
    select {
    case <-m.ctx.Done():
    case m.queue <- data:
    }
}

func (m *MsgResender) Pop(f func(key string, data *ResenderMsg) bool) bool {
    var ret bool
    select {
    case <-m.ctx.Done():
    case d := <-m.queue:
        ret = f(d.Uuid, d)
        m.dataMap.Delete(d.Uuid)
    }
    return ret
}

func (m *MsgResender) IsSend(key string) {
    v, ok := m.dataMap.Load(key)
    if !ok {
        log.Infof("MsgResender data not found:%v", key)
        return
    }

    msg, ok := v.(*ResenderMsg)
    if !ok {
        log.Errorf("MsgResender data type error:%v", reflect.TypeOf(v))
        return
    }
    msg.IsSendCancel()
    atomic.AddInt64(&m.unsendCount, -1)
}

func (m *MsgResender) Maintenance() {
    m.dataMap.Range(func(key, value interface{}) bool {
        data, ok := value.(*ResenderMsg)
        if !ok {
            log.Errorf("MsgResender Maintenance type error:%v, subId:%v", reflect.TypeOf(value), m.subId)
        }

        select {
        case <-data.IsSendCtx().Done():
            // 已经发送成功的消息跳过去
            return true
        case <-m.ctx.Done():
            return false
        default:
            if time.Now().Unix()-data.Timestamp > 10 {
                data.SendTimes++
                err := m.sender(data)
                if err != nil {
                    log.Errorf("MsgResender Maintenance send msg error:%v, subId:%v", err, m.subId)
                    return false
                }
            }
            return true
        }
    })
}`
	ret := AddCommentToText(text)
	log.Println(ret)
}
