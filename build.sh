#!/bin/bash

# echo 'building public sass/scss...';
# sass static/scss/public.scss public/public.css;
# echo 'minifying public css...';
# minify public/public.css -o public/public.min.css;

echo 'copying javascript files to public folder...';
cp ./assets/node_modules/flatpickr/dist/flatpickr.min.css ./public/js;

echo 'minifying style css...';
minify assets/style.css -o public/style.min.css;

echo 'minifying blog css...';
minify assets/blog.css -o public/blog.min.css;

# /usr/local/go/bin/go build -o ./tmp/main .
go build -o ./tmp/main .
