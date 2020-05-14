FROM golang:1.14.2-alpine3.11 as build
WORKDIR /go/src/gitlab.com/craftfoundry/dataclips_export_api_m
COPY . .
RUN apk update && apk add git
RUN go get
RUN go build -o app
FROM alpine:3.7
COPY --from=build /go/src/gitlab.com/craftfoundry/dataclips_export_api_m/app /usr/local/bin/app
RUN apk --no-cache add ca-certificates
ENTRYPOINT ["/usr/local/bin/app"]