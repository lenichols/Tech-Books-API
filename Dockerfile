FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/lenichols/Tech-Books-API/main
RUN cd /build && git clone github.com/lenichols/Tech-Books-API.git

RUN cd /build/Tech-Books-API/main && go build

EXPOSE 8080

ENTRYPOINT [ "/build/Tech-Books-API/main/main" ]