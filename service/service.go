package service

import (
	"go-space-station/data"
	"go-space-station/model"
)

func GetSpaceStationVisibility() bool {
	return true
}

func GetSpaceStationPosition() model.IssPosition {
	positionData := data.GetPositionData()
	return positionData
}
