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

run-w: 
	(cd ${WINDOWS_BUILD_PATH} && ./flappy-gopher.exe)
	
run-windows: build-windows run-w

run-o:
	(cd ${OSX_BUILD_PATH} && ./flappy-gopher)

run-osx: build-osx run-o

build-osx:
	rm -Rf ${OSX_BUILD_PATH}
	CGO_ENABLED=1 CC="gcc" GOOS="darwin" GOARCH="amd64" go build -o ${OSX_BUILD_PATH}/${EXECUTABLE_NAME} -tags static -ldflags "-s -w" ${CMD_PATH}
	mkdir ${OSX_BUILD_PATH}/game
	cp -R ${FONTS_PATH} ${OSX_BUILD_PATH}/game
	cp -R ${SPRITES_PATH} ${OSX_BUILD_PATH}/game

build-windows:
	rm -Rf ${WINDOWS_BUILD_PATH}
	CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" GOOS=windows GOARCH=amd64 go build -o ${WINDOWS_BUILD_PATH}/${EXECUTABLE_NAME}.exe -tags static -ldflags "-s -w" ${CMD_PATH}
	mkdir ${WINDOWS_BUILD_PATH}/game
	cp -R ${FONTS_PATH} ${WINDOWS_BUILD_PATH}/game
	cp -R ${SPRITES_PATH} ${WINDOWS_BUILD_PATH}/game
	cp -a ${WIN_SDL2_PATH}/. ${WINDOWS_BUILD_PATH}