# DeepCopy
[![codecov](https://codecov.io/gh/wangsquirrel/deepcopy/branch/main/graph/badge.svg?token=ZPW283W4QV)](https://codecov.io/gh/wangsquirrel/deepcopy)

## Usage

```golang
var a string = "a"
var b string
err := deepcopy.DeepCopy(&a, &b)
if err != nil {
}
```