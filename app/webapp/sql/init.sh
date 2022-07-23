#!/bin/sh

set -ex
cd `dirname $0`

export ISUCON_DB_HOST=${ISUCON_DB_HOST:-127.0.0.1}
export ISUCON_DB_PORT=${ISUCON_DB_PORT:-3306}
export ISUCON_DB_USER=${ISUCON_DB_USER:-isucon}
export ISUCON_DB_PASSWORD=${ISUCON_DB_PASSWORD:-isucon}
export ISUCON_DB_NAME=${ISUCON_DB_NAME:-isuports}

# MySQLを初期化
mysql -u"$ISUCON_DB_USER" \
		-p"$ISUCON_DB_PASSWORD" \
		--host "$ISUCON_DB_HOST" \
		--port "$ISUCON_DB_PORT" \
		"$ISUCON_DB_NAME" < init.sql &

# export id=1
# sh init_tenant_huge.sh &

for id in $(seq 2 100)
do
	export id=$id
	sh init_tenant.sh &
done

wait
