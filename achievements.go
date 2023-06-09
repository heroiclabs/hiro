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

// AchievementsConfig is the data definition for the TutorialsSystem type.
type AchievementsConfig struct {
	Achievements map[string]*AchievementsConfigAchievement `json:"achievements"`
}

type AchievementsConfigAchievement struct {
	AutoClaim            bool                                         `json:"auto_claim"`
	AutoClaimTotal       bool                                         `json:"auto_claim_total"`
	AutoReset            bool                                         `json:"auto_reset"`
	Category             string                                       `json:"category"`
	Count                int64                                        `json:"count"`
	Description          string                                       `json:"description"`
	ResetCronexpr        string                                       `json:"reset_cronexpr"`
	DurationSec          int64                                        `json:"duration_sec"`
	MaxCount             int64                                        `json:"max_count"`
	Name                 string                                       `json:"name"`
	PreconditionIDs      []string                                     `json:"precondition_ids"`
	Reward               *EconomyConfigReward                         `json:"reward"`
	TotalReward          *EconomyConfigReward                         `json:"total_reward"`
	SubAchievements      map[string]*AchievementsConfigSubAchievement `json:"sub_achievements"`
	AdditionalProperties map[string]string                            `json:"additional_properties"`
}

type AchievementsConfigSubAchievement struct {
	AutoClaim            bool                 `json:"auto_claim"`
	AutoReset            bool                 `json:"auto_reset"`
	Category             string               `json:"category"`
	Count                int64                `json:"count"`
	Description          string               `json:"description"`
	ResetCronexpr        string               `json:"reset_cronexpr"`
	DurationSec          int64                `json:"duration_sec"`
	MaxCount             int64                `json:"max_count"`
	Name                 string               `json:"name"`
	PreconditionIDs      []string             `json:"precondition_ids"`
	RewardConfig         *EconomyConfigReward `json:"reward_config"`
	AdditionalProperties map[string]string    `json:"additional_properties"`
}

// An AchievementsSystem is a gameplay system which represents one-off, repeat, preconditioned, and sub-achievements.
type AchievementsSystem interface {
	System

	// ClaimAchievements when one or more achievements whose progress has completed by their IDs.
	ClaimAchievements(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, achievementIDs []string, claimTotal bool) (map[string]*Achievement, map[string]*Achievement, error)

	// GetAchievements returns all achievements available to the user and progress on them.
	GetAchievements(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string) (map[string]*Achievement, map[string]*Achievement, error)

	// UpdateAchievements updates progress on one or more achievements by the same amount.
	UpdateAchievements(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, achievementIDs []string, amount int64) (map[string]*Achievement, map[string]*Achievement, error)

	// SetOnAchievementReward sets a custom reward function which will run after an achievement's reward is rolled.
	SetOnAchievementReward(fn OnReward[*AchievementsConfigAchievement])

	// SetOnSubAchievementReward sets a custom reward function which will run after a sub-achievement's reward is
	// rolled.
	SetOnSubAchievementReward(fn OnReward[*AchievementsConfigSubAchievement])

	// SetOnAchievementTotalReward sets a custom reward function which will run after an achievement's total reward is
	// rolled.
	SetOnAchievementTotalReward(fn OnReward[*AchievementsConfigAchievement])
}
