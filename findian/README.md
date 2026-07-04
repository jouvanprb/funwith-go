# findian.go
You will write a Go program that checks whether a string meets specific conditions and prints “Found!” or “Not Found!” accordingly.

The goal of this assignment is to practice working with **strings** in Go.

1. Prompt the user to enter a string.
2. Check whether the string:
3. The program should not be case-sensitive.

If all conditions are met, print:  `Found!`
Otherwise print: `Not Found!`

## Example Inputs

#### 1. These should print `Found!`
ian
Ian
iuiygaygn
I d skd a efju N

#### 2. These should print `Not Found!`
ihhhhhn
ina
xian

## Solution
Go to file `findian.go`

<img src="../img/findian/outputFindian.avif" alt="Ouput findian" width="400">


## Conclusion of package `strings`
```go
import "strings"

// 1. TrimSpace - Hapus whitespace di awal/akhir
strings.TrimSpace("  hello  ")  // "hello"

// 2. ToLower - Ubah ke huruf kecil
strings.ToLower("HELLO")  // "hello"

// 3. ReplaceAll - Ganti semua karakter
strings.ReplaceAll("hello world", " ", "")  // "helloworld"

// 4. HasPrefix - Cek awalan
strings.HasPrefix("hello", "he")  // true

// 5. HasSuffix - Cek akhiran
strings.HasSuffix("hello", "lo")  // true

// 6. Contains - Cek apakah mengandung
strings.Contains("hello", "ell")  // true
```


