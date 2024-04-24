package main

import (
	"fmt"
	"github.com/Zettablock/zsource/utils"
)

// HandleInitialization :
// Make sure initializationHandlers called in only one yml file to avoid unnecessary race condition.
// initializationHandlers are supposed to initialize pg, which can be done manually too.
// We prefer to create tables/schemas manually for now
func HandleInitialization(deps *utils.Deps) error {
	deps.DestinationDB.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %v.avs(
    id                     VARCHAR(100),
    avs_list               VARCHAR(100),
    restakeable_strategies VARCHAR(100) array,
    metadata_uri text,
    PRIMARY KEY (id));`, deps.DestinationDBSchema))
	return nil
}
