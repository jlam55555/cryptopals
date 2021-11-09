# cryptopals
My solutions to the Cryptopals Crypto Challenge

---

### Why Go?

(Because I wanted to use Go for at least one project this semester.)

Because I know it has encoding libraries in its stdlib that should make this relatively easy.

---

### Running tests
This checks the solutions against the test cases provided. (Note that there is no `main` function anywhere -- the tests are the driver code.)

```bash
$ go test ./src/set1  # or,
$ GOPATH=$(pwd):$GOPATH go test set1
```
