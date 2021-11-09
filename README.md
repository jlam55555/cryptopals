# cryptopals
My solutions to the Cryptopals Crypto Challenge

---

### Why Go?

Go has some nice libraries for encoding/encryption.

(But mostly because I wanted to use Go for at least one project this semester.)

---

### Build instructions
A standard installation of Go is required.

##### Running tests
This checks the solutions against the test cases provided. (Note that there is no `main` function anywhere -- the tests are the driver code.) See `./src/set1/set1_test.go` for the test cases.
```bash
$ go test ./src/set1  # or,
$ GOPATH=$(pwd):$GOPATH go test set1
```
