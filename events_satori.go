// Copyright 2023 GameUp Online, Inc.
// All rights reserved.
//
// NOTICE: All information contained herein is, and remains the property of GameUp
// Online, Inc. and its suppliers, if any. The intellectual and technical concepts
// contained herein are proprietary to GameUp Online, Inc. and its suppliers and
// may be covered by U.S. and Foreign Patents, patents in process, and are protected by
// trade secret or copyright law. Dissemination of this information or reproduction of
// this material is strictly forbidden unless prior written permission is obtained
// from GameUp Online, Inc.

package hiro

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strconv"
)

// source of obtaining/spending an item, currency, energy, or modifier.
const (
	eventSourceDonationClaimed           = "donation_claimed"
	eventSourceDonationGiven             = "donation_given"
	eventSourceDonationRequested         = "donation_requested"
	eventSourceEconomyGranted            = "economy_granted"
	eventSourceRewardGranted             = "reward_granted"
	eventSourceItemPurchased             = "item_purchased"
	eventSourcePlacementSucceeded        = "placement_succeeded"
	eventSourceAchievementUpdated        = "achievement_updated"
	eventSourceAchievementClaimed        = "achievement_claimed"
	eventSourceItemsConsumed             = "items_consumed"
	eventSourceItemsGranted              = "items_granted"
	eventSourceEnergySpent               = "energy_spent"
	eventSourceIncentiveSenderClaimed    = "incentive_sender_claimed"
	eventSourceIncentiveRecipientClaimed = "incentive_recipient_claimed"
	eventSourceEventLeaderboardClaimed   = "event_leaderboard_claimed"
	eventSourceProgressionPurchased      = "progression_purchased"
	eventSourceUnlockableUnlockPurchased = "unlockable_unlock_purchased"
	eventSourceUnlockableSlotPurchased   = "unlockable_slot_purchased"
	eventSourceUnlockableClaimed         = "unlockable_claimed"
	eventSourceInitializeUser            = "initialize_user"
	eventSourceChallengeClaimed          = "challenge_claimed"
)

