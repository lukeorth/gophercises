package main

import (
	"bytes"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	phonedb "github.com/lukeorth/gophercises/phone/db"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    dbname   = "gophercises_phone"
)

func main() {
    // gracefully handle errors and exit
    var err error
    defer func() {
        if err != nil {
            log.Fatalln(err)
        }
    }()

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s sslmode=disable", host, port, user)
    err = phonedb.Reset("postgres", psqlInfo, dbname)

    psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
    err = phonedb.Migrate("postgres", psqlInfo)

    var db *phonedb.DB
    db, err = phonedb.Open("postgres", psqlInfo)
    defer db.Close()
    
    err = db.Seed()

    var phones []phonedb.Phone
    phones, err = db.AllPhones()
    for _, p := range phones {
        fmt.Printf("Working on... %+v\n", p)
        number := normalize(p.Number)
        if number != p.Number {
            fmt.Println("Updating or removing...", number)
            var existing *phonedb.Phone
            existing, err = db.FindPhone(number)
            if existing != nil {
                err = db.DeletePhone(p.ID)
            } else {
                p.Number = number
                err = db.UpdatePhone(&p)
            }
        } else {
            fmt.Println("No changes required")
        }
    }
}

func normalize(phone string) string {
    var buf bytes.Buffer
    for _, ch := range phone {
        if ch >= '0' && ch <= '9' {
            buf.WriteRune(ch)
        }
    }
    return buf.String()
}
