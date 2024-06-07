package main

import (
	"github.com/Zettablock/zsource/dao/ethereum"
	"github.com/Zettablock/zsource/utils"
)

// HandleAVSMetadataURIUpdated : shouldRetry tell zrunner whether to retry on errors
// *utils.Deps contains *gorm.DB which can be used for CRUD
// *utils.Deps also contains Logger
func HandleAVSMetadataURIUpdated(log ethereum.Log, deps *utils.Deps) (bool, error) {
	//shouldRetry := false
	//var avs m.Avs
	//deps.DestinationDB.Table(deps.DestinationDBSchema+".avs").FirstOrCreate(&avs, m.Avs{ID: log.ArgumentValues[0]})
	//deps.Logger.Info("Got AVS:", avs)
	//avs.MetadataUri = log.ArgumentValues[1]
	//return shouldRetry, deps.DestinationDB.Table(deps.DestinationDBSchema + ".avs").Save(avs).Error
	return false, nil
}

func mapRegistrationStatus(status int) string {
	if status == 0 {
		return "UNREGISTERED"
	} else if status == 1 {
		return "REGISTERED"
	} else {
		return "UNREGISTERED"
	}
}

// HandleOperatorAVSRegistrationStatusUpdated : shouldRetry tell zrunner whether to retry on errors
// *utils.Deps contains *gorm.DB which can be used for CRUD
// *utils.Deps also contains Logger
func HandleOperatorAVSRegistrationStatusUpdated(log ethereum.Log, deps *utils.Deps) (bool, error) {
	return false, nil
}
