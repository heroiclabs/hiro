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
	"sync"
	"time"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

const StoragePersonalizerCollectionDefault = "hiro_datadefinitions"

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

func NewStoragePersonalizerDefault() *StoragePersonalizer {
	return &StoragePersonalizer{
		cache:       make(map[SystemType]*StoragePersonalizerCachedStorageObject, 20),
		cacheExpiry: 10 * time.Minute,
		collection:  StoragePersonalizerCollectionDefault,
	}
}

func NewStoragePersonalizer(logger runtime.Logger, cacheExpirySec int, collection string) *StoragePersonalizer {
	return &StoragePersonalizer{
		cache:       make(map[SystemType]*StoragePersonalizerCachedStorageObject, 20),
		cacheExpiry: time.Duration(cacheExpirySec) * time.Second,
		collection:  collection,
		logger:      logger,
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
			readOp = &runtime.StorageRead{Collection: p.collection, Key: "achievements"}
		case SystemTypeEconomy:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: "economy"}
		case SystemTypeEnergy:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: "energy"}
		case SystemTypeEventLeaderboards:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: "event_leaderboards"}
		case SystemTypeIncentives:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: "incentives"}
		case SystemTypeLeaderboards:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: "leaderboards"}
		case SystemTypeProgression:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: "progression"}
		case SystemTypeStats:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: "stats"}
		case SystemTypeTeams:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: "teams"}
		case SystemTypeTutorials:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: "tutorials"}
		case SystemTypeUnlockables:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: "unlockables"}
		case SystemTypeBase:
			readOp = &runtime.StorageRead{Collection: p.collection, Key: "base"}
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
