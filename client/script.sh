#!/bin/sh

curl -H "Content-Type: application/json" -d '{ "author": "occupytheweb", "title": "Network Basics for Hackers"}' -X POST http://goapi.apps.svc.cluster.local:50051/books
