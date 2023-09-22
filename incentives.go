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

type IncentivesConfig struct {
	Incentives map[string]*IncentivesConfigIncentive `json:"incentives"`
}

type IncentivesConfigIncentive struct {
	Type               IncentiveType        `json:"type"`
	Name               string               `json:"name"`
	Description        string               `json:"description"`
	MaxClaims          int                  `json:"max_claims"`
	MaxGlobalClaims    int                  `json:"max_global_claims"`
	MaxRecipientAgeSec int64                `json:"max_recipient_age_sec"`
	RecipientReward    *EconomyConfigReward `json:"recipient_reward"`
	SenderReward       *EconomyConfigReward `json:"sender_reward"`
	MaxConcurrent      int                  `json:"max_concurrent"`
	ExpiryDurationSec  int64                `json:"expiry_duration_sec"`
}

// The IncentivesSystem provides a gameplay system which can create and claim incentives and their associated rewards.
type IncentivesSystem interface {
	System

	SenderList(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string) ([]*Incentive, error)

	SenderCreate(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, incentiveID string) ([]*Incentive, error)

	SenderDelete(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, code string) ([]*Incentive, error)

	SenderClaim(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, code string, claimantIDs []string) ([]*Incentive, error)

	RecipientGet(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, code string) (*IncentiveInfo, error)

	RecipientClaim(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, code string) (*IncentiveInfo, error)

	// SetOnSenderReward sets a custom reward function which will run after an incentive sender's reward is rolled.
	SetOnSenderReward(fn OnReward[*IncentivesConfigIncentive])

	// SetOnRecipientReward sets a custom reward function which will run after an incentive recipient's reward is rolled.
	SetOnRecipientReward(fn OnReward[*IncentivesConfigIncentive])
}
