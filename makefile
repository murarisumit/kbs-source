SHELL := /usr/bin/env bash
local:
	hugo serve -p 1313

vbox:
	hugo serve -D --bind 192.168.78.100 --baseURL 192.168.78.100/kbs

build:
	hugo

