package elasticbindutil

import (
	"encoding/json"
	"errors"
	"io"
)

type NotFoundError struct {
	Error struct {
		RootCause []struct {
			Type   string `json:"type"`
			Reason string `json:"reason"`
		} `json:"root_cause"`
		Type   string `json:"type"`
		Reason string `json:"reason"`
	} `json:"error"`
	Status int `json:"status"`
}

type MethodNotAllowedError struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func HandleAndDecodeResponse[dest any](status int, resource io.ReadCloser) (dest, error) {
	var result dest

	// var debugAny map[string]interface{}
	// json.NewDecoder(resource).Decode(&debugAny)

	// log.Println(debugAny)

	if status == 404 {
		var err NotFoundError
		if err := json.NewDecoder(resource).Decode(&err); err != nil {
			return result, err
		}
		return result, errors.New(err.Error.Reason)
	} else if status == 405 {
		var err MethodNotAllowedError
		if err := json.NewDecoder(resource).Decode(&err); err != nil {
			return result, err
		}

		return result, errors.New(err.Error)
	}

	// else if status > 299 && status < 200 {

	// 	if err := json.NewDecoder(resource).Decode(&result); err != nil {
	// 		return result, err
	// 	}
	// }

	if err := json.NewDecoder(resource).Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}
