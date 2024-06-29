package elasticbindutil

import (
	"encoding/json"
	"errors"
	"fmt"
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
	if status == 404 {
		var err NotFoundError
		if err := json.NewDecoder(resource).Decode(&err); err != nil {
			return result, err
		}
		fmt.Println("err404", err)
		return result, errors.New(err.Error.Reason)
	} else if status == 405 {
		var err MethodNotAllowedError
		if err := json.NewDecoder(resource).Decode(&err); err != nil {
			return result, err
		}
		fmt.Println("err405", err)

		return result, errors.New(err.Error)
	}

	if err := json.NewDecoder(resource).Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}
