all:
	docker build -f Dockerfile -t orders:2.0 .

up:
	cd ./infra/compose/ && docker compose -f docker_compose.yaml -p compose up -d

down:
	cd ./infra/compose/ && docker compose -f docker_compose.yaml down

logs:
	cd ./infra/compose/ && docker compose -f docker_compose.yaml logs