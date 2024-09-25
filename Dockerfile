FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /go-docker

EXPOSE 8181

CMD [ "/go-docker" ]

