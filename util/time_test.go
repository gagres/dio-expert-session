package util_test

import (
	"testing"

	"github.com/gagres/dio-expert-session-finance/util"
)

func TestStringToTime(t *testing.T) {
	var convertedTime = util.StringToTime("2019-02-12T10:10:10")

	if convertedTime.Year() != 2019 {
		t.Errorf("expects year to be 2019")
	}
}
