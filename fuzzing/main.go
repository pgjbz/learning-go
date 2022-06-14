package main

import (
    "errors"
    "unicode/utf8"
)

func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}

func main() {
    input := "Socorram me subi no onibus em marrocos"
    rev, _ := Reverse(input)
    doubleRev, _ := Reverse(rev)
    println("Original:", input)
    println("Reverse:", rev)
    println("Double reverse:", doubleRev)
}