.PHONY: dev stop test test-backend test-frontend test-e2e

dev:
	docker compose up --build -d

stop:
	docker compose down

test: test-backend test-frontend

test-backend:
	cd backend && go test ./... -v

test-frontend:
	cd frontend && npm test

test-e2e:
	cd e2e && npx playwright test