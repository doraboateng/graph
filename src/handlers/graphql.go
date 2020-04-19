package handlers

import (
	"encoding/json"
	"net/http"
)

func GraphHandler(writer http.ResponseWriter, request *http.Request) {
	var params struct {
		Query     string                 `json:"query"`
		Operation string                 `json:"operationName"`
		Variables map[string]interface{} `json:"variables"`
	}

	if err := json.NewDecoder(request.Body).Decode(&params); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	response := Schema.Exec(request.Context(), params.Query, params.Operation, params.Variables)
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Write(jsonResponse)
}
