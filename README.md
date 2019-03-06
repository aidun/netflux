# netflux

Netflux is a simple daemon which collects weather metrics from the netatmo weather api and push this metrics to an influxdb instance.

Supported metrics:

- temperature
- hummidity

Netflux will check for all stations linked to your account and all modules linked to your stations. If a module support none of the metrics, it will not be recognized.

## Installation

### Docker

```
docker pull aidun/netflux
```

### From Source

In order to install netflux from source, you have to install go 1.11 with go modules support.

```
git clone https://github.com/aidun/netflux.git
cd netflux
go test ./...
go build
chmod +x netflux
cp netflux /usr/local/bin
```

## Usage

### Requirements

- influxdb with authentication
- Netatmo account developer account
- Netatmo application https://dev.netatmo.com/myaccount/createanapp

### Commandline

Netflux requires some commandline parameters:

```
netflux is a tool to push netatmo weather data to influxdb

Usage:
  netflux [flags]

Flags:
  -h, --help                          help for netflux
      --influxdb_database string      Database
      --influxdb_password string      Password of the influxdb user
      --influxdb_url string           API of the influxdb instance
      --influxdb_user string          User of the influxdb with write access
      --netatmo_clientid string       Client-ID to Netatmo Weather API
      --netatmo_clientsecret string   Client-Secret to Netatmo Weather API
      --netatmo_password string       Password to Netatmo Weather API
      --netatmo_user string           User to Netatmo Weather API
```

An example call looks like:

```
netflux
  --influxdb_database netatmo \
  --influxdb_user user_with_write_access_to_influxdb \
  --influxdb_password start123 \
  --influxdb_url http://localhost:8086 \
  --netatmo_user youraccount@example.com \
  --netatmo_password start123
  --netatmo_clientid 123456
  --netatmo_clientsecret 654321
```
