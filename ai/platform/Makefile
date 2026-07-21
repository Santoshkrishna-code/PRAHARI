.PHONY: install run test lint typecheck docker-up docker-down

install:
	poetry install

run:
	poetry run uvicorn app.main:app --reload --host 127.0.0.1 --port 8000

test:
	poetry run pytest -v --tb=short tests/

lint:
	poetry run ruff check .
	poetry run ruff format --check .

typecheck:
	poetry run mypy app

docker-up:
	docker compose up --build -d

docker-down:
	docker compose down -v
