### Day 1: 16 Aug 2025

> Application to manage codegym track

```sh
go get package-name
go mod tidy
```

```sh
go get -u github.com/go-chi/chi/v5
go install github.com/air-verse/air@latest
```

**Entities:**

- Day: date, completed/not-completed

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
