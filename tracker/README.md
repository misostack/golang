### Day 1: 16 Aug 2025

#### Getting started

```sh
cp .env.sample .env
cp config.sample.yml config.yml
```

> Application to manage codegym track

```sh
go get package-name
go mod tidy
```

```sh
go get -u github.com/go-chi/chi/v5
go install github.com/air-verse/air@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
go get github.com/joho/godotenv
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/rs/zerolog/log
```

```sh
goose
goose create add_tracker_table sql
goose up
goose status
goose create add_tracker_status_column sql
```

**Entities:**

- Tracker: TrackerId, Name
- TrackerDay: TrackerId, date, completed/not-completed

```json
{
  "date": "YYYY-mm-dd",
  "completed": 0
}
```

**Todo:**

- Init rest api project
- Add ORM
- Add service

### References

- [GoChi](https://go-chi.io/#/pages/getting_started) : golang http framework
- [Air](https://github.com/air-verse/air) : live reload golang project
- [Goose](https://pressly.github.io/goose/installation/#linux) : database migration
- [Godotenv](https://github.com/joho/godotenv): load env
- [GORM](https://gorm.io/docs/) : Golang ORM
- [Twelve Factors](https://12factor.net/)
- [Zerolog](https://github.com/rs/zerolog)
