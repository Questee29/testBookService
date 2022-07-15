#1 stage - build stage

FROM golang:alpine3.16 AS builder 
# Support CGO and SSL
WORKDIR /app 
COPY go.mod go.sum app.env ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/server/main.go 

#Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/app.env .
COPY --from=builder /app/main .

CMD [ "./main" ]