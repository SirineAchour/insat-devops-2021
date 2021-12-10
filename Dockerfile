FROM golang:1.16 as builder

COPY main.go main.go

RUN go get ./...
RUN CGO_ENABLED=0 go build -o /mybinary main.go


FROM scratch
COPY --from=builder /mybinary /app
CMD ["/app"]