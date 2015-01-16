gothue
===

This is the [Thue Programming Language](http://catseye.tc/node/Thue) processer written in Go.


What is Thue?
---

Thue is the esoteric programming language designed by Chris Pressey in 2000.
Thue is based on a string rewriting system, it is also turing complete!

Also see the [article of Wikipedia](http://en.wikipedia.org/wiki/Thue_%28programming_language%29).


How to install
---

You can use `go get` to install, following to:

```console
$ go get github.com/MakeNowJust/gothue
```

And if you had set `$PATH` to `$GOPATH/bin`, to be avaiable to `gothue` command.

```console
$ gothue -h
gothue - Thue interpreter written in Go

usage:
  gothue [-s=<seed>] [-lrdh] <file name>

option:
  -l         execute leftmost matches first
  -r         execute rightmost matches first
  -d         debug mode
  -s=<seed>  set a seed (default is current unix time)
  -h         show this help
$ gothue _example/hello.t # what text happen printing???  it is randomly.
```


License
---

This is released under the [MIT License](http://makenowjust.github.io/license/mit?2015).
