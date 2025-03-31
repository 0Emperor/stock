package tfm

import (
	"fmt"
	"stock/parser"
)

func Println(data *parser.BuildData) {
	fmt.Println("stock:")
	for thing, quantity := range data.Stock {
		fmt.Println(thing+" => ", quantity)
	}
}
