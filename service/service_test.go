package service

import "testing"

func TestSerivceResponse(t *testing.T) {
	got := GetSpaceStationVisibility()

	if !got {
		t.Errorf("Response was false")
	}
}
