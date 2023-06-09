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

type InventoryConfig struct {
	Items    map[string]*InventoryConfigItem `json:"items"`
	ItemSets map[string]map[string]bool      `json:"-"` // Auto-computed when the config is read or personalized.
}

type InventoryConfigItem struct {
	Name              string               `json:"name"`
	Description       string               `json:"description"`
	Category          string               `json:"category"`
	ItemSets          []string             `json:"item_sets"`
	MaxCount          int64                `json:"max_count"`
	Stackable         bool                 `json:"stackable"`
	Consumable        bool                 `json:"consumable"`
	ConsumeReward     *EconomyConfigReward `json:"consume_reward"`
	StringProperties  map[string]string    `json:"string_properties"`
	NumericProperties map[string]float64   `json:"numeric_properties"`
}

// The InventorySystem provides a gameplay system which can manage a player's inventory.
//
// A player can have items added via economy rewards, or directly.
type InventorySystem interface {
	System

	// List will return the items defined as well as the computed item sets for the user by ID.
	List(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string) (map[string]*InventoryConfigItem, map[string][]string, error)

	// ListInventoryItems will return the items which are part of a user's inventory by ID.
	ListInventoryItems(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, category string) (*Inventory, error)

	// ConsumeItems will deduct the item(s) from the user's inventory and run the consume reward for each one, if
	// defined.
	ConsumeItems(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, items map[string]int64, overConsume bool) (*Inventory, map[string][]*Reward, error)

	// GrantItems will add the item(s) to a user's inventory by ID.
	GrantItems(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, items map[string]int64) (*Inventory, error)

	// UpdateItems will update the properties which are stored on each item by ID for a user.
	UpdateItems(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, itemUpdates map[string]*InventoryUpdateItemProperties) (*Inventory, error)

	// SetOnConsumeReward sets a custom reward function which will run after an inventory items' consume reward is
	// rolled.
	SetOnConsumeReward(fn OnReward[*InventoryConfigItem])
}
