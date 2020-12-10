package lexoffice

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/joho/godotenv"
)

const (
	prettyPrintPrefix = ""
	prettyPrintIndent = "  "
)

// helper to get a nicely formatted json from an arbitrary struct
func prettyPrintJson(object interface{}) string {
	b, _ := json.MarshalIndent(object, prettyPrintPrefix, prettyPrintIndent)
	return string(b)
}

// helper to get a dump from an outgoing http request
func prettyPrintRequest(req *http.Request) string {
	b, _ := httputil.DumpRequestOut(req, true)
	return string(b)
}

// helper to get a dump from a http response
func prettyPrintResponse(res *http.Response) string {
	b, _ := httputil.DumpResponse(res, true)
	return string(b)
}

// helper to load the LEXOFFICE_API_KEY from environment and/or .env file
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
