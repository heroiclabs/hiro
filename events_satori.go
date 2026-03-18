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

package main

import (
	"encoding/json"
	"strconv"

	"github.com/heroiclabs/hiro"
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

// A donation was claimed.
func newDonationClaimedEvent(system hiro.System, donationID string, donationConfig *hiro.EconomyConfigDonation, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "donationClaimed",
		Id:   newUUIDv4(),
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
func newDonationGivenEvent(system hiro.System, donationID string, donationConfig *hiro.EconomyConfigDonation, recipientID string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "donationGiven",
		Id:   newUUIDv4(),
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
func newDonationRequestedEvent(system hiro.System, donationID string, donationConfig *hiro.EconomyConfigDonation, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "donationRequested",
		Id:   newUUIDv4(),
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
func newCurrencyGrantedEvent(system hiro.System, sourceID string, sourceConfig any, currencyID string, amount int64, source string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "currencyGranted",
		Id:   newUUIDv4(),
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

func newItemGrantedEvent(system hiro.System, sourceID string, sourceConfig any, itemID string, amount int64, source string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "itemsGranted",
		Id:   newUUIDv4(),
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

func newTeamItemGrantedEvent(system hiro.System, sourceID string, sourceConfig any, teamID, itemID string, amount int64, source string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "teamItemsGranted",
		Id:   newUUIDv4(),
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
func newEnergyGrantedEvent(system hiro.System, sourceID string, sourceConfig any, energyID string, amount int32, source string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "energyGranted",
		Id:   newUUIDv4(),
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
func newEnergyModiferGrantedEvent(system hiro.System, sourceID string, sourceConfig any, energyModifierID string, operator string, value int64, durationSec uint64, source string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "energyModifierGranted",
		Id:   newUUIDv4(),
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
func newRewardModifierGrantedEvent(system hiro.System, sourceID string, sourceConfig any, rewardModifierID string, modifierType string, operator string, value int64, durationSec uint64, source string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "rewardModifierGranted",
		Id:   newUUIDv4(),
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
func newCurrencySpentEvent(system hiro.System, sourceID string, sourceConfig any, currencyID string, amount int64, source string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "currencySpent",
		Id:   newUUIDv4(),
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
func newItemSpentEvent(system hiro.System, sourceID string, sourceConfig any, itemID string, amount int64, source string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "itemSpent",
		Id:   newUUIDv4(),
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
func newEnergySpentEvent(system hiro.System, energyID string, energyConfig *hiro.EnergyConfigEnergy, amount int32, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "energySpent",
		Id:   newUUIDv4(),
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
func newPurchaseIntentEvent(system hiro.System, storeItemID string, storeItem *hiro.EconomyConfigStoreItem, storeType hiro.EconomyStoreType, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "purchaseIntent",
		Id:   newUUIDv4(),
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
func newPurchaseCompletedEvent(system hiro.System, storeItemID string, storeItem *hiro.EconomyConfigStoreItem, currency string, amount float64, storeType hiro.EconomyStoreType, amountUSDCents int64, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "purchaseCompleted",
		Id:   newUUIDv4(),
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
func newAdPlacementStartedEvent(system hiro.System, placementID string, placementConfig *hiro.EconomyConfigPlacement, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "adPlacementStarted",
		Id:   newUUIDv4(),
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
func newAdPlacementSucceededEvent(system hiro.System, placementID string, placementConfig *hiro.EconomyConfigPlacement, maxRetries bool, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "adPlacementSucceeded",
		Id:   newUUIDv4(),
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
func newAdPlacementFailedEvent(system hiro.System, placementID string, placementConfig *hiro.EconomyConfigPlacement, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "adPlacementFailed",
		Id:   newUUIDv4(),
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
func newAchievementUpdatedEvent(system hiro.System, achievementID string, achievementConfig any, count int64, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "achievementUpdated",
		Id:   newUUIDv4(),
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
func newTeamAchievementUpdatedEvent(system hiro.System, teamID, achievementID string, achievementConfig any, count int64, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "achievementUpdated",
		Id:   newUUIDv4(),
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
func newAchievementClaimedEvent(system hiro.System, achievementID string, achievementConfig *hiro.AchievementsConfigAchievement, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "achievementClaimed",
		Id:   newUUIDv4(),
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
func newTeamAchievementClaimedEvent(system hiro.System, teamID, achievementID string, achievementConfig *hiro.AchievementsConfigAchievement, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "teamAchievementClaimed",
		Id:   newUUIDv4(),
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
func newProgressionPurchasedEvent(system hiro.System, progressionID string, progressionConfig *hiro.ProgressionConfigProgression, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "progressionPurchased",
		Id:   newUUIDv4(),
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
func newProgressionUpdatedEvent(system hiro.System, progressionID string, progressionConfig *hiro.ProgressionConfigProgression, countID string, count int64, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "progressionUpdated",
		Id:   newUUIDv4(),
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
func newProgressionResetEvent(system hiro.System, progressionID string, progressionConfig *hiro.ProgressionConfigProgression, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "progressionReset",
		Id:   newUUIDv4(),
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
func newItemsConsumedEvent(system hiro.System, itemID string, itemConfig *hiro.InventoryConfigItem, amount int64, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "itemsConsumed",
		Id:   newUUIDv4(),
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
func newTeamItemsConsumedEvent(system hiro.System, teamID, itemID string, itemConfig *hiro.InventoryConfigItem, amount int64, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "teamItemsConsumed",
		Id:   newUUIDv4(),
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
func newItemUpdatedEvent(system hiro.System, itemID string, itemConfig *hiro.InventoryConfigItem, stringProperties map[string]string, numericProperties map[string]float64, ts int64) (*hiro.PublisherEvent, error) {
	stringPropertiesJson, err := json.Marshal(stringProperties)
	if err != nil {
		return nil, err
	}

	numericPropertiesJson, err := json.Marshal(numericProperties)
	if err != nil {
		return nil, err
	}

	return &hiro.PublisherEvent{
		Name: "itemUpdated",
		Id:   newUUIDv4(),
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

func newTeamItemUpdatedEvent(system hiro.System, teamID, itemID string, itemConfig *hiro.InventoryConfigItem, stringProperties map[string]string, numericProperties map[string]float64, ts int64) (*hiro.PublisherEvent, error) {
	stringPropertiesJson, err := json.Marshal(stringProperties)
	if err != nil {
		return nil, err
	}

	numericPropertiesJson, err := json.Marshal(numericProperties)
	if err != nil {
		return nil, err
	}

	return &hiro.PublisherEvent{
		Name: "teamItemUpdated",
		Id:   newUUIDv4(),
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
func newStatUpdatedEvent(system hiro.System, name string, stat any, operator hiro.StatUpdateOperator, value int64, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
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
func newTeamStatUpdatedEvent(system hiro.System, name string, stat any, operator hiro.StatUpdateOperator, value int64, ts int64, teamID string) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "statUpdated",
		Id:   newUUIDv4(),
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
func newTutorialAcceptedEvent(system hiro.System, tutorialID string, tutorialConfig *hiro.TutorialsConfigTutorial, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "tutorialAccepted",
		Id:   newUUIDv4(),
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
func newTutorialDeclinedEvent(system hiro.System, tutorialID string, tutorialConfig *hiro.TutorialsConfigTutorial, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "tutorialDeclined",
		Id:   newUUIDv4(),
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
func newTutorialStartedEvent(system hiro.System, tutorialID string, tutorialConfig *hiro.TutorialsConfigTutorial, step int, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "tutorialStarted",
		Id:   newUUIDv4(),
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
func newTutorialAbandonedEvent(system hiro.System, tutorialID string, tutorialConfig *hiro.TutorialsConfigTutorial, step int32, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "tutorialAbandoned",
		Id:   newUUIDv4(),
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
func newTutorialStepCompletedEvent(system hiro.System, tutorialID string, tutorialConfig *hiro.TutorialsConfigTutorial, step int, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "tutorialStepCompleted",
		Id:   newUUIDv4(),
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
func newTutorialCompletedEvent(system hiro.System, tutorialID string, tutorialConfig *hiro.TutorialsConfigTutorial, step int, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "tutorialCompleted",
		Id:   newUUIDv4(),
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
func newTutorialResetEvent(system hiro.System, tutorialID string, tutorialConfig *hiro.TutorialsConfigTutorial, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "tutorialReset",
		Id:   newUUIDv4(),
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
func newTeamCreatedEvent(system hiro.System, teamID string, team *hiro.Team, open bool, maxCount int32, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "teamCreated",
		Id:   newUUIDv4(),
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
func newIncentiveCreatedEvent(system hiro.System, incentiveID string, incentiveConfig *hiro.IncentivesConfigIncentive, code string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "incentiveCreated",
		Id:   newUUIDv4(),
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
func newIncentiveDeletedEvent(system hiro.System, incentiveID string, incentiveConfig *hiro.IncentivesConfigIncentive, code string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "incentiveDeleted",
		Id:   newUUIDv4(),
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
func newIncentiveSenderClaimedEvent(system hiro.System, incentiveID string, incentiveConfig *hiro.IncentivesConfigIncentive, code string, clamaintID string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "incentiveSenderClaimed",
		Id:   newUUIDv4(),
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
func newIncentiveRecipientClaimedEvent(system hiro.System, incentiveID string, incentiveConfig *hiro.IncentivesConfigIncentive, code string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "incentiveRecipientClaimed",
		Id:   newUUIDv4(),
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
func newEventLeaderboardRolledEvent(system hiro.System, eventLeaderboardID string, eventLeaderboardConfig *hiro.EventLeaderboardsConfigLeaderboard, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "eventLeaderboardRolled",
		Id:   newUUIDv4(),
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
func newTeamEventLeaderboardRolledEvent(system hiro.System, teamID, eventLeaderboardID string, eventLeaderboardConfig *hiro.EventLeaderboardsConfigLeaderboard, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "teamEventLeaderboardRolled",
		Id:   newUUIDv4(),
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
func newEventLeaderboardUpdatedEvent(system hiro.System, eventLeaderboardID string, eventLeaderboardConfig *hiro.EventLeaderboardsConfigLeaderboard, score int64, subscore int64, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "eventLeaderboardUpdated",
		Id:   newUUIDv4(),
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
func newTeamEventLeaderboardUpdatedEvent(system hiro.System, teamID, eventLeaderboardID string, eventLeaderboardConfig *hiro.EventLeaderboardsConfigLeaderboard, score int64, subscore int64, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "teamEventLeaderboardUpdated",
		Id:   newUUIDv4(),
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
func newEventLeaderboardClaimedEvent(system hiro.System, eventLeaderboardID string, eventLeaderboardConfig *hiro.EventLeaderboardsConfigLeaderboard, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "eventLeaderboardClaimed",
		Id:   newUUIDv4(),
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
func newTeamEventLeaderboardClaimedEvent(system hiro.System, teamID, eventLeaderboardID string, eventLeaderboardConfig *hiro.EventLeaderboardsConfigLeaderboard, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "teamEventLeaderboardClaimed",
		Id:   newUUIDv4(),
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
func newUnlockableCreatedEvent(system hiro.System, unlockableID string, unlockableConfig *hiro.UnlockablesConfigUnlockable, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "unlockableCreated",
		Id:   newUUIDv4(),
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
func newUnlockableUnlockStartedEvent(system hiro.System, unlockableID string, unlockableConfig *hiro.UnlockablesConfigUnlockable, instanceID string, activeUnlockables int, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "unlockableUnlockStarted",
		Id:   newUUIDv4(),
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
func newUnlockableUnlockPurchasedEvent(system hiro.System, unlockableID string, unlockableConfig *hiro.UnlockablesConfigUnlockable, instanceID string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "unlockableUnlockPurchased",
		Id:   newUUIDv4(),
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
func newUnlockableSlotPurchasedEvent(system hiro.System, config *hiro.UnlockablesConfig, activeSlots int, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
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
func newUnlockableClaimedEvent(system hiro.System, unlockableID string, unlockableConfig *hiro.UnlockablesConfigUnlockable, instanceID string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "unlockableClaimed",
		Id:   newUUIDv4(),
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

func newAuctionCreatedEvent(system hiro.System, auctionID string, auctionConfig *hiro.AuctionsConfigAuction, templateID, conditionID string, itemIDs map[string]int64, ts int64) (*hiro.PublisherEvent, error) {
	itemIDsEncoded, err := json.Marshal(itemIDs)
	if err != nil {
		return nil, err
	}

	return &hiro.PublisherEvent{
		Name: "auctionCreated",
		Id:   newUUIDv4(),
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

func newAuctionCancelledEvent(system hiro.System, auctionID string, auction *hiro.Auction, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "auctionCancelled",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"auction_id": auctionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: auctionID,
		Source:   auction,
	}
}

func newAuctionBidEvent(system hiro.System, auctionID string, auction *hiro.Auction, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "auctionBid",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"auction_id": auctionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: auctionID,
		Source:   auction,
	}
}

func newAuctionClaimCreatedEvent(system hiro.System, auctionID string, auction *hiro.Auction, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "auctionClaimCreated",
		Id:   newUUIDv4(),
		Metadata: map[string]string{
			"auction_id": auctionID,
		},
		Timestamp: ts,

		System:   system,
		SourceId: auctionID,
		Source:   auction,
	}
}

func newAuctionClaimBidEvent(system hiro.System, auctionID string, auction *hiro.Auction, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "auctionClaimBid",
		Id:   newUUIDv4(),
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
func newChallengeCreatedEvent(system hiro.System, challengeId string, challengeConfig *hiro.ChallengesConfigChallenge, templateId string, isOpen bool, size int64, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "challengeCreated",
		Id:   newUUIDv4(),
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
func newChallengeInvitationSentEvent(system hiro.System, challengeId string, challengeConfig any, inviteeId string, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "challengeInvitationSent",
		Id:   newUUIDv4(),
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
func newChallengeInvitationAcceptedEvent(system hiro.System, challengeId string, challengeConfig any, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "challengeInvitationAccepted",
		Id:   newUUIDv4(),
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
func newChallengeJoinedEvent(system hiro.System, challengeId string, challengeConfig any, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "challengeJoined",
		Id:   newUUIDv4(),
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
func newChallengeUpdatedEvent(system hiro.System, challengeId string, challengeConfig any, score int64, subscore int64, oldRank int64, newRank int64, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "challengeUpdated",
		Id:   newUUIDv4(),
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
func newChallengeClaimedEvent(system hiro.System, challengeId string, challengeConfig any, score int64, subscore int64, rank int64, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "challengeClaimed",
		Id:   newUUIDv4(),
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
func newChallengeLeftEvent(system hiro.System, challengeId string, challengeConfig any, ts int64) *hiro.PublisherEvent {
	return &hiro.PublisherEvent{
		Name: "challengeLeft",
		Id:   newUUIDv4(),
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
func newRewardEvents(systems *Hiro, sourceID string, sourceConfig any, reward *hiro.Reward, source string, ts int64) []*hiro.PublisherEvent {
	if reward == nil {
		return nil
	}

	events := make([]*hiro.PublisherEvent, 0)

	if reward.Items != nil {
		for itemID, amount := range reward.Items {
			events = append(events, newItemGrantedEvent(systems.inventorySystem, sourceID, sourceConfig, itemID, amount, source, ts))
		}
	}

	if reward.Currencies != nil {
		for itemID, amount := range reward.Currencies {
			events = append(events, newCurrencyGrantedEvent(systems.economySystem, sourceID, sourceConfig, itemID, amount, source, ts))
		}
	}

	if reward.Energies != nil {
		for energyID, amount := range reward.Energies {
			events = append(events, newEnergyGrantedEvent(systems.economySystem, sourceID, sourceConfig, energyID, amount, source, ts))
		}
	}

	if reward.EnergyModifiers != nil {
		for _, modifier := range reward.EnergyModifiers {
			events = append(events, newEnergyModiferGrantedEvent(systems.economySystem, sourceID, sourceConfig, modifier.Id, modifier.Operator, modifier.Value, modifier.DurationSec, source, ts))
		}
	}

	if reward.RewardModifiers != nil {
		for _, modifier := range reward.RewardModifiers {
			events = append(events, newRewardModifierGrantedEvent(systems.economySystem, sourceID, sourceConfig, modifier.Id, modifier.Type, modifier.Operator, modifier.Value, modifier.DurationSec, source, ts))
		}
	}

	return events
}

// Helper for creating multiple cost-related events.
func newCostEvents(systems *Hiro, sourceID string, sourceConfig any, currencyCost map[string]int64, itemCost map[string]int64, source string, ts int64) []*hiro.PublisherEvent {
	events := make([]*hiro.PublisherEvent, 0)

	if len(currencyCost) > 0 {
		for currencyID, amount := range currencyCost {
			events = append(events, newCurrencySpentEvent(systems.economySystem, sourceID, sourceConfig, currencyID, amount, source, ts))
		}
	}

	if len(itemCost) > 0 {
		for itemID, amount := range itemCost {
			events = append(events, newItemSpentEvent(systems.inventorySystem, sourceID, sourceConfig, itemID, amount, source, ts))
		}
	}

	return events
}
