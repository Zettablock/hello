package main

import (
	m "example/user/hello/dao"
	"github.com/Zettablock/zsource/dao/ethereum"
	"github.com/Zettablock/zsource/utils"
	"strconv"
)

func HandleAVSMetadataURIUpdated(log ethereum.Log, deps *utils.Deps) (bool, error) {
	shouldRetry := false
	var avs m.Avs
	deps.DestinationDB.Table(deps.DestinationDBSchema+".avs").FirstOrCreate(&avs, m.Avs{ID: log.ArgumentValues[0]})
	deps.Logger.Info("Got AVS:", avs)
	avs.MetadataUri = log.ArgumentValues[1]
	return shouldRetry, deps.DestinationDB.Table(deps.DestinationDBSchema + ".avs").Save(avs).Error

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

func HandleOperatorAVSRegistrationStatusUpdated(log ethereum.Log, deps *utils.Deps) (bool, error) {
	shouldRetry := false
	avsOperatorID := log.ArgumentValues[1] + "-" + log.ArgumentValues[0]
	var avsOperator *m.AvsOperator
	deps.DestinationDB.Table(deps.DestinationDBSchema+".avs_operator").Where("ID = ?", avsOperatorID).First(avsOperator)
	if avsOperator == nil {
		var avs m.Avs
		deps.DestinationDB.Table(deps.DestinationDBSchema+".avs").FirstOrCreate(&avs, m.Avs{ID: log.ArgumentValues[1]})
		n, _ := strconv.Atoi(log.ArgumentValues[2])
		return shouldRetry, deps.DestinationDB.Table(deps.DestinationDBSchema + ".avs_operator").Save(&m.AvsOperator{
			ID:                 log.ArgumentValues[1] + "-" + log.ArgumentValues[0],
			Avs:                avs.ID,
			Operator:           log.ArgumentValues[0],
			RegistrationStatus: mapRegistrationStatus(n),
		}).Error
	}
	n, _ := strconv.Atoi(log.ArgumentValues[2])
	avsOperator.RegistrationStatus = mapRegistrationStatus(n)
	return shouldRetry, deps.DestinationDB.Table(deps.DestinationDBSchema + ".avs_operator").Save(avsOperator).Error
}