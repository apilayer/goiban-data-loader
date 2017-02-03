MIGRATE_BIN = $(GOPATH)/bin/goose
DATABASE_URL ?= root:root@tcp(localhost:3306)/goiban?charset=utf8

$(MIGRATE_BIN):
	go get bitbucket.org/liamstask/goose/cmd/goose

goiban-data-loader:
	go build

migrate: $(MIGRATE_BIN)
	$(MIGRATE_BIN) up

load: goiban-data-loader
	./goiban-data-loader bundesbank "$(DATABASE_URL)"
	./goiban-data-loader lu "$(DATABASE_URL)"
	./goiban-data-loader nl "$(DATABASE_URL)"
	./goiban-data-loader nbb "$(DATABASE_URL)"