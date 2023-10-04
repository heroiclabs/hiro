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
	EventLeaderboards map[string]*EventLeaderboardsConfigLeaderboard `json:"event_leaderboards"`
}

type EventLeaderboardsConfigLeaderboard struct {
	Name                 string                                                     `json:"name"`
	Description          string                                                     `json:"description"`
	Category             string                                                     `json:"category"`
	Ascending            bool                                                       `json:"ascending"`
	Operator             string                                                     `json:"operator"`
	ResetSchedule        string                                                     `json:"reset_schedule"`
	CohortSize           int                                                        `json:"cohort_size"`
	AdditionalProperties map[string]string                                          `json:"additional_properties"`
	MaxNumScore          int                                                        `json:"max_num_score"`
	RewardTiers          map[string][]*EventLeaderboardsConfigLeaderboardRewardTier `json:"reward_tiers"`
	ChangeZones          map[string]*EventLeaderboardsConfigChangeZone              `json:"change_zones"`
	Tiers                int                                                        `json:"tiers"`
	MaxIdleTierDrop      int                                                        `json:"max_idle_tier_drop"`
	StartTimeSec         int64                                                      `json:"start_time_sec"`
	EndTimeSec           int64                                                      `json:"end_time_sec"`
	Duration             int64                                                      `json:"duration"`
}

type EventLeaderboardsConfigLeaderboardRewardTier struct {
	RankMax    int                  `json:"rank_max"`
	RankMin    int                  `json:"rank_min"`
	Reward     *EconomyConfigReward `json:"reward"`
	TierChange int                  `json:"tier_change"`
}

type EventLeaderboardsConfigChangeZone struct {
	Promotion  float64 `json:"promotion"`
	Demotion   float64 `json:"demotion"`
	DemoteIdle bool    `json:"demote_idle"`
}

// An EventLeaderboardsSystem is a gameplay system which represents cohort-segmented, tier-based event leaderboards.
type EventLeaderboardsSystem interface {
	System

	// GetEventLeaderboard returns a specified event leaderboard's cohort for the user.
	GetEventLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, username, eventLeaderboardID string) (*EventLeaderboard, error)

	// RollEventLeaderboard places the user into a new cohort for the specified event leaderboard if possible.
	RollEventLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, username, eventLeaderboardID string) (*EventLeaderboard, error)

	// UpdateEventLeaderboard updates the user's score in the specified event leaderboard, and returns the user's updated cohort information.
	UpdateEventLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, username, eventLeaderboardID string, score, subscore int64) (*EventLeaderboard, error)

	// ClaimEventLeaderboard claims the user's reward for the given event leaderboard.
	ClaimEventLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, eventLeaderboardID string) (*EventLeaderboard, error)

	// SetOnEventLeaderboardsReward sets a custom reward function which will run after an event leaderboard's reward is rolled.
	SetOnEventLeaderboardsReward(fn OnReward[*EventLeaderboardsConfigLeaderboard])

	// SetOnEventLeaderboardCohortSelection sets a custom function that can replace the cohort or opponent selection feature of event leaderboards.
	SetOnEventLeaderboardCohortSelection(fn OnEventLeaderboardCohortSelection)
}

type OnEventLeaderboardCohortSelection func(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, eventID string, config *EventLeaderboardsConfigLeaderboard, userID string, tier int) (cohortID string, cohortUserIDs []string, err error)
