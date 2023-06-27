.PHONY: all test stop build down
all			:
		docker-compose up task_app
test		:
		docker-compose up testuserapp testtaskapp
stop		:
		docker-compose stop
build		:
		docker-compose build
down		:
		docker-compose down