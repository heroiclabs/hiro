// Copyright 2023 Heroic Labs & Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hiro

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

var (
	ErrLeaderboardNotFound = runtime.NewError("leaderboard not found", 3) // INVALID_ARGUMENT
)

// LeaderboardsConfig is the data definition for the LeaderboardsSystem type.
type LeaderboardsConfig struct {
	Leaderboards []*LeaderboardsConfigLeaderboard `json:"leaderboards,omitempty"`
}

type LeaderboardsConfigLeaderboard struct {
	Id            string            `json:"id,omitempty"`
	Name          string            `json:"name,omitempty"`
	Description   string            `json:"description,omitempty"`
	Category      string            `json:"category,omitempty"`
	SortOrder     string            `json:"sort_order,omitempty"`
	Operator      string            `json:"operator,omitempty"`
	StartTimeSec  int64             `json:"start_time_sec,omitempty"`
	EndTimeSec    int64             `json:"end_time_sec,omitempty"`
	Duration      int64             `json:"duration,omitempty"`
	ResetSchedule string            `json:"reset_schedule,omitempty"`
	Authoritative bool              `json:"authoritative,omitempty"`
	Regions       []string          `json:"regions,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty"`
}

// OnLeaderboardUpdate is a function called before or after a leaderboard score is updated.
type OnLeaderboardUpdate func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, userID, leaderboardID string, score, subscore int64, metadata map[string]interface{}) (int64, int64, map[string]interface{}, error)

// The LeaderboardsSystem defines a collection of leaderboards which can be defined as global or regional with Nakama server.
type LeaderboardsSystem interface {
	System

	// Deprecated: Use List instead. Get returns a list of available leaderboards for the user.
	Get(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string) (*LeaderboardConfigList, error)

	// ListLeaderboard returns a list of available leaderboards for the user.
	ListLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, categories []string) (*LeaderboardList, error)

	// GetLeaderboard returns a specified leaderboard with scores.
	GetLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, leaderboardID string) (*Leaderboard, error)

	// ListLeaderboardScores returns a list of scores for a specified leaderboard.
	ListLeaderboardScores(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, leaderboardID string, ownerIds []string, limit int, cursor string, expiry int64) (*LeaderboardScoreList, error)

	// ListLeaderboardScoresAroundOwner returns a list of scores for a specified leaderboard around the owner's score.
	ListLeaderboardScoresAroundOwner(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, leaderboardID string, ownerId string, limit int, cursor string, expiry int64) (*LeaderboardScoreList, error)

	// UpdateLeaderboard updates the user's score in the specified leaderboard.
	UpdateLeaderboard(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, userID, username, leaderboardID string, score, subscore int64, metadata map[string]interface{}, conditionalMetadataUpdate bool) (*LeaderboardScore, error)

	// SetOnBeforeUpdateScore sets a custom function which will run before a leaderboard score is updated.
	SetOnBeforeUpdateScore(fn OnLeaderboardUpdate)

	// SetOnAfterUpdateScore sets a custom function which will run after a leaderboard score is updated.
	SetOnAfterUpdateScore(fn OnLeaderboardUpdate)
}

// ValidateWriteScoreFn is a function used to validate the leaderboard score input.
type ValidateWriteScoreFn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, input *LeaderboardUpdate) *runtime.Error
