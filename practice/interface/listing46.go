package main

import "fmt"

type Duration int

func (d *Duration) pretty () string{
	return fmt.Sprintf("GoodDuration: %d", *d)
}

func main() {
	var aa *Duration
	*aa = 100
	fmt.Printf("---%v", aa)

	dd := Duration(aa).pretty()
	fmt.Printf("it is %v", dd)
	//Duration(32).pretty()
}
