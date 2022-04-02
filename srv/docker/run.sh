docker run -it -d -v ~/:/source -p 100:80 --name srv \
--env MYSQL_ADDRESS="172.17.0.2:3306" \
--env MYSQL_NET="tcp" \
--env MYSQL_USER="root" \
--env MYSQL_PASSWORD="my-secret-pw" \
--env MYSQL_DBNAME="dev" \
srv
