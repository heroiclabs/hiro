// Copyright 2026 Heroic Labs & Contributors
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
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strconv"
)

// source of obtaining/spending an item, currency, energy, or modifier.
const (
	EventSourceDonationClaimed           = "donationClaimed"
	EventSourceDonationGiven             = "donationGiven"
	EventSourceDonationRequested         = "donationRequested"
	EventSourceEconomyGranted            = "economyGranted"
	EventSourceRewardGranted             = "rewardGranted"
	EventSourceItemPurchased             = "itemPurchased"
	EventSourcePlacementSucceeded        = "placementSucceeded"
	EventSourceAchievementUpdated        = "achievementUpdated"
	EventSourceAchievementClaimed        = "achievementClaimed"
	EventSourceItemsConsumed             = "itemsConsumed"
	EventSourceItemsGranted              = "itemsGranted"
	EventSourceEnergySpent               = "energySpent"
	EventSourceIncentiveSenderClaimed    = "incentiveSenderClaimed"
	EventSourceIncentiveRecipientClaimed = "incentiveRecipientClaimed"
	EventSourceEventLeaderboardClaimed   = "eventLeaderboardClaimed"
	EventSourceProgressionPurchased      = "progressionPurchased"
	EventSourceUnlockableUnlockPurchased = "unlockableUnlockPurchased"
	EventSourceUnlockableSlotPurchased   = "unlockableSlotPurchased"
	EventSourceUnlockableClaimed         = "unlockableClaimed"
	EventSourceInitializeUser            = "initializeUser"
	EventSourceChallengeClaimed          = "challengeClaimed"
)

func newUUIDv4() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		panic(err)
	}

	// Set the appropriate version and variant
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10

	// Return the formatted UUID v4
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

// A donation was claimed.
func NewDonationClaimedEvent(system System, donationID string, donationConfig *EconomyConfigDonation, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "donationClaimed",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"donationId": donationID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: donationID,
		Source:   donationConfig,
	}
}

// A donation was given.
func NewDonationGivenEvent(system System, donationID string, donationConfig *EconomyConfigDonation, recipientID string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "donationGiven",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"donationId":  donationID,
			"recipientId": recipientID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: donationID,
		Source:   donationConfig,
	}
}

// A donation was requested.
func NewDonationRequestedEvent(system System, donationID string, donationConfig *EconomyConfigDonation, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "donationRequested",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"donationId": donationID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: donationID,
		Source:   donationConfig,
	}
}

