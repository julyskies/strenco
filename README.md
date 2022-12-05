## strenco

String encoding with modified [Caesar cipher](https://en.wikipedia.org/wiki/Caesar_cipher) for Golang.

Encoding is done using [rolling cipher](http://www.thealmightyguru.com/Wiki/index.php?title=Rolling_cipher).

Minimal required Go version: `v1.18`

### Install

```shell script
go get github.com/julyskies/strenco
```

### Usage

##### Encode plaintext string

```go
package main

import "github.com/julyskies/strenco"

func main() {
  encoded, encodingError := strenco.Encode("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
  if encodingError != nil {
    println(encodingError.Error())
  } else {
    println(encoded) // %%7z;1tsk1JA!NbAFwzj j&{r(3#"+L%m*{<D%Wrh@iyB(qM!eG)>B/ox/%%
  }
}
```

##### Decode encoded string

```go
package main

import "github.com/julyskies/strenco"

func main() {
  decoded, decodingError := strenco.Decode(`%%7z;1tsk1JA!NbAFwzj j&{r(3#"+L%m*{<D%Wrh@iyB(qM!eG)>B/ox/%%`)
  if decodingError != nil {
    println(decodingError.Error())
  } else {
    println(decoded) // Lorem ipsum dolor sit amet, consectetur adipiscing elit.
  }
}
```

### Security

This module is not secure, since its main goal is to make the original value unreadable. There are no additional keys, passphrases or secrets required, meaning that anyone can decode encoded value. It is not recommended to store sensitive data encoded with this module.

### Testing

Deploy project locally

```shell script
git clone https://github.com/julyskies/strenco
cd ./strenco
gvm use go1.18
```

Run tests

```shell script
go test -v
```

### License

[MIT](./LICENSE.md)
