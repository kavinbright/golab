#! /bin/bash
echo "start checkout fetch merge push and checkout back..."
git checkout master
git fetch
git merge origin/master
git merge origin/company
git push origin master
git checkout company
echo "end"
