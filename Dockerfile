FROM alpine:3.10
COPY ./config /config
CMD ["/config"]
