FROM golang:1.20

WORKDIR /app

COPY go.* input.txt *.go ./

RUN go mod download && go build

CMD [ "./advent-2023" ]

