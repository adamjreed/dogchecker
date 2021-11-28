.PHONY: up down run logs deploy

up:
	docker compose up --build -d

down:
	docker compose down

run:
	curl -XPOST "http://localhost:8080/2015-03-31/functions/function/invocations" -d '{}'

logs:
	docker compose logs lambda -f

deploy:
	//do some deploy stuff