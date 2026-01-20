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

// LeaderboardsConfig is the data definition for the LeaderboardsSystem type.
type LeaderboardsConfig struct {
	Leaderboards []*LeaderboardsConfigLeaderboard `json:"leaderboards,omitempty"`
}

type LeaderboardsConfigLeaderboard struct {
	Id            string            `json:"id,omitempty"`
	SortOrder     string            `json:"sort_order,omitempty"`
	Operator      string            `json:"operator,omitempty"`
	ResetSchedule string            `json:"reset_schedule,omitempty"`
	Authoritative bool              `json:"authoritative,omitempty"`
	Regions       []string          `json:"regions,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty"`
}

// OnLeaderboardUpdateScore is a function called before or after a leaderboard score is updated.
type OnLeaderboardUpdateScore func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, userID, leaderboardID, ownerID string, score, subscore int64, metadata map[string]any) (int64, int64, map[string]any, error)

// OnLeaderboardDeleteScore is a function called before or after a leaderboard score is deleted.
type OnLeaderboardDeleteScore func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, userID, leaderboardID, ownerID string) error

// OnLeaderboardReset is a function called when a leaderboard resets.
type OnLeaderboardReset func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, leaderboardID string, reset int64, config *LeaderboardsConfigLeaderboard) error

// The LeaderboardsSystem defines a collection of leaderboards which can be defined as global or regional with Nakama
// server.
type LeaderboardsSystem interface {
	System

	// Get returns a list of available leaderboards for the user.
	Get(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string) (*LeaderboardConfigList, error)

	// GetLeaderboard returns a specified leaderboard with scores.
	GetLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, leaderboardID, ownerID string, limit int, cursor string) (*Leaderboard, error)

	// Create creates a new leaderboard.
	Create(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, id, sortOrder, operator, resetSchedule string, authoritative bool, regions []string) error

	// Delete deletes an existing leaderboard.
	Delete(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, id string) error

	// UpdateScore updates a leaderboard score.
	UpdateScore(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, userID, leaderboardID, ownerID, username string, score, subscore int64, metadata map[string]any, operator Operator, conditionalMetadataUpdate bool) (*Leaderboard, error)

	// DeleteScore deletes a leaderboard score.
	DeleteScore(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, userID, leaderboardID, ownerID string) error

	// SetOnBeforeUpdateScore sets a custom function which will run before a leaderboard score is updated.
	SetOnBeforeUpdateScore(fn OnLeaderboardUpdateScore)

	// SetOnAfterUpdateScore sets a custom function which will run after a leaderboard score is updated.
	SetOnAfterUpdateScore(fn OnLeaderboardUpdateScore)

	// SetOnBeforeDeleteScore sets a custom function which will run before a leaderboard score is deleted.
	SetOnBeforeDeleteScore(fn OnLeaderboardDeleteScore)

	// SetOnAfterDeleteScore sets a custom function which will run after a leaderboard score is deleted.
	SetOnAfterDeleteScore(fn OnLeaderboardDeleteScore)

	// SetOnLeaderboardReset sets a custom function which will run when a leaderboard resets.
	SetOnLeaderboardReset(fn OnLeaderboardReset)
}

// ValidateWriteScoreFn is a function used to validate the leaderboard score input.
type ValidateWriteScoreFn func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, input *LeaderboardUpdate) *runtime.Error
