package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord,
	otherWord,
	otherWord,
	otherWord + " app",
	otherWord + " site",
	otherWord + " time",
	"get " + otherWord,
	"go " + otherWord,
	"lets " + otherWord,
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func writeToFile(content []byte, filename string) {
	if err := ioutil.WriteFile(filename, content, 0644); err != nil {
		panic(err)
	}
}

func writeAppendToFile(content []byte, filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	checkError(err)

	defer f.Close()

	if _, err := f.Write(content); err != nil {
		panic(err)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		content := []byte(strings.Replace(t, otherWord, s.Text(), -1))
		writeToFile(content, "result.out")
		writeAppendToFile(content, "result.append")
		fmt.Println(string(content))
	}
}
