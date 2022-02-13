FROM golang:1.17.4-alpine3.15 AS build

WORKDIR /src

ENV CGO_ENABLED=0 GOOS=linux
ENV GO_PROXY=https://proxy.golang.org,direct

COPY . .

RUN GOPROXY=${GO_PROXY} go mod download && \
    go build -o /out/app ./001/main.go

## -----------------------------------------------
## -----------------------------------------------

FROM alpine:3.15 AS runner


COPY --from=build /out/app /

EXPOSE 80 443 8080

ENTRYPOINT [ "/app" ]
