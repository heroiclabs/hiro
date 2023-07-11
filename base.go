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
	"errors"
	"github.com/heroiclabs/nakama-common/runtime"
	"google.golang.org/protobuf/encoding/protojson"
	"plugin"
)

// The BaseSystem provides various small features which aren't large enough to be in their own gameplay systems.
type BaseSystem interface {
	System

	// RateApp uses the SMTP configuration to receive feedback from players via email.
	RateApp(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, username string, score uint32, message string) error

	// SetDevicePrefs sets push notification tokens on a user's account so push messages can be received.
	SetDevicePrefs(ctx context.Context, logger runtime.Logger, db *sql.DB, userID, deviceID, pushTokenAndroid, pushTokenIos string, preferences map[string]bool) error
}

// BaseSystemConfig is the data definition for the BaseSystem type.
type BaseSystemConfig struct {
	RateAppSmtpAddr          string `json:"rate_app_smtp_addr"`            // "smtp.gmail.com"
	RateAppSmtpUsername      string `json:"rate_app_smtp_username"`        // "email@domain"
	RateAppSmtpPassword      string `json:"rate_app_smtp_password"`        // "password"
	RateAppSmtpEmailFrom     string `json:"rate_app_smtp_email_from"`      // "gamename-server@mmygamecompany.com"
	RateAppSmtpEmailFromName string `json:"rate_app_smtp_email_from_name"` // My Game Company
	RateAppSmtpEmailSubject  string `json:"rate_app_smtp_email_subject"`   // "RateApp Feedback"
	RateAppSmtpEmailTo       string `json:"rate_app_smtp_email_to"`        // "gamename-rateapp@mygamecompany.com"
	RateAppSmtpPort          int    `json:"rate_app_smtp_port"`            // 587

	RateAppTemplate string `json:"rate_app_template"` // HTML email template
}

// Hiro provides a type which combines all gameplay systems.
type Hiro interface {
	SetPersonalizer(Personalizer)

	GetAchievementsSystem() AchievementsSystem
	GetBaseSystem() BaseSystem
	GetEconomySystem() EconomySystem
	GetEnergySystem() EnergySystem
	GetInventorySystem() InventorySystem
	GetLeaderboardsSystem() LeaderboardsSystem
	GetStatsSystem() StatsSystem
	GetTeamsSystem() TeamsSystem
	GetTutorialsSystem() TutorialsSystem
	GetUnlockablesSystem() UnlockablesSystem
}

// The SystemType identifies each of the gameplay systems.
type SystemType uint

const (
	SystemTypeUnknown SystemType = iota
	SystemTypeBase
	SystemTypeEnergy
	SystemTypeUnlockables
	SystemTypeTutorials
	SystemTypeLeaderboards
	SystemTypeStats
	SystemTypeTeams
	SystemTypeInventory
	SystemTypeAchievements
	SystemTypeEconomy
)

// Init initializes a Hiro type with the configurations provided.
func Init(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, initializer runtime.Initializer, binPath string, configs ...SystemConfig) (Hiro, error) {
	// Open the plugin.
	binFile, err := nk.ReadFile(binPath)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer binFile.Close()

	p, err := plugin.Open(binFile.Name())
	if err != nil {
		return nil, err
	}

	// Look up the required initialisation function.
	f, err := p.Lookup("Init")
	if err != nil {
		return nil, err
	}

	// Ensure the function has the correct types.
	fn, ok := f.(func(context.Context, runtime.Logger, runtime.NakamaModule, runtime.Initializer, *protojson.MarshalOptions, *protojson.UnmarshalOptions, ...SystemConfig) (Hiro, error))
	if !ok {
		return nil, errors.New("error reading hiro-gdk.Init function in Go module")
	}

	marshaler := &protojson.MarshalOptions{
		UseEnumNumbers:  true,
		UseProtoNames:   true,
		EmitUnpopulated: false,
	}
	unmarshaler := &protojson.UnmarshalOptions{DiscardUnknown: false}

	return fn(ctx, logger, nk, initializer, marshaler, unmarshaler, configs...)
}

// The SystemConfig describes the configuration that each gameplay system must use to configure itself.
type SystemConfig interface {
	// GetType returns the runtime type of the gameplay system.
	GetType() SystemType

	// GetConfigFile returns the configuration file used for the data definitions in the gameplay system.
	GetConfigFile() string

	// GetRegister returns true if the gameplay system's RPCs should be registered with the game server.
	GetRegister() bool

	// GetExtra returns the extra parameter used to configure the gameplay system.
	GetExtra() any
}

