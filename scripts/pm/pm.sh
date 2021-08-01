#!/bin/bash

pm() {
	l=$(devnv list | jq -r '.[].name' | sort | fzf -1 --ansi --preview 'devnv get {}' --info=inline --border -q "$1")
	source <(devnv cd $l)
}
