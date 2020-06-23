package main

import (
    "log"
    "regexp"
    "strings"
)

func ParseFunc(text string) []*WillCommentLine {
    // func (r *ResenderMsg) IsSendCtx() context.Context {
    re, err := regexp.Compile(`^func\s+([A-Z][a-zA-Z0-9_-]+)\(|^func\s*\([a-zA-Z0-9]+\s+[*a-zA-Z0-9]+\)\s+([A-Z][a-zA-Z0-9_-]+)`)
    if err != nil {
        log.Println(err)
        return []*WillCommentLine{}
    }
    lines := strings.Split(text, "\n")
    willComments := make([]*WillCommentLine, 0)
    for i, line := range lines {
        ret := re.FindAllStringSubmatch(line, -1)
        if len(ret) >= 1 {
            matchArr := ret[0]
            if len(matchArr) >= 3 {
                var funcName string
                if matchArr[1] != "" {
                    funcName = matchArr[1]
                } else {
                    funcName = matchArr[2]
                }

                log.Printf("func name:%v\n", funcName)
                tp := new(WillCommentLine)
                tp.Name = funcName
                tp.OriginLineNo = i
                willComments = append(willComments, tp)
            }
        }
    }
    return willComments
}

func filterRegex(re *regexp.Regexp, lines []string) []*WillCommentLine {
    willComments := make([]*WillCommentLine, 0)
    for i, line := range lines {
        ret := re.FindAllStringSubmatch(line, -1)
        if len(ret) >= 1 {
            matchArr := ret[0]
            if len(matchArr) >= 2 {
                var funcName string
                if matchArr[1] != "" {
                    funcName = matchArr[1]
                } else {
                    if len(matchArr) >= 3 {
                        funcName = matchArr[2]
                    }
                }

                log.Printf("func name:%v\n", funcName)
                tp := new(WillCommentLine)
                tp.Name = funcName
                tp.OriginLineNo = i
                willComments = append(willComments, tp)
            }
        }
    }
    return willComments
}