var _ SystemConfig = &systemConfig{}

type systemConfig struct {
	systemType SystemType
	configFile string
	register   bool

	extra any
}

func (sc *systemConfig) GetType() SystemType {
	return sc.systemType
}
func (sc *systemConfig) GetConfigFile() string {
	return sc.configFile
}
func (sc *systemConfig) GetRegister() bool {
	return sc.register
}
func (sc *systemConfig) GetExtra() any {
	return sc.extra
}

// OnReward is a function which can be used by each gameplay system to provide an override reward.
type OnReward[T any] func(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string, source T, rewardConfig *EconomyConfigReward, reward *Reward) (*Reward, error)

// A System is a base type for a gameplay system.
type System interface {
	// GetType provides the runtime type of the gameplay system.
	GetType() SystemType

	// GetConfig returns the configuration type of the gameplay system.
	GetConfig() any
}

// UsernameOverrideFn can be used to provide a different username generation strategy from the default in Nakama server.
type UsernameOverrideFn func() string

// WithAchievementsSystem configures an AchievementsSystem type and optionally registers it's RPCs with the game server.
func WithAchievementsSystem(configFile string, register bool) SystemConfig {
	return &systemConfig{
		systemType: SystemTypeAchievements,
		configFile: configFile,
		register:   register,
	}
}

// WithBaseSystem configures a BaseSystem type and optionally registers it's RPCs with the game server.
func WithBaseSystem(configFile string, register bool, usernameOverride ...UsernameOverrideFn) SystemConfig {
	return &systemConfig{
		systemType: SystemTypeBase,
		configFile: configFile,
		register:   register,

		extra: usernameOverride,
	}
}

// WithEconomySystem configures an EconomySystem type and optionally registers it's RPCs with the game server.
func WithEconomySystem(configFile string, register bool, ironSrcPrivKey ...string) SystemConfig {
	return &systemConfig{
		systemType: SystemTypeEconomy,
		configFile: configFile,
		register:   register,

		extra: ironSrcPrivKey,
	}
}

// WithEnergySystem configures an EnergySystem type and optionally registers it's RPCs with the game server.
func WithEnergySystem(configFile string, register bool) SystemConfig {
	return &systemConfig{
		systemType: SystemTypeEnergy,
		configFile: configFile,
		register:   register,
	}
}

// WithInventorySystem configures an InventorySystem type and optionally registers it's RPCs with the game server.
func WithInventorySystem(configFile string, register bool) SystemConfig {
	return &systemConfig{
		systemType: SystemTypeInventory,
		configFile: configFile,
		register:   register,
	}
}

// WithLeaderboardsSystem configures a LeaderboardsSystem type.
func WithLeaderboardsSystem(validateWriteScore ...ValidateWriteScoreFn) SystemConfig {
	return &systemConfig{
		systemType: SystemTypeLeaderboards,

		extra: validateWriteScore,
	}
}

// WithStatsSystem configures a StatsSystem type and optionally registers it's RPCs with the game server.
func WithStatsSystem(register bool) SystemConfig {
	return &systemConfig{
		systemType: SystemTypeStats,
		register:   register,
	}
}

// WithTeamsSystem configures a TeamsSystem type and optionally registers it's RPCs with the game server.
func WithTeamsSystem(configFile string, register bool, validateCreateTeam ...ValidateCreateTeamFn) SystemConfig {
	return &systemConfig{
		systemType: SystemTypeTeams,
		configFile: configFile,
		register:   register,

		extra: validateCreateTeam,
	}
}

// WithTutorialsSystem configures a TutorialsSystem type and optionally registers it's RPCs with the game server.
func WithTutorialsSystem(configFile string, register bool) SystemConfig {
	return &systemConfig{
		systemType: SystemTypeTutorials,
		configFile: configFile,
		register:   register,
	}
}

// WithUnlockablesSystem configures an UnlockablesSystem type and optionally registers it's RPCs with the game server.
func WithUnlockablesSystem(configFile string, register bool) SystemConfig {
	return &systemConfig{
		systemType: SystemTypeUnlockables,
		configFile: configFile,
		register:   register,
	}
}
