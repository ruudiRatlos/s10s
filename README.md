# spacetracders.io Go SDK

This is a basic Go client SDK for [spacetraders.io](spacetraders.io).

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

## regenerate client from OpenAPI

```bash
rm -r api
mkdir api
openapi-generator generate -g go -i "https://stoplight.io/api/v1/projects/spacetraders/spacetraders/nodes/reference/SpaceTraders.json?fromExportButton=true&snapshotType=http_service&deref=optimizedBundle" --skip-validate-spec --package-name "api" -o api -c openapi-cfg.json --git-host github.com --git-repo-id s10s --git-user-id ruudiRatlos

# fix bug in openapi-generator when generating into a submodule
gfind api/ -type f -print0 | xargs -0 -n 1 perl -i -pe 's!openapiclient "github.com/ruudiRatlos/s10s"!openapiclient "github.com/ruudiRatlos/s10s/api"!'
```


