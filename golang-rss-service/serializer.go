package main

import (
	"encoding/json"
)

func Deserialize(data []byte, structure interface{}) error {

	if len(data) > 0 {

		err := json.Unmarshal(data, structure)

		if err != nil {

			return err
		}
	}

	return nil
}

func Serialize(structure interface{}) ([]byte, error) {

	if structure != nil {

		data, err := json.Marshal(structure)

		if err != nil {

			return nil, err
		}

		return data, nil
	}

	return nil, nil
}
