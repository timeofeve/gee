#!/bin/bash
cat words.txt | tr -s ' ' '\n'|sort|uniq -c|awk '{print $2" "$1}'|sort -nrk2