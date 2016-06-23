github-dice
---
[![Build Status](https://travis-ci.org/s-tajima/github-dice.svg?branch=master)](https://travis-ci.org/s-tajima/github-dice)

A useful tool for assigning someone to GitHub Issue, like a rolling dice.


```
$ ./github-dice
2016/06/23 19:49:45 Candidates are s-tajima, chocopie116, dkkoma
2016/06/23 19:49:46 Assigned s-tajima to #41 (some awsome issue)
2016/06/23 19:49:46 Assigned chocopie116 to #42 (some bothersome issue )
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
  -q, --query=    query strings for search issue/pull-request. (default: is:issue)
  -n, --dry-run   show candidates and list issues, without assign.
  -f, --force     if true, reassign even if already assigned.
  -o, --run-once  if true, assign just once issue.
  -d, --debug

Help Options:
  -h, --help      Show this help message
```

## License

[MIT](./LICENSE)

## Author

[Satoshi Tajima](https://github.com/s-tajima)
