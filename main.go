package main

import (
	"fmt"
	"os"

	optimizer "stock/optimize"
	"stock/parser"
	tfm "stock/stock"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("usage:\n\n./stock_exchange <file> <waiting_time>")
		return
	}
	waiting_time, err := strconv.Atoi(args[2])
	if err != nil || waiting_time == 0 {
		fmt.Println("invalid waiting time")
		return
	}
	file := os.Args[1]
	data, err := parser.ParseFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		time.Sleep(time.Duration(waiting_time * 1000000000))
		tfm.Println(data)
		os.Exit(0)
	}()
	optimizer.Optimze(data)
}
