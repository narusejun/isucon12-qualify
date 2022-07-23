for i in (seq 1 100)
	echo "SET autocommit = 0;" > $i.sql
	~/webapp/sql/sqlite3-to-sql $i.db >> $i.sql
end

for id in $(seq 1 100)
do
	echo "
		DROP DATABASE IF EXISTS base_isuports_tenant_$id;
		CREATE DATABASE base_isuports_tenant_$id;
	" | mysql -uroot -proot
	mysql -uroot -proot \
		"base_isuports_tenant_$id" < "$HOME/initial_data/$id.sql"
done
