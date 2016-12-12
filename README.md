github-dice
---
[![Build Status](https://travis-ci.org/s-tajima/github-dice.svg?branch=master)](https://travis-ci.org/s-tajima/github-dice)

A useful tool for assigning someone to GitHub Issue, like a rolling dice.


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
export GITHUB_TEAM=
```
You can use .env file as well.


## Usage

```
Usage:
  github-dice [OPTIONS]

Application Options:
  -q, --query         query strings for search issue/pull-request. (default: "type:pr is:open")
  -c, --comment       issues's comment when assigned. (default: ":game_die:")
  -n, --dry-run       show candidates and list issues, without assign.
  -f, --force         if true, reassign even if already assigned.
  -o, --run-once      if true, assign just once issue.
  -a, --assign-author if true, issue/pr's author also assigns.
  -l, --limit         maximum number of issues per running command. (default: 0)
  -e, --exempt-users  user names separated by comma who exempt assignee.
  -d, --debug

Help Options:
  -h, --help      Show this help message
```

## License

[MIT](./LICENSE)

## Author

[Satoshi Tajima](https://github.com/s-tajima)

