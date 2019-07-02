# Download Slack Emojis

Go script to download all emojis from slack

## How to use it

Get a [slack token issued](https://api.slack.com/custom-integrations/legacy-tokens) and run:

```bash
SLACK_TOKEN=<TOKEN> go run emoji-download.go
```

That will pull all the emojis into the `emojis/` folder.