// Currency was granted.
func NewCurrencyGrantedEvent(system System, sourceID string, sourceConfig any, currencyID string, amount int64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "currencyGranted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"currencyId": currencyID,
			"source":     source,
		},
		Value:     strconv.FormatInt(amount, 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

func NewItemGrantedEvent(system System, sourceID string, sourceConfig any, itemID string, amount int64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "itemsGranted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"itemId": itemID,
			"source": source,
		},
		Value:     strconv.FormatInt(amount, 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

func NewTeamItemGrantedEvent(system System, sourceID string, sourceConfig any, teamID, itemID string, amount int64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamItemsGranted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"teamId": teamID,
			"itemId": itemID,
			"source": source,
		},
		Value:     strconv.FormatInt(amount, 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

// Energy was granted.
func NewEnergyGrantedEvent(system System, sourceID string, sourceConfig any, energyID string, amount int32, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "energyGranted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"energyId": energyID,
			"source":   source,
		},
		Timestamp: ts,
		Value:     strconv.FormatInt(int64(amount), 10),

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

// A reward modifier was granted.
func NewEnergyModifierGrantedEvent(system System, sourceID string, sourceConfig any, energyModifierID string, operator string, value int64, durationSec uint64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "energyModifierGranted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"energyModifierId": energyModifierID,
			"operator":         operator,
			"durationSec":      strconv.FormatUint(durationSec, 10),
			"source":           source,
		},
		Value:     strconv.FormatInt(value, 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

// A reward modifier was granted.
func NewRewardModifierGrantedEvent(system System, sourceID string, sourceConfig any, rewardModifierID string, modifierType string, operator string, value int64, durationSec uint64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "rewardModifierGranted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"rewardModifierId": rewardModifierID,
			"type":             modifierType,
			"operator":         operator,
			"durationSec":      strconv.FormatUint(durationSec, 10),
			"source":           source,
		},
		Value:     strconv.FormatInt(value, 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

// A currency was spent.
func NewCurrencySpentEvent(system System, sourceID string, sourceConfig any, currencyID string, amount int64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "currencySpent",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"currencyId": currencyID,
			"source":     source,
		},
		Value:     strconv.FormatInt(int64(amount), 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

// An item was spent.
func NewItemSpentEvent(system System, sourceID string, sourceConfig any, itemID string, amount int64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "itemSpent",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"itemId": itemID,
			"source": source,
		},
		Value:     strconv.FormatInt(int64(amount), 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

// Energy was spent.
func NewEnergySpentEvent(system System, energyID string, energyConfig *EnergyConfigEnergy, amount int32, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "energySpent",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"energyId": energyID,
		},
		Value:     strconv.FormatInt(int64(amount), 10),
		Timestamp: ts,

		System:   system,
		SourceId: energyID,
		Source:   energyConfig,
	}
}

// A purchase intent was placed.
func NewPurchaseIntentEvent(system System, storeItemID string, storeItem *EconomyConfigStoreItem, storeType EconomyStoreType, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "purchaseIntent",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"storeItemId": storeItemID,
			"storeType":   storeType.String(),
		},
		Timestamp: ts,

		System:   system,
		SourceId: storeItemID,
		Source:   storeItem,
	}
}

// A SKU purchase was completed.
func NewPurchaseCompletedEvent(system System, storeItemID string, test bool, storeItem *EconomyConfigStoreItem, currency string, amount float64, storeType EconomyStoreType, amountUSDCents int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "purchaseCompleted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"id":        storeItemID, // keep named as "id" rather than "storeItemId" for backwards compatibility.
			"test":      strconv.FormatBool(test),
			"currency":  currency,
			"amount":    strconv.FormatFloat(amount, 'f', 2, 64),
			"storeType": storeType.String(),
		},
		Value:     strconv.FormatInt(amountUSDCents, 10),
		Timestamp: ts,

		System:   system,
		SourceId: storeItemID,
		Source:   storeItem,
	}
}

// An ad placement started.
func NewAdPlacementStartedEvent(system System, placementID string, placementConfig *EconomyConfigPlacement, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "adPlacementStarted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"placementId": placementID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: placementID,
		Source:   placementConfig,
	}
}

// An ad placement succeeded.
func NewAdPlacementSucceededEvent(system System, placementID string, placementConfig *EconomyConfigPlacement, maxRetries bool, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "adPlacementSucceeded",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"placementId": placementID,
			"maxRetries":  strconv.FormatBool(maxRetries),
		},
		Timestamp: ts,

		System:   system,
		SourceId: placementID,
		Source:   placementConfig,
	}
}

// An ad placement failed.
func NewAdPlacementFailedEvent(system System, placementID string, placementConfig *EconomyConfigPlacement, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "adPlacementFailed",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"placementId": placementID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: placementID,
		Source:   placementConfig,
	}
}

// An achievement was updated.
func NewAchievementUpdatedEvent(system System, achievementID string, achievementConfig any, count int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "achievementUpdated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"achievementId": achievementID,
		},
		Value:     strconv.FormatInt(count, 10),
		Timestamp: ts,

		System:   system,
		SourceId: achievementID,
		Source:   achievementConfig,
	}
}

// A team achievement was updated.
func NewTeamAchievementUpdatedEvent(system System, teamID, achievementID string, achievementConfig any, count int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "achievementUpdated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"teamId":        teamID,
			"achievementId": achievementID,
		},
		Value:     strconv.FormatInt(count, 10),
		Timestamp: ts,

		System:   system,
		SourceId: achievementID,
		Source:   achievementConfig,
	}
}

