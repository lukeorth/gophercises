package main

import (
	"bytes"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_ParseProblems(t *testing.T) {
    tests := map[string]struct {
        input   bytes.Buffer
        want    problem
        err     error
    }{
        "simple":           {input: *bytes.NewBufferString("1+1,2"), want: problem{question: "1+1", answer: "2"}, err: nil},
        "wrong sep":        {input: *bytes.NewBufferString("1+1/2"), want: problem{question: "1+1", answer: "2"}, err: ErrFormat},
        "no sep":           {input: *bytes.NewBufferString("1+1 2"), want: problem{question: "1+1", answer: "2"}, err: ErrFormat},
        "trailing comma":   {input: *bytes.NewBufferString("1+1,2,"), want: problem{question: "1+1", answer: "2"}, err: nil},
        "comma in string":  {input: *bytes.NewBufferString(`"what's 1+1, sir?",2`), want: problem{question: "what's 1+1, sir?", answer: "2"}, err: nil},
        "extra data":       {input: *bytes.NewBufferString("1+1,2,3+3,6"), want: problem{question: "1+1", answer: "2"}, err: nil},
        "not enough data":  {input: *bytes.NewBufferString("1+1"), want: problem{question: "1+1", answer: ""}, err: ErrFormat},
    }

    for name, tc := range tests {
        t.Run(name, func(t *testing.T) {
            p, err := parseProblems(&tc.input)

            if !errors.Is(err, tc.err) {
                t.Fatalf("unexpected error: %v", err)
            }
            
            if err == nil {
                diff := cmp.Diff(tc.want, p[0], cmp.AllowUnexported(problem{}))
                if diff != "" {
                    t.Fatalf(diff)
                }
            }
        })
    }
}
