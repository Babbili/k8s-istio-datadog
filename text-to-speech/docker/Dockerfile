FROM golang:1.20-alpine as build
WORKDIR /work
COPY . .
RUN go mod download && go mod verify
RUN go build -o app

FROM scratch as runtime

COPY --from=build /work/app /usr/local/bin/

EXPOSE 5000
CMD ["app"]