// An achievement was claimed.
func NewAchievementClaimedEvent(system System, achievementID string, achievementConfig *AchievementsConfigAchievement, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "achievementClaimed",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"achievementId": achievementID,
		},
		Timestamp: ts,
		// No value.

		System:   system,
		SourceId: achievementID,
		Source:   achievementConfig,
	}
}

// A team achievement was claimed.
func NewTeamAchievementClaimedEvent(system System, teamID, achievementID string, achievementConfig *AchievementsConfigAchievement, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamAchievementClaimed",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"teamId":        teamID,
			"achievementId": achievementID,
		},
		Timestamp: ts,
		// No value.

		System:   system,
		SourceId: achievementID,
		Source:   achievementConfig,
	}
}

// A progression was purchased.
func NewProgressionPurchasedEvent(system System, progressionID string, progressionConfig *ProgressionConfigProgression, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "progressionPurchased",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"progressionId": progressionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: progressionID,
		Source:   progressionConfig,
	}
}

// A progression was updated.
func NewProgressionUpdatedEvent(system System, progressionID string, progressionConfig *ProgressionConfigProgression, countID string, count int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "progressionUpdated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"progressionId": progressionID,
			"countId":       countID,
		},
		Value:     strconv.FormatInt(count, 10),
		Timestamp: ts,

		System:   system,
		SourceId: progressionID,
		Source:   progressionConfig,
	}
}

// A progression was reset.
func NewProgressionResetEvent(system System, progressionID string, progressionConfig *ProgressionConfigProgression, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "progressionReset",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"progressionId": progressionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: progressionID,
		Source:   progressionConfig,
	}
}

// Inventory items were consumed.
func NewItemsConsumedEvent(system System, itemID string, itemConfig *InventoryConfigItem, amount int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "itemsConsumed",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"itemId": itemID,
		},
		Value:     strconv.FormatInt(amount, 10),
		Timestamp: ts,

		System:   system,
		SourceId: itemID,
		Source:   itemConfig,
	}
}

// Team inventory items were consumed.
func NewTeamItemsConsumedEvent(system System, teamID, itemID string, itemConfig *InventoryConfigItem, amount int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamItemsConsumed",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"itemId": itemID,
			"teamId": teamID,
		},
		Value:     strconv.FormatInt(amount, 10),
		Timestamp: ts,

		System:   system,
		SourceId: itemID,
		Source:   itemConfig,
	}
}

// An inventory item was updated.
func NewItemUpdatedEvent(system System, itemID string, itemConfig *InventoryConfigItem, stringProperties map[string]string, numericProperties map[string]float64, ts int64) (*PublisherEvent, error) {
	stringPropertiesJson, err := json.Marshal(stringProperties)
	if err != nil {
		return nil, err
	}

	numericPropertiesJson, err := json.Marshal(numericProperties)
	if err != nil {
		return nil, err
	}

	return &PublisherEvent{
		Name: "itemUpdated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"itemId":            itemID,
			"stringProperties":  string(stringPropertiesJson),
			"numericProperties": string(numericPropertiesJson),
		},
		Timestamp: ts,

		System:   system,
		SourceId: itemID,
		Source:   itemConfig,
	}, nil
}

func NewTeamItemUpdatedEvent(system System, teamID, itemID string, itemConfig *InventoryConfigItem, stringProperties map[string]string, numericProperties map[string]float64, ts int64) (*PublisherEvent, error) {
	stringPropertiesJson, err := json.Marshal(stringProperties)
	if err != nil {
		return nil, err
	}

	numericPropertiesJson, err := json.Marshal(numericProperties)
	if err != nil {
		return nil, err
	}

	return &PublisherEvent{
		Name: "teamItemUpdated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"teamId":            teamID,
			"itemId":            itemID,
			"stringProperties":  string(stringPropertiesJson),
			"numericProperties": string(numericPropertiesJson),
		},
		Timestamp: ts,

		System:   system,
		SourceId: itemID,
		Source:   itemConfig,
	}, nil
}

