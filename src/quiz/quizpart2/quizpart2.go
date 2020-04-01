package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	path, err := os.Getwd()
	check(err)
	p := filepath.Join(path, "/src/quiz/quizpart2/problems.csv")

	right := 0
	questions := 0

	f, err := os.Open(p)
	check(err)
	defer f.Close()

	r := csv.NewReader(bufio.NewReader(f))
	records, err := r.ReadAll()
	check(err)

	s := bufio.NewScanner(os.Stdin)

	fmt.Println("Click Enter when you are ready to start")
	s.Scan()
	timeout := 30
	if s.Text() != "" {
		timeout, err = strconv.Atoi(s.Text())
		check(err)
	}
	fmt.Println("Time for the test", timeout)

	fmt.Println("Enter Y/y if you want to shuffle the questions", timeout)
	s.Scan()
	if strings.ToLower(strings.TrimSpace(s.Text())) == "y" {
		fmt.Println("Shuffle")
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

	go func() {
		timer1 := time.NewTimer(time.Duration(timeout) * time.Second)
		<-timer1.C
		fmt.Println("Time Out")
		printResult(right, len(records))
		os.Exit(2)
	}()

	for _, record := range records {
		fmt.Println(record[0])
		questions++
		s.Scan()
		resp := strings.TrimSpace(s.Text())
		if resp == strings.TrimSpace(record[1]) {
			right++
		}
	}

	printResult(right, len(records))
}

func printResult(right, questions int) {
	fmt.Printf("%v right answers out of %v questions\n", right, questions)
}
