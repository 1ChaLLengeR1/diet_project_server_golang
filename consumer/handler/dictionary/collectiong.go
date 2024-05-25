package dictionary

import (
	dictionary_data "myInternal/consumer/data/dictionary"
	database "myInternal/consumer/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCollectionDictionary struct {
	Collection []dictionary_data.Collection `json:"collection"`
	Status     int                          `json:"status"`
	Error      string                       `json:"error"`
}

func HandlerCollectionDictionary(c *gin.Context){
	
	collection, err := CollectionDictionary()
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseCollectionDictionary{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseCollectionDictionary{
		Collection: collection.Collection,
		Status: collection.Status,
		Error: collection.Error,
	})
}


func CollectionDictionary()(ResponseCollectionDictionary, error){

	var dictionaryCollection []dictionary_data.Collection
	db, err := database.ConnectToDataBase()
    if err != nil {
        return ResponseCollectionDictionary{}, err
    }

	query := `SELECT * FROM dictionary;`
	rows, err := db.Query(query)
    if err != nil {
        return ResponseCollectionDictionary{}, err
    }
	defer rows.Close()

	for rows.Next() {
		var dictionary dictionary_data.Collection
		if err := rows.Scan(&dictionary.Id, &dictionary.Key, &dictionary.Translation, &dictionary.CreatedUp, &dictionary.UpdateUp); err != nil {
			return ResponseCollectionDictionary{}, err
		}
		dictionaryCollection = append(dictionaryCollection, dictionary)
	}

	return ResponseCollectionDictionary{
		Collection: dictionaryCollection,
		Status: http.StatusOK,
		Error: "",
	}, nil
}