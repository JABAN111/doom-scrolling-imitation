container_runtime := $(shell which docker || which podman)
# TODO: потестить на подмане

tools:
	make -C doom-scrolling tools

lint: tools
	make -C doom-scrolling lint

up:
	${container_runtime} compose up --build -d

start-all:
	rm -fr ./volumes
	${container_runtime} compose down -v
	${container_runtime} compose up -d
	@echo wait cluster to start && sleep 10
	make -C ./configs/couchbase cluster-up

down:
	${container_runtime} compose down

run-tests: 
	${container_runtime} run --rm --network=host tests:latest

# TODO тесты нужно в ci добавить
tests:
	make down
	make up
	@echo wait cluster to start && sleep 2
	make run-tests
	make down
	@echo "test finished"
