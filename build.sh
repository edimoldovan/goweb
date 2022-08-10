#!/bin/bash

# echo 'building public sass/scss...';
# sass static/scss/public.scss public/public.css;
# echo 'minifying public css...';
# minify public/public.css -o public/public.min.css;

# echo 'building private sass/scss...';
# sass static/scss/private.scss public/private.css;
# echo 'minifying private css...';
# minify public/private.css -o public/private.min.css;

echo 'minifying style css...';
minify assets/style.css -o public/style.min.css;

echo 'minifying blog css...';
minify assets/blog.css -o public/blog.min.css;

# /usr/local/go/bin/go build -o ./tmp/main .
go build -o ./tmp/main .
