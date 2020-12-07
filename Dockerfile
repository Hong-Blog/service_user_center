FROM golang:1.15 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -tags netgo -o app .

RUN mkdir publish && cp app publish && \
    cp -r docs publish


FROM alpine
WORKDIR /app
COPY --from=builder /app/publish .
ENV GIN_MODE=release \
    PORT=8081

EXPOSE 8081
ENTRYPOINT ["./app"]