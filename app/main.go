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
  // check for slack bot token
  if os.Getenv("SLACK_BOT_TOKEN")=="" {
    fmt.Fprintf(os.Stderr, "Missing SLACK_BOT_TOKEN")
    os.Exit(1)
  }

  // connect to db
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

  // start slackbot
  fmt.Println("slackbot starting...")

  bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), slacker.WithDebug(getEnvAsBool("DEBUG", false)))

  definition := &slacker.CommandDefinition{
    Description: "Affirmator affirms you!",
    Example:     "@Affirmator Tell me I'm beautiful",
    Handler: func(request slacker.Request, response slacker.ResponseWriter) {

      id, err_uuid := uuid.NewUUID()
      if err_uuid !=nil {
        fmt.Fprintf(os.Stderr, "Unable to generate UUID: %v\n", err_uuid)
        os.Exit(1)
      }
      fmt.Printf(id.String())

      _, err_pq := conn.Exec("insert into affirmations(transaction_id, affirmation) values($1, $2)", id, request.Param("affirmation"))

      if err_pq != nil {
        fmt.Fprintf(os.Stderr, "Unable to add task: %v\n", err_pq)
        os.Exit(1)
      }
      affirmation := fmt.Sprintf("Dangit, you are %v!", request.Param("affirmation"))

      response.Reply(affirmation)
    },
  }

  bot.Command("tell me i'm <affirmation>", definition)
  bot.Command("tell me i am <affirmation>", definition)
  bot.Command("tell me I'm <affirmation>", definition)
  bot.Command("tell me I am <affirmation>", definition)
  bot.Command("Tell me I'm <affirmation>", definition)
  bot.Command("Tell me I am <affirmation>", definition)

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
