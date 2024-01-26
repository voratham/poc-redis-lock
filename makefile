some-redis:
	docker run -d  --name some-redis --network shared_network -p 6379:6379 -e ALLOW_EMPTY_PASSWORD=yes bitnami/redis:latest


new-network:
	docker network create shared_network