package otherpkg

import (
  "somepkg"
  "fmt"
)

func OtherFunc() {
     somepkg.SomeFunc()
     somepkg.SomeVar = 5
}

