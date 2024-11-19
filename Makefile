#extract(year from date)=2025 Change these variables as necessary.
MAIN_PACKAGE_PATH := ./
BINARY_NAME := main

## TODO: Documenting

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 $(go list ./... | grep -v /api/)
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
# for now, remove integration test
# go test -race -buildvcs -vet=off ./...


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## generate: generate all codes
.PHONY: generate
generate:
	go generate -v ./...

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## build: build the application
.PHONY: build
build: #generate
	# Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run svc=$1 env=$2: run the application
.PHONY: run
run: build
	/tmp/bin/${BINARY_NAME} ${svc} --env=${env}

## run/live svc=$1 env=$2: run the application with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "/tmp/bin/${BINARY_NAME} ${svc} --env=${env}" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"


# ==================================================================================== #
# SQL MIGRATIONS
# ==================================================================================== #

## migrations/up: apply all up databases migrations
.PHONY: migrations/up
migrations/up: confirm user/migrations/up 

## migrations/down: apply all down databases migrations
.PHONY: migrations/down
migrations/down: user/migrations/down 

# ====================================== #
# SQL MIGRATIONS - USER
# ====================================== #

## user/migrations/new name=$1: create a new database migration
.PHONY: user/migrations/new
user/migrations/new:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest create -seq -ext=.sql -dir=./internal/repositories/userrepo/migrations ${name}

## user/migrations/up: apply all up database migrations
.PHONY: user/migrations/up
user/migrations/up:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./internal/repositories/userrepo/migrations -database="postgres://${USER_DB_DSN}" up

## user/migrations/down: apply all down database migrations
.PHONY: user/migrations/down
user/migrations/down:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./internal/repositories/userrepo/migrations -database="postgres://${USER_DB_DSN}" down

## user/migrations/goto version=$1: migrate to a specific version number
.PHONY: user/migrations/goto
user/migrations/goto:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./internal/repositories/userrepo/migrations -database="postgres://${USER_DB_DSN}" goto ${version}

## user/migrations/force version=$1: force database migration
.PHONY: user/migrations/force
user/migrations/force:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./internal/repositories/userrepo/migrations -database="postgres://${USER_DB_DSN}" force ${version}

## user/migrations/version: print the current in-use migration version
.PHONY: user/migrations/version
user/migrations/version:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./internal/repositories/userrepo/migrations -database="postgres://${USER_DB_DSN}" version

# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## push: push changes to the remote Git repository
.PHONY: push
push: tidy audit no-dirty
	git push

## development/build: build and push development image to artifact registry
.PHONY: development/build
development/build: confirm tidy no-dirty audit
	docker buildx build --platform=linux/amd64 -f Dockerfile -t region-docker.pkg.dev/path/of/image . && \
	docker push region-docker.pkg.dev/project-path

## development/deploy/auth: deploy the auth service to development
.PHONY: development/deploy/auth
development/deploy/auth: development/build
	gcloud --project=project-name run deploy service-name --image=asia-southeast1-docker.pkg.dev/path/of/image:latest --region=asia-southeast1

## development/deploy/core: deploy the core service to development
.PHONY: development/deploy/core
development/deploy/core: development/build
	gcloud --project=project-name run deploy service-name --image=asia-southeast1-docker.pkg.dev/path/of/image:latest --region=asia-southeast1

## development/deploy: deploy the all services to development
.PHONY: development/deploy
development/deploy: development/build
	gcloud --project=project-name run deploy service-name --image=asia-southeast1-docker.pkg.dev/path/of/image:latest --region=asia-southeast1 && \
	gcloud --project=project-name run deploy service-name --image=asia-southeast1-docker.pkg.dev/path/of/image:latest --region=asia-southeast1 && \
