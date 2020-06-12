generate:
	go run github.com/99designs/gqlgen generate .

gqlinit:
	go run github.com/99designs/gqlgen init

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/graphql handlers/graphql/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/migrate handlers/migrate/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/pre-signup handlers/pre-signup/main.go

run:
	# build the package
	make build
	cp serverless.default.yaml serverless.yml
	# generate sam template for local development
	make run-local

deploy:
	make build

	cp serverless.default.yaml serverless.yml

	npx serverless deploy --stage $(STAGE) -v --region $(REGION)

	#  TODO will do it later make migrate

remove:
	sls remove --stage $(STAGE) -v --region $(REGION)

migrate:
	sls invoke -f migrate --stage $(STAGE) -v --region $(REGION)

migrate-deploy:
	rm -rf bin/
	env GOOS=linux go build -ldflags="-s -w" -o bin/migrate handlers/migrate/main.go
	sls deploy -f migrate --stage $(STAGE) -v --region $(REGION)

run-local-db:
	# run database locally
	docker-compose -f  docker-compose.local.yaml up --build

run-local-server:
	# run sam server locally
	#sam local start-api
	sls offline start --useDocker --host $(DATABASE_HOSTNAME)

migrate-local:
	sls invoke local -f migrate --stage dev

run-local: run-local-db run-local-server migrate-local

lint:
	golangci-lint run

dependency:
	cd di && go run github.com/google/wire/cmd/wire
