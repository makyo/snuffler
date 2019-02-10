---
title: Snuffler
layout: home
---

Snuffler is a Golang package for managing cascading configuration files.
[![GoDoc reference](https://godoc.org/github.com/makyo/snuffler?status.svg)](https://godoc.org/github.com/makyo/snuffler)

Snuffler will snuffle through all of the paths and globs you provide to look for config files (YAML, TOML, and JSON, for the moment), and use those to populate the config object you provide. It will clobber existing keys, but that's often what you want when generating user specific config which has the ability to override global config. Simply provide the paths/glob where config might live and it will root through them in order for config files.

For instance, consider the following scenario:

* Your program has a bunch of global configuration. This includes sensible defaults, such as how much input to keep in memory and some details as to how the UI works.
    * This includes some very basic details in a master config file.
    * It also includes some other bits broken down by purpose into separate files.
* The user will also be able to define some of their own configuration.
    * Some of this takes the form of the user's own master config file in the form of an 'rc' file, which is likely overrides on stuff in the global config files.
    * But say you also have a bunch of open-ended stuff that can be configured, as well. The user might have their own collection of files broken down by purpose.

Or, in real-world terms, consider [Stimmtausch](https://stimmtausch.com):

* Stimmtausch has a bunch of global configuration. This includes sensitivle defaults, such as how many lines of the connection to keep in memory and how the UI works.
    * This includes some very basic details such as the config version in `/etc/stimmtausch/st.yaml`
    * It also includes some other bits broken down by purpose into separate files, such as `/etc/stimmtausch/client.yaml` and `/etc/stimmtausch/servers.yaml`.
* The user will also be able to define some of their own configuration.
    * Some of this takes the form of the user's own master config file in the form of `~/.strc`, which will contain overrides for servers and client settings.
    * But the user may also have a bunch of triggers (hilites, gags, macros, scripts, etc) defined in any number of files broken down by purpose in `~/.config/stimmtausch/`

Rather than put any thought or logic into overrides, it's often preferable to treat it as a cascade, where more specific config overrides more general config. This is doubly easy in Go where a lot of the structured file loaders (YAML, TOML, and JSON specifically, which Snuffler supports), will do this automatically by virtue of being passed a pointer to a `struct` to be filled out. It makes for a very fast, very lightweight library for configuration file management.

[![Support me on Patreon](https://img.shields.io/badge/patreon-support-%23222222.svg)](https://patreon.com/makyo)
[![Buy me a Coffee](https://img.shields.io/badge/kofi-support-%23222222.svg)](https://ko-fi.com/drabmakyo)
[![Support me on LiberaPay](https://img.shields.io/badge/liberapay-support-%23222222.svg)](https://liberapay.com/makyo)

