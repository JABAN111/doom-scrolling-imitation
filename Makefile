container_runtime := $(shell which docker || which podman)
# TODO: потестить на подмане

tools:
	make -C doom-scrolling tools

lint: tools
	make -C doom-scrolling lint

up:
	${container_runtime} compose up --build -d

tart-all:
	rm -fr ./volumes
	${container_runtime} compose down -v
	make -C ./configs/clickhouse config
	${container_runtime} compose up -d
	@echo wait luster to start && sleep 15
	make -C ./configs/couchbase luster-up

start-all:
	rm -fr ./volumes
	${container_runtime} compose down -v
	chmod 640 configs/neof4j/neo4j.conf
	make -C ./configs/clickhouse config
	${container_runtime} compose up -d
	@echo wait sluster to start && sleep 15
	make -C ./configs/couchbase cluster-up
	sleep 3
	${container_runtime} compose up -d --build doom-scrolling

down:
	${container_runtime} compose down

run-tests: 
	${container_runtime} run --rm --network=host tests:latest

tests:
	make down
	make up
	@echo wait cluster to start && sleep 2
	make run-tests
	make down
	@echo "test finished"
