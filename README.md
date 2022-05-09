# DeepCopy
![Build and Test](https://github.com/wangsquirrel/deepcopy/workflows/go/badge.svg?branch=master&event=push)
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
