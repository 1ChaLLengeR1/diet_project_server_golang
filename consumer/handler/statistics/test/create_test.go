package test

import (
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
	createStatisticOption, err := statistics_functions.CreateStatisticOption(params)
	if err != nil {
		t.Fatalf("Error create statistic options function: %v", err)
	}

	if len(createStatisticOption.Statistics) == 0{
		t.Fatalf("Error len function statistic options is 0")
	}

	createStatistics := statistics_functions.CollectionStatistics(createStatisticOption.Statistics)
	if len(createStatistics) == 0{
		t.Fatalf("Error len function statistic is 0")
	}
	
}