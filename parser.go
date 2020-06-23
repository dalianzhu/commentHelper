package main

import (
    "fmt"
    "log"
    "sort"
    "strings"
)

// WillCommentLine 这一行需要加注释
type WillCommentLine struct {
    OriginLineNo int
    Name         string
}

func ExtractText(text string) string {
    originTextLines := strings.Split(text, "\n")

    isInCommentMap := make(map[int]bool)
    isInComment := false
    for i, line := range originTextLines {
        if isInComment {
            isInCommentMap[i] = true
        }
        trimLine := strings.Trim(line, " ")
        if strings.HasPrefix(trimLine, "//") {
            log.Printf("line:%v is in comment", i)
            isInCommentMap[i] = true
        }
        if strings.HasPrefix(trimLine, "/*") {
            // 进入大注释
            isInCommentMap[i] = true
            isInComment = true
        }
        if strings.HasSuffix(trimLine, "*/") {
            isInCommentMap[i] = true
            isInComment = false
        }
    }

    willComments := ParseFunc(text)
    willComments = append(willComments, ParseStruct(text)...)
    willComments = append(willComments, ParseType(text)...)
    sort.Slice(willComments, func(i, j int) bool {
        return willComments[i].OriginLineNo < willComments[j].OriginLineNo
    })
    offset := 0
    for _, c := range willComments {
        willInsertLine := c.OriginLineNo + offset
        log.Printf("willInsertLine:%v", willInsertLine)
        _, ok := isInCommentMap[c.OriginLineNo-1]
        if !ok {
            comment := fmt.Sprintf("// %v ...", c.Name)
            originTextLines = InsertSlice(originTextLines, willInsertLine, comment)
            offset++
        }
    }

    return strings.Join(originTextLines, "\n")
}

func InsertSlice(ss []string, index int, inserted string) []string {
    rear := append([]string{}, ss[index:]...)
    ss = append(ss[0:index], inserted)
    ss = append(ss, rear...)
    return ss
}
