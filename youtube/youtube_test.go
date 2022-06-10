package youtube

import (
	"testing"
	"time"
)

func TestGetVideoDuration(t *testing.T) {
	duration, err := GetVideoDuration("https://www.youtube.com/watch?v=UkgK8eUdpAo")
	expect := 3*time.Minute + 40*time.Second
	if err != nil {
		t.Error(err)
	}
	if duration != expect {
		t.Error("got", duration, "expected", expect)
	}
	duration, err = GetVideoDuration("my youtube link https://www.youtube.com/watch?v=khn0rV_Svlc https://www.youtube.com/watch?v=UkgK8eUdpAo")
	expect = 30*time.Minute + 2*time.Second
	if err != nil {
		t.Error(err)
	}
	if duration != expect {
		t.Error("got", duration, "expected", expect)
	}
}
