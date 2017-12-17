NAME = computorv1

RED=\033[1;31m
GREEN=\033[1;32m
NC=\033[0m

.SILENT:

all: $(NAME)

$(NAME):
	go build src/*.go
	@printf "$(GREEN)[✓]$(NC) Executable $(NAME) ready!\n"

clean:
	@rm -rf $(NAME)
	@printf "$(RED)[-]$(NC) Executable $(NAME) deleted\n"

re: clean all

.PHONY: re clean all