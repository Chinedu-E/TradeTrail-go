FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go build -o main ./cmd/http/

EXPOSE 8000

ENV GO_ENV=production

CMD [ "/app/main" ]