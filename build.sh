#!/bin/bash

echo 'copying javascript files to public folder...';
cp assets/node_modules/flatpickr/dist/flatpickr.min.js public/js;

cp assets/node_modules/solid-js/dist/solid.js public/js/solid;
cp assets/node_modules/solid-js/html/dist/html.js public/js/solid;
cp assets/node_modules/solid-js/web/dist/web.js public/js/solid;

cp assets/js/web-component.js public/js/;
cp assets/js/is-land.js public/js/;

echo 'copying fonts...'
cp assets/fonts/* public/fonts/

go build -o ./tmp/main .
