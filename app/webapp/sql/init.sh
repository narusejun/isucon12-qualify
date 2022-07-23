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

for id in $(seq 1 100)
do
	echo "
		USE base_isuports_tenant_$id;
		UNLOCK TABLES;
		FLUSH TABLES competition FOR EXPORT;
		UNLOCK TABLES;
		FLUSH TABLES player FOR EXPORT;
		UNLOCK TABLES;
		FLUSH TABLES player_score FOR EXPORT;
	" | mysql -u"$ISUCON_DB_USER" \
			-p"$ISUCON_DB_PASSWORD" \
			--host "$ISUCON_DB_HOST" \
			--port "$ISUCON_DB_PORT"

	echo "
		DROP DATABASE IF EXISTS isuports_tenant_$id;
		CREATE DATABASE isuports_tenant_$id;
		USE isuports_tenant_$id;
		$(cat tenant/10_schema.sql)
		ALTER TABLE competition DISCARD TABLESPACE;
		ALTER TABLE player DISCARD TABLESPACE;
		ALTER TABLE player_score DISCARD TABLESPACE;
	" | mysql -u"$ISUCON_DB_USER" \
			-p"$ISUCON_DB_PASSWORD" \
			--host "$ISUCON_DB_HOST" \
			--port "$ISUCON_DB_PORT"

	sudo rm -rf /var/lib/mysql/isuports_tenant_$id
	sudo cp -ra /var/lib/mysql/base_isuports_tenant_$id /var/lib/mysql/isuports_tenant_$id

	echo "
		ALTER TABLE competition IMPORT TABLESPACE;
		ALTER TABLE player IMPORT TABLESPACE;
		ALTER TABLE player_score IMPORT TABLESPACE;
		$(cat tenant/11_index.sql)
	" | mysql -u"$ISUCON_DB_USER" \
			-p"$ISUCON_DB_PASSWORD" \
			--host "$ISUCON_DB_HOST" \
			--port "$ISUCON_DB_PORT" \
			"isuports_tenant_$id"
done
