package main

import "fmt"
import "rsc.io/quote" //import module from pkg.go.dev to get quotes

func main(){
	fmt.Println(quote.Go())
}
