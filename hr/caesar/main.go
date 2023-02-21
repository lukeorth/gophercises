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
            rot := 'A' + (int(ch) - 'A' + k) % 26
            cipher += string(rot)
            continue
        }
        if unicode.IsLower(ch) {
            rot := 'a' + (int(ch) - 'a' + k) % 26
            cipher += string(rot)
            continue
        }
        cipher += string(ch)
    }

    fmt.Println(cipher)
}
