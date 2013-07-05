test_nodes:
	@cd ./test/members/nodes && go test -i && go test

test_members:
	@cd ./test/members && go test -i && go test

test: test_members test_nodes
