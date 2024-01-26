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
	"database/sql"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

const StoragePersonalizerCollectionDefault = "hiro_datadefinitions"
const storagePersonalizerKeyAchievements = "achievements"
const storagePersonalizerKeyEconomy = "economy"
const storagePersonalizerKeyEnergy = "energy"
const storagePersonalizerKeyEventLeaderboards = "event_leaderboards"
const storagePersonalizerKeyIncentives = "incentives"
const storagePersonalizerKeyLeaderboards = "leaderboards"
const storagePersonalizerKeyProgression = "progression"
const storagePersonalizerKeyStats = "stats"
const storagePersonalizerKeyTeams = "teams"
const storagePersonalizerKeyTutorials = "tutorials"
const storagePersonalizerKeyUnlockables = "unlockables"
const storagePersonalizerKeyBase = "base"

var _ Personalizer = (*StoragePersonalizer)(nil)

type StoragePersonalizerCachedStorageObject struct {
	object      *api.StorageObject
	refreshTime time.Time
	expiryTime  time.Time
}

type StoragePersonalizer struct {
	sync.RWMutex
	cache       map[SystemType]*StoragePersonalizerCachedStorageObject
	cacheExpiry time.Duration
	collection  string
	logger      runtime.Logger
}

type storagePersonalizerUploadRequest struct {
	Achievements     *AchievementsConfig      `json:"achievements"`
	Economy          *EconomyConfig           `json:"economy"`
	Energy           *EnergyConfig            `json:"energy"`
	EventLeaderboard *EventLeaderboardsConfig `json:"event_leaderboard"`
	Incentives       *IncentivesConfig        `json:"incentives"`
	Leaderboards     *LeaderboardConfig       `json:"leaderboards"`
	Progression      *ProgressionConfig       `json:"progression"`
	Stats            *StatsConfig             `json:"stats"`
	Teams            *TeamsConfig             `json:"teams"`
	Tutorials        *TutorialsConfig         `json:"tutorials"`
	Unlockables      *UnlockablesConfig       `json:"unlockables"`
	Base             *BaseSystemConfig        `json:"base"`
}

func NewStoragePersonalizerDefault(initializer *runtime.Initializer) *StoragePersonalizer {
	return &StoragePersonalizer{
		cache:       make(map[SystemType]*StoragePersonalizerCachedStorageObject, 20),
		cacheExpiry: 10 * time.Minute,
		collection:  StoragePersonalizerCollectionDefault,
	}
}

func NewStoragePersonalizer(logger runtime.Logger, cacheExpirySec int, collection string, initializer runtime.Initializer, register bool) *StoragePersonalizer {
	personalizer := &StoragePersonalizer{
		cache:       make(map[SystemType]*StoragePersonalizerCachedStorageObject, 20),
		cacheExpiry: time.Duration(cacheExpirySec) * time.Second,
		collection:  collection,
		logger:      logger,
	}

	if register {
		err := initializer.RegisterRpc(RpcId_RPC_ID_STORAGE_PERSONALIZER_UPLOAD.String(), rpcStoragePersonalizerUpload(initializer, personalizer))
		if err != nil {
			logger.WithField("error", err.Error()).Error("Error registering storage personalizer upload RPC.")
		}
	}

	return personalizer
}

