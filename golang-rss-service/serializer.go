package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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

func Response(c *gin.Context, code int, data []interface{}) {

	var response ResponseItem

	response.Items = data

	c.JSON(code, response)
}