func newUUID() string {
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
func newDonationClaimedEvent(system System, donationID string, donationConfig *EconomyConfigDonation, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "donationClaimed",
		Id:   newUUID(),
		Metadata: map[string]string{
			"donation_id": donationID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: donationID,
		Source:   donationConfig,
	}
}

// A donation was given.
func newDonationGivenEvent(system System, donationID string, donationConfig *EconomyConfigDonation, recipientID string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "donationGiven",
		Id:   newUUID(),
		Metadata: map[string]string{
			"donation_id":  donationID,
			"recipient_id": recipientID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: donationID,
		Source:   donationConfig,
	}
}

// A donation was requested.
func newDonationRequestedEvent(system System, donationID string, donationConfig *EconomyConfigDonation, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "donationRequested",
		Id:   newUUID(),
		Metadata: map[string]string{
			"donation_id": donationID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: donationID,
		Source:   donationConfig,
	}
}

// Currency was granted.
func newCurrencyGrantedEvent(system System, sourceID string, sourceConfig any, currencyID string, amount int64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "currencyGranted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"currency_id": currencyID,
			"source":      source,
		},
		Value:     strconv.FormatInt(amount, 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

func newItemGrantedEvent(system System, sourceID string, sourceConfig any, itemID string, amount int64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "itemsGranted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"item_id": itemID,
			"source":  source,
		},
		Value:     strconv.FormatInt(amount, 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

func newTeamItemGrantedEvent(system System, sourceID string, sourceConfig any, teamID, itemID string, amount int64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamItemsGranted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"team_id": teamID,
			"item_id": itemID,
			"source":  source,
		},
		Value:     strconv.FormatInt(amount, 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

// Energy was granted.
func newEnergyGrantedEvent(system System, sourceID string, sourceConfig any, energyID string, amount int32, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "energyGranted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"energy_id": energyID,
			"source":    source,
		},
		Timestamp: ts,
		Value:     strconv.FormatInt(int64(amount), 10),

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

// A reward modifier was granted.
func newEnergyModiferGrantedEvent(system System, sourceID string, sourceConfig any, energyModifierID string, operator string, value int64, durationSec uint64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "energyModifierGranted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"energy_modifier_id": energyModifierID,
			"operator":           operator,
			"duration_sec":       strconv.FormatUint(durationSec, 10),
			"source":             source,
		},
		Value:     strconv.FormatInt(value, 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

// A reward modifier was granted.
func newRewardModifierGrantedEvent(system System, sourceID string, sourceConfig any, rewardModifierID string, modifierType string, operator string, value int64, durationSec uint64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "rewardModifierGranted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"reward_modifier_id": rewardModifierID,
			"type":               modifierType,
			"operator":           operator,
			"duration_sec":       strconv.FormatUint(durationSec, 10),
			"source":             source,
		},
		Value:     strconv.FormatInt(value, 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

// A currency was spent.
func newCurrencySpentEvent(system System, sourceID string, sourceConfig any, currencyID string, amount int64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "currencySpent",
		Id:   newUUID(),
		Metadata: map[string]string{
			"currency_id": currencyID,
			"source":      source,
		},
		Value:     strconv.FormatInt(int64(amount), 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

// An item was spent.
func newItemSpentEvent(system System, sourceID string, sourceConfig any, itemID string, amount int64, source string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "itemSpent",
		Id:   newUUID(),
		Metadata: map[string]string{
			"item_id": itemID,
			"source":  source,
		},
		Value:     strconv.FormatInt(int64(amount), 10),
		Timestamp: ts,

		System:   system,
		SourceId: sourceID,
		Source:   sourceConfig,
	}
}

// Energy was spent.
func newEnergySpentEvent(system System, energyID string, energyConfig *EnergyConfigEnergy, amount int32, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "energySpent",
		Id:   newUUID(),
		Metadata: map[string]string{
			"energy_id": energyID,
		},
		Value:     strconv.FormatInt(int64(amount), 10),
		Timestamp: ts,

		System:   system,
		SourceId: energyID,
		Source:   energyConfig,
	}
}

// A purchase intent was placed.
func newPurchaseIntentEvent(system System, storeItemID string, storeItem *EconomyConfigStoreItem, storeType EconomyStoreType, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "purchaseIntent",
		Id:   newUUID(),
		Metadata: map[string]string{
			"store_item_id": storeItemID,
			"store_type":    storeType.String(),
		},
		Timestamp: ts,

		System:   system,
		SourceId: storeItemID,
		Source:   storeItem,
	}
}

// A SKU purchase was completed.
func newPurchaseCompletedEvent(system System, storeItemID string, storeItem *EconomyConfigStoreItem, currency string, amount float64, storeType EconomyStoreType, amountUSDCents int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "purchaseCompleted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"id":         storeItemID, // keep named as "id" rather than "store_item_id" for backwards compatibility.
			"currency":   currency,
			"amount":     strconv.FormatFloat(amount, 'f', 2, 64),
			"store_type": storeType.String(),
		},
		Value:     strconv.FormatInt(amountUSDCents, 10),
		Timestamp: ts,

		System:   system,
		SourceId: storeItemID,
		Source:   storeItem,
	}
}

// An ad placement started.
func newAdPlacementStartedEvent(system System, placementID string, placementConfig *EconomyConfigPlacement, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "adPlacementStarted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"placement_id": placementID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: placementID,
		Source:   placementConfig,
	}
}

// An ad placement succeeded.
func newAdPlacementSucceededEvent(system System, placementID string, placementConfig *EconomyConfigPlacement, maxRetries bool, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "adPlacementSucceeded",
		Id:   newUUID(),
		Metadata: map[string]string{
			"placement_id": placementID,
			"max_retries":  strconv.FormatBool(maxRetries),
		},
		Timestamp: ts,

		System:   system,
		SourceId: placementID,
		Source:   placementConfig,
	}
}

// An ad placement failed.
func newAdPlacementFailedEvent(system System, placementID string, placementConfig *EconomyConfigPlacement, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "adPlacementFailed",
		Id:   newUUID(),
		Metadata: map[string]string{
			"placement_id": placementID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: placementID,
		Source:   placementConfig,
	}
}

// An achievement was updated.
func newAchievementUpdatedEvent(system System, achievementID string, achievementConfig any, count int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "achievementUpdated",
		Id:   newUUID(),
		Metadata: map[string]string{
			"achievement_id": achievementID,
		},
		Value:     strconv.FormatInt(count, 10),
		Timestamp: ts,

		System:   system,
		SourceId: achievementID,
		Source:   achievementConfig,
	}
}

