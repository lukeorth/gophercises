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
    k := 27

    var cipher string

    for _, ch := range input {
        if unicode.IsUpper(ch) {
            rot := 65 + ((k - (90 - int(ch) + 1) + 26) % 26)
            cipher += string(rot)
            continue
        }
        if unicode.IsLower(ch) {
            rot := 97 + ((k - (122 - int(ch) + 1) + 26) % 26)
            cipher += string(rot)
            continue
        }
        cipher += string(ch)
    }

    fmt.Println(cipher)
}
