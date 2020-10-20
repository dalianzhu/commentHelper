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
    log "github.com/sirupsen/logrus"
    "reflect"
    "sync"
    "sync/atomic"
    "time"
)

// 一个滑动窗口的resender

// NewResendMsg ...
func NewResendMsg(ctx context.Context, uuid string, timestamp int64, data interface{}) *ResenderMsg {
    r := &ResenderMsg{}
    r.Uuid = uuid
    r.Timestamp = timestamp
    r.Data = data
    r.isSend, r.isSendCancel = context.WithCancel(ctx)
    return r
}

// 重新发送未确认的msg
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

// IsSendCancel ...
func (r *ResenderMsg) IsSendCancel() {
    r.isSendCancel()
}

// NewMsgResender ...
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

// MsgResender ...
// MsgResender ...
type MsgResender struct {
    subId       int
    dataMap     sync.Map
    unsendCount int64
    sender      func(data *ResenderMsg) error
    queue       chan *ResenderMsg
    ctx         context.Context
}

// UnSendCount ...
func (m *MsgResender) UnSendCount() int64 {
    return m.unsendCount
}

// Wait ...
func (m *MsgResender) Wait() {
    for m.unsendCount != 0 {
        time.Sleep(time.Second)
    }
}

// Put ...
func (m *MsgResender) Put(key string, data *ResenderMsg) {
    data.Uuid = key
    atomic.AddInt64(&m.unsendCount, 1)
    m.dataMap.Store(key, data)
    select {
    case <-m.ctx.Done():
    case m.queue <- data:
    }

}

// Pop ...
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

// IsSend ...
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

// Maintenance ...
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

    text = `
// 一个滑动窗口的resender

// NewResendMsg ...
func NewResendMsg(ctx context.Context, uuid string, timestamp int64, data interface{}) *ResenderMsg {
    r := &ResenderMsg{}
    r.Uuid = uuid
    r.Timestamp = timestamp
    r.Data = data
    r.isSend, r.isSendCancel = context.WithCancel(ctx)
    return r
}

`
    willComments := ParseFunc(text)
    for _,item:=range willComments{
        log.Printf("line:%v, funcName:=%v\n", item.OriginLineNo, item.Name)
    }
}
