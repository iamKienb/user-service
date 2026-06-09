GO ?= go

run-all:
	@$(MAKE) -j3 run-command run-query run-worker
