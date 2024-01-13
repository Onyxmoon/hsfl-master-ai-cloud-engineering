FROM alpine:latest

RUN apk --no-cache add curl jq

COPY --chmod=755 init.sql /scripts/
COPY --chmod=755 check_and_insert.sh /scripts/

RUN chmod +x /scripts/check_and_insert.sh
CMD ["/scripts/check_and_insert.sh"]