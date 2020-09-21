FROM golang:alpine AS build

WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags='-s -w' -o /bin/goindex-server ./cmd/goindex-server

FROM scratch

# sqlite?
COPY --from=build /etc/services /etc/services
COPY --from=build /etc/protocols /etc/protocols
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /bin/goindex-server /bin/

ENTRYPOINT ["/bin/goindex-server"]
