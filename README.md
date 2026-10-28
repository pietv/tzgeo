# tzgeo

[![Go Reference](https://pkg.go.dev/badge/github.com/pietv/tzgeo.svg)](https://pkg.go.dev/github.com/pietv/tzgeo)

**tzgeo** provides functions to convert geographical locations in the form
of IANA timezones (for example: `Europe/Berlin`) to a pair of latitude and
longitude coördinates. The current timezone is usually set in the `TZ`
environment variable:

```bash
$ printenv TZ
```

## Coordinates

Latitude and longitude are positive and negative `float64` numbers:
* positive latitude is **N**orth, negative is **S**outh
* positive longitude is **E**ast, negative is **W**est.

```go
lat, lon, ok := tzgeo.Lookup("Europe/Berlin") // 52.5, 13.3667, true
```

## Data source

The data comes from `zone1970.tab`, part of the IANA time zone database.
On Linux it is usually found in the `tzdata` package (under
`/usr/share/zoneinfo/zone1970.tab`).

Fresh timezone data can be pulled with

```bash
$ curl -O https://data.iana.org/time-zones/data/zone1970.tab
```

and the `tzgeo` package rebuilt with it using the provided generaor (`generate_zones.go`),
in the package root, it regenerates the `zones.go` file:

```bash
$ go generate
```

## License

This project is licensed under the [MIT License](LICENSE).