// A team achievement was updated.
func newTeamAchievementUpdatedEvent(system System, teamID, achievementID string, achievementConfig any, count int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "achievementUpdated",
		Id:   newUUID(),
		Metadata: map[string]string{
			"team_id":        teamID,
			"achievement_id": achievementID,
		},
		Value:     strconv.FormatInt(count, 10),
		Timestamp: ts,

		System:   system,
		SourceId: achievementID,
		Source:   achievementConfig,
	}
}

// An achievement was claimed.
func newAchievementClaimedEvent(system System, achievementID string, achievementConfig *AchievementsConfigAchievement, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "achievementClaimed",
		Id:   newUUID(),
		Metadata: map[string]string{
			"achievement_id": achievementID,
		},
		Timestamp: ts,
		// No value.

		System:   system,
		SourceId: achievementID,
		Source:   achievementConfig,
	}
}

// A team achievement was claimed.
func newTeamAchievementClaimedEvent(system System, teamID, achievementID string, achievementConfig *AchievementsConfigAchievement, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamAchievementClaimed",
		Id:   newUUID(),
		Metadata: map[string]string{
			"team_id":        teamID,
			"achievement_id": achievementID,
		},
		Timestamp: ts,
		// No value.

		System:   system,
		SourceId: achievementID,
		Source:   achievementConfig,
	}
}

// A progression was purchased.
func newProgressionPurchasedEvent(system System, progressionID string, progressionConfig *ProgressionConfigProgression, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "progressionPurchased",
		Id:   newUUID(),
		Metadata: map[string]string{
			"progression_id": progressionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: progressionID,
		Source:   progressionConfig,
	}
}

// A progression was updated.
func newProgressionUpdatedEvent(system System, progressionID string, progressionConfig *ProgressionConfigProgression, countID string, count int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "progressionUpdated",
		Id:   newUUID(),
		Metadata: map[string]string{
			"progression_id": progressionID,
			"count_id":       countID,
		},
		Value:     strconv.FormatInt(count, 10),
		Timestamp: ts,

		System:   system,
		SourceId: progressionID,
		Source:   progressionConfig,
	}
}

// A progression was reset.
func newProgressionResetEvent(system System, progressionID string, progressionConfig *ProgressionConfigProgression, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "progressionReset",
		Id:   newUUID(),
		Metadata: map[string]string{
			"progression_id": progressionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: progressionID,
		Source:   progressionConfig,
	}
}

// Inventory items were consumed.
func newItemsConsumedEvent(system System, itemID string, itemConfig *InventoryConfigItem, amount int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "itemsConsumed",
		Id:   newUUID(),
		Metadata: map[string]string{
			"item_id": itemID,
		},
		Value:     strconv.FormatInt(amount, 10),
		Timestamp: ts,

		System:   system,
		SourceId: itemID,
		Source:   itemConfig,
	}
}

// Team inventory items were consumed.
func newTeamItemsConsumedEvent(system System, teamID, itemID string, itemConfig *InventoryConfigItem, amount int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamItemsConsumed",
		Id:   newUUID(),
		Metadata: map[string]string{
			"item_id": itemID,
			"team_id": teamID,
		},
		Value:     strconv.FormatInt(amount, 10),
		Timestamp: ts,

		System:   system,
		SourceId: itemID,
		Source:   itemConfig,
	}
}

// An inventory item was updated.
func newItemUpdatedEvent(system System, itemID string, itemConfig *InventoryConfigItem, stringProperties map[string]string, numericProperties map[string]float64, ts int64) (*PublisherEvent, error) {
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
		Id:   newUUID(),
		Metadata: map[string]string{
			"item_id":            itemID,
			"string_properties":  string(stringPropertiesJson),
			"numeric_properties": string(numericPropertiesJson),
		},
		Timestamp: ts,

		System:   system,
		SourceId: itemID,
		Source:   itemConfig,
	}, nil
}

func newTeamItemUpdatedEvent(system System, teamID, itemID string, itemConfig *InventoryConfigItem, stringProperties map[string]string, numericProperties map[string]float64, ts int64) (*PublisherEvent, error) {
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
		Id:   newUUID(),
		Metadata: map[string]string{
			"team_id":            teamID,
			"item_id":            itemID,
			"string_properties":  string(stringPropertiesJson),
			"numeric_properties": string(numericPropertiesJson),
		},
		Timestamp: ts,

		System:   system,
		SourceId: itemID,
		Source:   itemConfig,
	}, nil
}

