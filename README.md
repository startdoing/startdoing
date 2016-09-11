# StartDoing

StartDoing is a simple project and issue tracking system.  It should
have all the basic features required for a small to medium sized team.
Setting up the system for production usage will be very easy.  These
are the key technologies used in this project:

- [Go](https://golang.org/)
- [EmberJS](http://emberjs.com/)
- [BoltDB](https://github.com/boltdb/bolt)

Visit the [wiki for
documentation](https://github.com/startdoing/startdoing/wiki) to read more
about the usage, architecture etc.,


## Development

If you are interested to contribute to this project, please follow the
instruction given here.

1. Install [Go
compiler](http://muthukadan.net/golang/an-introduction-to-go-programming.html)
2. Run this command to get source code: `go get github.com/startdoing/startdoing`
   ([Git](http://git-scm.com/) is a prerequisite)

The code will be available under `$GOPATH/src/github.com/startdoing/startdoing`
You will also a get an executable binary under `$GOPATH/bin/startdoing`.

You can fork the project from
[Github](https://github.com/startdoing/startdoing) and push your
changes there.  Later you can send pull request with your changes.
Before sending the pull request, make sure the tests are running
locally using this command:

    go test ./...

You cal also run Go lint tools using [gometalinter](https://github.com/alecthomas/gometalinter),
to install it run these two commands:

    go get -u github.com/alecthomas/gometalinter
    gometalinter --install

Before commiting your changes, run gometalinter like this:

    gometalinter ./...

## License

Copyright 2016 Baiju Muthukadan

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

## Author

[Baiju Muthukadan](http://muthukadan.net)

## Contributors

1. [Your Name](http://example.org)
1. [Your Name](https://twitter.com/twitterhandle)
