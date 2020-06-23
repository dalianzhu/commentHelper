package main

import (
    "log"
    "regexp"
    "strings"
)

func ParseType(text string) []*WillCommentLine {
    re, err := regexp.Compile(`var\s+([A-Z][a-zA-Z0-9]*)\s+[a-zA-Z0-9*=]+|const\s+([A-Z][a-zA-Z0-9]*)\s+\=`)
    if err != nil {
        log.Println(err)
        return []*WillCommentLine{}
    }
    lines := strings.Split(text, "\n")
    willComments := filterRegex(re, lines)
    return willComments
}
