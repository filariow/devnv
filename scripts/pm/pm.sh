#!/bin/zsh

_pm() {
    l=$(devnv list | jq -r '.[] | .name +":"+ .folder' | sort | xargs -I % echo "'%'" | tr "\n" " ")
    _describe 'command' "($l)"
}

if [ "$funcstack[1]" = "_pm" ]; then
    _pm
else 
    compdef '_pm' pm
fi

pm() {
	l=$(devnv list | jq -r '.[].name' | sort | fzf -1 --ansi --preview 'devnv get {}' --info=inline --border -q "$1")
	source <(devnv cd $l)
}

