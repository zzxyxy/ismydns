FROM debian:bookworm

COPY entry.sh /
COPY ismydns /

RUN chmod +x /entry.sh
RUN chmod +x /ismydns

ENTRYPOINT /entry.sh
