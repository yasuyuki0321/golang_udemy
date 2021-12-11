package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// bufio

func main() {
	// 標準入力を行単位に読み込む
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(os.Stderr, "読み込みエラー", err)
	}
}
