FROM golang:1.17.0-alpine3.14  AS builder

RUN mkdir -p /src/admin-alarm/internal
WORKDIR /src/admin-alarm/

ENV GO111MODULE=on
COPY go.* /src/admin-alarm/
COPY main.go /src/admin-alarm/
COPY internal /src/admin-alarm/internal
RUN ls -als /src/
RUN CGO_ENABLED=0 go build -o /bin/admin-alarm

FROM scratch
COPY --from=builder /bin/admin-alarm /bin/admin-alarm
ENTRYPOINT ["/bin/admin-alarm"]
