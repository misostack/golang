# Config

## Config values

Cloud Config

```sh

```

```yml
server:
  host: "0.0.0.0"
  port: "8080"

database:
  host: "localhost"
  port: "5432"
  user: "postgres"
  password: "123456"
  dbname: "go_tracker"
  sslmode: "disable"

log:
  level: "info"
```

zerolog allows for logging at the following levels (from highest to lowest):

panic (zerolog.PanicLevel, 5)
fatal (zerolog.FatalLevel, 4)
error (zerolog.ErrorLevel, 3)
warn (zerolog.WarnLevel, 2)
info (zerolog.InfoLevel, 1)
debug (zerolog.DebugLevel, 0)
trace (zerolog.TraceLevel, -1)

## Usage

## How to add new config

## Cloud Support

### AWS

### Google Cloud

### Azure
