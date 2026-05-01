# fizzy-go

A Go client library for the [Fizzy](https://fizzy.do) API.

The package is `fizzy` and is consumed as `github.com/rogeriopvl/fizzy-go`. It exposes a `Client` (see `client.go`) constructed via `NewClient(accountSlug, accessToken, opts...)`, with functional options `WithBoard`, `WithHTTPClient`, `WithBaseURL`. Per-resource files (`boards.go`, `cards.go`, `columns.go`, `comments.go`, `identity.go`, `notifications.go`, `reactions.go`, `steps.go`, `tags.go`, `users.go`, `webhooks.go`, `pins.go`, `account.go`) hold the methods that hit the corresponding API endpoints. Shared types live in `types.go` and pagination helpers in `pagination.go`.

## API specs

The official Fizzy API specs live in `docs/api/`:

- `docs/api/README.md` — index/overview of the API
- `docs/api/sections/*.md` — one file per resource (boards, cards, columns, comments, identity, notifications, reactions, steps, tags, users, webhooks, pins, account, activities, authentication, exports, rich_text)

**Always read the relevant spec file in `docs/api/sections/` before implementing or modifying any feature.** The specs are the source of truth for endpoint paths, request/response shapes, query parameters, and error semantics — do not infer them from existing Go code alone, since the upstream API may have evolved.

The specs are synced from the upstream `basecamp/fizzy` repo via `make sync-api-spec`. Run that before starting work if the local copy may be stale.

## Conventions

- Each resource lives in its own file with a matching `_test.go` (e.g. `boards.go` / `boards_test.go`).
- Tests use `gotestsum` — run them with `make test`.
- New endpoints should follow the existing patterns: method on `*Client`, accept `context.Context` as the first arg, return typed structs from `types.go`.
- Every new feature must ship with tests in the matching `_test.go` file.
- Every feature gets its own commit using [Conventional Commits](https://www.conventionalcommits.org/) (`feat:`, `fix:`, `chore:`, `docs:`, etc., with an optional scope like `feat(board): ...`). Don't bundle unrelated features into a single commit.
- Do not add `Co-Authored-By` trailers (or any other AI-attribution trailer) to commit messages.
