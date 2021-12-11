package main

import (
	"log"
	"os"
)

func main() {
	// 出力先を設定
	// log.SetOutput(os.Stdout)

	/*
		log.Print("Log\n")
		log.Println("Log2")
		log.Printf("Log%d\n", 3)
	*/

	// log.Fatal("Log\n")
	// log.Fatalln("Log2")
	// log.Fatalf("Log%d\n", 3)

	// log.Panic("Log\n")
	// log.Panicln("Log2")
	// log.Panicf(("Log%d\n"), 3)

	// ファイルに出力

	// f, err := os.Create("test.log")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.SetOutput(f)
	// log.Println("ファイルに書き込む")

	/*
		// ログのフォーマット
		log.SetOutput(os.Stdout)
		log.SetFlags(log.LstdFlags)
		log.Println("A")

		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("B")

		log.SetFlags(log.Ltime | log.Lshortfile)
		log.Print("C")

		log.SetFlags(log.LstdFlags)
		log.SetPrefix("[LOG]")
		log.Print("D")

		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
		log.Print("E")
	*/

	// ロガーの生成
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("A")
	log.Println("B")

	_, err := os.Open("aaa")
	if err != nil {
		logger.Fatal(err)
		log.Fatal(err)
	}

}
