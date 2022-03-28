package main

import (
	"flag"
	"fmt"

	"github.com/kctboy/poortsniffer/poort"
)

var ipAdres string

func init() {
	flag.StringVar(&ipAdres, "IP", "8.8.8.8", "The IPadress")
	flag.Parse()
}
func main() {

	//open := poort.Scanport("tcp","localhost", 1313 )
	//fmt.Printf("portm open: %t\n", open)

	results := poort.InitialScan(ipAdres)
	fmt.Println(results)
}
