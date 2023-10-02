# wiremock-cli

Wiremock CLI is a convenient way to interact with the [Wiremock](https://github.com/wiremock/wiremock) admin API

> Note: this software is currently under active development: anything can change at any time, API and UI must be
> considered unstable until we release version 1.0.0. At the moment this has only been tested on a locally running 
> Wiremock server, via http (not https) and with no authentication.

## Docs

For guidance on the functionality provided by Wiremock administration API, please refer to
the [Administration API](https://wiremock.org/docs/standalone/administration/) on the Wiremock website.

## Current status

Wiremock CLI is currently in alpha stage, but still provides a number of useful features:

```
./bin/wm help
wm is a simple command line tool to access the Wiremock admin API

██     ██ ██ ██████  ███████ ███    ███  ██████   ██████ ██   ██
██     ██ ██ ██   ██ ██      ████  ████ ██    ██ ██      ██  ██
██  █  ██ ██ ██████  █████   ██ ████ ██ ██    ██ ██      █████
██ ███ ██ ██ ██   ██ ██      ██  ██  ██ ██    ██ ██      ██  ██
 ███ ███  ██ ██   ██ ███████ ██      ██  ██████   ██████ ██   ██

----------------------------------------------------------------
|               Cloud: https://wiremock.io/cloud               |
|                                                              |
|               Slack: https://slack.wiremock.org              |
----------------------------------------------------------------

Usage:
  wm [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Gets resources from the Wiremock server
  help        Help about any command
  reset       Reset mappings & journal
  shutdown    Shutdown the Wiremock server

Flags:
  -a, --admin-prefix string   Wiremock admin url prefix (default "__admin")
  -h, --help                  help for wm
  -H, --host string           Wiremock host (default "http://localhost")
  -p, --port string           Wiremock port (default "8080")
  -v, --version               version for wm

Use "wm [command] --help" for more information about a command.
```

### Get resources

We currently support getting stub mappings, requests and scenarios.

```
./bin/wm get help
Gets resources from the Wiremock server

Usage:
  wm get [command]

Available Commands:
  mappings    Get stub mappings
  requests    Get all requests in journal
  scenarios   Get all scenarios

Flags:
  -h, --help   help for get

Global Flags:
  -a, --admin-prefix string   Wiremock admin url prefix (default "__admin")
  -H, --host string           Wiremock host (default "http://localhost")
  -p, --port string           Wiremock port (default "8080")

Use "wm get [command] --help" for more information about a command.
```

We also support getting stub mappings and requests by id:

```
./bin/wm get mappings -h
Get stub mappings - if an id is specified, only that mapping is returned

Usage:
  wm get mappings [flags]

Examples:
wm get mappings
wm get mappings --limit 5
wm get mappings --limit 5 --offset 10
wm get mappings --id 0baca68a-0112-4f26-8529-ac12d8eb3720


Flags:
  -h, --help         help for mappings
  -i, --id string    Specify the id of the specific mapping you want to return
  -l, --limit int    Limit the number of mappings returned (only used if no id is specified) (default 10)
  -o, --offset int   Offset the returned mappings by this number (only used if no id is specified)

Global Flags:
  -a, --admin-prefix string   Wiremock admin url prefix (default "__admin")
  -H, --host string           Wiremock host (default "http://localhost")
  -p, --port string           Wiremock port (default "8080")
```

```
./bin/wm get requests -h
Get all requests in journal - if an id is specified, only that request is returned

Usage:
  wm get requests [flags]

Examples:
wm get requests
wm get requests --limit 5
wm get requests --id 0baca68a-0112-4f26-8529-ac12d8eb3720


Flags:
  -h, --help        help for requests
  -i, --id string   Specify the id of the specific request you want to return
  -l, --limit int   Limit the number of requests returned (only used if no id is specified) (default 10)

Global Flags:
  -a, --admin-prefix string   Wiremock admin url prefix (default "__admin")
  -H, --host string           Wiremock host (default "http://localhost")
  -p, --port string           Wiremock port (default "8080")
```