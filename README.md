# flappy-gopher

A game made using Go that mimics the flappy-bird game with Gopher style

## Dependencies

- Go
- SDL2
  - SDL2_img
  - SDL2_ttf
- MinGW

``` bash
brew install go
brew install sdl2
brew install sdl2_img
brew install sdl2_ttf
brew install mingw-w64
```

## Compiling

``` bash
make build-osx
make build-windows
```

## Running

### From Code

``` bash
make run
```
This command runs `go run` to execute the `./cmd/flappy-gopher/main.go`.

### OS X

``` bash
make run-osx
```

This command compiles the OS X version and runs this compiled version

### Windows

``` bash
make run-windows
```

This command compiles the Windows version and runs this compiled version
