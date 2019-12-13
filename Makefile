build:
	docker-compose build
install:
	docker-compose run --rm tasker bundle
	docker-compose run --rm tasker yarn
tasker:
	docker-compose up tasker
down:
	docker-compose down
migrate:
	docker-compose run --rm tasker rails db:migrate
rspec:
	docker-compose run --rm tasker rspec

