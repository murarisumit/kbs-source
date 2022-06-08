SHELL := /usr/bin/env bash
local:
	hugo serve -p 1313 --buildDrafts --navigateToChanged

local-nodrafts:
	hugo serve -p 1313 --navigateToChanged

vbox:
	hugo serve -D --bind 192.168.78.100 --baseURL 192.168.78.100/kbs

build: 
	hugo

clean:
	rm -rf public
	
update-submodule:
	git submodule update --init --recursive

tmux:
	tmuxp load ./tmuxp.yaml
