FROM alpine:latest

RUN apk --no-cache add curl jq

COPY ./init.sql ./init.sql
COPY ./check_and_insert.sh ./check_and_insert.sh

RUN chmod +x ./check_and_insert.sh
CMD ["./check_and_insert.sh"]