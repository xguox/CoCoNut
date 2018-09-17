## CoCoNut

##### 主要依赖

- Go 1.11
- [https://github.com/oxequa/realize](realize)
- ruby (db migration)
- `db/` [standalone_migrations](https://github.com/thuss/standalone-migrations) gem install standalone_migrations

##### 运行

```
rake db:create
rake db:migrate
go mod tidy
realize start --run
```

##### TODO

- [x] CATEGORY
- [x] PRODUCT
- [x] USER & AUTH
- [x] TAGGING
- [ ] OPTIONS
- [ ] VARIANT
- [ ] CUSTOMER
- [ ] ADDRESS
- [ ] COLLECTION
- [ ] ORDER
- [ ] ORDER_ITEM
- [ ] PICTURES
- [ ] WISH_LIST
- [ ] GROUPON, COUPON
- [ ] TESTING
