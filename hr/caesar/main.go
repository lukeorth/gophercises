package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
    // handle input
    s := bufio.NewScanner(os.Stdin)
    s.Scan()
    input := s.Text()
    k := 3

    var cipher string

    for _, ch := range input {
        if unicode.IsUpper(ch) {
            rot := 65 + (int(ch) - 65 + k) % 26
            cipher += string(rot)
            continue
        }
        if unicode.IsLower(ch) {
            rot := 97 + (int(ch) - 97 + k) % 26
            cipher += string(rot)
            continue
        }
        cipher += string(ch)
    }

    fmt.Println(cipher)
}
