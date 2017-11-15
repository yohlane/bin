#!/bin/bash
function join_by { local d="$1"; shift; echo -n "$1"; shift; printf "%s" "${@/#/$d}"; }

folders=("$( ls -aR ~/ | grep -i ".*:"  | grep -i "$(join_by ".*" "$@")" | sed s'/://' )");

min=999
newfolders=()
while read -r line; do
  depth="$(echo $line | awk -F/ '{print NF-1}')"
  if [ "$depth" -lt "$min" ]; then
    min=$depth
    newfolders=()
  fi
  if [ "$depth" -eq "$min" ]; then
    newfolders+=("$line")
  fi
done <<< "$folders"

if [ "${#newfolders[@]}" -eq "1" ]; then
    cd "${newfolders[1]}"
else
    i=0
    for line in "${newfolders[@]}"; do
        i=$(( $i+1 ))
        echo "$i - $line"
    done
    echo "------------------"
    read -r c
    cd "${newfolders[$c]}"
fi