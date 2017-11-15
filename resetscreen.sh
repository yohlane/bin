#!/bin/sh
xrandr --listactivemonitors | awk -F' ' '{print $2}' | sed 's/\+//' | tail -n +3 | xargs -I R xrandr --output R --off
