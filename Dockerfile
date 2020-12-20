FROM golang:1.15 as build
ENV CGO_ENABLED="0" GOOS="linux" GOARCH="amd64"
RUN ["go", "get", "-u", "github.com/jteeuwen/go-bindata/..."]
COPY [".", "/opt/src/go-tick-yourself"]
WORKDIR "/opt/src/go-tick-yourself"
RUN ["go", "generate", "./..."]
RUN ["go", "build", "-o", "bin/go-tick-yourself", "main.go"]

FROM scratch as final
EXPOSE 8080/tcp
COPY --from=0 ["/opt/src/go-tick-yourself/bin/go-tick-yourself", "/go-tick-yourself"]
CMD ["/go-tick-yourself"]
