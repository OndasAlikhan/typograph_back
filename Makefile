dev:
	docker compose -f docker-compose.local.yml up

deploy_prod:
	docker --context tp_remote compose -f docker-compose.prod.yml up -d