// Stats were updated.
func NewStatUpdatedEvent(system System, name string, stat any, operator StatUpdateOperator, value int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "statUpdated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"name":     name,
			"operator": operator.String(),
		},
		Value:     strconv.FormatInt(value, 10),
		Timestamp: ts,

		System:   system,
		SourceId: name,
		Source:   stat,
	}
}

// Team stats were updated.
func NewTeamStatUpdatedEvent(system System, name string, stat any, operator StatUpdateOperator, value int64, ts int64, teamID string) *PublisherEvent {
	return &PublisherEvent{
		Name: "statUpdated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"name":     name,
			"operator": operator.String(),
			"teamId":   teamID,
		},
		Value:     strconv.FormatInt(value, 10),
		Timestamp: ts,

		System:   system,
		SourceId: name,
		Source:   stat,
	}
}

// A tutorial was accepted.
func NewTutorialAcceptedEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialAccepted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"tutorialId": tutorialID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// A tutorial was declined.
func NewTutorialDeclinedEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialDeclined",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"tutorialId": tutorialID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// A tutorial started.
func NewTutorialStartedEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, step int, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialStarted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"tutorialId": tutorialID,
		},
		Value:     strconv.FormatInt(int64(step), 10),
		Timestamp: ts,

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// A tutorial was abandoned.
func NewTutorialAbandonedEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, step int32, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialAbandoned",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"tutorialId": tutorialID,
		},
		Timestamp: ts,
		Value:     strconv.FormatInt(int64(step), 10),

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// A tutorial step was completed.
func NewTutorialStepCompletedEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, step int, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialStepCompleted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"tutorialId": tutorialID,
		},
		Value:     strconv.FormatInt(int64(step), 10),
		Timestamp: ts,

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// A tutorial was completed.
func NewTutorialCompletedEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, step int, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialCompleted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"tutorialId": tutorialID,
		},
		Value:     strconv.FormatInt(int64(step), 10),
		Timestamp: ts,

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// One or more tutorials were reset.
func NewTutorialResetEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialReset",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"tutorialId": tutorialID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// A team was created.
func NewTeamCreatedEvent(system System, teamID string, team *Team, open bool, maxCount int32, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamCreated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"teamId": teamID,
			"open":   strconv.FormatBool(open),
		},
		Timestamp: ts,
		Value:     strconv.FormatInt(int64(maxCount), 10),

		System:   system,
		SourceId: teamID,
		Source:   team,
	}
}

// An incentive was created.
func NewIncentiveCreatedEvent(system System, incentiveID string, incentiveConfig *IncentivesConfigIncentive, code string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "incentiveCreated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"incentiveId": incentiveID,
			"code":        code,
		},
		Timestamp: ts,

		System:   system,
		SourceId: incentiveID,
		Source:   incentiveConfig,
	}
}

// An incentive was deleted.
func NewIncentiveDeletedEvent(system System, incentiveID string, incentiveConfig *IncentivesConfigIncentive, code string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "incentiveDeleted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"incentiveId": incentiveID,
			"code":        code,
		},
		Timestamp: ts,

		System:   system,
		SourceId: incentiveID,
		Source:   incentiveConfig,
	}
}

// An incentive was claimed by the sender.
func NewIncentiveSenderClaimedEvent(system System, incentiveID string, incentiveConfig *IncentivesConfigIncentive, code string, claimantID string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "incentiveSenderClaimed",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"incentiveId": incentiveID,
			"code":        code,
			"claimantId":  claimantID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: incentiveID,
		Source:   incentiveConfig,
	}
}

