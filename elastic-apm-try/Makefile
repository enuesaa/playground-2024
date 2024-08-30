run:
	ELASTIC_APM_SERVER_URL=https://localhost:8200 ELASTIC_APM_SECRET_TOKEN=secrettokengoeshere ELASTIC_APM_SERVER_CERT=../../elastdocker/secrets/certs/apm-server/apm-server.crt  CGO_LDFLAGS="-Wl,-no_warn_duplicate_libraries" go run .
