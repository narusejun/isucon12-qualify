echo "
	DROP DATABASE IF EXISTS isuports_tenant_$id;
	CREATE DATABASE isuports_tenant_$id;
	USE isuports_tenant_$id;
	BEGIN;
	$(cat tenant/10_schema.sql)
	INSERT INTO competition SELECT * FROM base_isuports_tenant_$id.competition;
	INSERT INTO player SELECT * FROM base_isuports_tenant_$id.player;
	INSERT INTO player_score SELECT * FROM base_isuports_tenant_$id.player_score;
	$(cat tenant/11_index.sql)
	COMMIT;
" | mysql -u"$ISUCON_DB_USER" \
		-p"$ISUCON_DB_PASSWORD" \
		--host "$ISUCON_DB_HOST" \
		--port "$ISUCON_DB_PORT"
