# wiremock-cli

Wiremock CLI is a convenient way to interact with the [Wiremock](https://github.com/wiremock/wiremock) admin API

> Note: this software is currently under active development: anything can change at any time, API and UI must be
> considered unstable until we release version 1.0.0. At the moment this has only been tested on a locally running 
> Wiremock server, via http (not https) and with no authentication.

## Docs

For guidance on the functionality provided by Wiremock administration API, please refer to
the [Administration API](https://wiremock.org/docs/standalone/administration/) on the Wiremock website.

## Current status

Wiremock CLI is currently in alpha stage, but still provides a number of useful features.  Full documentation can be
found [here](docs/wm.md).

## Feature parity with Wiremock admin API

### Mappings
* [x] GET - /__admin/mappings 
* [x] GET - /__admin/mappings/{id}
* [x] DELETE - /__admin/mappings/{id}
* [x] DELETE - /__admin/mappings
* [ ] POST - /__admin/mappings
* [ ] POST - /__admin/mappings/import
* [ ] POST - /__admin/mappings/save
* [ ] POST - /__admin/mappings/reset
* [ ] PUT - /__admin/mappings/{id}

### Requests

* [x] GET - /__admin/requests 
* [x] GET - /__admin/requests/{id}
* [x] DELETE - /__admin/requests
* [x] DELETE - /__admin/requests/{id}
* [ ] POST - /__admin/requests/count
* [ ] POST - /__admin/requests/remove
* [ ] POST - /__admin/requests/remove-by-metadata
* [ ] POST - /__admin/requests/find
* [ ] POST - /__admin/requests/find-by-metadata
* [ ] GET - /__admin/requests/unmatched

### Near Misses

* [ ] GET - /__admin/requests/unmatched/near-misses
* [ ] POST - /__admin/near-misses/request
* [ ] POST - /__admin/near-misses/request-pattern

### Recordings

* [ ] GET - /__admin/recordings/status
* [ ] POST - /__admin/recordings/start
* [ ] POST - /__admin/recordings/stop
* [ ] POST - /__admin/recordings/snapshot

### Scenarios

* [x] GET - /__admin/scenarios
* [ ] POST - /__admin/scenarios/reset

### System

* [x] POST - /__admin/shutdown
* [x] POST - /__admin/reset
* [ ] POST - /__admin/settings
