FROM golang:1.20-alpine as build
WORKDIR /work
COPY . .
RUN go mod download && go mod verify
RUN go build -o app

FROM alpine as runtime

RUN apk add curl

RUN mkdir -p /gc-service-account/creds
RUN chmod 771 /gc-service-account/creds

COPY --from=build /work/ops/sh/script.sh ops/sh/script.sh
COPY --from=build /work/ops/sh/curl.sh ops/sh/curl.sh
RUN chmod +x ops/sh/script.sh
RUN chmod +x ops/sh/curl.sh
COPY --from=build /work/app /usr/local/bin/

EXPOSE 1323
CMD ["app"]
