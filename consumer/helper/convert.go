package helper

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func BindJSONToMap(c *gin.Context, obj interface{}) (map[string]interface{}, error) {

	if err := c.BindJSON(obj); err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &jsonMap); err != nil {
		return nil, err
	}

	return jsonMap, nil
}