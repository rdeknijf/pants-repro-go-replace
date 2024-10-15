# `go replace` problem reproduction repo

This repo is an example that reproduces the following:

    16:46:37.10 [ERROR] 1 Exception encountered:

    Engine traceback:
    in `test` goal

    KeyError: 'Version'

Any command, like `pants test ::` etc. Will result in this error.