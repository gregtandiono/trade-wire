# Trade Wire

**Trade Wire** is an open-source inquiry and trade administration system built in Go, React, Redux.

## How to run it on your machine

Prequisites:

- Golang
- NodeJS
- Postgresql
- [Realize](https://github.com/tockins/realize) (optional)

*I like to use Realize because it watches for file changes and rebuilds*

run from root directory:
```
realize run
```

## Config Files

Be sure to create an `config/app.toml` file that resembles `config/app.sample.toml` structure.

## Test

I'm debating whether to do a full-blown unit test for controllers and models or not. Right now I'm focusing on endpoint testing for now. (Not a big fan of TDD, but more towards BDD)

**How to run tests**

```
ENV=TEST go test ./...
```

## Roadmap

### Server

- [x] User model
- [x] Commodity model
- [x] Variety model
- [x] Company model
- [ ] Trade model
- [x] Contact model
- [ ] Vessel model
- [x] User controller
- [x] Commodity controller
- [x] Variety controller
- [x] Company controller
- [ ] Trade controller
- [x] Contact controller
- [ ] Vessel controller
- [ ] Model validation for incoming requests
- [x] Auth
- [ ] Unit tests
- [x] e2e tests

### Client

Will come back to this when I'm done with the backend

### Ops

Haven't given this much thought.. containers?

### Others

- [ ] Project website
- [ ] UX
- [ ] UI
- [ ] More polished README
