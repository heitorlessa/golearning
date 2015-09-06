package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tlds tld

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

//var tlds = []string{"com", "net"}

type tld []string

// String method required for flag.Value interface that must return a string
func (s *tld) String() string {
	return fmt.Sprint(*s)
}

// Set method required for flag.Value interface that must return error
func (s *tld) Set(value string) error {

	if len(*s) > 5 {
		return errors.New("Too many domains - try up to 5 only")
	}

	// accept multiple values separated by comma instead of passing multiple times
	for _, values := range strings.Split(value, ",") {
		*s = append(*s, values)
	}
	return nil
}

func main() {

	flag.Var(&tlds, "tld", "List of domains separated by comma")
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(2)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}

			if !strings.ContainsRune(allowedChars, r) {
				continue
			}

			newText = append(newText, r)
		}
		fmt.Println(string(newText) + tlds[rand.Intn(len(tlds))])
	}
}
