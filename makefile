all:
	docker build -f Dockerfile -t orders:2.0 .

compose-up:
	cd ./infra/compose/ && docker compose -f docker_compose.yaml -p compose up -d

compose-down:
	cd ./infra/compose/ && docker compose -f docker_compose.yaml down

compose-logs:
	cd ./infra/compose/ && docker compose -f docker_compose.yaml logs