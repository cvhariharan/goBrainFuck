package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		dat, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Println(err)
		}

		code := string(dat)

		if isValid(code) {
			c := NewCPU(code)
			c.execute()
		}
	} else {
		log.Fatal("File not specified")
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
