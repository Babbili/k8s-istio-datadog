FROM golang:1.20-alpine as build
WORKDIR /work
COPY . .
RUN go mod download && go mod verify
RUN go build -o app

FROM alpine as runtime

RUN mkdir -p /gcsa
RUN chmod 771 /gcsa
COPY --from=build /work/app /usr/local/bin/

EXPOSE 1323
CMD ["app"]
