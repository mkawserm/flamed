package graphql

import (
	"encoding/json"
)

type RequestOptions struct {
	Query         string                 `json:"query" url:"query" schema:"query"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
	OperationName string                 `json:"operationName" url:"operationName" schema:"operationName"`
}

// a workaround for getting`variables` as a JSON string
type RequestOptionsCompatibility struct {
	Query         string `json:"query" url:"query" schema:"query"`
	Variables     string `json:"variables" url:"variables" schema:"variables"`
	OperationName string `json:"operationName" url:"operationName" schema:"operationName"`
}

func ParseGraphQLQuery(data []byte) *RequestOptions {
	var opts RequestOptions
	err := json.Unmarshal(data, &opts)
	if err != nil {
		// Probably `variables` was sent as a string instead of an object.
		// So, we try to be polite and try to parse that as a JSON string
		var optsCompatible RequestOptionsCompatibility
		_ = json.Unmarshal(data, &optsCompatible)

		_ = json.Unmarshal([]byte(optsCompatible.Variables), &opts.Variables)
	}

	return &opts
}
