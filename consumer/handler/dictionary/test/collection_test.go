package test

import (
	dictionary_function "myInternal/consumer/handler/dictionary"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCollection(t *testing.T) {

	env.LoadEnv("./../../../../.env")
	dictionaryCollection, err := dictionary_function.CollectionDictionary()
	if err != nil {
		t.Fatalf("Error collection dictionary function: %v", err)
	}

	if(len(dictionaryCollection.Collection) < 3){
		t.Fatalf("Error collection dictionary is smaller then three, error: %v", err)
	}

}