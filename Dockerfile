FROM golang:alpine

WORKDIR /app
# COPY backend/go.mod .
# COPY backend/go.sum .

RUN apk update
RUN apk add git vim
RUN go get github.com/oxequa/realize
# RUN go mod download

# CMD ["go", "run", "main.go"]