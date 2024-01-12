#!/bin/sh

RQLITE_URL="http://${RQLITE_HOST}:${RQLITE_PORT}"
CHECK_SQL="SELECT COUNT(*) FROM user;"

TABLE_EMPTY=$(curl -u "${RQLITE_USER}":"${RQLITE_PASSWORD}" -s -G "${RQLITE_URL}/db/query" --data-urlencode "q=${CHECK_SQL}" | jq '.results[0].values[0][0]' -r)

if echo "$TABLE_EMPTY" | grep -qE '^[0-9]+$' ; then
    if [ "$TABLE_EMPTY" -gt 0 ]; then
        echo "Database not reachable or not untouched. No changes will be made."
        exit 1
    fi
fi
echo "Adding test data to database:"
curl -u "${RQLITE_USER}":"${RQLITE_PASSWORD}" -s -XPOST "${RQLITE_URL}/db/load" -H "Content-type: text/plain" --data-binary @init.sql
exit 0