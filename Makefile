dev:
	docker compose -f docker-compose.local.yml up

prod:
	docker compose -f docker-compose.prod.yml up

deploy_prod:
	docker --context tp_remote compose -f docker-compose.prod.yml up -d

renew_cert:
	docker --context tp_remote compose -f docker-compose.prod.yml exec nginx certbot renew

generate_cert_on_server:
	docker --context tp_remote compose -f docker-compose.prod.yml run --rm  certbot certonly --webroot --webroot-path /var/www/certbot/ -d api.typograph.kz   