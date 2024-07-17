#!/bin/bash
newman run $1 --insecure --reporters htmlextra --reporter-htmlextra-export $2

