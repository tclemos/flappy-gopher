CMD_PATH=./cmd/flappy-gopher/main.go
FONTS_PATH=./game/fonts
SPRITES_PATH="./game/sprites"

BUILDS_PATH=./builds
WINDOWS_BUILD_PATH:=${BUILDS_PATH}/windows
OSX_BUILD_PATH:=${BUILDS_PATH}/osx
LINUX_BUILD_PATH:=${BUILDS_PATH}/linux

WIN_SDL2_PATH=./win-sdl2

EXECUTABLE_NAME=flappy-gopher

run: 
	go run ${CMD_PATH}

build:
	sudo docker build . -t flappy-gopher-builder:latest
	sudo docker run --rm --name=flappy-gopher-build -it flappy-gopher-builder bash

# BEGIN LINUX BUILD
build-lin-to-lin:
	rm -Rf ${LINUX_BUILD_PATH}
	CGO_ENABLED=1 go build -o ${LINUX_BUILD_PATH}/${EXECUTABLE_NAME} -tags static -ldflags "-s -w" ${CMD_PATH}
	mkdir ${LINUX_BUILD_PATH}/game
	cp -R ${FONTS_PATH} ${LINUX_BUILD_PATH}/game
	cp -R ${SPRITES_PATH} ${LINUX_BUILD_PATH}/game

build-lin-to-osx:
	echo "Not implemented yet"

build-lin-to-win: build-unix-to-win
# END LINUX BUILD

# BEGIN OSX BUILD
build-osx-to-lin: 
	echo "Not implemented yet"

build-osx-to-osx:
	rm -Rf ${OSX_BUILD_PATH}
	CGO_ENABLED=1 CC=gcc GOOS=darwin GOARCH=amd64 go build -o ${OSX_BUILD_PATH}/${EXECUTABLE_NAME} -tags static -ldflags "-s -w" ${CMD_PATH}
	mkdir ${OSX_BUILD_PATH}/game
	cp -R ${FONTS_PATH} ${OSX_BUILD_PATH}/game
	cp -R ${SPRITES_PATH} ${OSX_BUILD_PATH}/game

build-osx-to-win: build-unix-to-win
# END OSX BUILD

# BEGIN WINDOWS BUILD
build-win-to-lin: 
	echo "Not implemented yet"

build-win-to-osx:
	echo "Not implemented yet"

build-win-to-win:
	rm -Rf ${WINDOWS_BUILD_PATH}
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o ${WINDOWS_BUILD_PATH}/${EXECUTABLE_NAME}.exe -tags static -ldflags "-s -w" ${CMD_PATH}
	mkdir ${WINDOWS_BUILD_PATH}/game
	cp -R ${FONTS_PATH} ${WINDOWS_BUILD_PATH}/game
	cp -R ${SPRITES_PATH} ${WINDOWS_BUILD_PATH}/game
	cp -a ${WIN_SDL2_PATH}/. ${WINDOWS_BUILD_PATH}
# END WINDOWS BUILD

# BEGIN SHARED BUILD
build-unix-to-win:
	rm -Rf ${WINDOWS_BUILD_PATH}
	CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" GOOS=windows GOARCH=amd64 go build -o ${WINDOWS_BUILD_PATH}/${EXECUTABLE_NAME}.exe -tags static -ldflags "-s -w" ${CMD_PATH}
	mkdir ${WINDOWS_BUILD_PATH}/game
	cp -R ${FONTS_PATH} ${WINDOWS_BUILD_PATH}/game
	cp -R ${SPRITES_PATH} ${WINDOWS_BUILD_PATH}/game
	cp -a ${WIN_SDL2_PATH}/. ${WINDOWS_BUILD_PATH}
# END SHARED BUILD