#!/bin/sh

set -ex
cd `dirname $0`

ISUCON_DB_HOST=${ISUCON_DB_HOST:-127.0.0.1}
ISUCON_DB_PORT=${ISUCON_DB_PORT:-3306}
ISUCON_DB_USER=${ISUCON_DB_USER:-isucon}
ISUCON_DB_PASSWORD=${ISUCON_DB_PASSWORD:-isucon}
ISUCON_DB_NAME=${ISUCON_DB_NAME:-isuports}

# MySQLを初期化
mysql -u"$ISUCON_DB_USER" \
		-p"$ISUCON_DB_PASSWORD" \
		--host "$ISUCON_DB_HOST" \
		--port "$ISUCON_DB_PORT" \
		"$ISUCON_DB_NAME" < init.sql

for id in $(seq 2 100)
do
	echo "
		DROP DATABASE IF EXISTS isuports_tenant_$id;
		CREATE DATABASE isuports_tenant_$id;
		USE isuports_tenant_$id;
		SET autocommit = 0;
		BEGIN;
		$(cat tenant/10_schema.sql)
		INSERT INTO competition SELECT * FROM base_isuports_tenant_$id.competition;
		INSERT INTO player SELECT * FROM base_isuports_tenant_$id.player;
		INSERT INTO player_score SELECT * FROM base_isuports_tenant_$id.player_score;
		COMMIT;
	" | mysql -u"$ISUCON_DB_USER" \
			-p"$ISUCON_DB_PASSWORD" \
			--host "$ISUCON_DB_HOST" \
			--port "$ISUCON_DB_PORT" &
done

wait
