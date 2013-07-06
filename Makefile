test_nodes:
	@cd ./test/nodes && go test -i && go test

test: test_nodes
