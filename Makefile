dev:
	docker compose -f docker-compose.local.yml up

prod:
	docker compose -f docker-compose.prod.yml down
	docker compose -f docker-compose.prod.yml up --force-recreate

deploy_prod:
	docker --context tp_remote compose -f docker-compose.prod.yml down
	docker --context tp_remote compose -f docker-compose.prod.yml up -d --force-recreate
