#!/bin/bash
[ "$(id -u -n)" == "root" ] && SUDO="" || SUDO="sudo"

if [ ! -e "/etc/alpine-release" ]; then
    [[ ((! -x "/usr/local/bin/phantomjs")) && (( -z $(which phantomjs)  )) ]]     && ( echo "Installing phantomjs ..."     && ${SUDO} npm install -g phantomjs-prebuilt@^2.1.14 )     || ( echo "Found phantomjs in PATH ..." )
    else
    echo "PhantomJS is not compatible with alpine-release"
fi
