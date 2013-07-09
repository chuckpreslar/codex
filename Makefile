test_nodes:
	@cd ./test/nodes && go test -i && go test

test_visitors:
	@cd ./test/visitors && go test -i && go test

test_managers:
	@cd ./test/managers && go test -i && go test

test: test_nodes test_visitors test_managers
