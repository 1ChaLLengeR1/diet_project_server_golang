package test

import (
	"fmt"
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	statistics_functions "myInternal/consumer/handler/statistics"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCreateStatistics(t *testing.T) {

	params := params_data.Params{
		Header: common_test.UserTest,
		Param: "c13cc5a5-2f7a-4584-a78b-6f85b8b614bf",
	}

	env.LoadEnv("./.env")
	createStatistics, err := statistics_functions.CreateStatistics(params)
	if err != nil {
		t.Fatalf("Error create statistics function: %v", err)
	}

	fmt.Println(createStatistics)
}