package main

import (
	"fmt"
	"nilchan/feature1"
	"nilchan/feature2"
	simpleconnection "nilchan/feature_postgres/simple_connection"
)

func main() {
	fmt.Println("Hello Git!")
	feature1.Feature1()
	feature2.Feature2()
	simpleconnection.CheckConnection()
}
