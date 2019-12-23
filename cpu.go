package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type CPU interface {
	load(string)
	execute()
	input()
	output()
}

type cpu struct {
	ProgMem []string
	DataMem []int
	pc      int
	dc      int
}

func NewCPU(code string) *cpu {
	c := new(cpu)
	c.DataMem = make([]int, 30000)
	c.load(code)
	return c
}

func (c *cpu) load(code string) {
	c.ProgMem = strings.Split(code, "")
}

func (c *cpu) test() {
	fmt.Println(findMatching(c.ProgMem, 5, false))
}

func (c *cpu) execute() {
	// fmt.Println("U:", len(c.ProgMem))
	for c.pc < len(c.ProgMem) {
		opcode := c.ProgMem[c.pc]
		// fmt.Println(c.dc)
		switch opcode {
		case "+":
			c.DataMem[c.dc] += 1
			c.pc++
		case "-":
			c.DataMem[c.dc] -= 1
			c.pc++
		case ">":
			c.dc++
			c.pc++
		case "<":
			c.dc--
			c.pc++
		case ".":
			c.output()
			c.pc++
		case ",":
			c.input()
			c.pc++
		case "[":
			if c.DataMem[c.dc] == 0 {
				c.pc = findMatching(c.ProgMem, c.pc, true) + 1

			} else {
				c.pc++
			}
		case "]":
			if c.DataMem[c.dc] != 0 {
				c.pc = findMatching(c.ProgMem, c.pc, false) + 1
			} else {
				c.pc++
			}
		default:
			c.pc++
		}
		// time.Sleep(time.Millisecond * 1000)
	}

	fmt.Println()
}

func findMatching(code []string, pos int, dir bool) int {
	count := 1
	if dir {
		for i := pos + 1; i < len(code); i++ {
			if code[i] == "]" {
				count--
				if count == 0 {
					return i
				}
			} else if code[i] == "[" {
				count++
			}
		}
	} else {
		for i := pos - 1; i >= 0; i-- {
			if code[i] == "[" {
				count--
				if count == 0 {
					return i
				}
			} else if code[i] == "]" {
				count++
			}
		}
	}

	return -1
}

func (c *cpu) input() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	c.DataMem[c.dc] = int(char)
}

func (c *cpu) output() {
	fmt.Print(string(c.DataMem[c.dc]))
}
