# Change Log
All notable changes to this project are documented below.

The format is based on [keep a changelog](http://keepachangelog.com) and this project uses [semantic versioning](http://semver.org).

:warning: This server code is versioned separately to the download of the [Hiro game framework](https://heroiclabs.com/hiro/). :warning:

## [Unreleased]
### Added
- Add switches for core and authenticate events to be sent by the SatoriPersonalizer.

### Changed
- Update nakama-common to v1.30.0 release.

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
