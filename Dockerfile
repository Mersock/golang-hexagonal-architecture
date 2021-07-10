FROM golang:1.16.5-alpine

RUN mkdir /app

WORKDIR /app 

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go get -u github.com/cosmtrek/air

EXPOSE 8080

ENTRYPOINT ["air","-c","air.toml"]
