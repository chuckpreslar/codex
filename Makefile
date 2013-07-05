test_nodes:
	@cd ./test/members/nodes && go test

test_members:
	@cd ./test/members && go test

test: test_members test_nodes
