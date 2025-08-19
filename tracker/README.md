### Day 1: 16 Aug 2025

> Application to manage codegym track

```sh
go get package-name
go mod tidy
```

```sh
go get -u github.com/go-chi/chi/v5
go install github.com/air-verse/air@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
```

```sh
goose
goose create add_tracker_table sql
goose up
goose status
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
