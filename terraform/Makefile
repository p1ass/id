.PHONY:deps
deps:
	brew install terraform terragrunt direnv

.PHONY:fmt
fmt:
	terragrunt hclfmt && terraform fmt -recursive

.PHONY: plan
plan:
	terragrunt run-all plan

.PHONY: apply
apply:
	terragrunt run-all apply
