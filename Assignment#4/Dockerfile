


FROM golang:1.20 as build-stage
# ADD . /cmd/api
WORKDIR /assignment#4

COPY go.mod go.sum ./

RUN go mod download
RUN apt-get update
RUN apt-get -y install postgresql-client
COPY . .

RUN go build -o myapp ./cmd/api

EXPOSE 8080
EXPOSE 5432
# EXPOSE 

CMD ["./myapp"]

