# Gator CLI

**Gator** is a command-line RSS feed aggregator and reader with user management and feed following features. It uses a PostgreSQL database for storage and is written in Go.

---

## Requirements

- **Go** (version 1.24+)
- **PostgreSQL** (running and accessible)

---

## Installation

1. **Clone the repository:**

   ```
   git clone https://github.com/d-darac/gator.git
   cd gator
   ```

2. **Install the CLI:**

   ```
   go install ./...
   ```

   This will build and install the `gator` CLI binary to your `$GOPATH/bin` (make sure this is in your `$PATH`).

---

## Configuration

Before running `gator`, you need a config file specifying your PostgreSQL connection string and current user.

1. **Create a config file in your HOME directory** (`~/.gatorconfig.json`):

   ```json
   {
    "db_url": "postgres://username:password@localhost:5432/gator_db?sslmode=disable",
    "current_user_name": ""
   }
   ```

   - Replace `username`, `password`, and `gator_db` with your actual PostgreSQL credentials and database name.
   - `current_user_name` value will be set when you register a new user.

2. **Run database migrations**  
   (You can use a tool like [goose](https://github.com/pressly/goose) or run the SQL files in `sql/schema/` manually.)

---

## Usage

Run the CLI with:

```
gator <command> [args...]
```

### Example Commands

- **Register a new user:**

  ```
  gator register alice
  ```

- **Login as a user:**

  ```
  gator login alice
  ```

- **Add a new feed:**

  ```
  gator addfeed "Feed Name" "https://example.com/rss"
  ```

- **Follow an existing feed:**

  ```
  gator follow "https://example.com/rss"
  ```

- **Aggregate feeds (fetch new posts periodically):**

  ```
  gator agg 10m
  ```

  (Fetches new posts every 10 minutes.)

---

- **Browse your feed posts:**

  ```
  gator browse 5
  ```

- **List all users:**

  ```
  gator users
  ```

## Notes

- You must be logged in to add feeds, follow/unfollow feeds, or browse posts.
- The config file tracks the current user.
- All data is stored in PostgreSQL; make sure your database is running and accessible.
