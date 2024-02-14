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

// UnlockablesConfig is the data definition for a UnlockablesSystem type.
type UnlockablesConfig struct {
	ActiveSlots    int                                     `json:"active_slots,omitempty"`
	MaxActiveSlots int                                     `json:"max_active_slots,omitempty"`
	Slots          int                                     `json:"slots,omitempty"`
	SlotCost       *UnlockablesConfigSlotCost              `json:"slot_cost,omitempty"`
	Unlockables    map[string]*UnlockablesConfigUnlockable `json:"unlockables,omitempty"`

	UnlockableProbabilities []string `json:"-"`
}

type UnlockablesConfigSlotCost struct {
	Items      map[string]int64 `json:"items,omitempty"`
	Currencies map[string]int64 `json:"currencies,omitempty"`
}

type UnlockablesConfigUnlockable struct {
	Probability          int                                   `json:"probability,omitempty"`
	Category             string                                `json:"category,omitempty"`
	Cost                 *UnlockablesConfigUnlockableCost      `json:"cost,omitempty"`
	CostUnitTimeSec      int                                   `json:"cost_unit_time_sec,omitempty"`
	Description          string                                `json:"description,omitempty"`
	Name                 string                                `json:"name,omitempty"`
	StartCost            *UnlockablesConfigUnlockableStartCost `json:"start_cost,omitempty"`
	Reward               *EconomyConfigReward                  `json:"reward,omitempty"`
	WaitTimeSec          int                                   `json:"wait_time_sec,omitempty"`
	AdditionalProperties map[string]string                     `json:"additional_properties,omitempty"`
}

type UnlockablesConfigUnlockableCost struct {
	Items      map[string]int64 `json:"items,omitempty"`
	Currencies map[string]int64 `json:"currencies,omitempty"`
}

type UnlockablesConfigUnlockableStartCost struct {
	Items      map[string]int64 `json:"items,omitempty"`
	Currencies map[string]int64 `json:"currencies,omitempty"`
}

// The UnlockablesSystem is a gameplay system which provides slots to store rewards which can be unlocked over time.
type UnlockablesSystem interface {
	System

	// Create will place a new unlockable into a slot by ID.
	Create(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, unlockableID string) (*UnlockablesList, error)

	// Get returns all unlockables active for a user by ID.
	Get(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string) (*UnlockablesList, error)

	// UnlockStart will begin an unlock of an unlockable by instance ID for a user.
	UnlockStart(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, instanceID string) (*UnlockablesList, error)

	// PurchaseUnlock will immediately unlock an unlockable with the specified instance ID for a user.
	PurchaseUnlock(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, instanceID string) (*UnlockablesList, error)

	// PurchaseSlot will create a new slot for a user by ID.
	PurchaseSlot(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string) (*UnlockablesList, error)

	// Claim an unlockable which has been unlocked by instance ID for the user.
	Claim(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, instanceID string) (*UnlockablesReward, error)

	// SetOnClaimReward sets a custom reward function which will run after an unlockable's reward is rolled.
	SetOnClaimReward(fn OnReward[*UnlockablesConfigUnlockable])
}
