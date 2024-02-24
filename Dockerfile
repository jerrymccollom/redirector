# syntax=docker/dockerfile:1
FROM golang:1.18

ADD . /app

# Set destination for COPY
WORKDIR /app

# Download Go modules
RUN go mod download

# Build
RUN go build -o /docker-main .

# Listen on port 80
EXPOSE 80

# Environment
ENV REDIRECT_URL https://evotravelagent.com/sugartrails

# Run
CMD ["/docker-main"]

