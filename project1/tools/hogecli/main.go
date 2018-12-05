package main

import (
	"flag"
	"fmt"

	"github.com/kyokomi/lottery"
)

func lot() string {
	if lottery.NewDefault().LotOf(1, 40) {
		return "あたり"
	}

	return "はずれ"
}

func main() {
	count := flag.Int("count", 1, "")
	flag.Parse()

	for i := 0; i < *count; i++ {
		fmt.Println(lot())
	}
}
