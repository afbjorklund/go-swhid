#!/bin/sh

checksum() {
  pigz -d <"$1" | sha1sum - | sed -e "s,-,$1,"
}

for f; do
  checksum "$f"
done
