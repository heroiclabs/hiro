# Change Log
All notable changes to this project are documented below.

The format is based on [keep a changelog](http://keepachangelog.com) and this project uses [semantic versioning](http://semver.org).

:warning: This server code is versioned separately to the download of the [Hiro game framework](https://heroiclabs.com/hiro/). :warning:

## [Unreleased]


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
