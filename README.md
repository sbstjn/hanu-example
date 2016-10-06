# hanu example

Example `Go` Slack bot using the [hanu](https://github.com/sbstjn/hanu) framework.

## Supported commands

* `help` - Automated response by [hanu](https://github.com/sbstjn/hanu)
* `hi` - Bot will reply with `Hi yourself!`
* `shout <word>` - Bot will reply passed word in uppercase
* `uptime` - Bot will reply with it's uptime
* `version` - Bot will reply with it's version
* `whisper <word>` - Bot will reply passed word in lowercase

## Usage

```bash
$ > git clone git@github.com:sbstjn/hanu-example.git
$ > cd hanu-example
$ > HANU_EXAMPLE_SLACK_TOKEN=YOUR_TOKEN go run main.go
```

## Heroku

```bash
$ > git clone git@github.com:sbstjn/hanu-example.git
$ > cd hanu-example
$ > heroku create
$ > heroku config:set HANU_EXAMPLE_SLACK_TOKEN=YOUR_TOKEN
$ > git push heroku master
$ > heroku ps:scale worker=1

Scaling dynos... done, now running worker at 1:Hobby
```
