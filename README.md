github-dice
---
[![Build Status](https://travis-ci.org/s-tajima/github-dice.svg?branch=master)](https://travis-ci.org/s-tajima/github-dice)

A useful tool for assigning someone to GitHub Issues/Pull Requests, like a rolling dice.

```
$ ./github-dice -d
2016/06/23 19:49:45 Candidates are [s-tajima chocopie116 dkkoma]
2016/06/23 19:49:46 #9999 https://github.com/s-tajima/github-dice/issues/9999 issue's title => author:s-tajima assigned:chocopie116
```

## Index

* [Requirements](#requirements)
* [Installation](#installation)
* [Configure](#configure)
* [Usage](#usage)       
* [License](#license)    

## Requirements

github-dice requires the following to run:

* Golang

## Installation

```
$ go get github.com/s-tajima/github-dice
```

## Configure

Set your configuration as Environment Variables.
```
export GITHUB_ACCESS_TOKEN=
export GITHUB_ORGANIZATION=
export GITHUB_REPO=
export GITHUB_TEAM=   # github-dice works only Issues/Pull Requests that be created by this team members.
```
You can use .env file as well.


## Usage

```
Usage:
  github-dice [OPTIONS]

Application Options:
  -q, --query         Query strings. For search Issues/Pull Requests. (default: "type:pr is:open")
  -c, --comment       Comment. Would be posted before assigned. (default: ":game_die:")
  -n, --dry-run       If true, show candidates and list Issues, without assign.
  -f, --force         If true, reassign even if already assigned.
  -o, --run-once      If true, assign assign only one Issue.
  -a, --assign-author If true, Issue's author also assigns.
  -l, --limit         A maximum number of assign Issues. (default: 0)
  -e, --exempt-users  User names separated by comma who exempt assignee.
  -d, --debug

Help Options:
  -h, --help      Show this help message
```

## License

[MIT](./LICENSE)

## Author

[Satoshi Tajima](https://github.com/s-tajima)
