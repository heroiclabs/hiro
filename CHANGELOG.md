# Change Log
All notable changes to this project are documented below.

The format is based on [keep a changelog](http://keepachangelog.com) and this project uses [semantic versioning](http://semver.org).

:warning: This server code is versioned separately to the download of the [Hiro game framework](https://heroiclabs.com/hiro/). :warning:

## [1.11.0] - 2024-05-19
### Added
- Add "UnregisterRpc" to clear the implementation of one or more of the RPCs registered by gameplay systems in Hiro.
- Add helper function to convert "AvailableRewardsContents" type to an "EconomyConfigRewardContents" type.
- The "SatoriPersonalizer" can now cache data definitions within the lifecycle of a request if enabled.

### Changed
- The "ForceNetworkProbe" can now be switched between true/false at runtime.
- The collection name used to store data per player can be (optionally) set.
- Explicitly include "unordered_map" in CPP generated output for Windows platform.
- Run Economy initialize user before any custom after authentication hook.

### Fixed
- Fix how currencies are decoded when values are larger than "int32.MaxSize".
- Fix incorrect WARN message at startup with some Economy reward data definition validations.
- Add "type" field to JSON schemas for Incentives, Progressions, and Stats.
- Add "max_count" field to JSON schema in Economy.

## [1.10.0] - 2024-04-12
### Added
- (Unity) Add function to write score to regional leaderboards.

### Changed
- Use "omitempty" in marshaler settings with data definition structs.
- Improve error response codes in inventory and economy operations.
- "max_repeat_rolls" is now returned in the "AvailableRewards" type.
- Update to nakama-common v1.31.0 to be compatible with newer Nakama releases.
- Inventory "GrantItems" now returns the modified inventory and also the specific item instances which were granted.
- Use unsigned integers with the reward range type.
- Inventory items granted as part of a reward can now have their instance properties rolled at the same time.
- (Unity) Expose client and session types in the SatoriSystem type.
- (Unity) Update to latest Nakama and Satori SDK dependencies.

### Fixed
- Fixed unrecognized Inventory system type in storage personalizer.
- Restore behaviour where inventory items inherit their properties from the definition and those property keys are not stored in storage.
- (Unity) Fixed batch update function with player stats.
- Satori integration to publish analytics events correctly reads configuration parameters.
- (Unity) Detect UnityPurchasing 'fake' store and warn prices will have mock values.

## [1.9.0] - 2024-02-04
### Added
- New option "max_repeat_rolls" to set how many duplicate rows of rolled rewards can occur.
- The "StoragePersonalizer" can now update data definitions with a S2S RPC function.
- Progressions can now be programmatically reset.

### Changed
- The "SatoriPersonalizer" can optionally send analytics events for each gameplay system.

### Fixed
- (Unity) Fix visibility modifier with "StatUpdate" class.
- Set energy modifiers into server response with Energies spend function.
- Fix item properties not set when items are granted as part of user initialization.
- Fix unlockable slots populated in the wrong order when overflow slots are enabled.

## [1.8.1] - 2024-01-20
### Added
- New "UnmarshalWallet" function to get a Hiro wallet from a Nakama "\*api.Account" type.

### Changed
- Use clearer error messages in Personalizer middleware.
- Apply Satori identity authorization before Economy initialize user is processed.

### Fixed
- Use stable order when inter-dependent achievement progress updates are counted.
- Don't throw an error on reward grants if Energies system is uninitialized.

## [1.8.0] - 2023-12-27
### Added
- Add switches for core and authenticate events to be sent by the "SatoriPersonalizer".
- Add "instance_id" field to response in Inventory Item type.
- Allow the "Personalizer" type to be added as a chain of transforms to each gameplay's data definition.
- Achievement updates can now be sent as a batch to change different counts on multiple achievements at the same time.
- Progressions can now define a reset schedule similar to Achievements.
- New "StoragePersonalizer" type which can use Nakama's storage engine to manage gameplay data definitions.
- Progression "Reset" can be used to manually reset progress on a progression node (i.e. to reset a quest).
- (Unity) VContainer DI example is now packaged with the Unity package.
- (Unity) Add "IsClaimed" computed field to Achievement type.
- (Unity) Wrap "Satori.IClient" methods in "SatoriSystem" type for simpler code.
- Stats can update multiple different stats in a single request.
- (Unity) Progression IDs can optionally be sent to receive deltas for a portion of the progression graph.

### Changed
- Update nakama-common to v1.30.1 release.
- (Unreal) Update "HiroClient" with newest features.
- (TypeScript) Update "HiroClient" with newest features.
- Return instanced item rewards in response type when consumed.
- The "refill" and "refill_sec" fields are always populated in an Energy type (even if at max value).
- The builtin "SatoriPersonalizer" now (optionally) uses Satori Live Events to configure Event Leaderboards.
- Economy "Grant" now takes an optional wallet metadata input to record a reason in the Nakama ledger.
- A user who has not submitted any score to an Event Leaderboard is not eligible for rewards or promotions.
- Use Nakama's builtin Facebook Instant purchase validation function in the Economy system.
- If Satori is configured and enabled always authenticate server-side (rather than just new players).

### Fixed
- Some outdated or missing definitions and schemas have been updated.
- Don't throw an error when the sender claim has no reward defined.
- (Unity) Add the Preserve attribute to some types at the class level to avoid code stripping issues in Unity IL2CPP.
- (Unity) Notify observers should not be called twice in the Progression system.
- Energies granted in rewards should be returned immediately rather than the previous stale value.
- (Unity) Don't throw an error if Achievement category is unset or empty.
- (Unity) Use platform specific preprocessor statements with Unity Mobile Notifications system.
- Fix variable shadow error with how data definition of sub-achievements are populated in responses.
- Economy weighted table rewards should escape early if a valid reward row has already been granted.

## [1.7.0] - 2023-10-24
### Added
- New error type "ErrItemsNotConsumable" for Inventory items which are not consumable.

### Changed
- Energies "Grant" now returns a player's updated energies.
- "Get" will return an empty state for an Event Leaderboard when a player has never had a previous cohort.
- Add "locked" field to the storage engine index used with Event Leaderboard cohort generation.
- (Unity) Improve "InventorySystem" to use observer pattern.

### Fixed
- (Unity) Use "PurchaseFailureDescription.reason" with Unity IAP package for error messages.
- Sender claim uses the newer internal operation in Incentives system.
- Do not shadow parent Reward when it is created to be granted in Achievements system.
- (Unity) Use an async pattern in "IStoreListener.ProcessPurchase" with Unity IAP package.

## [1.6.0] - 2023-10-15
### Added
- Add fields for "is_active", "can_claim", and "can_roll" for simpler client code with Event Leaderboards.
- (Unity) Add "IncentivesSystem".
- New "max_overflow" field to the data definition for Energies.

### Changed
- (Unity) Allow both "IEconomyListStoreItem" and "IEconomyLocalizedStoreItem" to be used in purchase flows.

### Fixed
- Use Inventory after the Progression purchase has been applied to calculate the latest Progression deltas.
- Energy counts granted as an Economy Reward are kept as overflow.
- Fix panic in progression precondition comparison.
- Batch economy changes which resolve to items removed are now marked correctly.
- (Unity) Serialize the input for Inventory update items request correctly to JSON.
- Fix to progression deltas computations.

## [1.5.0] - 2023-10-04
### Added
- Add server interface for the Incentives gameplay system.
- Cohort selection in Event Leaderboards can now be overridden with a custom function.

### Changed
- "Get" in the Progression gameplay system now returns a delta of Progression Nodes who's state has changed if a previous graph is passed to it.

## [1.4.0] - 2023-09-14
### Added
- New function to "Roll" a new cohort with an Event Leaderboard in an active phase.
- Each Progression Node can now contain multiple counts for local progress to be expressed.

### Fixed
- Update Event Leaderboard cohort search for Nakama 3.17.1 release.

## [1.3.0] - 2023-09-07
### Added
- New gameplay system called Progression to express Saga Maps, RPG Quests, and other game mechanics.
- Event Leaderboards can now express promotion and demotion zones with percentages.

### Changed
- An Event Leaderboard which is active but no cohort has been assigned now returns a precondition failed on claim.

## [1.2.0] - 2023-08-29
### Added
- Add server interface for Stats gameplay system.

### Changed
- Pin dependencies to compatible versions of Nakama common at v1.28.1.
- Return all Reward Tiers when an Event Leaderboard is fetched for the current user.

### Fixed
- Fix weighted reward error when definition is empty (instead of nil).

## [1.1.0] - 2023-08-23
### Added
- Add server interface for Event Leaderboards gameplay system.

## [1.0.4] - 2023-08-22
### Changed
- Add ChannelMessageAck message to proto definition.

### Fixed
- Expose server functions for reward and roll in Hiro.

## [1.0.3] - 2023-08-10
### Added
- Add enum value options to proto definition as code generation hints for Unreal Engine.

### Changed
- Update to Nakama 3.17.0 release.

## [1.0.2] - 2023-07-11
### Changed
- Find the binary lookup path relative to Nakama modules dir.

## [1.0.1] - 2023-07-11
### Changed
- Pin dependencies to compatible versions of Nakama common at v1.27.0.

## [1.0.0] - 2023-07-10
### Added
- Initial public commit.
