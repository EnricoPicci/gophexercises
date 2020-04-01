package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	right := 0
	f, err := os.Open("./src/quiz/quizpart1/problems.csv")
	check(err)
	defer f.Close()

	r := csv.NewReader(bufio.NewReader(f))
	s := bufio.NewScanner(os.Stdin)

	questions := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)

		fmt.Println(record[0])
		questions++
		s.Scan()
		if s.Text() == record[1] {
			right++
		}
	}

	fmt.Printf("%v right answers out of %v questions.", right, questions)
}
