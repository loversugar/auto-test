package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Data struct {
	OperationType string
	ElementId string
	Data string
}

func ParseHttpRequest(request *http.Request) *Data {
	var data *Data
	if request.Method == "POST" {
		if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
			fmt.Println(err)

		}
	}

	return data
}

func ResponseWithOrigin(w http.ResponseWriter, r *http.Request, json []byte) {
		// Set CORS headers
		header := w.Header()
		header.Set("Content-Type", "application/json")
		header.Set("Accept", "application/json")
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		header.Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		header.Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		header.Set("Access-Control-Allow-Credentials", "true")

	// Adjust status code to 204
	if strings.EqualFold(r.Method, "options") {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Write(json)
	}
}
