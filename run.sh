#!/bin/bash

go build -o bin/app src/*.go
cd bin && ./app
