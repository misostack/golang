# Each Domain

```
internal/
└── task/
    ├── handler.go       # HTTP handlers for task endpoints
    ├── service.go       # business logic
    ├── repository.go    # DB access
    ├── model.go         # struct Task {...}
    └── route.go         # chi subrouter for tasks
└── user/
    ├── handler.go
    ├── service.go
    ├── repository.go
    └── model.go
```

## Layer pattern

```
[ Handler ]  ->  [ Service ]  ->  [ Repository ]  ->  [ Model ]
   (HTTP)         (Business)       (Data access)       (Entity)
```
