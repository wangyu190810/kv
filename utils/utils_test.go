package utils

import (
	"testing"
	"time"
)

func TestName(t *testing.T) {
	dateTime := time.Unix(int64(1559788818), 0)
	if dateTime.Add(60 * time.Minute).After(time.Now()) {
		t.Fatalf("timeout")
	}else{
		t.Log("now")
	}
}
