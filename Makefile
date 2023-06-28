.PHONY: all test stop build down rebuild clean
all			:
		-docker-compose up task_app user_app db migrate
build		:
		-docker-compose build task_app user_app db migrate
down		:
		-docker-compose down
test		:
		-docker-compose up testuserapp testtaskapp
stop		:
		-docker-compose stop
retest		:
		-docker-compose build testuserapp testtaskapp
rebuild		:	clean all

clean		:	down
		-docker image rm mytest-user_app mytest-task_app