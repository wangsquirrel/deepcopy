# DeepCopy
![example workflow](https://github.com/wangsquirrel/deepcopy/actions/workflows/main.yml/badge.svg)
[![codecov](https://codecov.io/gh/wangsquirrel/deepcopy/branch/main/graph/badge.svg?token=ZPW283W4QV)](https://codecov.io/gh/wangsquirrel/deepcopy)


---

Deep copy most of the golang types except `unsafe.Pointer` and `chan` with good performance

## Usage

```golang
var a string = "a"
var b string
err := deepcopy.DeepCopy(&a, &b)
if err != nil {
}
```
