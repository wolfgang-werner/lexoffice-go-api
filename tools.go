package lexoffice

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
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
