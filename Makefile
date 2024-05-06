backend_dev:
	air -c .air.toml

frontend_dev:
	cd frontend && npm run dev

frontend_build:
	cd frontend && npm run build