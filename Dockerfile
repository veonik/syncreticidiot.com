FROM golang:buster AS build

WORKDIR /build

COPY . .

RUN go get -v ./...

RUN go build -o sidiot ./cmd/sidiot/*.go


FROM debian:buster-slim

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y ca-certificates

COPY --from=build /build/sidiot /bin/sidiot

CMD /bin/sidiot
