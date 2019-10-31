FROM minextu/sdl2-cross-platform

# INSTALL GO 
RUN wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz \
 && tar -xvf go1.13.3.linux-amd64.tar.gz                      \
 && mv go /usr/local                                          
ENV GOROOT=/usr/local/go                               
ENV GOPATH=$HOME/go                                    
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# COPY CODE TO BUILD
RUN mkdir src \
 && mkdir src/builds
COPY /cmd /src/cmd
COPY /game /src/game
COPY /win-sdl2 /src/win-sdl2
COPY Makefile /src
COPY go.mod /src
COPY go.sum /src