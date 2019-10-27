# flappy-gopher

[![Build Status](https://img.shields.io/travis/tclemos/flappy-gopher/master.svg?style=flat-square)](https://travis-ci.org/tclemos/flappy-gopher)

A game that mimics the flappy-bird game with Gopher style.

This game uses [go-sdl2](https://github.com/veandco/go-sdl2), that is a wrapper for the SDL2. 

> **_If you want to help with the development or run it locally from the code, the original SDL2 installation is required._**

> _If you just want to play the game, there is no need to install SDL2, just download the zip files provided by the releases, extract it and run the executable._

## Development dependencies

- [Go](https://golang.org/dl/) >= **1.12**
- [SDL2](https://www.libsdl.org/download-2.0.php)  **v2.0.10**
  - [SDL2_img](http://www.libsdl.org/projects/SDL_image/) >= **v2.0.5**
  - [SDL2_ttf](http://www.libsdl.org/projects/SDL_ttf/) >= **v2.0.15**

## Running

``` bash
make run
```
This command runs the game based on the local code, useful during the development phase.

Make sure to know it will focus the execution using your current OS as the target version, since it uses `go run` to execute the game.

If you want to run a compiled version, run one of the compile commands and then check the `./builds` directory.

## Deployment dependencies

If you are using `Linux` or `OS X` and want to compile a `Windows` version, you must install the `mingw-64` to be able to compile it.

Linux(Ubuntu/Debian)
```
sudo apt install -y mingw-w64
```

OS X
```
brew install mingw-w64
```

## Compiling
Make sure to use the correct command accordingly to your `development OS` and the `target OS`, otherwise it will not work properly.

### From Linux
``` bash
make build-lin-to-lin
make build-lin-to-osx
make build-lin-to-win
```

### From Windows
``` bash
make build-win-to-lin
make build-win-to-osx
make build-win-to-win
```

### From OS X
``` bash
make build-osx-to-lin
make build-osx-to-osx
make build-osx-to-win
```


