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

__choices() {
    l=`devnv list | jq -r '.[].name' | sort`
    o=`echo $l | grep -E "^$1$"`
    if [ ! -z $o ]; then
        echo "$o"
        return
    fi
    echo "$l"
}

pm() {
    l=$(__choices "$1" | sort | fzf -1 --ansi --preview 'devnv get {} && echo "" && devnv script {}' --info=inline --border -q "$1")
	source <(devnv cd $l)
}

