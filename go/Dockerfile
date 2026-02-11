FROM alpine:3.21.0

WORKDIR /app

COPY ./backend-binary /app/backend-binary
COPY ./env.empty /app/env

EXPOSE 8080 9090

RUN addgroup -S golang && adduser -S golang -G golang
RUN chown -R golang:golang /app

USER golang

ENTRYPOINT ["/app/backend-binary"]