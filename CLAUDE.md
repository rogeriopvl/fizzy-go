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

## Releases

Releases follow [SemVer](https://semver.org/) and are driven from the [Conventional Commits](https://www.conventionalcommits.org/) made since the previous release tag.

1. Make sure the working tree is clean and `make test` passes. Don't release a red build.
2. Inspect commits since the last `vX.Y.Z` tag and pick the new version:
   - `BREAKING CHANGE` footer or `!` after the type → major bump
   - `feat:` → minor bump
   - `fix:` / `chore:` / `docs:` / other types → patch bump
3. Update `CHANGELOG.md` with a new section for the chosen version, summarizing the changes. Review the diff before staging.
4. Commit the changelog update with `chore(release): vX.Y.Z`.
5. Create an annotated tag: `git tag -a vX.Y.Z -m "vX.Y.Z"`.
6. Push the release commit and tag together: `git push --follow-tags`. Tagging a commit that isn't on the remote yet will make the next step (and pkg.go.dev) fail.
7. Create the GitHub release with `gh release create vX.Y.Z`, using the new `CHANGELOG.md` section as the body (e.g. `--notes-file` pointing at an extracted snippet, or `--notes "$(...)"`).
8. Warm the Go module proxy so pkg.go.dev picks up the version: `curl https://proxy.golang.org/github.com/rogeriopvl/fizzy-go/@v/vX.Y.Z.info`.
