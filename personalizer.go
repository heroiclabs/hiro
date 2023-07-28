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
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

// The Personalizer describes an intermediate server or service which can be used to personalize the base data
// definitions defined for the gameplay systems.
type Personalizer interface {
	// GetValue returns a config which has been modified for a gameplay system.
	GetValue(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, system System, identity string) (any, error)
}

var _ Personalizer = &SatoriPersonalizer{}

type SatoriPersonalizer struct {
}

func (p *SatoriPersonalizer) GetValue(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, system System, userID string) (any, error) {
	var flagName string
	switch system.GetType() {
	case SystemTypeAchievements:
		flagName = "Hiro-Achievements"
	case SystemTypeBase:
		flagName = "Hiro-Base"
	case SystemTypeEconomy:
		flagName = "Hiro-Economy"
	case SystemTypeEnergy:
		flagName = "Hiro-Energy"
	case SystemTypeInventory:
		flagName = "Hiro-Inventory"
	case SystemTypeLeaderboards:
		flagName = "Hiro-Leaderboards"
	case SystemTypeTeams:
		flagName = "Hiro-Teams"
	case SystemTypeTutorials:
		flagName = "Hiro-Tutorials"
	case SystemTypeUnlockables:
		flagName = "Hiro-Unlockables"
	case SystemTypeStats:
		flagName = "Hiro-Stats"
	case SystemTypeProgression:
		flagName = "Hiro-Progression"
	default:
		return nil, runtime.NewError("hiro system type unknown", 3)
	}

	flagList, err := nk.GetSatori().FlagsList(ctx, userID, flagName)
	if err != nil {
		logger.WithField("userID", userID).WithField("error", err.Error()).Error("error requesting Satori flag list")
		return nil, err
	}

	// If this caller doesn't have the given flag, return no override.
	if len(flagList.Flags) < 1 {
		return nil, nil
	}

	config := system.GetConfig()
	if err := json.Unmarshal([]byte(flagList.Flags[0].Value), config); err != nil {
		logger.WithField("userID", userID).WithField("error", err.Error()).Error("error merging Satori flag value")
		return nil, err
	}

	return config, nil
}
