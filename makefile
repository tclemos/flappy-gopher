CMD_PATH=./cmd/flappy-gopher/main.go
FONTS_PATH=./game/fonts
SPRITES_PATH="./game/sprites"

BUILDS_PATH=./builds
WINDOWS_BUILD_PATH:=${BUILDS_PATH}/windows
OSX_BUILD_PATH:=${BUILDS_PATH}/osx
LINUX_BUILD_PATH:=${BUILDS_PATH}/linux

WIN_SDL2_PATH=./win-sdl2

EXECUTABLE_NAME=flappy-gopher

.PHONY: clean
clean:
	rm -Rf ${BUILDS_PATH}

.PHONY: configure
configure: clean
	mkdir ${BUILDS_PATH}
	mkdir ${WINDOWS_BUILD_PATH}
	mkdir ${OSX_BUILD_PATH}
	mkdir ${LINUX_BUILD_PATH}

.PHONE: build
build: configure
	docker build -t builder:latest .
	docker run --rm -v $$(pwd)/builds:/src/builds -it builder:latest bash -c "cd src && make build-lin && make build-osx && make build-win"

.PHONY: run
run: 
	go run ${CMD_PATH}

.PHONY: build-lin
build-lin:
	echo "not implemented yet"

.PHONY: build-osx
build-osx:
	echo "not implemented yet"

.PHONY: build-win
build-win:
	$(MAKE) BUILD_PATH=${WINDOWS_BUILD_PATH} CC="x86_64-w64-mingw32-gcc" GOOS=windows GOARCH=amd64 EXECUTABLE=${EXECUTABLE_NAME}.exe go-build 

.PHONY: go-build
go-build: 
	CGO_ENABLED=1 CC=${CC} GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${BUILD_PATH}/${EXECUTABLE} -ldflags "-s -w" ${CMD_PATH}
	mkdir ${BUILD_PATH}/game
	cp -R ${FONTS_PATH} ${BUILD_PATH}/game
	cp -R ${SPRITES_PATH} ${BUILD_PATH}/game
	
go-x1:
	$(MAKE) go-x2
go-x2: 
	echo x2

	

# .PHONY: build-linux
# build-linux: 
# 	cmd
# 	cmd
# 	$(MAKE) BUILD_PATH=${BUILDS_PATH}/${OS} build
# 	cmd
# 	$(MAKE) clean




# func deploy(path, exename) {
# 	gamepath := ${path}/game
# 	CGO_ENABLED=1 go build -o ${path}/${exename} -tags static -ldflags "-s -w" ${CMD_PATH}
# 	mkdir ${gamepath}
# 	cp -R ${FONTS_PATH} ${gamepath}
# 	cp -R ${SPRITES_PATH} ${gamepath}
# }

# # BEGIN LINUX BUILD
# build-lin-to-lin:
# 	deploy(${LINUX_BUILD_PATH}, ${EXECUTABLE_NAME})

# build-lin-to-osx:
# 	echo "Not implemented yet"

# build-lin-to-win: 
# 	deploy(${WINDOWS_BUILD_PATH}, ${EXECUTABLE_NAME}.exe)
# 	cp -a ${WIN_SDL2_PATH}/. ${WINDOWS_BUILD_PATH}

# # END LINUX BUILD

# # BEGIN OSX BUILD
# build-osx-to-lin: 
# 	echo "Not implemented yet"

# build-osx-to-osx:
# 	rm -Rf ${OSX_BUILD_PATH}
# 	CGO_ENABLED=1 CC=gcc GOOS=darwin GOARCH=amd64 go build -o ${OSX_BUILD_PATH}/${EXECUTABLE_NAME} -tags static -ldflags "-s -w" ${CMD_PATH}
# 	mkdir ${OSX_BUILD_PATH}/game
# 	cp -R ${FONTS_PATH} ${OSX_BUILD_PATH}/game
# 	cp -R ${SPRITES_PATH} ${OSX_BUILD_PATH}/game

# build-osx-to-win: build-unix-to-win
# # END OSX BUILD

# # BEGIN WINDOWS BUILD
# build-win-to-lin: 
# 	echo "Not implemented yet"

# build-win-to-osx:
# 	echo "Not implemented yet"

# build-win-to-win:
# 	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o ${WINDOWS_BUILD_PATH}/${EXECUTABLE_NAME}.exe -tags static -ldflags "-s -w" ${CMD_PATH}
# 	mkdir ${WINDOWS_BUILD_PATH}/game
# 	cp -R ${FONTS_PATH} ${WINDOWS_BUILD_PATH}/game
# 	cp -R ${SPRITES_PATH} ${WINDOWS_BUILD_PATH}/game
# 	cp -a ${WIN_SDL2_PATH}/. ${WINDOWS_BUILD_PATH}
# # END WINDOWS BUILD

# # BEGIN SHARED BUILD
# build-unix-to-win:
# 	CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" GOOS=windows GOARCH=amd64 go build -o ${WINDOWS_BUILD_PATH}/${EXECUTABLE_NAME}.exe -tags static -ldflags "-s -w" ${CMD_PATH}
# 	mkdir ${WINDOWS_BUILD_PATH}/game
# 	cp -R ${FONTS_PATH} ${WINDOWS_BUILD_PATH}/game
# 	cp -R ${SPRITES_PATH} ${WINDOWS_BUILD_PATH}/game
# 	cp -a ${WIN_SDL2_PATH}/. ${WINDOWS_BUILD_PATH}
# # END SHARED BUILD
