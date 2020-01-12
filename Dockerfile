FROM minextu/mingw-sdl2-debian:latest

RUN echo "$PWD"

# install sdl
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        libsdl2-dev                            \
        libsdl2-ttf-dev                        \
        libsdl2-image-dev                   && \
    rm -rf /var/lib/apt/lists/*

# install build tools
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        zip                                 && \
    rm -rf /var/lib/apt/lists/*

VOLUME "$(pwd)/builds" builds


CMD /bin/bash