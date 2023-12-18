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
	"strings"

	"github.com/heroiclabs/nakama-common/runtime"
)

// The Personalizer describes an intermediate server or service which can be used to personalize the base data
// definitions defined for the gameplay systems.
type Personalizer interface {
	// GetValue returns a config which has been modified for a gameplay system,
	// or nil if the config is not being adjusted by this personalizer.
	GetValue(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, system System, identity string) (any, error)
}

var _ Personalizer = &SatoriPersonalizer{}

type SatoriPersonalizer struct {
	publishAuthenticateRequest bool
	publishCoreEvents          bool
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
	case SystemTypeEventLeaderboards:
		flagName = "Hiro-Event-Leaderboards"
	case SystemTypeProgression:
		flagName = "Hiro-Progression"
	case SystemTypeIncentives:
		flagName = "Hiro-Incentives"
	default:
		return nil, runtime.NewError("hiro system type unknown", 3)
	}

	flagList, err := nk.GetSatori().FlagsList(ctx, userID, flagName)
	if err != nil {
		logger.WithField("userID", userID).WithField("error", err.Error()).Error("error requesting Satori flag list")
		return nil, err
	}

	var config any
	var found bool

	if len(flagList.Flags) >= 1 {
		config = system.GetConfig()
		decoder := json.NewDecoder(strings.NewReader(flagList.Flags[0].Value))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(config); err != nil {
			logger.WithField("userID", userID).WithField("error", err.Error()).Error("error merging Satori flag value")
			return nil, err
		}
		found = true
	}

	if system.GetType() == SystemTypeEventLeaderboards {
		// If looking at event leaderboards, also load live events.
		liveEventsList, err := nk.GetSatori().LiveEventsList(ctx, userID)
		if err != nil {
			logger.WithField("userID", userID).WithField("error", err.Error()).Error("error requesting Satori live events list")
			return nil, err
		}
		if len(liveEventsList.LiveEvents) > 0 {
			if config == nil {
				config = system.GetConfig()
			}
			for _, liveEvent := range liveEventsList.LiveEvents {
				decoder := json.NewDecoder(strings.NewReader(liveEvent.Value))
				decoder.DisallowUnknownFields()
				if err := decoder.Decode(config); err != nil {
					// The live event may be intended for a different purpose, do not log or return an error here.
					continue
				}
				found = true
			}
		}
	}

	// If this caller doesn't have the given flag (or live events) return the nil to indicate no change to the config.
	if !found {
		return nil, nil
	}

	return config, nil
}

func (p *SatoriPersonalizer) IsPublishAuthenticateRequest() bool {
	return p.publishAuthenticateRequest
}

func (p *SatoriPersonalizer) IsPublishCoreEvents() bool {
	return p.publishCoreEvents
}

func NewSatoriPersonalizer(publishAuthenticateEvent, publishCoreEvents bool) *SatoriPersonalizer {
	return &SatoriPersonalizer{
		publishAuthenticateRequest: publishAuthenticateEvent,
		publishCoreEvents:          publishCoreEvents,
	}
}
