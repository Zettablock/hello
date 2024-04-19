package main

import (
	"github.com/Zettablock/zsource/utils"
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
