package main

import "fmt"

const (
	MAIN_PORT = "8001"
)

func main() {

	const a = 10
	fmt.Println(a)

	const (
		ipAdress = "1.1.100.1"
		port     = "8080"
		dbName   = "postgres"
	)
	fmt.Println(ipAdress)

	fmt.Println(checkPortIsValid())
}

func checkPortIsValid() bool {
	if MAIN_PORT == "8002" {
		return true
	}
	return false
}