// Stats were updated.
func newStatUpdatedEvent(system System, name string, stat any, operator StatUpdateOperator, value int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "statUpdated",
		Id:   newUUID(),
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
func newTeamStatUpdatedEvent(system System, name string, stat any, operator StatUpdateOperator, value int64, ts int64, teamID string) *PublisherEvent {
	return &PublisherEvent{
		Name: "statUpdated",
		Id:   newUUID(),
		Metadata: map[string]string{
			"name":     name,
			"operator": operator.String(),
			"team_id":  teamID,
		},
		Value:     strconv.FormatInt(value, 10),
		Timestamp: ts,

		System:   system,
		SourceId: name,
		Source:   stat,
	}
}

// A tutorial was accepted.
func newTutorialAcceptedEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialAccepted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"tutorial_id": tutorialID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// A tutorial was declined.
func newTutorialDeclinedEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialDeclined",
		Id:   newUUID(),
		Metadata: map[string]string{
			"tutorial_id": tutorialID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// A tutorial started.
func newTutorialStartedEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, step int, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialStarted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"tutorial_id": tutorialID,
		},
		Value:     strconv.FormatInt(int64(step), 10),
		Timestamp: ts,

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// A tutorial was abandoned.
func newTutorialAbandonedEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, step int32, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialAbandoned",
		Id:   newUUID(),
		Metadata: map[string]string{
			"tutorial_id": tutorialID,
		},
		Timestamp: ts,
		Value:     strconv.FormatInt(int64(step), 10),

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// A tutorial step was completed.
func newTutorialStepCompletedEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, step int, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialStepCompleted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"tutorial_id": tutorialID,
		},
		Value:     strconv.FormatInt(int64(step), 10),
		Timestamp: ts,

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// A tutorial was completed.
func newTutorialCompletedEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, step int, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialCompleted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"tutorial_id": tutorialID,
		},
		Value:     strconv.FormatInt(int64(step), 10),
		Timestamp: ts,

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// One or more tutorials were reset.
func newTutorialResetEvent(system System, tutorialID string, tutorialConfig *TutorialsConfigTutorial, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "tutorialReset",
		Id:   newUUID(),
		Metadata: map[string]string{
			"tutorial_id": tutorialID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: tutorialID,
		Source:   tutorialConfig,
	}
}

// A team was created.
func newTeamCreatedEvent(system System, teamID string, team *Team, open bool, maxCount int32, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamCreated",
		Id:   newUUID(),
		Metadata: map[string]string{
			"team_id": teamID,
			"open":    strconv.FormatBool(open),
		},
		Timestamp: ts,
		Value:     strconv.FormatInt(int64(maxCount), 10),

		System:   system,
		SourceId: teamID,
		Source:   team,
	}
}

// An incentive was created.
func newIncentiveCreatedEvent(system System, incentiveID string, incentiveConfig *IncentivesConfigIncentive, code string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "incentiveCreated",
		Id:   newUUID(),
		Metadata: map[string]string{
			"incentive_id": incentiveID,
			"code":         code,
		},
		Timestamp: ts,

		System:   system,
		SourceId: incentiveID,
		Source:   incentiveConfig,
	}
}

// An incentive was deleted.
func newIncentiveDeletedEvent(system System, incentiveID string, incentiveConfig *IncentivesConfigIncentive, code string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "incentiveDeleted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"incentive_id": incentiveID,
			"code":         code,
		},
		Timestamp: ts,

		System:   system,
		SourceId: incentiveID,
		Source:   incentiveConfig,
	}
}

// An incentive was claimed by the sender.
func newIncentiveSenderClaimedEvent(system System, incentiveID string, incentiveConfig *IncentivesConfigIncentive, code string, clamaintID string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "incentiveSenderClaimed",
		Id:   newUUID(),
		Metadata: map[string]string{
			"incentive_id": incentiveID,
			"code":         code,
			"claimaint_id": clamaintID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: incentiveID,
		Source:   incentiveConfig,
	}
}