// An incentive was claimed by the recipient.
func NewIncentiveRecipientClaimedEvent(system System, incentiveID string, incentiveConfig *IncentivesConfigIncentive, code string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "incentiveRecipientClaimed",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"incentiveId": incentiveID,
			"code":        code,
		},
		Timestamp: ts,

		System:   system,
		SourceId: incentiveID,
		Source:   incentiveConfig,
	}
}

// An event leaderboard was rolled.
func NewEventLeaderboardRolledEvent(system System, eventLeaderboardID string, eventLeaderboardConfig *EventLeaderboardsConfigLeaderboard, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "eventLeaderboardRolled",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"eventLeaderboardId": eventLeaderboardID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: eventLeaderboardID,
		Source:   eventLeaderboardConfig,
	}
}

// A team leaderboard was rolled.
func NewTeamEventLeaderboardRolledEvent(system System, teamID, eventLeaderboardID string, eventLeaderboardConfig *EventLeaderboardsConfigLeaderboard, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamEventLeaderboardRolled",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"teamId":             teamID,
			"eventLeaderboardId": eventLeaderboardID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: eventLeaderboardID,
		Source:   eventLeaderboardConfig,
	}
}

// An event leaderboard was updated.
func NewEventLeaderboardUpdatedEvent(system System, eventLeaderboardID string, eventLeaderboardConfig *EventLeaderboardsConfigLeaderboard, score int64, subscore int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "eventLeaderboardUpdated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"eventLeaderboardId": eventLeaderboardID,
			"subscore":           strconv.FormatInt(subscore, 10),
		},
		Timestamp: ts,
		Value:     strconv.FormatInt(score, 10),

		System:   system,
		SourceId: eventLeaderboardID,
		Source:   eventLeaderboardConfig,
	}
}

// A team event leaderboard was updated.
func NewTeamEventLeaderboardUpdatedEvent(system System, teamID, eventLeaderboardID string, eventLeaderboardConfig *EventLeaderboardsConfigLeaderboard, score int64, subscore int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamEventLeaderboardUpdated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"teamId":             teamID,
			"eventLeaderboardId": eventLeaderboardID,
			"subscore":           strconv.FormatInt(subscore, 10),
		},
		Timestamp: ts,
		Value:     strconv.FormatInt(score, 10),

		System:   system,
		SourceId: eventLeaderboardID,
		Source:   eventLeaderboardConfig,
	}
}

// An event leaderboard was claimed.
func NewEventLeaderboardClaimedEvent(system System, eventLeaderboardID string, eventLeaderboardConfig *EventLeaderboardsConfigLeaderboard, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "eventLeaderboardClaimed",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"eventLeaderboardId": eventLeaderboardID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: eventLeaderboardID,
		Source:   eventLeaderboardConfig,
	}
}

// A team event leaderboard was claimed.
func NewTeamEventLeaderboardClaimedEvent(system System, teamID, eventLeaderboardID string, eventLeaderboardConfig *EventLeaderboardsConfigLeaderboard, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamEventLeaderboardClaimed",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"teamId":             teamID,
			"eventLeaderboardId": eventLeaderboardID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: eventLeaderboardID,
		Source:   eventLeaderboardConfig,
	}
}

// An unlockable was created.
func NewUnlockableCreatedEvent(system System, unlockableID string, unlockableConfig *UnlockablesConfigUnlockable, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "unlockableCreated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"unlockableId": unlockableID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: unlockableID,
		Source:   unlockableConfig,
	}
}

// An unlockable's unlock was started.
func NewUnlockableUnlockStartedEvent(system System, unlockableID string, unlockableConfig *UnlockablesConfigUnlockable, instanceID string, activeUnlockables int, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "unlockableUnlockStarted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"unlockableId": unlockableID,
			"instanceId":   instanceID,
		},
		Value:     strconv.FormatInt(int64(activeUnlockables), 10),
		Timestamp: ts,

		System:   system,
		SourceId: unlockableID,
		Source:   unlockableConfig,
	}
}

