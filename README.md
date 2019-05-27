Glanceable
==========

[![Go Report Card](https://goreportcard.com/badge/github.com/YuukiARIA/glanceable)](https://goreportcard.com/report/github.com/YuukiARIA/glanceable)

Glanceable is an unofficial tool for formatting [Concourse](https://concourse-ci.org/) team configurations.

Currently `fly set-team` shows following output, but it is hard to understand what will change.

<div align="center">
  <img src="https://github.com/YuukiARIA/glanceable/blob/master/doc/set-team.png" width="60%" />
</div>

Glanceable is a workaround for this.

<div align="center">
  <img src="https://github.com/YuukiARIA/glanceable/blob/master/doc/get-team-glanceable.png" width="60%" />
</div>

## Install

```
go get github.com/YuukiARIA/glanceable
```

### Docker

```
docker pull yuukiaria/glanceable
```

## Usage

`glanceable` command requires two inputs:

- existing team configuration in JSON by stdin
- YAML file of new team configuration by `-c` option

Existing team configuration in JSON is provided by `fly get-team` with `--json` (`-j`).

So, usage pattern is basically as follows:

```
fly -t [target] get-team -n [team] -j | glanceable -c [team-config.yml]
```

### Docker

Note that `-i` is required to open stdin and mount workspace to make `[team-config.yml]` visible.

```
fly -t [target] get-team -n [team] -j | docker run --rm -i -v $(pwd):/work -w /work yuukiaria/glanceable -c [team-config.yml]
```

## Licence

Copyright (c) 2019 YuukiARIA

This sofware is released under the MIT License, see [LICENSE](https://github.com/YuukiARIA/glanceable/blob/master/LICENSE).