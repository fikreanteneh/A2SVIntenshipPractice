package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordCounter(s string) map[string]int {
    var counter = make(map[string]int)
    for _, letter := range s {
        if !unicode.IsLetter(letter) {
            continue
        }
        counter[strings.ToLower(string(letter))]++
    }
    return counter
}

func main() {
	fmt.Println(wordCounter("Hello, World!"))

}