package youtube

import (
	"testing"
)

func TestGetVideoDuration(t *testing.T) {
	str, err := getVideoDurationString("VnEI8d2rG-A")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(str)
}
