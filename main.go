package main

import "github.com/kapi1023/licencjat/api/serwis"

func main() {
	port := musiGetEnv("PORT")
	serwis.Start(port)
}
