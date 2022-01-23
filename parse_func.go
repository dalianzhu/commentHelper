package main

import (
	"log"
	"regexp"
	"strings"
)

// ParseFunc ...
func ParseFunc(text string) []*NeedCommentLine {
	// func (r *ResenderMsg) IsSendCtx() context.Context {
	// func (s *DefaultPBConverter[T, U]) Convert(
	re, err := regexp.Compile(
		`^func\s+([A-Z][a-zA-Z0-9_-]+)\[*.+\(|^func\s*\([a-zA-Z0-9]+\s+[ *a-zA-Z0-9,\[\]\.]+\)\s+([A-Z][a-zA-Z0-9_-]+)`)
	if err != nil {
		log.Println(err)
		return []*NeedCommentLine{}
	}
	lines := strings.Split(text, "\n")
	willComments := make([]*NeedCommentLine, 0)
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

				// log.Printf("func name:%v\n", funcName)
				tp := new(NeedCommentLine)
				tp.Name = funcName
				tp.OriginLineNo = i
				willComments = append(willComments, tp)
			}
		}
	}
	return willComments
}

func filterRegex(re *regexp.Regexp, lines []string) []*NeedCommentLine {
	willComments := make([]*NeedCommentLine, 0)
	for i, line := range lines {
		ret := re.FindAllStringSubmatch(line, -1)
		if len(ret) >= 1 {
			matchArr := ret[0]
			var funcName string
			if len(matchArr) >= 2 {
				for _, v := range matchArr[1:] {
					if strings.Trim(v, "\t\r\n ") != "" {
						funcName = v
						break
					}
				}

				// log.Printf("func name:%v\n", funcName)
				tp := new(NeedCommentLine)
				tp.Name = funcName
				tp.OriginLineNo = i
				willComments = append(willComments, tp)
			}
		}
	}
	return willComments
}
