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

var (
	ErrTeamNotFound        = runtime.NewError("team not found", 3)         // INVALID_ARGUMENT
	ErrTeamMaxSizeExceeded = runtime.NewError("team max size exceeded", 3) // INVALID_ARGUMENT
)

// TeamsConfig is the data definition for a TeamsSystem type.
type TeamsConfig struct {
	InitialMaxTeamSize int `json:"initial_max_team_size,omitempty"`
	MaxTeamSize        int `json:"max_team_size,omitempty"`

	Wallet            *TeamsWalletConfig           `json:"wallet,omitempty"`
	Stats             *TeamsStatsConfig            `json:"stats,omitempty"`
	Inventory         *TeamsInventoryConfig        `json:"inventory,omitempty"`
	Achievements      *TeamsAchievementsConfig     `json:"achievements,omitempty"`
	EventLeaderboards *TeamEventLeaderboardsConfig `json:"event_leaderboards,omitempty"`
}

type TeamsWalletConfig struct {
	Currencies map[string]int64 `json:"currencies,omitempty"`
}

type TeamsStatsConfig struct {
	StatsPublic  map[string]*StatsConfigStat `json:"stats_public,omitempty"`
	StatsPrivate map[string]*StatsConfigStat `json:"stats_private,omitempty"`
}

type TeamsInventoryConfig struct {
	Items    map[string]*InventoryConfigItem `json:"items,omitempty"`
	Limits   *InventoryConfigLimits          `json:"limits,omitempty"`
	ItemSets map[string]map[string]bool      `json:"-"` // Auto-computed when the config is read or personalized.

	ConfigSource ConfigSource[*InventoryConfigItem] `json:"-"` // Not included in serialization, set dynamically.
}

type TeamsAchievementsConfig struct {
	Achievements map[string]*AchievementsConfigAchievement `json:"achievements,omitempty"`
}

type TeamEventLeaderboardsConfig struct {
	EventLeaderboards map[string]*EventLeaderboardsConfigLeaderboard `json:"event_leaderboards,omitempty"`
}

