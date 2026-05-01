# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.2.0] - 2026-05-01

### Added

- feat(account): implement join codes (GetAccountJoinCode, UpdateAccountJoinCode, ResetAccountJoinCode)
- feat(users): implement avatar removal and email address change (DeleteUserAvatar, RequestUserEmailChange, ConfirmUserEmailChange)
- feat(notifications): implement settings management (GetNotificationSettings, UpdateNotificationSettings)
- feat(identity): implement timezone update (UpdateMyTimezone)
- feat(columns): implement card listing per column (GetColumnCards)
- feat(boards): implement board accesses listing (GetBoardAccesses)
- feat(webhooks): implement delivery history listing (GetWebhookDeliveries)
- feat(exports): implement account and user data exports (CreateAccountExport, GetAccountExport, CreateUserDataExport, GetUserDataExport)
- feat(activities): implement account activity feed (GetActivities)

### Fixed

- fix(types): add missing `postponed` field on `Card` so it decodes from `GET /columns/:id/cards` and activity-feed responses

## [1.1.0] - 2026-03-16

### Added

- feat(board): add PublishBoard, UnpublishBoard and UpdateBoardEntropy
- feat(identity): add CreateAccessToken
- feat: add webhook endpoints
- feat: add GetAccountSettings and UpdateAccountEntropy

## [1.0.0] - 2026-02-21

### Added

- feat(api): implement pins - pin/unpin cards and get pinned cards

## [0.1.0] - 2026-02-18

### Added

- feat: first release

[1.2.0]: https://github.com/rogeriopvl/fizzy-go/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/rogeriopvl/fizzy-go/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/rogeriopvl/fizzy-go/compare/v0.1.0...v1.0.0
[0.1.0]: https://github.com/rogeriopvl/fizzy-go/releases/tag/v0.1.0
