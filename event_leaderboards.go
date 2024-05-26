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
	"github.com/heroiclabs/nakama-common/runtime"
)

// EventLeaderboardsConfig is the data definition for the EventLeaderboardsSystem type.
type EventLeaderboardsConfig struct {
	EventLeaderboards map[string]*EventLeaderboardsConfigLeaderboard `json:"event_leaderboards,omitempty"`
}

type EventLeaderboardsConfigLeaderboard struct {
	Name                 string                                                     `json:"name,omitempty"`
	Description          string                                                     `json:"description,omitempty"`
	Category             string                                                     `json:"category,omitempty"`
	Ascending            bool                                                       `json:"ascending,omitempty"`
	Operator             string                                                     `json:"operator,omitempty"`
	ResetSchedule        string                                                     `json:"reset_schedule,omitempty"`
	CohortSize           int                                                        `json:"cohort_size,omitempty"`
	AdditionalProperties map[string]string                                          `json:"additional_properties,omitempty"`
	MaxNumScore          int                                                        `json:"max_num_score,omitempty"`
	RewardTiers          map[string][]*EventLeaderboardsConfigLeaderboardRewardTier `json:"reward_tiers,omitempty"`
	ChangeZones          map[string]*EventLeaderboardsConfigChangeZone              `json:"change_zones,omitempty"`
	Tiers                int                                                        `json:"tiers,omitempty"`
	MaxIdleTierDrop      int                                                        `json:"max_idle_tier_drop,omitempty"`
	StartTimeSec         int64                                                      `json:"start_time_sec,omitempty"`
	EndTimeSec           int64                                                      `json:"end_time_sec,omitempty"`
	Duration             int64                                                      `json:"duration,omitempty"`
}

type EventLeaderboardsConfigLeaderboardRewardTier struct {
	Name       string               `json:"name,omitempty"`
	RankMax    int                  `json:"rank_max,omitempty"`
	RankMin    int                  `json:"rank_min,omitempty"`
	Reward     *EconomyConfigReward `json:"reward,omitempty"`
	TierChange int                  `json:"tier_change,omitempty"`
}

type EventLeaderboardsConfigChangeZone struct {
	Promotion  float64 `json:"promotion,omitempty"`
	Demotion   float64 `json:"demotion,omitempty"`
	DemoteIdle bool    `json:"demote_idle,omitempty"`
}

// An EventLeaderboardsSystem is a gameplay system which represents cohort-segmented, tier-based event leaderboards.
type EventLeaderboardsSystem interface {
	System

	// GetEventLeaderboard returns a specified event leaderboard's cohort for the user.
	GetEventLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, eventLeaderboardID string) (*EventLeaderboard, error)

	// RollEventLeaderboard places the user into a new cohort for the specified event leaderboard if possible.
	RollEventLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, eventLeaderboardID string, tier *int, matchmakerProperties map[string]interface{}) (*EventLeaderboard, error)

	// UpdateEventLeaderboard updates the user's score in the specified event leaderboard, and returns the user's updated cohort information.
	UpdateEventLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, username, eventLeaderboardID string, score, subscore int64) (*EventLeaderboard, error)

	// ClaimEventLeaderboard claims the user's reward for the given event leaderboard.
	ClaimEventLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, eventLeaderboardID string) (*EventLeaderboard, error)

	// SetOnEventLeaderboardsReward sets a custom reward function which will run after an event leaderboard's reward is rolled.
	SetOnEventLeaderboardsReward(fn OnReward[*EventLeaderboardsConfigLeaderboard])

	// SetOnEventLeaderboardCohortSelection sets a custom function that can replace the cohort or opponent selection feature of event leaderboards.
	SetOnEventLeaderboardCohortSelection(fn OnEventLeaderboardCohortSelection)
}

type OnEventLeaderboardCohortSelection func(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, storageIndex string, eventID string, config *EventLeaderboardsConfigLeaderboard, userID string, tier int, matchmakerProperties map[string]interface{}) (cohortID string, cohortUserIDs []string, err error)
