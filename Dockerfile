FROM golang:1.22.5-bullseye

RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME=/go/src/mathapp
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
EXPOSE 6143
CMD ["bee", "run"]
