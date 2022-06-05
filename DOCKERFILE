FROM golang:1.18.3-alpine3.16 AS buider

WORKDIR /go/src/github.com/DmitryIT/pipeline
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o my_pipeline -a



FROM alpine:3.16

RUN apk --no-cache add ca-certificates
WORKDIR /root

EXPOSE 8080

COPY --from=buider /go/src/github.com/DmitryIT/pipeline/my_pipeline ./
RUN chmod +x my_pipeline
CMD [ "./my_pipeline" ]


