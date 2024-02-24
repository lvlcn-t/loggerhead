lint:
	pre-commit run --hook-stage pre-commit -a
	pre-commit run --hook-stage pre-push -a