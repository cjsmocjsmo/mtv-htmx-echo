for f in /media/charliepi/HD/MTVSERVER/mtv_svelte/src/lib/images/*.jpg; do
cwebp -q 95 -resize 600 0 "$f" -o "${f%.jpg}.webp"
done
