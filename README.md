Glanceable
==========

Glanceable is an unofficial tool for formatting [Concourse](https://concourse-ci.org/) team configurations.

Currently `fly set-team` shows following output, but it is hard to understand what will change.

<div align="center">
  <img src="doc/set-team.png" width="60%" />
</div>

Glanceable is a workaround for this.

<div align="center">
  <img src="doc/get-team-glanceable.png" width="60%" />
</div>

## Usage

`glanceable` command requires two inputs:

- existing team configuration in JSON by stdin
- YAML file of new team configuration by `-c` option

Existing team configuration in JSON is provided by`fly get-team` with `--json` (`-j`).

So, usage pattern is basically as follows:

```
fly -t <target> get-team -n <team> -j | glanceable -c <team-config.yml>
```

## Licence

Copyright (c) 2019 YuukiARIA

This sofware is released under the MIT License, see [LICENSE](https://github.com/YuukiARIA/glanceable/blob/master/LICENSE).