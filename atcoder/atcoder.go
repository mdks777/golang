package main

import(
	"fmt"
) 

func main() {
	var a,b,c int
	var d string
	fmt.Scan(&a)
	fmt.Scan(&a, &c)
	fmt.Scan(d)
	fmt.Printf("%d %s",a+b+c, d)
}