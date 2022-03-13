package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
	"strconv"
	"time"
)

var dateResp struct {
	Payload string `json:"payload"`
}

func server_date(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	t := time.Now().UTC()
	day := strconv.Itoa(t.Day())
	hour := strconv.Itoa(t.Hour())
	isdst := strconv.FormatBool(t.IsDST())
	min := strconv.Itoa(t.Minute())
	month := strconv.Itoa(int(t.Month()))
	sec := strconv.Itoa(t.Second())
	wday := strconv.Itoa(int(t.Weekday()))
	yday := strconv.Itoa(t.YearDay())
	year := strconv.Itoa(t.Year())

	date := "{\"day\":" + day + ",\"hour\":" + hour
	date += ",\"isdst\":" + isdst + ",\"min\":" + min
	date += ",\"month\":" + month + ",\"sec\":" + sec
	date += ",\"wday\":" + wday + ",\"yday\":" + yday + ",\"year\":" + year + "}"
	dateResp.Payload = date

	out, err := json.Marshal(dateResp)
	if err != nil {
		logger.Error("Marshal error: %v", err)
		return "", errMarshal
	}

	logger.Debug("serverdate resp: %v", string(out))
	return string(out), nil
}
