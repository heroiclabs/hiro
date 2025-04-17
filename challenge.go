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
	Ascending            bool                                   `json:"ascending,omitempty"`
	Operator             string                                 `json:"operator,omitempty"`
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

	// GetTemplates lists all available challenge configurations that can be used to create new challenges.
	GetTemplates(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userID string) (*ChallengeTemplates, error)

	// GetChallenge returns a challenge the user has been invited to or which is participating in.
	GetChallenge(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, challengeId, userId string) (*Challenge, error)

	// ListChallenges Lists all the user's pending or joined challenges.
	ListChallenges(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userId string, categories []string) ([]*Challenge, error)

	// CreateChallenge a new challenge for a list of users.
	CreateChallenge(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userId, templateId, name, description, category string, startDelaySec, durationSec int64, invitees []string, maxPlayers int64) (*Challenge, error)

	// JoinChallenge Joins a challenge the user's been invited to.
	JoinChallenge(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userId, challengeId string) (*Challenge, error)

	// LeaveChallenge rejects a challenge invitation or abandons a joined challenge.
	LeaveChallenge(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userId, challengeId string) (*Challenge, error)

	// SubmitChallengeScore submits a new score to the challenge.
	SubmitChallengeScore(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userId, challengeId string, score, subscore int64, metadata map[string]any) (challenge *Challenge, err error)

	// ClaimChallenge claims a reward of a challenge, if any.
	ClaimChallenge(ctx context.Context, logger runtime.Logger, nk runtime.NakamaModule, userId, challengeId string) (*Challenge, error)

	// SetOnChallengeReward sets a custom reward function which will run after a challenge reward has been claimed.
	SetOnChallengeReward(fn OnReward[*Challenge])
}