// An incentive was claimed by the recipient.
func newIncentiveRecipientClaimedEvent(system System, incentiveID string, incentiveConfig *IncentivesConfigIncentive, code string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "incentiveRecipientClaimed",
		Id:   newUUID(),
		Metadata: map[string]string{
			"incentive_id": incentiveID,
			"code":         code,
		},
		Timestamp: ts,

		System:   system,
		SourceId: incentiveID,
		Source:   incentiveConfig,
	}
}

// An event leaderboard was rolled.
func newEventLeaderboardRolledEvent(system System, eventLeaderboardID string, eventLeaderboardConfig *EventLeaderboardsConfigLeaderboard, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "eventLeaderboardRolled",
		Id:   newUUID(),
		Metadata: map[string]string{
			"event_leaderboard_id": eventLeaderboardID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: eventLeaderboardID,
		Source:   eventLeaderboardConfig,
	}
}

// A team leaderboard was rolled.
func newTeamEventLeaderboardRolledEvent(system System, teamID, eventLeaderboardID string, eventLeaderboardConfig *EventLeaderboardsConfigLeaderboard, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamEventLeaderboardRolled",
		Id:   newUUID(),
		Metadata: map[string]string{
			"team_id":              teamID,
			"event_leaderboard_id": eventLeaderboardID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: eventLeaderboardID,
		Source:   eventLeaderboardConfig,
	}
}

// An event leaderboard was updated.
func newEventLeaderboardUpdatedEvent(system System, eventLeaderboardID string, eventLeaderboardConfig *EventLeaderboardsConfigLeaderboard, score int64, subscore int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "eventLeaderboardUpdated",
		Id:   newUUID(),
		Metadata: map[string]string{
			"event_leaderboard_id": eventLeaderboardID,
			"subscore":             strconv.FormatInt(subscore, 10),
		},
		Timestamp: ts,
		Value:     strconv.FormatInt(score, 10),

		System:   system,
		SourceId: eventLeaderboardID,
		Source:   eventLeaderboardConfig,
	}
}

// A team event leaderboard was updated.
func newTeamEventLeaderboardUpdatedEvent(system System, teamID, eventLeaderboardID string, eventLeaderboardConfig *EventLeaderboardsConfigLeaderboard, score int64, subscore int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamEventLeaderboardUpdated",
		Id:   newUUID(),
		Metadata: map[string]string{
			"team_id":              teamID,
			"event_leaderboard_id": eventLeaderboardID,
			"subscore":             strconv.FormatInt(subscore, 10),
		},
		Timestamp: ts,
		Value:     strconv.FormatInt(score, 10),

		System:   system,
		SourceId: eventLeaderboardID,
		Source:   eventLeaderboardConfig,
	}
}

// An event leaderboard was claimed.
func newEventLeaderboardClaimedEvent(system System, eventLeaderboardID string, eventLeaderboardConfig *EventLeaderboardsConfigLeaderboard, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "eventLeaderboardClaimed",
		Id:   newUUID(),
		Metadata: map[string]string{
			"event_leaderboard_id": eventLeaderboardID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: eventLeaderboardID,
		Source:   eventLeaderboardConfig,
	}
}

// A team event leaderboard was claimed.
func newTeamEventLeaderboardClaimedEvent(system System, teamID, eventLeaderboardID string, eventLeaderboardConfig *EventLeaderboardsConfigLeaderboard, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "teamEventLeaderboardClaimed",
		Id:   newUUID(),
		Metadata: map[string]string{
			"team_id":              teamID,
			"event_leaderboard_id": eventLeaderboardID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: eventLeaderboardID,
		Source:   eventLeaderboardConfig,
	}
}

// An unlockable was created.
func newUnlockableCreatedEvent(system System, unlockableID string, unlockableConfig *UnlockablesConfigUnlockable, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "unlockableCreated",
		Id:   newUUID(),
		Metadata: map[string]string{
			"unlockable_id": unlockableID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: unlockableID,
		Source:   unlockableConfig,
	}
}

// An unlockable's unlock was started.
func newUnlockableUnlockStartedEvent(system System, unlockableID string, unlockableConfig *UnlockablesConfigUnlockable, instanceID string, activeUnlockables int, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "unlockableUnlockStarted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"unlockable_id": unlockableID,
			"instance_id":   instanceID,
		},
		Value:     strconv.FormatInt(int64(activeUnlockables), 10),
		Timestamp: ts,

		System:   system,
		SourceId: unlockableID,
		Source:   unlockableConfig,
	}
}

