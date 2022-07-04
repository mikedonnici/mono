#!/bin/sh

echo "Copying generated python to corresponding package..."

for N in ./gen/*/; do
  if [ -d "${N}" ]; then
    name=$(basename "$N")

    for V in "$N"/*/; do
      if [ -d "${V}" ]; then

        version=$(basename "$V")
        pkg="$name$version"
        source_dir="./gen/$name/$version"
        target_dir="./pkg/$pkg/$pkg"

        if [ -d "$target_dir" ]; then
          echo "copying all from $source_dir -> $target_dir"
          cp -r "$source_dir"/*.py "$target_dir"
        else
          echo "no target dir $target_dir - might need to create the python package first"
        fi

      fi
    done

  fi
done
