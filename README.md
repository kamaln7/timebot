# timebot

timebot is a Slack slash command that prints out the current time for every user in the config file.

## Usage

1. `go get github.com/kamaln7/timebot/cmd/timebot`
2. Copy [config.toml](/config/config.toml) and edit it to match your settings
3. Run `timebot` in the same directory as your `config.toml`. Use an nginx reverse proxy, cloudflare, let's encrypt, or anything else to serve it via HTTPS with a valid SSL certificate (Slack requires that).
4. [Add a Slash Command](https://team.slack.com/apps/new/A0F82E8CA-slash-commands) with the URL set to your server's URL.

## Notes

* If `InChannel` is set to `true`, usernames will be ṁunged in order to prevent people from getting notified.

## License

See [LICENSE](/LICENSE)
