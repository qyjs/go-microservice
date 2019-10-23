FROM golang:1.12.0-alpine3.9 as BUILD
COPY . /app/
WORKDIR /app/
RUN go build -mod=vendor -o app .


FROM alpine:3.9
WORKDIR /app/
COPY --from=BUILD /app/app .

CMD ["./app"]
