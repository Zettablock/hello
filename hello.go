package main

import (
	m "example/user/hello/dao"
	"github.com/Zettablock/zsource/dao/ethereum"
	"github.com/Zettablock/zsource/utils"
	"strconv"
)

func HandleInitialization(deps *utils.Deps) error {
	deps.DestinationDB.Exec(`CREATE TABLE IF NOT EXISTS public.avs(
    id                     VARCHAR(100),
    avs_list               VARCHAR(100),
    restakeable_strategies VARCHAR(100) array,
    metadata_uri text,
    PRIMARY KEY (id));`)
	return nil
}

func HandleAVSMetadataURIUpdated(log ethereum.Log, deps *utils.Deps) (bool, error) {
	shouldRetry := false
	var avs m.Avs
	deps.DestinationDB.FirstOrCreate(&avs, m.Avs{ID: log.ArgumentValues[0]})
	deps.Logger.Info("Got AVS:", avs)
	avs.MetadataUri = log.ArgumentValues[1]
	return shouldRetry, deps.DestinationDB.Save(avs).Error

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
	deps.DestinationDB.Where("ID = ?", avsOperatorID).First(avsOperator)
	if avsOperator == nil {
		var avs m.Avs
		deps.DestinationDB.FirstOrCreate(&avs, m.Avs{ID: log.ArgumentValues[1]})
		n, _ := strconv.Atoi(log.ArgumentValues[2])
		return shouldRetry, deps.DestinationDB.Save(&m.AvsOperator{
			ID:                 log.ArgumentValues[1] + "-" + log.ArgumentValues[0],
			Avs:                avs.ID,
			Operator:           log.ArgumentValues[0],
			RegistrationStatus: mapRegistrationStatus(n),
		}).Error
	}
	n, _ := strconv.Atoi(log.ArgumentValues[2])
	avsOperator.RegistrationStatus = mapRegistrationStatus(n)
	return shouldRetry, deps.DestinationDB.Save(avsOperator).Error
}
