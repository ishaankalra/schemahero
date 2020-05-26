SHELL := /bin/bash
DATABASE_IMAGE_NAME := schemahero/database
DATABASE_CONTAINER_NAME := schemahero-database
DRIVER := mysql
URI := schemahero:password@tcp(localhost:13306)/schemahero?tls=false

.PHONY: run
run:
	# Fixtures
	docker pull mysql:$(MYSQL_VERSION)
	docker build -t $(DATABASE_IMAGE_NAME) .
	@-docker rm -f $(DATABASE_CONTAINER_NAME) > /dev/null 2>&1 ||:
	docker run -p 13306:3306 --rm -d --name $(DATABASE_CONTAINER_NAME) $(DATABASE_IMAGE_NAME)
	while ! docker exec $(DATABASE_CONTAINER_NAME) mysqladmin ping --silent; do sleep 1; done
	@sleep 10

	# Plan
	../../../../bin/schemahero plan --driver=$(DRIVER) --uri="$(URI)" --spec-file $(SPEC_FILE) > out.sql

	# Verify
	@echo Verifying results for $(TEST_NAME)
	diff -B expect.sql out.sql

	# Cleanup
	@-sleep 5
	# rm ./out.sql
	@-docker rm -f $(DATABASE_CONTAINER_NAME)
