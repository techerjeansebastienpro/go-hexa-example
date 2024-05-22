FROM golang:1.22

WORKDIR /usr/src/app



COPY . .
RUN go mod download
RUN go run github.com/steebchen/prisma-client-go prefetch
ENV GENERATOR_PROVIDER "go run github.com/steebchen/prisma-client-go"
ENV GENERATOR_OUTPUT "./pkg/models"
RUN go run github.com/steebchen/prisma-client-go generate

ENTRYPOINT go mod tidy && go run /usr/src/app/cmd/api/api.go