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
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

var (
	ErrEconomyNoItem            = runtime.NewError("item not found", 3)                        // INVALID_ARGUMENT
	ErrEconomyNoSku             = runtime.NewError("sku not found", 3)                         // INVALID_ARGUMENT
	ErrEconomySkuInvalid        = runtime.NewError("invalid sku", 3)                           // INVALID_ARGUMENT
	ErrEconomyNotEnoughCurrency = runtime.NewError("not enough currency for purchase", 3)      // INVALID_ARGUMENT
	ErrEconomyNotEnoughItem     = runtime.NewError("not enough item", 3)                       // INVALID_ARGUMENT
	ErrEconomyReceiptInvalid    = runtime.NewError("invalid receipt", 3)                       // INVALID_ARGUMENT
	ErrEconomyReceiptDuplicate  = runtime.NewError("duplicate receipt", 3)                     // INVALID_ARGUMENT
	ErrEconomyReceiptMismatch   = runtime.NewError("mismatched product receipt", 3)            // INVALID_ARGUMENT
	ErrEconomyNoPlacement       = runtime.NewError("placement not found", 3)                   // INVALID_ARGUMENT
	ErrEconomyNoDonation        = runtime.NewError("donation not found", 3)                    // INVALID_ARGUMENT
	ErrEconomyMaxDonation       = runtime.NewError("donation maximum contribution reached", 3) // INVALID_ARGUMENT
	ErrEconomyClaimedDonation   = runtime.NewError("donation already claimed", 3)              // INVALID_ARGUMENT

	ErrInventoryNotInitialized = runtime.NewError("inventory not initialized for batch", 13) // INTERNAL
	ErrItemsNotConsumable      = runtime.NewError("items not consumable", 3)                 // INVALID_ARGUMENT
	ErrItemsInsufficient       = runtime.NewError("insufficient items", 9)                   // FAILED_PRECONDITION
	ErrCurrencyInsufficient    = runtime.NewError("insufficient currency", 9)                // FAILED_PRECONDITION
)

// EconomyConfig is the data definition for the EconomySystem type.
type EconomyConfig struct {
	InitializeUser *EconomyConfigInitializeUser       `json:"initialize_user,omitempty"`
	Donations      map[string]*EconomyConfigDonation  `json:"donations,omitempty"`
	StoreItems     map[string]*EconomyConfigStoreItem `json:"store_items,omitempty"`
	Placements     map[string]*EconomyConfigPlacement `json:"placements,omitempty"`
}

type EconomyConfigDonation struct {
	Cost                     *EconomyConfigDonationCost `json:"cost,omitempty"`
	Count                    int64                      `json:"count,omitempty"`
	Description              string                     `json:"description,omitempty"`
	DurationSec              int64                      `json:"duration_sec,omitempty"`
	MaxCount                 int64                      `json:"max_count,omitempty"`
	Name                     string                     `json:"name,omitempty"`
	RecipientReward          *EconomyConfigReward       `json:"recipient_reward,omitempty"`
	ContributorReward        *EconomyConfigReward       `json:"contributor_reward,omitempty"`
	UserContributionMaxCount int64                      `json:"user_contribution_max_count,omitempty"`
	AdditionalProperties     map[string]string          `json:"additional_properties,omitempty"`
}

type EconomyConfigDonationCost struct {
	Currencies map[string]int64 `json:"currencies,omitempty"`
	Items      map[string]int64 `json:"items,omitempty"`
}

type EconomyConfigInitializeUser struct {
	Currencies map[string]int64 `json:"currencies,omitempty"`
	Items      map[string]int64 `json:"items,omitempty"`
}

type EconomyConfigPlacement struct {
	Reward               *EconomyConfigReward `json:"reward,omitempty"`
	AdditionalProperties map[string]string    `json:"additional_properties,omitempty"`
}

type EconomyConfigReward struct {
	Guaranteed     *EconomyConfigRewardContents   `json:"guaranteed,omitempty"`
	Weighted       []*EconomyConfigRewardContents `json:"weighted,omitempty"`
	MaxRolls       int64                          `json:"max_rolls,omitempty"`
	MaxRepeatRolls int64                          `json:"max_repeat_rolls,omitempty"`
	TotalWeight    int64                          `json:"total_weight,omitempty"`
}

type EconomyConfigRewardContents struct {
	Items           map[string]*EconomyConfigRewardItem     `json:"items,omitempty"`
	ItemSets        []*EconomyConfigRewardItemSet           `json:"item_sets,omitempty"`
	Currencies      map[string]*EconomyConfigRewardCurrency `json:"currencies,omitempty"`
	Energies        map[string]*EconomyConfigRewardEnergy   `json:"energies,omitempty"`
	EnergyModifiers []*EconomyConfigRewardEnergyModifier    `json:"energy_modifiers,omitempty"`
	RewardModifiers []*EconomyConfigRewardRewardModifier    `json:"reward_modifiers,omitempty"`
	Weight          int64                                   `json:"weight,omitempty"`
}

type EconomyConfigRewardCurrency struct {
	EconomyConfigRewardRangeInt64
}

type EconomyConfigRewardEnergy struct {
	EconomyConfigRewardRangeInt32
}

type EconomyConfigRewardEnergyModifier struct {
	Id          string                         `json:"id,omitempty"`
	Operator    string                         `json:"operator,omitempty"`
	Value       *EconomyConfigRewardRangeInt64 `json:"value,omitempty"`
	DurationSec *EconomyConfigRewardRangeInt64 `json:"duration_sec,omitempty"`
}

