package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filePath string) []Person {
	f, err := os.Open(filePath)

	s := []Person{}

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")

		if len(parts) < 3 {
			log.Println("Dữ liệu không hợp lệ:", line)
			continue
		}

		year, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Fatal(err)
		}
		p := Person{name: parts[0], job: parts[1], year: year}
		s = append(s, p)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return s
}
