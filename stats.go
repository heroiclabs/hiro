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

// StatsConfig is the data definition for a StatsSystem type.
type StatsConfig struct {
	Whitelist []string `json:"whitelist"`
}

type StatsSystem interface {
	System

	// List all private stats for one or more users.
	List(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userIDs []string) (map[string]*StatList, error)

	// Update private stats for a particular user.
	Update(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, publicStats []*StatUpdate, privateStats []*StatUpdate) (*StatList, error)
}
