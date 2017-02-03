MIGRATE_BIN = $(GOPATH)/bin/goose

$(MIGRATE_BIN):
	go get bitbucket.org/liamstask/goose/cmd/goose

migrate: $(MIGRATE_BIN)
	$(MIGRATE_BIN) up