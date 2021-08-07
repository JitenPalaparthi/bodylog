FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
#CMD mkdir app
COPY --from=builder /go/bin/medicalResearch /app

ENTRYPOINT ./app
LABEL Name=medicalResearch Version=0.0.1
LABEL maintainer="readyGo team. JitenP@Outlook.Com"
EXPOSE 50061