type EconomyConfigRewardItem struct {
	EconomyConfigRewardRangeInt64
}

type EconomyConfigRewardItemSet struct {
	EconomyConfigRewardRangeInt64

	MaxRepeats int64    `json:"max_repeats,omitempty"`
	Set        []string `json:"set,omitempty"`
}

type EconomyConfigRewardRangeInt32 struct {
	Min      int32 `json:"min,omitempty"`
	Max      int32 `json:"max,omitempty"`
	Multiple int32 `json:"multiple,omitempty"`
}

type EconomyConfigRewardRangeInt64 struct {
	Min      int64 `json:"min,omitempty"`
	Max      int64 `json:"max,omitempty"`
	Multiple int64 `json:"multiple,omitempty"`
}

type EconomyConfigRewardRewardModifier struct {
	Id          string                         `json:"id,omitempty"`
	Type        string                         `json:"type,omitempty"`
	Operator    string                         `json:"operator,omitempty"`
	Value       *EconomyConfigRewardRangeInt64 `json:"value,omitempty"`
	DurationSec *EconomyConfigRewardRangeInt64 `json:"duration_sec,omitempty"`
}

type EconomyConfigStoreItem struct {
	Category             string                      `json:"category,omitempty"`
	Cost                 *EconomyConfigStoreItemCost `json:"cost,omitempty"`
	Description          string                      `json:"description,omitempty"`
	Name                 string                      `json:"name,omitempty"`
	Reward               *EconomyConfigReward        `json:"reward,omitempty"`
	AdditionalProperties map[string]string           `json:"additional_properties,omitempty"`
}

type EconomyConfigStoreItemCost struct {
	Currencies map[string]int64 `json:"currencies,omitempty"`
	Sku        string           `json:"sku,omitempty"`
}

// The EconomySystem is the foundation of a game's economy.
//
// It provides functionality for 4 different reward types: basic, gacha, weighted table, and custom. These rolled
// rewards are available to generate in all other gameplay systems and can be generated manually as well.
type EconomySystem interface {
	System

	// RewardCreate prepares a new reward configuration to be filled in and used later.
	RewardCreate() *EconomyConfigReward

	// RewardRoll takes a reward configuration and rolls an actual reward from it, applying all appropriate rules.
	RewardRoll(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, rewardConfig *EconomyConfigReward) (*Reward, error)

	// RewardGrant updates a user's economy, inventory, and/or energy models with the contents of a rolled reward.
	RewardGrant(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, reward *Reward, metadata map[string]interface{}) error

	// DonationClaim will claim donation rewards for a user and the given donation IDs.
	DonationClaim(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, donationIDs []string) (*EconomyDonationsList, error)

	// DonationGet will get all donations for the given list of user IDs.
	DonationGet(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userIDs []string) (*EconomyDonationsByUserList, error)

	// DonationGive will contribute to a particular donation for a user ID.
	DonationGive(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, donationID, fromUserID string) (map[string]int64, *Inventory, *Reward, error)

	// DonationRequest will create a donation request for a given donation ID and user ID.
	DonationRequest(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, donationID string) (*EconomyDonation, bool, error)

	// List will get the defined store items and placements within the economy system.
	List(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string) (map[string]*EconomyConfigStoreItem, map[string]*EconomyConfigPlacement, error)

	// Grant will add currencies, and reward modifiers to a user's economy by ID.
	Grant(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, currencies map[string]int64, items map[string]int64, modifiers []*RewardModifier, walletMetadata map[string]interface{}) (map[string]int64, error)

	// UnmarshalWallet unmarshals and returns the account's wallet as a map[string]int64.
	UnmarshalWallet(account *api.Account) (map[string]int64, error)

	// PurchaseIntent will create a purchase intent for a particular store item for a user ID.
	PurchaseIntent(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, itemID string, store EconomyStoreType, sku string) error

	// PurchaseItem will validate a purchase and give the user ID the appropriate rewards.
	PurchaseItem(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, itemID string, store EconomyStoreType, receipt string) (map[string]int64, *Inventory, *Reward, bool, error)

	// PurchaseRestore will process a restore attempt for the given user, based on a set of restore receipts.
	PurchaseRestore(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, store EconomyStoreType, receipts []string) error

	// PlacementStatus will get the status of a specified placement.
	PlacementStatus(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, rewardID, placementID string, retryCount int) (*EconomyPlacementStatus, error)

	// PlacementStart will indicate that a user ID has begun viewing an ad placement.
	PlacementStart(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, placementID string) (*EconomyPlacementStatus, error)

	// PlacementSuccess will indicate that the user ID has successfully viewed an ad placement and provide the appropriate reward.
	PlacementSuccess(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, rewardID, placementID string) (*Reward, error)

	// PlacementFail will indicate that the user ID has failed to successfully view the ad placement.
	PlacementFail(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, rewardID, placementID string) error

	// SetOnDonationClaimReward sets a custom reward function which will run after a donation's reward is rolled.
	SetOnDonationClaimReward(fn OnReward[*EconomyConfigDonation])

	// SetOnDonationContributorReward sets a custom reward function which will run after a donation's sender reward is rolled.
	SetOnDonationContributorReward(fn OnReward[*EconomyConfigDonation])

	// SetOnPlacementReward sets a custom reward function which will run after a placement's reward is rolled.
	SetOnPlacementReward(fn OnReward[*EconomyConfigPlacement])

	// SetOnStoreItemReward sets a custom reward function which will run after store item's reward is rolled.
	SetOnStoreItemReward(fn OnReward[*EconomyConfigStoreItem])
}
