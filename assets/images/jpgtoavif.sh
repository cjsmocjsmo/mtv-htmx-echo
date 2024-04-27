#!/bin/sh
# imagemagick must be installed


for image in *.jpg; do
    convert "$image" -resize 539x800 "$image"
    convert "$image" "${image%.*}.avif"
    #rm "$image";
done
