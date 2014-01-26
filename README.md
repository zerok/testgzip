# TestGZIP

This is a simple helper script for checking if a webserver sends a response
for a set of URLs gzipped/deflated.

The order of the output messages doesn't necessarily represents the order
of the commandline arguments, but instead mostly relies on the speed in which
each server responded.

**Please note:** This was created as part of my #golang learning process in
which I ported some of my rather old helper scripts from Python, Ruby, Bash,
etc. to Go. Because of that I tried to cram as much stuff as possible into
this which normally perhaps wouldn't necessarily make sense.

[![baby-gopher](https://raw2.github.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)
