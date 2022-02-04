FROM golang:1.15-alpine as dev

WORKDIR /work

FROM golang:1.15-alpine as build

WORKDIR /cmd
COPY ./cmd/* /cmd/
RUN go build -o parser

FROM alpine as runtime 
COPY --from=build /cmd/parser /
CMD ./parser