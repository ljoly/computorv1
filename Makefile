# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: ljoly <ljoly@student.42.fr>                +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2017/12/21 14:19:03 by ljoly             #+#    #+#              #
#    Updated: 2017/12/21 14:21:47 by ljoly            ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

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
	@rm -f $(NAME)
	@printf "$(RED)[-]$(NC) Executable $(NAME) deleted\n"

re: clean all

.PHONY: re clean all