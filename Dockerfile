FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
# COPY *.go ./
COPY . .
RUN go build -o /app/build/accessibility-backend-v1 .

EXPOSE 8080

CMD [ "/app/build/accessibility-backend-v1" ]
