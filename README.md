# DeepCopy
![example workflow](https://github.com/wangsquirrel/deepcopy/actions/workflows/go.yml/badge.svg)
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

## Performance

```
goos: linux
goarch: amd64
pkg: github.com/wangsquirrel/deepcopy
cpu: AMD Ryzen 7 2700 Eight-Core Processor          
Benchmark_Copy-16                 399337              2790 ns/op            1176 B/op         23 allocs/op
Benchmark_SonicCopy-16            134254              8994 ns/op            8386 B/op         30 allocs/op
Benchmark_GOBCopy-16                8893            134633 ns/op           31186 B/op        689 allocs/op
Benchmark_ReflectCopy-16           78482             15120 ns/op            3576 B/op        145 allocs/op
Benchmark_JsonCopy-16              29007             40924 ns/op            4326 B/op        118 allocs/op
Benchmark_CPYCopy-16               81296             14695 ns/op            2952 B/op         43 allocs/op
Benchmark_RawCopy-16             2407807               500.7 ns/op           272 B/op         16 allocs/op
PASS
ok      github.com/wangsquirrel/deepcopy        10.555s
```