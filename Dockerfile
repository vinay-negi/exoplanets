# Dockerfile
FROM golang:1.20-alpine

RUN mkdir /app
ADD . /app/

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# COPY *.go ./

RUN go build -o exoplanet-service .

EXPOSE 8080

CMD ["./exoplanet-service"]