// A TeamsSystem is a gameplay system which wraps the groups system in Nakama server.
type TeamsSystem interface {
	System

	// Create makes a new team (i.e. Nakama group) with additional metadata which configures the team.
	Create(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, req *TeamCreateRequest) (team *Team, err error)

	// List will return a list of teams which the user can join.
	List(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, req *TeamListRequest) (teams *TeamList, err error)

	// Search for teams based on given criteria.
	Search(ctx context.Context, db *sql.DB, logger runtime.Logger, nk runtime.NakamaModule, req *TeamSearchRequest) (teams *TeamList, err error)

	// WriteChatMessage sends a message to the user's team even when they're not connected on a realtime socket.
	WriteChatMessage(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, req *TeamWriteChatMessageRequest) (resp *ChannelMessageAck, err error)

	// UpdateMaxSize sets a new maximum team size for the selected team.
	UpdateMaxSize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, userID, teamID string, count int, delta bool) error

	// WalletGet fetches the wallet for a specified team.
	WalletGet(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID string) (*TeamWallet, error)

	// WalletGrant grants currencies to the specified team's wallet.
	WalletGrant(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID string, currencies map[string]int64) (*TeamWallet, error)

	// InventoryList will return the items defined as well as the computed item sets for the team by ID.
	InventoryList(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID, category string) (items map[string]*InventoryConfigItem, itemSets map[string][]string, err error)

	// InventoryListInventoryItems will return the items which are part of a team's inventory by ID.
	InventoryListInventoryItems(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID, category string) (inventory *Inventory, err error)

	// InventoryConsumeItems will deduct the item(s) from the team's inventory and run the consume reward for each one, if defined.
	InventoryConsumeItems(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID string, itemIDs, instanceIDs map[string]int64, overConsume bool) (updatedInventory *Inventory, rewards map[string][]*Reward, instanceRewards map[string][]*Reward, err error)

	// InventoryGrantItems will add the item(s) to a team's inventory by ID.
	InventoryGrantItems(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID string, itemIDs map[string]int64, ignoreLimits bool) (updatedInventory *Inventory, newItems map[string]*InventoryItem, updatedItems map[string]*InventoryItem, notGrantedItemIDs map[string]int64, err error)

	// InventoryUpdateItems will update the properties which are stored on each item by instance ID for a team.
	InventoryUpdateItems(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID string, instanceIDs map[string]*InventoryUpdateItemProperties) (updatedInventory *Inventory, err error)

	// SetOnInventoryConsumeReward sets a custom reward function which will run after a team inventory item consume reward is rolled.
	SetOnInventoryConsumeReward(fn OnReward[*InventoryConfigItem])

	// SetInventoryConfigSource sets a custom additional config lookup function.
	SetInventoryConfigSource(fn ConfigSource[*InventoryConfigItem])

	// ClaimAchievements when one or more achievements whose progress has completed by their IDs.
	ClaimAchievements(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID string, achievementIDs []string, claimTotal bool) (achievements map[string]*Achievement, repeatAchievements map[string]*Achievement, err error)

	// GetAchievements returns all achievements available to the user and progress on them.
	GetAchievements(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID string) (achievements map[string]*Achievement, repeatAchievements map[string]*Achievement, err error)

	// UpdateAchievements updates progress on one or more achievements by the same amount.
	UpdateAchievements(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID string, achievementUpdates map[string]int64) (achievements map[string]*Achievement, repeatAchievements map[string]*Achievement, err error)

	// SetOnAchievementReward sets a custom reward function which will run after an achievement's reward is rolled.
	SetOnAchievementReward(fn OnReward[*AchievementsConfigAchievement])

	// SetOnSubAchievementReward sets a custom reward function which will run after a sub-achievement's reward is rolled.
	SetOnSubAchievementReward(fn OnReward[*AchievementsConfigSubAchievement])

	// SetOnAchievementTotalReward sets a custom reward function which will run after an achievement's total reward is rolled.
	SetOnAchievementTotalReward(fn OnReward[*AchievementsConfigAchievement])

	// ListEventLeaderboard returns available event leaderboards for the team.
	ListEventLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID string, withScores bool, categories []string) (eventLeaderboards []*EventLeaderboard, err error)

	// GetEventLeaderboard returns a specified event leaderboard's cohort for the team.
	GetEventLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID, eventLeaderboardID string) (eventLeaderboard *EventLeaderboard, err error)

	// RollEventLeaderboard places the team into a new cohort for the specified event leaderboard if possible.
	RollEventLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID, eventLeaderboardID string, tier *int, matchmakerProperties map[string]interface{}) (eventLeaderboard *EventLeaderboard, err error)

	// UpdateEventLeaderboard updates the team's score in the specified event leaderboard, and returns the team's updated cohort information.
	UpdateEventLeaderboard(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, userID, teamID, username, eventLeaderboardID string, score, subscore int64, metadata map[string]interface{}, alwaysUpdateMetadata bool) (eventLeaderboard *EventLeaderboard, err error)

	// ClaimEventLeaderboard claims the team's reward for the given event leaderboard.
	ClaimEventLeaderboard(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID, eventLeaderboardID string) (eventLeaderboard *EventLeaderboard, err error)

	// SetOnEventLeaderboardsReward sets a custom reward function which will run after a team event leaderboard's reward is rolled.
	SetOnEventLeaderboardsReward(fn OnReward[*EventLeaderboardsConfigLeaderboard])

	// SetOnEventLeaderboardCohortSelection sets a custom function that can replace the cohort or opponent selection feature of team event leaderboards.
	SetOnEventLeaderboardCohortSelection(fn OnTeamEventLeaderboardCohortSelection)

	// DebugFill fills the user's current cohort with dummy teams for all remaining available slots.
	DebugFill(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID, eventLeaderboardID string, targetCount int) (eventLeaderboard *EventLeaderboard, err error)

	// DebugRandomScores assigns random scores to the participants of the team's current cohort, except to the team itself.
	DebugRandomScores(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, teamID, eventLeaderboardID string, scoreMin, scoreMax, subscoreMin, subscoreMax int64, operator *int) (eventLeaderboard *EventLeaderboard, err error)
}

// ValidateCreateTeamFn allows custom rules or velocity checks to be added as a precondition on whether a team is created or not.
type ValidateCreateTeamFn func(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, req *TeamCreateRequest) *runtime.Error

type OnTeamEventLeaderboardCohortSelection func(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, storageIndex string, eventID string, config *EventLeaderboardsConfigLeaderboard, userID, teamID string, tier int, matchmakerProperties map[string]interface{}) (cohortID string, cohortUserIDs []string, newCohort *EventLeaderboardCohortConfig, err error)
