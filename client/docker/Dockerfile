FROM alpine:edge

RUN apk add curl

WORKDIR /app

COPY . .

RUN chmod +x script.sh 

EXPOSE 5000

ENTRYPOINT [ "/bin/sh" ]