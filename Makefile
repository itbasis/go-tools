go-all: go-update go-all-tests
go-all-tests: go-generate go-lint go-test

go-dependencies:
	.itbasis/builder dependencies

go-update: go-dependencies
	.itbasis/builder update

go-generate:
	.itbasis/builder generate

go-lint:
	.itbasis/builder lint

go-test: go-lint
	.itbasis/builder unit-test
