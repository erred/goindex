FROM golang:alpine AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags='-s -w' -o /bin/goindex-server ./cmd/goindex-server

FROM scratch

# sqlite?
COPY --from=build /etc/services /etc/services
COPY --from=build /etc/protocols /etc/protocols

COPY --from=build /bin/goindex-server /bin/

ENTRYPOINT ["/bin/goindex-server"]
