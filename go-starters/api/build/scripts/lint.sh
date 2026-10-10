#!/usr/bin/env bash

set -e
set -x

if golangci-lint run; then \
		echo "golangci-lint ran successfully"; \
	else \
		echo "golangci-lint failed to lint project files"; \
		read -rp "Would you like to try running it again with --fix (y/n)? " selection; \
		if [ "$selection" = "y" ]; then \
			echo "Running golangci-lint with --fix..."; \
			if golangci-lint run --fix; then \
				echo "golangci-lint ran successfully with --fix"; \
			else \
				echo "golangci-lint still failed to lint project files."; \
			fi; \
		else \
			echo "Skipping fix attempt. Please review the linting errors."; \
		fi; \
	fi
