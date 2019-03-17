## CoCoNut

##### 主要依赖

- Go 1.11+
- [https://github.com/erikdubbelboer/realize](realize)
- ruby (db migration)
- `db/` [standalone_migrations](https://github.com/thuss/standalone-migrations) gem install standalone_migrations

##### 运行

```
cp config/conf.example.yml config/conf.yml
rake db:create
rake db:migrate
go mod tidy
realize start --run
```

##### or Docker

```
cp config/conf.example.yml config/conf.docker.yml
cp docker-compose-example-yml docker-compose.yml

docker-compose up --build
```

##### API docs

http://localhost:9876/swagger/index.html

##### TODO

- [x] CATEGORY
- [x] PRODUCT
- [x] USER & AUTH
- [x] TAGGING
- [x] OPTIONS
- [x] VARIANT
- [ ] CUSTOMER
- [ ] ADDRESS
- [ ] COLLECTION
- [ ] ORDER
- [ ] ORDER_ITEM
- [ ] PICTURES
- [ ] WISH_LIST
- [ ] GROUPON, COUPON
- [x] DOCKERIZED
- [ ] TESTING
