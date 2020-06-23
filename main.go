package main

import (
    "io/ioutil"
    "log"
    "os"
)

func main() {
    path := os.Args[1]
    textBytes, err := ioutil.ReadFile(path)
    if err != nil {
        log.Println(err)
        return
    }
    text := string(textBytes)
    text = ExtractText(text)
    err = ioutil.WriteFile(path, []byte(text), 0666)
    if err != nil {
        log.Println(err)
        return
    }
}
