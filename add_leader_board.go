package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
)

func add_leader_board(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	var input struct {
		Id       string                 `json:"id"`
		Metadata map[string]interface{} `json:"metadata"`
		Sort     string                 `json:"sort"`
		Operator string                 `json:"operator"`
		Reset    string                 `json:"reset"`
	}

	if err := json.Unmarshal([]byte(payload), &input); err != nil {
		return "", err
	}

	id := input.Id
	authoritative := false
	sort := input.Sort
	operator := input.Operator
	metadata := input.Metadata
	reset := input.Reset
	if err := nk.LeaderboardCreate(ctx, id, authoritative, sort, operator, reset, metadata); err != nil {
		return "", err
	}

	return payload, nil
}
