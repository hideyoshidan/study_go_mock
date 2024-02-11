pull-mockery:
	docker pull vektra/mockery

mock:
	docker run -v "$PWD":/src -w /src vektra/mockery --all
