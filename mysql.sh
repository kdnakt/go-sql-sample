docker run --rm -d -e MYSQL_ALLOW_EMPTY_PASSWORD=yes \
  -p $MYSQL_PORT:3306 --name mysql_tmp mysql:5.7