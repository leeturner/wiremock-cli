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
* [x] GET - /__admin/mappings - Get all stub mappings
* [x] GET - /__admin/mappings/{id} - Get a stub mapping by ID
* [x] DELETE - /__admin/mappings - Delete all stub mappings
* [x] DELETE - /__admin/mappings/{id} - Delete a stub mapping by ID
* [ ] POST - /__admin/mappings - Create a new stub mapping
* [ ] POST - /__admin/mappings/import - Import stub mappings
* [ ] POST - /__admin/mappings/save - Save stub mappings
* [ ] POST - /__admin/mappings/reset - Reset stub mappings
* [ ] PUT - /__admin/mappings/{id} - Update a stub mapping by ID

### Requests
* [x] GET - /__admin/requests - Get all requests in journal
* [x] GET - /__admin/requests/{id} - Get a request by ID
* [x] DELETE - /__admin/requests - Delete all requests in journal
* [x] DELETE - /__admin/requests/{id} - Delete a request by ID
* [ ] POST - /__admin/requests/count - Count requests by criteria
* [ ] POST - /__admin/requests/remove - Remove requests by criteria
* [ ] POST - /__admin/requests/remove-by-metadata - Remove requests by matching metadata
* [ ] POST - /__admin/requests/find - Find requests by criteria
* [ ] POST - /__admin/requests/find-by-metadata - Find requests by matching metadata
* [ ] GET - /__admin/requests/unmatched - Get all unmatched requests in journal

### Near Misses
* [ ] GET - /__admin/requests/unmatched/near-misses - Get all near misses for unmatched requests
* [ ] POST - /__admin/near-misses/request - Find near misses matching specific request
* [ ] POST - /__admin/near-misses/request-pattern - Find near misses matching specific request pattern
* 
### Recordings
* [ ] GET - /__admin/recordings/status - Get recording status
* [ ] POST - /__admin/recordings/start - Start recording
* [ ] POST - /__admin/recordings/stop - Stop recording
* [ ] POST - /__admin/recordings/snapshot - Take a snapshot of the current recording

### Scenarios
* [x] GET - /__admin/scenarios - Get all scenarios
* [ ] POST - /__admin/scenarios/reset - Reset the state of all scenarios

### System
* [x] POST - /__admin/shutdown - Shutdown WireMock
* [x] POST - /__admin/reset - Reset mappings and request journal
* [ ] POST - /__admin/settings - Update global WireMock settings
