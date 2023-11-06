FROM golang:1.21.3

WORKDIR /app

ADD . /templates
ADD . /public

COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod download

COPY ./templates ./templates
COPY ./public ./public

RUN go build .

CMD [ "./darius" ]