// An unlockable's unlock was purchased.
func NewUnlockableUnlockPurchasedEvent(system System, unlockableID string, unlockableConfig *UnlockablesConfigUnlockable, instanceID string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "unlockableUnlockPurchased",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"unlockableId": unlockableID,
			"instanceId":   instanceID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: unlockableID,
		Source:   unlockableConfig,
	}
}

// An unlockable slot was purchased.
func NewUnlockableSlotPurchasedEvent(system System, config *UnlockablesConfig, activeSlots int, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name:      "unlockableSlotPurchased",
		Id:        newUUIDv4(),
		Value:     strconv.FormatInt(int64(activeSlots), 10),
		Timestamp: ts,

		System:   system,
		SourceId: "",
		Source:   config,
	}
}

// An unlockable was claimed.
func NewUnlockableClaimedEvent(system System, unlockableID string, unlockableConfig *UnlockablesConfigUnlockable, instanceID string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "unlockableClaimed",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"unlockableId": unlockableID,
			"instanceId":   instanceID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: unlockableID,
		Source:   unlockableConfig,
	}
}

func NewAuctionCreatedEvent(system System, auctionID string, auctionConfig *AuctionsConfigAuction, templateID, conditionID string, itemIDs map[string]int64, ts int64) (*PublisherEvent, error) {
	itemIDsEncoded, err := json.Marshal(itemIDs)
	if err != nil {
		return nil, err
	}

	return &PublisherEvent{
		Name: "auctionCreated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"auctionId":   auctionID,
			"templateId":  templateID,
			"conditionId": conditionID,
			"itemIds":     string(itemIDsEncoded),
		},
		Timestamp: ts,

		System:   system,
		SourceId: auctionID,
		Source:   auctionConfig,
	}, nil
}

func NewAuctionCancelledEvent(system System, auctionID string, auction *Auction, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "auctionCancelled",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"auctionId": auctionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: auctionID,
		Source:   auction,
	}
}

func NewAuctionBidEvent(system System, auctionID string, auction *Auction, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "auctionBid",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"auctionId": auctionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: auctionID,
		Source:   auction,
	}
}

func NewAuctionClaimCreatedEvent(system System, auctionID string, auction *Auction, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "auctionClaimCreated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"auctionId": auctionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: auctionID,
		Source:   auction,
	}
}

func NewAuctionClaimBidEvent(system System, auctionID string, auction *Auction, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "auctionClaimBid",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"auctionId": auctionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: auctionID,
		Source:   auction,
	}
}

// New challenge created by a player
func NewChallengeCreatedEvent(system System, challengeId string, challengeConfig *ChallengesConfigChallenge, templateId string, isOpen bool, size int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeCreated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"challengeId": challengeId,
			"templateId":  templateId,
			"isOpen":      strconv.FormatBool(isOpen),
			"maxSize":     strconv.FormatInt(size, 10),
		},
		Timestamp: ts,

		System:   system,
		SourceId: challengeId,
		Source:   challengeConfig,
	}
}

// New invitation to a challenge is sent by a player:
// 1. Invitees added while creating challenge
// 2. An invitation can be sent after the challenge is created.
func NewChallengeInvitationSentEvent(system System, challengeId string, challengeConfig any, inviteeId string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeInvitationSent",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"challengeId": challengeId,
			"inviteeId":   inviteeId,
		},
		Timestamp: ts,

		System:   system,
		SourceId: challengeId,
		Source:   challengeConfig,
	}
}

// New invitation accepted event - Sent when a player accepted the challenge invitation
func NewChallengeInvitationAcceptedEvent(system System, challengeId string, challengeConfig any, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeInvitationAccepted",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"challengeId": challengeId,
		},
		Timestamp: ts,

		System:   system,
		SourceId: challengeId,
		Source:   challengeConfig,
	}
}

// Player joined to an open challenge
func NewChallengeJoinedEvent(system System, challengeId string, challengeConfig any, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeJoined",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"challengeId": challengeId,
		},
		Timestamp: ts,

		System:   system,
		SourceId: challengeId,
		Source:   challengeConfig,
	}
}

