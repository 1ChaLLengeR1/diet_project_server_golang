package consumer_database

import (
	env "internal/consumer/initializers"
	"testing"
)

func TestConnectToDataBase(t *testing.T) {

	env.LoadEnv("./../../.env")
	err := ConnectToDataBase()
	if err != nil {
		t.Errorf("not connect do database: %s", err.Error())
	}
}