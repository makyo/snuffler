# snuffle

Snuffle your way through config files to the configuration your app needs.

## Goals

* Be able to build a config object from many sources
    * Global config files (e.g: in `/etc/appname`)
    * User config files (e.g: in `$HOME/.config/appname`)
    * An expected path (e.g: `$HOME/.appnamerc`)
    * A specified file (e.g: via `appname -c conf.yaml`)
    * Environment variables
* Merge all these sources into a single config object - more specific clobbers more general

## Design decisions

* YAML or TOML at first (pick one), to keep Dev fast and this project small

