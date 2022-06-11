package youtube

import (
	"testing"
)

func TestGetIdFromText(t *testing.T) {
	tests := [...][2]string{{"my youtube link https://www.youtube.com/watch?v=khn0rV_Svlc https://www.youtube.com/watch?v=UkgK8eUdpAo", "khn0rV_Svlc"},
		{"youtu.be/UkgK8eUdpAo", "UkgK8eUdpAo"},
	}
	for _, test := range tests {
		id, err := GetIdFromText(test[0])
		expect := test[1]
		if err != nil {
			t.Error(err)
		} else if expect != id {
			t.Error("expect", expect, "got", id)
		}
	}
	// no url
	_, err := GetIdFromText("here is text without url")
	if err != nil {
		_, ok := err.(*IdNotFound)
		if !ok {
			t.Error(err)
		}
	} else {
		t.Error("must be error")
	}
}
