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
	"github.com/heroiclabs/nakama-common/runtime"
)

// TeamsConfig is the data definition for a TeamsSystem type.
type TeamsConfig struct {
	MaxTeamSize int `json:"max_team_size,omitempty"`
}

// A TeamsSystem is a gameplay system which wraps the groups system in Nakama server.
type TeamsSystem interface {
	System

	// Create makes a new team (i.e. Nakama group) with additional metadata which configures the team.
	Create(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, req *TeamCreateRequest) (*Team, error)

	// List will return a list of teams which the user can join.
	List(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, req *TeamListRequest) (*TeamList, error)

	// Search
	Search(ctx context.Context, db *sql.DB, logger runtime.Logger, nk runtime.NakamaModule, req *TeamSearchRequest) (*TeamList, error)

	// WriteChatMessage sends a message to the user's team even when they're not connected on a realtime socket.
	WriteChatMessage(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, req *TeamWriteChatMessageRequest) (*ChannelMessageAck, error)
}

// ValidateCreateTeamFn allows custom rules or velocity checks to be added as a precondition on whether a team is
// created or not.
type ValidateCreateTeamFn func(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, req *TeamCreateRequest) *runtime.Error
