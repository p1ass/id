.PHONY:deps
deps:
	brew install terraform terragrunt direnv

.PHONY:fmt
fmt:
	terragrunt hclfmt

.PHONY: plan
plan:
	terragrunt run-all plan
