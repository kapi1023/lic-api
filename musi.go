package main

import (
	"log"
	"os"
)

func musiGetEnv(env string, musi ...bool) string {
	wartosc := os.Getenv(env)
	if wartosc == "" {
		if len(musi) != 0 && !musi[0] {
			return ""
		}
		log.Fatalf("błąd: nie ustawiono zmiennej środowiskowej \"%s\"\n", env)
	}
	return wartosc
}
