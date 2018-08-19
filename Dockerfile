FROM golang:alpine AS build-env
RUN apk add --no-cache git

RUN mkdir -p /go/src/github.com/kamaln7/timebot/
WORKDIR /go/src/github.com/kamaln7/timebot/
ADD . .

RUN go get -v -d ./...

RUN mkdir /opt && go build -o /opt/timebot github.com/kamaln7/timebot/cmd/timebot

# final image
FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=build-env /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=build-env /opt/timebot /opt/timebot

ENTRYPOINT ["/opt/timebot"]
