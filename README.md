# DeepCopy

## Usage

```golang
var a string = "a"
var b string
err := deepcopy.DeepCopy(&a, &b)
if err != nil {
}
```