func rpcStoragePersonalizerUpload(initializer runtime.Initializer, p *StoragePersonalizer) func(context.Context, runtime.Logger, *sql.DB, runtime.NakamaModule, string) (string, error) {
	return func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
		_, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
		if ok {
			return "", ErrSessionUser
		}

		req := &storagePersonalizerUploadRequest{}
		err := json.Unmarshal([]byte(payload), req)
		if err != nil {
			return "", ErrPayloadDecode
		}

		writes := []*runtime.StorageWrite{}

		if req.Achievements != nil {
			write, err := p.newStorageWrite(req.Achievements, storagePersonalizerKeyAchievements)
			if err != nil {
				logger.WithField("error", err.Error()).Error("Error creating achievements storage object.")
				return "", ErrInternal
			}

			writes = append(writes, write)
		}

		if req.Economy != nil {
			write, err := p.newStorageWrite(req.Economy, storagePersonalizerKeyEconomy)
			if err != nil {
				logger.WithField("error", err.Error()).Error("Error creating economy storage object.")
				return "", ErrInternal
			}

			writes = append(writes, write)
		}

		if req.Energy != nil {
			write, err := p.newStorageWrite(req.Energy, storagePersonalizerKeyEnergy)
			if err != nil {
				logger.WithField("error", err.Error()).Error("Error creating energy storage object.")
				return "", ErrInternal
			}

			writes = append(writes, write)
		}

		if req.EventLeaderboard != nil {
			write, err := p.newStorageWrite(req.EventLeaderboard, storagePersonalizerKeyEventLeaderboards)
			if err != nil {
				logger.WithField("error", err.Error()).Error("Error creating event leaderboard storage object.")
				return "", ErrInternal
			}

			writes = append(writes, write)
		}

		if req.Incentives != nil {
			write, err := p.newStorageWrite(req.Incentives, storagePersonalizerKeyIncentives)
			if err != nil {
				logger.WithField("error", err.Error()).Error("Error creating incentives storage object.")
				return "", ErrInternal
			}

			writes = append(writes, write)
		}

		if req.Leaderboards != nil {
			write, err := p.newStorageWrite(req.Leaderboards, storagePersonalizerKeyLeaderboards)
			if err != nil {
				logger.WithField("error", err.Error()).Error("Error creating leaderboards storage object.")
				return "", ErrInternal
			}

			writes = append(writes, write)
		}

		if req.Progression != nil {
			write, err := p.newStorageWrite(req.Progression, storagePersonalizerKeyProgression)
			if err != nil {
				logger.WithField("error", err.Error()).Error("Error creating progression storage object.")
				return "", ErrInternal
			}

			writes = append(writes, write)
		}

		if req.Stats != nil {
			write, err := p.newStorageWrite(req.Stats, storagePersonalizerKeyStats)
			if err != nil {
				logger.WithField("error", err.Error()).Error("Error creating stats storage object.")
				return "", ErrInternal
			}

			writes = append(writes, write)
		}

		if req.Teams != nil {
			write, err := p.newStorageWrite(req.Teams, storagePersonalizerKeyTeams)
			if err != nil {
				logger.WithField("error", err.Error()).Error("Error creating teams storage object.")
				return "", ErrInternal
			}

			writes = append(writes, write)
		}

		if req.Tutorials != nil {
			write, err := p.newStorageWrite(req.Tutorials, storagePersonalizerKeyTutorials)
			if err != nil {
				logger.WithField("error", err.Error()).Error("Error creating tutorials storage object.")
				return "", ErrInternal
			}

			writes = append(writes, write)
		}

		if req.Unlockables != nil {
			write, err := p.newStorageWrite(req.Unlockables, storagePersonalizerKeyUnlockables)
			if err != nil {
				logger.WithField("error", err.Error()).Error("Error creating unlockables storage object.")
				return "", ErrInternal
			}

			writes = append(writes, write)
		}

		if req.Base != nil {
			write, err := p.newStorageWrite(req.Base, storagePersonalizerKeyBase)
			if err != nil {
				logger.WithField("error", err.Error()).Error("Error creating base storage object.")
				return "", ErrInternal
			}

			writes = append(writes, write)
		}

		_, err = nk.StorageWrite(ctx, writes)
		if err != nil {
			return "", err
		}

		return "{}", nil
	}
}

func (p *StoragePersonalizer) GetValue(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, system System, userID string) (any, error) {
	now := time.Now().UTC()
	systemType := system.GetType()

	p.RLock()
	cached, found := p.cache[systemType]
	p.RUnlock()

	if !found || now.After(cached.expiryTime) {
		var readOp *runtime.StorageRead
		switch systemType {
		case SystemTypeAchievements:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: storagePersonalizerKeyAchievements}
		case SystemTypeEconomy:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: storagePersonalizerKeyEconomy}
		case SystemTypeEnergy:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: storagePersonalizerKeyEnergy}
		case SystemTypeEventLeaderboards:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: storagePersonalizerKeyEventLeaderboards}
		case SystemTypeIncentives:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: storagePersonalizerKeyIncentives}
		case SystemTypeLeaderboards:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: storagePersonalizerKeyLeaderboards}
		case SystemTypeProgression:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: storagePersonalizerKeyProgression}
		case SystemTypeStats:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: storagePersonalizerKeyStats}
		case SystemTypeTeams:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: storagePersonalizerKeyTeams}
		case SystemTypeTutorials:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: storagePersonalizerKeyTutorials}
		case SystemTypeUnlockables:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: storagePersonalizerKeyUnlockables}
		case SystemTypeBase:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: storagePersonalizerKeyBase}
		default:
			return nil, runtime.NewError("hiro system type unknown", 3)
		}

		objects, err := nk.StorageRead(ctx, []*runtime.StorageRead{readOp})
		if err != nil {
			logger.WithField("error", err.Error()).Error("nk.StorageRead error")
			return nil, err
		}
		cached = &StoragePersonalizerCachedStorageObject{
			refreshTime: now,
			expiryTime:  now.Add(p.cacheExpiry),
		}
		if len(objects) > 0 {
			cached.object = objects[0]
		}
		found = true
		p.Lock()
		p.cache[systemType] = cached
		p.Unlock()
	}

	if !found || cached.object == nil {
		// No personalization found for this system type.
		return nil, nil
	}

	config := system.GetConfig()
	decoder := json.NewDecoder(strings.NewReader(cached.object.Value))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(config); err != nil {
		logger.WithField("userID", userID).WithField("error", err.Error()).Error("error merging storage value")
		return nil, err
	}

	return config, nil
}

func (p *StoragePersonalizer) newStorageWrite(config any, storageKey string) (*runtime.StorageWrite, error) {
	json, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	return &runtime.StorageWrite{
		Collection:      p.collection,
		Key:             storageKey,
		Value:           string(json),
		PermissionRead:  0,
		PermissionWrite: 0,
	}, nil
}
