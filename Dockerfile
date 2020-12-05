FROM scratch

EXPOSE 8080/tcp

ADD ./bin/go-tick-yourself-linux-amd64 /go-tick-yourself

CMD ["/go-tick-yourself"]