#!/bin/sh

echo "Copying generated Go code to corresponding package..."

for N in ./gen/*/; do
  if [ -d "${N}" ]; then
    name=$(basename "$N")

    for V in "$N"/*/; do
      if [ -d "${V}" ]; then

        version=$(basename "$V")
        source_dir="./gen/$name/$version"

        pkg_root="./internal/$name"
        if ! [ -d "$pkg_root" ]; then
          echo "no pkg root dir $pkg_root - might need to create the package first"
          exit 1
        fi

        pb_root="$pkg_root/pb"
        if ! [ -d "$pb_root" ]; then
          mkdir "$pb_root"
        fi

        ver_root="$pb_root/$version"
        if ! [ -d "$ver_root" ]; then
          mkdir "$ver_root"
        fi

        echo "copying all from $source_dir -> $ver_root"
        cp -r "$source_dir"/*.go "$ver_root"

      fi
    done

  fi
done
