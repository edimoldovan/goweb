#!/bin/bash

# echo 'building public sass/scss...';
# sass static/scss/public.scss public/public.css;
# echo 'minifying public css...';
# minify public/public.css -o public/public.min.css;

echo 'copying javascript files to public folder...';
cp assets/node_modules/flatpickr/dist/flatpickr.min.js public/js;

echo 'building and minifying global css with rust-based postcss...'
cd assets/css-toolchain/; npx parcel-css --bundle -m --nesting src/css/global.css -o ../../public/global.min.css; cd ..; cd ..;

# echo 'minifying style css...';
# minify assets/style.css -o public/style.min.css;

# echo 'minifying blog css...';
# minify assets/blog.css -o public/blog.min.css;

# /usr/local/go/bin/go build -o ./tmp/main .
go build -o ./tmp/main .
