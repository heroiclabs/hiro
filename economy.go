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

// EconomyConfig is the data definition for the EconomySystem type.
type EconomyConfig struct {
	InitializeUser *EconomyConfigInitializeUser       `json:"initialize_user"`
	Donations      map[string]*EconomyConfigDonation  `json:"donations"`
	StoreItems     map[string]*EconomyConfigStoreItem `json:"store_items"`
	Placements     map[string]*EconomyConfigPlacement `json:"placements"`
}

type EconomyConfigDonation struct {
	Cost                     *EconomyConfigDonationCost `json:"cost"`
	Count                    int64                      `json:"count"`
	Description              string                     `json:"description"`
	DurationSec              int64                      `json:"duration_sec"`
	MaxCount                 int64                      `json:"max_count"`
	Name                     string                     `json:"name"`
	RecipientRewardConfig    *EconomyConfigReward       `json:"recipient_reward_config"`
	ContributorRewardConfig  *EconomyConfigReward       `json:"contributor_reward_config"`
	UserContributionMaxCount int64                      `json:"user_contribution_max_count"`
	AdditionalProperties     map[string]string          `json:"additional_properties"`
}

type EconomyConfigDonationCost struct {
	Currencies map[string]int64 `json:"currencies"`
	Items      map[string]int64 `json:"items"`
}

type EconomyConfigInitializeUser struct {
	Currencies map[string]int64 `json:"currencies"`
	Items      map[string]int64 `json:"items"`
}

type EconomyConfigPlacement struct {
	RewardConfig         *EconomyConfigReward `json:"reward_config"`
	AdditionalProperties map[string]string    `json:"additional_properties"`
}

type EconomyConfigReward struct {
	Guaranteed  *EconomyConfigRewardContents   `json:"guaranteed"`
	Weighted    []*EconomyConfigRewardContents `json:"weighted"`
	MaxRolls    int64                          `json:"max_rolls"`
	TotalWeight int64                          `json:"total_weight"`
}

type EconomyConfigRewardContents struct {
	Items           map[string]*EconomyConfigRewardItem     `json:"items"`
	ItemSets        []*EconomyConfigRewardItemSet           `json:"item_sets"`
	Currencies      map[string]*EconomyConfigRewardCurrency `json:"currencies"`
	Energies        map[string]*EconomyConfigRewardEnergy   `json:"energies"`
	EnergyModifiers []*EconomyConfigRewardEnergyModifier    `json:"energy_modifiers"`
	RewardModifiers []*EconomyConfigRewardRewardModifier    `json:"reward_modifiers"`
	Weight          int64                                   `json:"weight"`
}

type EconomyConfigRewardCurrency struct {
	EconomyConfigRewardRangeInt64
}

type EconomyConfigRewardEnergy struct {
	EconomyConfigRewardRangeInt32
}

type EconomyConfigRewardEnergyModifier struct {
	Id          string                         `json:"id"`
	Operator    string                         `json:"operator"`
	Value       *EconomyConfigRewardRangeInt64 `json:"value"`
	DurationSec *EconomyConfigRewardRangeInt64 `json:"duration_sec"`
}

type EconomyConfigRewardItem struct {
	EconomyConfigRewardRangeInt64
}

type EconomyConfigRewardItemSet struct {
	EconomyConfigRewardRangeInt64

	MaxRepeats int64    `json:"max_repeats"`
	Set        []string `json:"set"`
}

type EconomyConfigRewardRangeInt32 struct {
	Min      int32 `json:"min"`
	Max      int32 `json:"max"`
	Multiple int32 `json:"multiple"`
}

type EconomyConfigRewardRangeInt64 struct {
	Min      int64 `json:"min"`
	Max      int64 `json:"max"`
	Multiple int64 `json:"multiple"`
}

type EconomyConfigRewardRewardModifier struct {
	Id          string                         `json:"id"`
	Type        string                         `json:"type"`
	Operator    string                         `json:"operator"`
	Value       *EconomyConfigRewardRangeInt64 `json:"value"`
	DurationSec *EconomyConfigRewardRangeInt64 `json:"duration_sec"`
}

type EconomyConfigStoreItem struct {
	Category             string                      `json:"category"`
	Cost                 *EconomyConfigStoreItemCost `json:"cost"`
	Description          string                      `json:"description"`
	Name                 string                      `json:"name"`
	RewardConfig         *EconomyConfigReward        `json:"reward_config"`
	AdditionalProperties map[string]string           `json:"additional_properties"`
}

type EconomyConfigStoreItemCost struct {
	Currencies map[string]int64 `json:"currencies"`
	Sku        string           `json:"sku"`
}

// The EconomySystem is the foundation of a game's economy.
//
// It provides functionality for 4 different reward types: basic, gacha, weighted table, and custom. These rolled
// rewards are available to generate in all other gameplay systems and can be generated manually as well.
type EconomySystem interface {
	System

	// Grant will add currencies, and reward modifiers to a user's economy by ID.
	Grant(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, currencies map[string]int64, modifiers []*RewardModifier) (map[string]int64, error)

	// SetOnDonationClaimReward sets a custom reward function which will run after a donation's reward is rolled.
	SetOnDonationClaimReward(fn OnReward[*EconomyConfigDonation])

	// SetOnDonationContributorReward sets a custom reward function which will run after a donation's sender reward is rolled.
	SetOnDonationContributorReward(fn OnReward[*EconomyConfigDonation])

	// SetOnPlacementReward sets a custom reward function which will run after a placement's reward is rolled.
	SetOnPlacementReward(fn OnReward[*EconomyConfigPlacement])

	// SetOnStoreItemReward sets a custom reward function which will run after store item's reward is rolled.
	SetOnStoreItemReward(fn OnReward[*EconomyConfigStoreItem])
}
