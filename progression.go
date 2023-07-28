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

// ProgressionConfig is the data definition for a ProgressionSystem type.
type ProgressionConfig struct {
	Progressions map[string]*ProgressionConfigProgression `json:"progressions"`
}

type ProgressionConfigProgression struct {
	Id                   string                            `json:"id"`
	Name                 string                            `json:"name"`
	Description          string                            `json:"description"`
	Category             string                            `json:"category"`
	PreconditionIDs      []string                          `json:"precondition_ids"`
	Cost                 *ProgressionConfigProgressionCost `json:"cost"`
	Count                int64                             `json:"count"`
	MaxCount             int64                             `json:"max_count"`
	Reward               *EconomyConfigReward              `json:"reward"`
	AdditionalProperties map[string]string                 `json:"additional_properties"`
}

type ProgressionConfigProgressionCost struct {
	Items      map[string]int64 `json:"items"`
	Currencies map[string]int64 `json:"currencies"`
}

// A ProgressionSystem is a gameplay system which represents a sequence of progression steps.
type ProgressionSystem interface {
	System

	// GetProgression returns all progressions visible to the user, even if they are not available yet.
	GetProgression(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string) (map[string]*Progression, error)

	// UpdateProgression updates the count of a progression, spending any associated cost, and optionally claiming any reward if possible.
	UpdateProgression(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, progressionIDs []string, amount int64, autoClaim bool) (map[string]*Progression, error)

	// ClaimProgression claims one or more progressions which have met the conditions to be claimed and have their rewards granted.
	ClaimProgression(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, progressionIDs []string) (map[string]*Progression, error)

	// SetOnProgressionReward sets a custom reward function which will run after an progression's reward is rolled.
	SetOnProgressionReward(fn OnReward[*ProgressionConfigProgression])
}
