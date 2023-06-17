# iBook Backend

```shell
go test -v ./test/
```

```shell
docker build -t ibook-backend:latest .
```

```shell
docker run -d \
  --restart always \
  --publish 18080:8080 \
  -e MYSQL_HOST=localhost \
  -e MYSQL_PORT=3306 \
  -e MYSQL_USERNAME=user \
  -e MYSQL_PASSWORD=secret \
  -e MYSQL_DB=db \
  -e JWT_KEY=key \
  ibook-backend:latest
```
