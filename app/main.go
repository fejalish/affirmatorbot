package main

import (
  "context"
  "fmt"
  "github.com/shomali11/slacker"
  "log"
  "os"
  "strconv"
)

func main() {
  fmt.Println("slackbot starting...")

  // TODO: check for slack token and exit if doesn't exist

  bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), slacker.WithDebug(getEnvAsBool("DEBUG", false)))

  definition := &slacker.CommandDefinition{
    Handler: func(request slacker.Request, response slacker.ResponseWriter) {
      response.Reply("pong")
    },
  }

  bot.Command("ping", definition)

  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()

  err := bot.Listen(ctx)
  if err != nil {
    log.Fatal(err)
  }
}

func getEnvAsBool(name string, defaultVal bool) bool {
    valStr := os.Getenv(name)
    if val, err := strconv.ParseBool(valStr); err == nil {
  return val
    }

    return defaultVal
}
