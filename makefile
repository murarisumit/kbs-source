SHELL := /usr/bin/env bash
run:
	hugo serve -D --bind 192.168.78.100 --baseURL 192.168.78.100/kbs

build:
	hugo