// An unlockable's unlock was purchased.
func newUnlockableUnlockPurchasedEvent(system System, unlockableID string, unlockableConfig *UnlockablesConfigUnlockable, instanceID string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "unlockableUnlockPurchased",
		Id:   newUUID(),
		Metadata: map[string]string{
			"unlockable_id": unlockableID,
			"instance_id":   instanceID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: unlockableID,
		Source:   unlockableConfig,
	}
}

// An unlockable slot was purchased.
func newUnlockableSlotPurchasedEvent(system System, config *UnlockablesConfig, activeSlots int, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name:      "unlockableSlotPurchased",
		Id:        newUUID(),
		Value:     strconv.FormatInt(int64(activeSlots), 10),
		Timestamp: ts,

		System:   system,
		SourceId: "",
		Source:   config,
	}
}

// An unlockable was claimed.
func newUnlockableClaimedEvent(system System, unlockableID string, unlockableConfig *UnlockablesConfigUnlockable, instanceID string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "unlockableClaimed",
		Id:   newUUID(),
		Metadata: map[string]string{
			"unlockable_id": unlockableID,
			"instance_id":   instanceID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: unlockableID,
		Source:   unlockableConfig,
	}
}

func newAuctionCreatedEvent(system System, auctionID string, auctionConfig *AuctionsConfigAuction, templateID, conditionID string, itemIDs map[string]int64, ts int64) (*PublisherEvent, error) {
	itemIDsEncoded, err := json.Marshal(itemIDs)
	if err != nil {
		return nil, err
	}

	return &PublisherEvent{
		Name: "auctionCreated",
		Id:   newUUID(),
		Metadata: map[string]string{
			"auction_id":   auctionID,
			"template_id":  templateID,
			"condition_id": conditionID,
			"item_ids":     string(itemIDsEncoded),
		},
		Timestamp: ts,

		System:   system,
		SourceId: auctionID,
		Source:   auctionConfig,
	}, nil
}

func newAuctionCancelledEvent(system System, auctionID string, auction *Auction, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "auctionCancelled",
		Id:   newUUID(),
		Metadata: map[string]string{
			"auction_id": auctionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: auctionID,
		Source:   auction,
	}
}

func newAuctionBidEvent(system System, auctionID string, auction *Auction, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "auctionBid",
		Id:   newUUID(),
		Metadata: map[string]string{
			"auction_id": auctionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: auctionID,
		Source:   auction,
	}
}

func newAuctionClaimCreatedEvent(system System, auctionID string, auction *Auction, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "auctionClaimCreated",
		Id:   newUUID(),
		Metadata: map[string]string{
			"auction_id": auctionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: auctionID,
		Source:   auction,
	}
}

func newAuctionClaimBidEvent(system System, auctionID string, auction *Auction, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "auctionClaimBid",
		Id:   newUUID(),
		Metadata: map[string]string{
			"auction_id": auctionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: auctionID,
		Source:   auction,
	}
}

// New challenge created by a player
func newChallengeCreatedEvent(system System, challengeId string, challengeConfig *ChallengesConfigChallenge, templateId string, isOpen bool, size int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeCreated",
		Id:   newUUID(),
		Metadata: map[string]string{
			"challenge_id": challengeId,
			"template_id":  templateId,
			"is_open":      strconv.FormatBool(isOpen),
			"max_size":     strconv.FormatInt(size, 10),
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
func newChallengeInvitationSentEvent(system System, challengeId string, challengeConfig any, inviteeId string, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeInvitationSent",
		Id:   newUUID(),
		Metadata: map[string]string{
			"challenge_id": challengeId,
			"invitee_id":   inviteeId,
		},
		Timestamp: ts,

		System:   system,
		SourceId: challengeId,
		Source:   challengeConfig,
	}
}

// New invitation accepted event - Sent when a player accepted the challenge invitation
func newChallengeInvitationAcceptedEvent(system System, challengeId string, challengeConfig any, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeInvitationAccepted",
		Id:   newUUID(),
		Metadata: map[string]string{
			"challenge_id": challengeId,
		},
		Timestamp: ts,

		System:   system,
		SourceId: challengeId,
		Source:   challengeConfig,
	}
}

// Player joined to an open challenge
func newChallengeJoinedEvent(system System, challengeId string, challengeConfig any, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeJoined",
		Id:   newUUID(),
		Metadata: map[string]string{
			"challenge_id": challengeId,
		},
		Timestamp: ts,

		System:   system,
		SourceId: challengeId,
		Source:   challengeConfig,
	}
}

