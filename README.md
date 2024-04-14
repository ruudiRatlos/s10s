
## regenerate client from OpenAPI

```bash
rm -r api
mkdir api
openapi-generator generate -g go -i "https://stoplight.io/api/v1/projects/spacetraders/spacetraders/nodes/reference/SpaceTraders.json?fromExportButton=true&snapshotType=http_service&deref=optimizedBundle" --skip-validate-spec --package-name "api" -o api -c openapi-cfg.json --git-host github.com --git-repo-id s10s --git-user-id ruudiRatlos

# fix bug in openapi-generator when generating into a submodule
gfind api/ -type f -print0 | xargs -0 -n 1 perl -i -pe 's!openapiclient "github.com/ruudiRatlos/s10s"!openapiclient "github.com/ruudiRatlos/s10s/api"!'
```

## API Deviations

- AllSystems instead of ListSystems
- Get Ship Cargo missing
- GetScrapShip -> GetScrapShipValue
