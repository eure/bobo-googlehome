bobo-googlehome
----

[![GoDoc][1]][2] [![License: MIT][3]][4] [![Release][5]][6] [![Build Status][7]][8] [![Co decov Coverage][11]][12] [![Go Report Card][13]][14] [![Code Climate][19]][20] [![BCH compliance][21]][22] [![Downloads][15]][16]

[1]: https://godoc.org/github.com/eure/bobo-googlehome?status.svg
[2]: https://godoc.org/github.com/eure/bobo-googlehome
[3]: https://img.shields.io/badge/License-MIT-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/eure/bobo-googlehome.svg
[6]: https://github.com/eure/bobo-googlehome/releases/latest
[7]: https://travis-ci.org/eure/bobo-googlehome.svg?branch=master
[8]: https://travis-ci.org/eure/bobo-googlehome
[9]: https://coveralls.io/repos/eure/bobo-googlehome/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/eure/bobo-googlehome?branch=master
[11]: https://codecov.io/github/eure/bobo-googlehome/coverage.svg?branch=master
[12]: https://codecov.io/github/eure/bobo-googlehome?branch=master
[13]: https://goreportcard.com/badge/github.com/eure/bobo-googlehome
[14]: https://goreportcard.com/report/github.com/eure/bobo-googlehome
[15]: https://img.shields.io/github/downloads/eure/bobo-googlehome/total.svg?maxAge=1800
[16]: https://github.com/eure/bobo-googlehome/releases
[17]: https://img.shields.io/github/stars/eure/bobo-googlehome.svg
[18]: https://github.com/eure/bobo-googlehome/stargazers
[19]: https://codeclimate.com/github/eure/bobo-googlehome/badges/gpa.svg
[20]: https://codeclimate.com/github/eure/bobo-googlehome
[21]: https://bettercodehub.com/edge/badge/eure/bobo-googlehome?branch=master
[22]: https://bettercodehub.com/



Google Home Commands for [eure/bobo](https://github.com/eure/bobo) (Slackbot)


# Install

```bash
$ go get -u github.com/eure/bobo-googlehome
```

# Build

```bash
$ make build
```

for Raspberry Pi

```bash
$ make build-arm6
```

# Run

```bash
SLACK_RTM_TOKEN=xoxb-0000... ./bin/bobo
```

## Environment variables

|Name|Description|
|:--|:--|
| `SLACK_RTM_TOKEN` | [Slack Bot Token](https://slack.com/apps/A0F7YS25R-bots) |
| `SLACK_BOT_TOKEN` | [Slack Bot Token](https://slack.com/apps/A0F7YS25R-bots) |
| `SLACK_TOKEN` | [Slack Bot Token](https://slack.com/apps/A0F7YS25R-bots) |
| `BOBO_DEBUG` | Flag for debug logging. Set [boolean like value](https://golang.org/pkg/strconv/#ParseBool). |
| `GOOGLE_HOME_HOST` | Hostname or IP address of Google Home for speech feature. |
| `GOOGLE_HOME_PORT` | Port number of Google Home. Default is `8009`. |
| `GOOGLE_HOME_LANG` | Speaking language of Google Home. Default is `en`. |
| `GOOGLE_HOME_ACCENT` | Speaking accent of Google Home. Default is `us`. |


## Supported Tasks

- Play (`say` command)
- Volume (`volume` command)
