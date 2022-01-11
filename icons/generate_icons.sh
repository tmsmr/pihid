#!/usr/bin/env bash

ICONS="
solid/power-off.svg,Power,power.go
solid/fan.svg,Fan,fan.go
solid/thermometer-half.svg,Temperature,temperature.go
solid/microchip.svg,Cpu,cpu.go
solid/hdd.svg,Hdd,hdd.go
solid/network-wired.svg,Network,network.go
solid/bong.svg,Ondemand,ondemand.go
solid/bolt.svg,Performance,performance.go
"

git submodule update --init
docker build --pull -t gen-icon-src generator

for line in $ICONS; do
  IFS=, read -r -a icon <<< "$line"
  echo "generating ${icon[2]}"
  docker run --rm -ti -v "$(pwd)/Font-Awesome/svgs:/svgs" gen-icon-src:latest /gen_icon_src.py "/svgs/${icon[0]}" "${icon[1]}" > "${icon[2]}"
done
