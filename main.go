package main

import (
	"io/ioutil"
	"log"
	"strings"
)

var Memory [30000]int

func main() {
	dat, err := ioutil.ReadFile("hello.bf")
	if err != nil {
		log.Println(err)
	}

	code := string(dat)

	if isValid(code) {
		c := NewCPU(code)
		c.execute()
	}

}

// Check for the correct placement of []
func isValid(code string) bool {
	count := 0
	for _, op := range strings.Split(code, "") {
		if op == "]" {
			count--
			if count < 0 {
				// ] is before [
				log.Println("] before [")
				return false
			}
		} else if op == "[" {
			count++
		}
	}

	if count < 0 {
		log.Println("Mismatched []")
		return false
	}

	return true
}
