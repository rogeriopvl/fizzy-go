# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.2.2] - 2026-05-04

### Fixed

- fix(identity): route `UpdateMyTimezone` under `:account_slug` — the method was building its URL from `BaseURL`, which 302-redirects to `/session/menu` in production and yields 406 once the client follows the redirect. Same class of bug as the eight methods fixed in v1.2.1; this one was missed

## [1.2.1] - 2026-05-02

### Fixed

- fix: route account-scoped endpoints under `:account_slug` — `GetMyPins`, `GetAccountSettings`, `UpdateAccountEntropy`, `GetAccountJoinCode`, `UpdateAccountJoinCode`, `ResetAccountJoinCode`, `CreateAccountExport`, and `GetAccountExport` were building URLs from `BaseURL`, which 302-redirects to `/session/menu` in production and yields 406 once the client follows the redirect

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

[1.2.1]: https://github.com/rogeriopvl/fizzy-go/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/rogeriopvl/fizzy-go/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/rogeriopvl/fizzy-go/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/rogeriopvl/fizzy-go/compare/v0.1.0...v1.0.0
[0.1.0]: https://github.com/rogeriopvl/fizzy-go/releases/tag/v0.1.0
