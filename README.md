Abdi
-----
Hide!  [![Build Status](https://travis-ci.org/fortytw2/abdi.svg)](https://travis-ci.org/fortytw2/abdi) [![GoDoc](https://godoc.org/github.com/fortytw2/abdi?status.svg)](http://godoc.org/github.com/fortytw2/abdi)

Dead simple library for verifying and hashing passwords, following the Mozilla
Password guidelines found [here](https://wiki.mozilla.org/WebAppSec/Secure_Coding_Guidelines#Password_Complexity)
using the prescribed salted bcrypt + hmac.

Contains a 1000 word password Blacklist derived from [here](https://xato.net/passwords/more-top-worst-passwords/)

Usage
------

Install with `go get github.com/fortytw2/abdi`

Example -

```go
package main

import (
  "fmt"
  "github.com/fortyw2/abdi"
)

func main() {
  // by default, abdi.Hash enforces an 8 character minimum password length
  // change this by changing abdi.MinPasswordLength

  // to change the default Blacklist, simply edit abdi.Blacklist, a []string

  abdi.Key = "thisismyhmackey"

  hash, err := abdi.Hash("thispassword")
  if err != nil {
    panic(err)
  }

  if err = abdi.Check("thispassword", hash); err == nil {
    fmt.Println("Password looks good to me :)")
  }
}

```

LICENSE
------
abdi.go and abdi_test.go are Public Domain, see UNLICENSE

common_passwords.txt
