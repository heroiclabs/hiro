// Copyright 2024 Heroic Labs & Contributors
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

var ErrStreakResetInvalid = runtime.NewError("streak reset schedule invalid", 13)

// StreaksConfig is the data definition for a StreaksSystem type.
type StreaksConfig struct {
	Streaks map[string]*StreaksConfigStreak `json:"streaks,omitempty"`
}

type StreaksConfigStreak struct {
	AutoClaim            bool                         `json:"auto_claim,omitempty"`
	Name                 string                       `json:"name,omitempty"`
	Description          string                       `json:"description,omitempty"`
	Count                int64                        `json:"count,omitempty"`
	MaxCount             int64                        `json:"max_count,omitempty"`
	MaxCountCurrentReset int64                        `json:"max_count_current_reset,omitempty"`
	IdleCountDecayReset  int64                        `json:"idle_count_decay_reset,omitempty"`
	MaxIdleCountDecay    int64                        `json:"max_idle_count_decay,omitempty"`
	ResetCronexpr        string                       `json:"reset_cronexpr,omitempty"`
	TimeOffsetSec        int64                        `json:"time_offset_sec,omitempty"`
	GraceWindowSec       int64                        `json:"grace_window_sec,omitempty"`
	Rewards              []*StreaksConfigStreakReward `json:"rewards,omitempty"`
	StartTimeSec         int64                        `json:"start_time_sec,omitempty"`
	EndTimeSec           int64                        `json:"end_time_sec,omitempty"`
	Disabled             bool                         `json:"disabled,omitempty"`
}

type StreaksConfigStreakReward struct {
	CountMin   int64                `json:"count_min,omitempty"`
	CountMax   int64                `json:"count_max,omitempty"`
	Reward     *EconomyConfigReward `json:"reward,omitempty"`
	Repeatable bool                 `json:"repeatable,omitempty"`
}

// StreakChange reports a streak's state before (nil if it did not exist) and after a change.
type StreakChange struct {
	ID  string
	Old *StreakSnapshot
	New *StreakSnapshot
}

// OnStreakChange is invoked after streaks are persisted, with the streaks changed by the operation.
type OnStreakChange func(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, changes []*StreakChange)

// StreakSnapshot is a point-in-time copy of a streak's stored state.
type StreakSnapshot struct {
	Count             int64           `json:"count,omitempty"`
	CountCurrentReset int64           `json:"count_current_reset,omitempty"`
	MaxCountReached   int64           `json:"max_count_reached,omitempty"`
	ClaimCount        int64           `json:"claim_count,omitempty"`
	CreateTimeSec     int64           `json:"create_time_sec,omitempty"`
	UpdateTimeSec     int64           `json:"update_time_sec,omitempty"`
	ClaimTimeSec      int64           `json:"claim_time_sec,omitempty"`
	ClaimedRewards    []*StreakReward `json:"claimed_rewards,omitempty"`
}

type StreaksSystem interface {
	System

	// List all streaks and their current state and progress for a given user.
	List(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string) (streaks map[string]*Streak, err error)

	// Update one or more streaks with the indicated counts for the given user.
	Update(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, streakIDs map[string]int64) (streaks map[string]*Streak, err error)

	// Claim rewards for one or more streaks for the given user.
	Claim(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, streakIDs []string) (streaks map[string]*Streak, err error)

	// Reset progress on selected streaks for the given user.
	Reset(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, streakIDs []string) (streaks map[string]*Streak, err error)

	// Set overwrites streak counts; grantRewards=false marks rewards up to the count claimed, true keeps them claimable.
	Set(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, counts map[string]int64, grantRewards bool) (streaks map[string]*Streak, err error)

	// SetState overwrites streaks with the given snapshots, stashing the prior state for revert.
	SetState(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, snapshots map[string]*StreakSnapshot) (streaks map[string]*Streak, err error)

	// Revert restores each streak's last stashed snapshot, swapping it with the current state.
	Revert(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, streakIDs []string) (streaks map[string]*Streak, err error)

	// SetOnClaimReward sets a custom reward function which will run after a streak's reward is rolled.
	SetOnClaimReward(fn OnReward[*StreaksConfigStreak])

	// SetOnStreakChange sets a custom function which will run after streaks are persisted, with the streaks changed by the operation.
	SetOnStreakChange(fn OnStreakChange)
}
