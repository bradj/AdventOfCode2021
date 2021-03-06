package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func d2p1() {
	bcontent, err := ioutil.ReadFile("2.input")

	if err != nil {
		log.Fatal(err)
	}

	content := string(bcontent)
	items := strings.Split(content, "\n")

	x := 0
	y := 0

	for _, v := range items {
		instruction := strings.Split(v, " ")

		if len(instruction) != 2 {
			continue
		}

		inc, err := strconv.Atoi(instruction[1])

		if err != nil {
			fmt.Printf("Error: %s - %v - %v", v, instruction, err)
			continue
		}

		switch instruction[0] {
		case "forward":
			x += inc
			break
		case "down":
			y += inc
			break
		case "up":
			y -= inc
			break
		}
	}

	fmt.Printf("final position: %d\n", x*y)
}
