package route

import (
	"io/ioutil"
	"log"
	"os"
)

const envPath = "etc/config/"

func Start() *Ruter {
	r := &Ruter{}

	if apiPrefix, ok := getEnv("API_PREFIX"); ok {
		r.apiPrefix = apiPrefix
	}

	log.Println("Wystartowano ruter z konfiguracją: ", r.apiPrefix)

	return r
}

func getEnv(env string) (wartosc string, ok bool) {
	wartosc, ok = getEnvOs(env)
	if !ok {
		wartosc, ok = getEnvFile(env)
	}
	return
}

func getEnvOs(env string) (string, bool) {
	wartosc := os.Getenv(env)
	return wartosc, wartosc != ""
}

func getEnvFile(env string) (string, bool) {
	wartosc, err := ioutil.ReadFile(envPath + env)
	if err != nil {
		return "", false
	}

	return string(wartosc), true
}
