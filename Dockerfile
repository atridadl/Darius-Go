FROM golang:1.21.3

WORKDIR /app

ADD . /templates
ADD . /pages
ADD . /api
ADD . /lib
ADD . /public

COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod download

COPY ./templates ./templates
COPY ./public ./public
COPY ./pages ./pages
COPY ./api ./api
COPY ./lib ./lib

RUN go build .

CMD [ "./darius" ]