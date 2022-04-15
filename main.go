package main

import (
	"example/user/hello/mascot"
	"fmt"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
}

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Best" {
		t.Fatal("wrong mascot")
	}
}

