package server

import (
	"fmt"
	"go-space-station/service"
	"net/http"
)

const (
	isVisibleEndpoit     = "/is-visible"
	positionDataEndpoint = "/position"
)

type SpaceStationServer struct {
	http.Handler
}

func NewSpaceStationServer() *SpaceStationServer {
	server := new(SpaceStationServer)

	router := http.NewServeMux()
	router.Handle(isVisibleEndpoit, http.HandlerFunc(server.spaceStationVisibilityHandler))
	router.Handle(positionDataEndpoint, http.HandlerFunc(server.positionDataHandler))

	server.Handler = router

	return server
}

func (server *SpaceStationServer) spaceStationVisibilityHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, service.GetSpaceStationVisibility())
}

func (server *SpaceStationServer) positionDataHandler(w http.ResponseWriter, r *http.Request) {
	lat := service.GetSpaceStationPosition().Latitude
	lon := service.GetSpaceStationPosition().Longitude
	fmt.Fprintf(w, "{%f, %f}", lat, lon)
	// fmt.Fprint(w, service.GetSpaceStationPosition())
}
