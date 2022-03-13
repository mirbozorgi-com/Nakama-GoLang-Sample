package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
	"strconv"
	"time"
)

var timeResp struct {
	Payload string `json:"payload"`
}
func server_time(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error){

	utc := time.Now().UTC().Unix()
    timeResp.Payload=strconv.Itoa(int(utc))
	out, err := json.Marshal(timeResp)
	if err != nil {
		logger.Error("Marshal error: %v", err)
		return "", errMarshal
	}

	logger.Debug("servertime resp: %v", string(out))
	return string(out), nil

}
