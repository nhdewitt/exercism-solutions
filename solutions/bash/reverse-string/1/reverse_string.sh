#!/usr/bin/env bash

reverse() {
    local s="$1"
    local rev=""
    local i

    for (( i=${#s}-1; i>=0; i-- )); do
        rev+="${s:$i:1}"
    done

    printf "%s\n" "$rev"
}

main () {
    reverse "$1"
}

main "$@"