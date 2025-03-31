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
	timeout := time.After(time.Duration(1000000000*waiting_time))
	file := os.Args[1]
	data, err := parser.ParseFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	done :=make(chan int, 1)
	go optimizer.Optimize(data,done)
	select{
	case cycle:=<-done:
		fmt.Println("no more work ",cycle)
		tfm.Println(data)
	case <-timeout:
		fmt.Println("more work left after timeout")
	}
}