// Player joined to an open challenge
func newChallengeUpdatedEvent(system System, challengeId string, challengeConfig any, score int64, subscore int64, oldRank int64, newRank int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeUpdated",
		Id:   newUUID(),
		Metadata: map[string]string{
			"challenge_id": challengeId,
			"score":        strconv.FormatInt(score, 10),
			"subscore":     strconv.FormatInt(subscore, 10),
			"old_rank":     strconv.FormatInt(oldRank, 10),
			"new_rank":     strconv.FormatInt(newRank, 10),
		},
		Value:     strconv.FormatInt(score, 10),
		Timestamp: ts,

		System:   system,
		SourceId: challengeId,
		Source:   challengeConfig,
	}
}

// The player claims challenge rewards
func newChallengeClaimedEvent(system System, challengeId string, challengeConfig any, score int64, subscore int64, rank int64, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeClaimed",
		Id:   newUUID(),
		Metadata: map[string]string{
			"challenge_id": challengeId,
			"score":        strconv.FormatInt(score, 10),
			"subscore":     strconv.FormatInt(subscore, 10),
			"rank":         strconv.FormatInt(rank, 10),
		},
		Value:     strconv.FormatInt(score, 10),
		Timestamp: ts,

		System:   system,
		SourceId: challengeId,
		Source:   challengeConfig,
	}
}

// Player left the challenge
func newChallengeLeftEvent(system System, challengeId string, challengeConfig any, ts int64) *PublisherEvent {
	return &PublisherEvent{
		Name: "challengeLeft",
		Id:   newUUID(),
		Metadata: map[string]string{
			"challenge_id": challengeId,
		},
		Timestamp: ts,

		System:   system,
		SourceId: challengeId,
		Source:   challengeConfig,
	}
}

// Helper for creating multiple reward-related events.
func newRewardEvents(systems Hiro, sourceID string, sourceConfig any, reward *Reward, source string, ts int64) []*PublisherEvent {
	if reward == nil {
		return nil
	}

	events := make([]*PublisherEvent, 0)

	if reward.Items != nil {
		for itemID, amount := range reward.Items {
			events = append(events, newItemGrantedEvent(systems.GetInventorySystem(), sourceID, sourceConfig, itemID, amount, source, ts))
		}
	}

	if reward.Currencies != nil {
		for itemID, amount := range reward.Currencies {
			events = append(events, newCurrencyGrantedEvent(systems.GetEconomySystem(), sourceID, sourceConfig, itemID, amount, source, ts))
		}
	}

	if reward.Energies != nil {
		for energyID, amount := range reward.Energies {
			events = append(events, newEnergyGrantedEvent(systems.GetEconomySystem(), sourceID, sourceConfig, energyID, amount, source, ts))
		}
	}

	if reward.EnergyModifiers != nil {
		for _, modifier := range reward.EnergyModifiers {
			events = append(events, newEnergyModiferGrantedEvent(systems.GetEconomySystem(), sourceID, sourceConfig, modifier.Id, modifier.Operator, modifier.Value, modifier.DurationSec, source, ts))
		}
	}

	if reward.RewardModifiers != nil {
		for _, modifier := range reward.RewardModifiers {
			events = append(events, newRewardModifierGrantedEvent(systems.GetEconomySystem(), sourceID, sourceConfig, modifier.Id, modifier.Type, modifier.Operator, modifier.Value, modifier.DurationSec, source, ts))
		}
	}

	return events
}

// Helper for creating multiple cost-related events.
func newCostEvents(systems Hiro, sourceID string, sourceConfig any, currencyCost map[string]int64, itemCost map[string]int64, source string, ts int64) []*PublisherEvent {
	events := make([]*PublisherEvent, 0)

	if len(currencyCost) > 0 {
		for currencyID, amount := range currencyCost {
			events = append(events, newCurrencySpentEvent(systems.GetEconomySystem(), sourceID, sourceConfig, currencyID, amount, source, ts))
		}
	}

	if len(itemCost) > 0 {
		for itemID, amount := range itemCost {
			events = append(events, newItemSpentEvent(systems.GetInventorySystem(), sourceID, sourceConfig, itemID, amount, source, ts))
		}
	}

	return events
}
