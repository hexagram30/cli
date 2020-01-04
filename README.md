# cli

*The unified CLI tool for Hexagram30*

[![][logo]][logo-large]


## Set Up Environment

Whatever means you use to set up a Golang programming environment. Here's an
example:

```shell
$ export GOPATH=$GOPATH:`pwd`
$ export GOBIN=$GOPATH/bin:`pwd`/bin
$ export PATH=$PATH:$GOBIN
$ mkdir -p `pwd`/src/github.com/hexagram30
$ cd ./src/github.com/hexagram30
$ git clone https://github.com/hexagram30/cli.git
$ cd cli
```

## Building

To install pre-built binaries, you currently need to have Go installed (in the
future, [hxgm30 binaries will be published to github](https://github.com/hexagram30/cli/issues/5)).

Having done the steps above for setting up an environment (or something similar):

```shell
$ make build
$ hxgm30 help
```

## Installation

Same steps as in the build, and then:

```shell
$ make install
```

## Dice Roller

Single roll:

```shell
$ hxgm30 roll d20
```
```
12
```

Multiple rolls of the same die:

```shell
$ hxgm30 roll d6 10
```
```
d6: [1 5 2 2 4 2 2 2 6 1]
```

## License

```
Copyright © 2019-2020, Hexagram30 <hexagram30@cnbb.games>

Apache License, Version 2.0
```

<!-- Named page links below: /-->

[logo]: https://raw.githubusercontent.com/hexagram30/resources/master/branding/logo/h30-logo-2-long-with-text-x695.png
[logo-large]: https://raw.githubusercontent.com/hexagram30/resources/master/branding/logo/h30-logo-2-long-with-text-x3440.png
