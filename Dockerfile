FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=auto
RUN go get github.com/lenichols/Tech-Books-API
RUN cd /build && git clone github.com/lenichols/Tech-Books-API.git

RUN cd /build/Books-API/ && go build

EXPOSE 8080

ENTRYPOINT [ "/build/Tech-Books-API/main/" ]