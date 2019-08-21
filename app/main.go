package main

import (
  "context"
  "fmt"
  "github.com/google/uuid"
  "github.com/jackc/pgx"
  "github.com/shomali11/slacker"
  "log"
  "os"
  "strconv"
)

var conn *pgx.Conn

func main() {

  config, err_pq := pgx.ParseEnvLibpq()

  if err_pq != nil {
    fmt.Fprintln(os.Stderr, "Unable to parse environment:", err_pq)
    os.Exit(1)
  }

  conn, err_pq = pgx.Connect(config)

  if err_pq != nil {
    fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err_pq)
    os.Exit(1)
  }

  fmt.Println("slackbot starting...")

  // TODO: check for slack token and exit if doesn't exist

  bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), slacker.WithDebug(getEnvAsBool("DEBUG", false)))

  definition := &slacker.CommandDefinition{
    Handler: func(request slacker.Request, response slacker.ResponseWriter) {


      id, err_uuid := uuid.NewUUID()
      if err_uuid !=nil {
        fmt.Fprintf(os.Stderr, "Unable to generate UUID: %v\n", err_uuid)
        os.Exit(1)
      }
      fmt.Printf(id.String())

      _, err_pq := conn.Exec("insert into affirmations(transaction_id, affirmation) values($1, $2)", id, "ping")

      if err_pq != nil {
        fmt.Fprintf(os.Stderr, "Unable to add task: %v\n", err_pq)
        os.Exit(1)
      }

      response.Reply("pong")
    },
  }

  bot.Command("ping", definition)

  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  err_bot := bot.Listen(ctx)
  if err_bot != nil {
    log.Fatal(err_bot)
  }
}

func getEnvAsBool(name string, defaultVal bool) bool {
    valStr := os.Getenv(name)
    if val, err := strconv.ParseBool(valStr); err == nil {
  return val
    }

    return defaultVal
}
