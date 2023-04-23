package main

import (
	"log"
	"net/http"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func handleRequests(url string, user string, password string) {
	handleWoLRequest := func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		mac := query.Get("mac")
		targetIf := query.Get("if")
		succeeded := sendWoL(url, user, password, mac, targetIf)
		if succeeded {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}

	http.HandleFunc("/wol", handleWoLRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	var configDir = os.Getenv("CONFIG_DIR")
	var configFile = path.Join(configDir, ".env")
	var myEnv map[string]string
	myEnv, err := godotenv.Read(configFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := myEnv["PFSENSE_URL"]
	user := myEnv["PFSENSE_USER"]
	password := myEnv["PFSENSE_PASSWORD"]

	handleRequests(url, user, password)
}
