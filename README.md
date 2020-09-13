# goslack

`goslack` is a simple go module to send Slack message.

## Installation

```bash
go get -u github.com/ilho-tiger/goslack
```

## Usage

As a CLI application,

```bash
Usage of /tmp/go-build709585003/b001/exe/goslack:
  -message string
        a message to send to Slack Incoming Webhook
  -url string
        a Slack Incoming Webhook URL to use
```

e.g.
```bash
$ goslack -message "Hello Slackers!"
```

As a Go package in your own Go applcation,

```go
package main

import "github.com/ilho-tiger/goslack/slack"

func main() {
	slack.SetWebhookURL("https://hooks.slack.com/...")
	slack.SendMessage("wow")
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[Apache License 2.0](./LICENSE)