package lexoffice

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

const (
	prettyPrintPrefix = ""
	prettyPrintIndent = "  "
)

func prettyPrintJson(object interface{}) string {
	b, _ := json.MarshalIndent(object, prettyPrintPrefix, prettyPrintIndent)
	return string(b)
}

func prettyPrintRequest(req *http.Request) string {
	b, _ := httputil.DumpRequestOut(req, true)
	return string(b)
}

func prettyPrintResponse(res *http.Response) string {
	b, _ := httputil.DumpResponse(res, true)
	return string(b)
}

func getApiKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	apiKey, defined := os.LookupEnv("LEXOFFICE_API_KEY")
	if !defined || apiKey == "" {
		log.Fatalf("define non-empty environment variable LEXOFFICE_API_KEY (may be located in .env file")
		return ""
	}

	return apiKey
}
