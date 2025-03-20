# TODO: потестить на подмане
container_runtime := $(shell which docker || which podman)


$(info using ${container_runtime})

up:
	${container_runtime} compose up --build -d

down:
	${container_runtime} compose down

run-tests: 
	${container_runtime} run --rm --network=host tests:latest

# test: TODO: если хватит времени на интеграционные тесты
# 	make down
# 	make up
# 	make run-tests
# 	make down
# 	@echo "test finished"

