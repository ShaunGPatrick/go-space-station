package data

import (
	"fmt"
	"go-space-station/model"
	"io"
	"net/http"
	"strings"
	"testing"
)

type StubHttpClient HttpClient

func (c StubHttpClient) Do(req *http.Request) (*http.Response, error) {
	fmt.Println("We are in the stub!")
	data := `{"iss_position": {"longitude": "31.0058", "latitude": "-2.5085"}, "message": "success", "timestamp": 1648548448}`
	stringReader := strings.NewReader(data)
	stringReadCloser := io.NopCloser(stringReader)
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       stringReadCloser,
	}, nil
}

func TestGetSpaceStationPosition(t *testing.T) {

	tests := []struct {
		name string
		got  float64
		want float64
	}{
		{
			name: "Test lat",
			got:  GetLat(),
			want: 1.123,
		},
		{
			name: "Test long",
			got:  GetLong(),
			want: 14.123,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if test.got != test.want {
				t.Errorf("got %f, want %f", test.got, test.want)
			}
		})
	}

}

func TestPasePositionData(t *testing.T) {

	t.Run("Test successful parse", func(t *testing.T) {
		stubClient := StubHttpClient{}
		got := GetPositionDataFromUrl(stubClient, model.SpaceStationData{}, "blah")
		want := model.IssPosition{Longitude: 31.0058, Latitude: -2.5085}

		if got != want {
			t.Errorf("got %f but wanted %f", got, want)
		}
	})

	t.Run("Test something bad happening", func(t *testing.T) {

	})
}