// Player's challenge score or rank was updated
func NewChallengeUpdatedEvent(system System, challengeId string, challengeConfig any, score int64, subscore int64, oldRank int64, newRank int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeUpdated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"challengeId": challengeId,
			"score":       strconv.FormatInt(score, 10),
			"subscore":    strconv.FormatInt(subscore, 10),
			"oldRank":     strconv.FormatInt(oldRank, 10),
			"newRank":     strconv.FormatInt(newRank, 10),
		},
		Value:     strconv.FormatInt(score, 10),
		Timestamp: ts,

		System:   system,
		SourceId: challengeId,
		Source:   challengeConfig,
	}
}

// The player claims challenge rewards
func NewChallengeClaimedEvent(system System, challengeId string, challengeConfig any, score int64, subscore int64, rank int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeClaimed",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"challengeId": challengeId,
			"score":       strconv.FormatInt(score, 10),
			"subscore":    strconv.FormatInt(subscore, 10),
			"rank":        strconv.FormatInt(rank, 10),
		},
		Value:     strconv.FormatInt(score, 10),
		Timestamp: ts,

		System:   system,
		SourceId: challengeId,
		Source:   challengeConfig,
	}
}

// Player left the challenge
func NewChallengeLeftEvent(system System, challengeId string, challengeConfig any, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeLeft",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"challengeId": challengeId,
		},
		Timestamp: ts,

		System:   system,
		SourceId: challengeId,
		Source:   challengeConfig,
	}
}

// Helper for creating multiple reward-related events.
func NewRewardEvents(systems Hiro, sourceID string, sourceConfig any, reward *Reward, source string, ts int64) []*PublisherEvent {
	if reward == nil {
		return nil
	}

	events := make([]*PublisherEvent, 0)

	if reward.Items != nil {
		for itemID, amount := range reward.Items {
			events = append(events, NewItemGrantedEvent(systems.GetInventorySystem(), sourceID, sourceConfig, itemID, amount, source, ts))
		}
	}

	if reward.Currencies != nil {
		for itemID, amount := range reward.Currencies {
			events = append(events, NewCurrencyGrantedEvent(systems.GetEconomySystem(), sourceID, sourceConfig, itemID, amount, source, ts))
		}
	}

	if reward.Energies != nil {
		for energyID, amount := range reward.Energies {
			events = append(events, NewEnergyGrantedEvent(systems.GetEconomySystem(), sourceID, sourceConfig, energyID, amount, source, ts))
		}
	}

	if reward.EnergyModifiers != nil {
		for _, modifier := range reward.EnergyModifiers {
			events = append(events, NewEnergyModifierGrantedEvent(systems.GetEconomySystem(), sourceID, sourceConfig, modifier.Id, modifier.Operator, modifier.Value, modifier.DurationSec, source, ts))
		}
	}

	if reward.RewardModifiers != nil {
		for _, modifier := range reward.RewardModifiers {
			events = append(events, NewRewardModifierGrantedEvent(systems.GetEconomySystem(), sourceID, sourceConfig, modifier.Id, modifier.Type, modifier.Operator, modifier.Value, modifier.DurationSec, source, ts))
		}
	}

	return events
}

// Helper for creating multiple cost-related events.
func NewCostEvents(systems Hiro, sourceID string, sourceConfig any, currencyCost map[string]int64, itemCost map[string]int64, source string, ts int64) []*PublisherEvent {
	events := make([]*PublisherEvent, 0)

	if len(currencyCost) > 0 {
		for currencyID, amount := range currencyCost {
			events = append(events, NewCurrencySpentEvent(systems.GetEconomySystem(), sourceID, sourceConfig, currencyID, amount, source, ts))
		}
	}

	if len(itemCost) > 0 {
		for itemID, amount := range itemCost {
			events = append(events, NewItemSpentEvent(systems.GetInventorySystem(), sourceID, sourceConfig, itemID, amount, source, ts))
		}
	}

	return events
}
