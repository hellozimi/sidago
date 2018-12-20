#!/usr/bin/env bash

VERSION=0.0.1
TIME=$(date)
sed -i "s/Version = \"[^\"]*\"/Version = \"${VERSION}\"" sidaversion/version.go