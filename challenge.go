package hiro

import (
	"context"
	"github.com/heroiclabs/nakama-common/runtime"
)

type ChallengesConfig struct {
	Challenges map[string]*ChallengesConfigChallenge `json:"challenges,omitempty"`
}

type ChallengesConfigChallenge struct {
	RewardTiers          []*ChallengesConfigChallengeRewardTier `json:"reward_tiers,omitempty"`
	AdditionalProperties map[string]string                      `json:"additional_properties,omitempty"`
	MaxNumScore          int64                                  `json:"max_num_score,omitempty"`
	StartDelayMaxSec     int64                                  `json:"start_delay_max_sec,omitempty"`
	Open                 bool                                   `json:"open,omitempty"`
	Duration             *ChallengeDuration                     `json:"duration,omitempty"`
	Players              *ChallengesConfigPlayers               `json:"players,omitempty"`
}

type ChallengesConfigChallengeRewardTier struct {
	RankMax int64                `json:"rank_max,omitempty"`
	RankMin int64                `json:"rank_min,omitempty"`
	Reward  *EconomyConfigReward `json:"reward,omitempty"`
}

type ChallengeDuration struct {
	MinSec int64 `json:"min_sec,omitempty"`
	MaxSec int64 `json:"max_sec,omitempty"`
}

type ChallengesConfigPlayers struct {
	Min int64 `json:"min,omitempty"`
	Max int64 `json:"max,omitempty"`
}

type ChallengeSystem interface {
	System

	// GetTemplates lists all available auction configurations that can be used to create auction listings.
	GetTemplates(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string) (*ChallengeTemplates, error)

	GetChallenge(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, challengeId string) (*Challenge, error)

	ListChallenges(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string) ([]*Challenge, error)

	CreateChallenge(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID, name, description string, startTimeSec, durationSec int64, invitees []string, ascending, open bool, maxParticipants, maxScores int64) (*Challenge, error)

	JoinChallenge(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userId, challengeId string) (*Challenge, error)

	LeaveChallenge(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userId, challengeId string) (*Challenge, error)

	SubmitChallengeScore(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userId, username, displayName, avatarUrl, challengeId string, score, subscore int64, metadata map[string]any) (challenge *Challenge, err error)

	ClaimChallenge(ctx context.Context, userId, challengeId string) (*Challenge, error)
}
