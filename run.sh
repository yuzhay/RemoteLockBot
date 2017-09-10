#!/bin/bash

go build -o bin/rlbot src/*.go
cd bin && ./rlbot
