package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "csv file with questions and answers")
	duration := flag.Int("time", 30, "quiz time limit")
	randomize := flag.Bool("random", false, "randomize questions")

	//Parse
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	questions := loadRecords(*csvFile)
}

func loadRecords(csvFile string) [][]string {

}