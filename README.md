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
- [ ] Buyer model
- [ ] Commodity model
- [ ] Variety model
- [ ] Supplier model
- [ ] Trade model
- [ ] Contact model
- [ ] Tracking model
- [x] User controller
- [ ] Buyer controller 
- [ ] Commodity controller
- [ ] Variety controller
- [ ] Supplier controller
- [ ] Trade controller
- [ ] Contact controller
- [ ] Tracking controller
- [ ] Model validation for incoming requests
- [x] Auth
- [ ] Unit tests
- [ ] e2e tests

### Client

Will come back to this when I'm done with the backend

### Ops

Haven't given this much thought.. containers?

### Others

- [ ] Project website
- [ ] UX
- [ ] UI
- [ ] More polished README
