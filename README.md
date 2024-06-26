# spacetraders.io Go SDK

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ruudiRatlos/s10s?style=flat-square)
[![GoDoc](https://godoc.org/github.com/ruudiRatlos/s10s?status.svg)](https://pkg.go.dev/github.com/ruudiRatlos/s10s)
[![Go Report Card](https://goreportcard.com/badge/github.com/ruudiRatlos/s10s)](https://goreportcard.com/report/github.com/ruudiRatlos/s10s)


This is a basic Go client SDK for [spacetraders.io](https://spacetraders.io).

## API Deviations

The API is pretty similiar compared to the published [OpenAPI specification](https://stoplight.io/api/v1/projects/spacetraders/spacetraders/nodes/reference/SpaceTraders.json?fromExportButton=true&snapshotType=http_service&deref=optimizedBundle).

A few things are different, though:

### SystemSymbol and WaypointSymbol

To diffentiate the two, SystemSymbol and WaypointSymbol have been introduced as
seperate types and are used by many of the Calls.

### AllSystems instead of ListSystems

Currently approximately 8500 systems exist in the game. Instead of a slice of
these, a channel is used, so you can use them right away.

### GetCargo is missing
Use GetMyShipFromValue instead and profit from all the other infos you also get
with one API call.

### GetScrapShip
GET https://api.spacetraders.io/v2/my/ships/{shipSymbol}/scrap was renamed to GetScrapShipValue

## regenerate client from OpenAPI specification

```bash
rm -r openapi
mkdir openapi
openapi-generator generate -g go -i "https://stoplight.io/api/v1/projects/spacetraders/spacetraders/nodes/reference/SpaceTraders.json?fromExportButton=true&snapshotType=http_service&deref=optimizedBundle" --skip-validate-spec -o openapi -c openapi-cfg.json --git-host github.com --git-repo-id s10s --git-user-id ruudiRatlos
gofmt -s -w ./openapi
```

## Release

```bash
export TAG=v1.0.x
git tag $TAG
git push origin $TAG
GOPROXY=proxy.golang.org go list -m github.com/ruudiRatlos/s10s@$TAG

