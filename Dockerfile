FROM golang:1.23

WORKDIR /usr/src/app

COPY . .

RUN go mod download

RUN go build -o /usr/bin/app

EXPOSE 8181

CMD [ "/godocker" ]

