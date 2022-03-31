package data

import (
	"encoding/json"
	"fmt"
	"go-space-station/model"
	"io/ioutil"
	"log"
	"net/http"
)

type IHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type HttpClient struct {
	// http.Client
}

const issUrl = "http://api.open-notify.org/iss-now.json"

func GetLat() float64 {
	return 1.123
}

func GetLong() float64 {
	return 14.123
}

func parsaePositionData() (float64, float64) {
	spaceStationData := model.SpaceStationData{}
	data := `{"iss_position": {"longitude": "31.0058", "latitude": "-2.5085"}, "message": "success", "timestamp": 1648548448}`
	textBytes := []byte(data)
	err := json.Unmarshal(textBytes, &spaceStationData)

	if err != nil {
		fmt.Println(err)
		return 0.0, 0.0
	}

	return spaceStationData.PositionData.Longitude, spaceStationData.PositionData.Latitude
}

func GetPositionDataFromUrl(client IHttpClient, dataObject model.SpaceStationData, url string) model.IssPosition {

	req, err := http.NewRequest(http.MethodGet, issUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	// req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// spaceStationData := model.SpaceStationData{}
	jsonErr := json.Unmarshal(body, &dataObject)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return dataObject.PositionData
}
