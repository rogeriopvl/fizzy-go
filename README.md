# fizzy-go

A Go client library for the [Fizzy](https://fizzy.do) API.

## Installation

```bash
go get github.com/rogeriopvl/fizzy-go
```

## Usage

### Basic Setup

```go
package main

import (
    "context"
    "log"
    "os"

    fizzy "github.com/rogeriopvl/fizzy-go"
)

func main() {
    // Get token from environment, config file, or secure store
    token := os.Getenv("FIZZY_ACCESS_TOKEN")
    if token == "" {
        log.Fatal("FIZZY_ACCESS_TOKEN not set")
    }

    // Create client with account slug and access token
    client, err := fizzy.NewClient("/my-account-slug", token)
    if err != nil {
        log.Fatal(err)
    }

    // List boards
    ctx := context.Background()
    boards, err := client.GetBoards(ctx)
    if err != nil {
        log.Fatal(err)
    }

    for _, board := range boards {
        log.Printf("Board: %s (%s)\n", board.Name, board.ID)
    }
}
```

### Working with Boards

Some operations require a board context. You can set it at client creation or later:

```go
// Set board at creation time
client, err := fizzy.NewClient("/my-account-slug", token, fizzy.WithBoard("board-id"))

// Or set it later
client.SetBoard("board-id")

// Now you can use board-specific operations
columns, err := client.GetColumns(ctx)
```

### Options

#### WithBoard

Sets the board ID for board-specific operations (columns, creating cards, etc.).

```go
client, err := fizzy.NewClient("/my-account-slug", token, fizzy.WithBoard("board-123"))
```

#### WithHTTPClient

Provides a custom HTTP client (useful for custom timeouts, transport, proxies, etc.).

```go
customHTTPClient := &http.Client{
    Timeout: 60 * time.Second,
}
client, err := fizzy.NewClient("/my-account-slug", token, fizzy.WithHTTPClient(customHTTPClient))
```

#### WithBaseURL

Overrides the default base URL (useful for testing or self-hosted instances).

```go
client, err := fizzy.NewClient("/my-account-slug", token, fizzy.WithBaseURL("https://custom.fizzy.do"))
```

## API Coverage

- **Identity**: Get current user identity and accounts
- **Boards**: List, get, create, update, delete
- **Cards**: List, get, create, update, delete, close, reopen, postpone, triage, watch, assign, tag, golden
- **Columns**: List, get, create, update, delete
- **Comments**: List, get, create, update, delete
- **Reactions**: List, create, delete
- **Steps**: Get, create, update, delete
- **Tags**: List
- **Users**: List, get, update, deactivate
- **Notifications**: List, get, mark read/unread, mark all read

## License

MIT
