# flappy-gopher

[![Build Status](https://img.shields.io/travis/tclemos/flappy-gopher/master.svg?style=flat-square)](https://travis-ci.org/tclemos/flappy-gopher)

A game made using Go that mimics the flappy-bird game with Gopher style

## Dependencies

- Go
- SDL2
  - SDL2_img
  - SDL2_ttf
- MinGW

### OS X
``` bash
brew install go
brew install sdl2
brew install sdl2_img
brew install sdl2_ttf
brew install mingw-w64
```

### Ubuntu/Debian
``` bash
sudo apt-get install -y golang
sudo apt-get install -y libsdl2-dev
sudo apt-get install -y libsdl2-image-dev
sudo apt-get install -y libsdl2-ttf-dev
sudo apt-get install -y mingw-w64
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
