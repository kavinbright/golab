#! /bin/bash

git checkout $1
git fetch
git merge origin/$1
git merge origin/$2
git push origin $1
git checkout $2
