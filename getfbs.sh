#! /usr/bin/env sh

curl -O https://raw.githubusercontent.com/tensorflow/tensorflow/master/tensorflow/lite/schema/schema.fbs
flatc -g schema.fbs
