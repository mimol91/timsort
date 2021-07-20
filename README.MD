[![Go Report Card](https://goreportcard.com/badge/github.com/mimol91/timsort)](https://goreportcard.com/report/github.com/mimol91/timsort)
[![Build Status](https://travis-ci.org/mimol91/timsort.svg?branch=master)](https://travis-ci.org/mimol91/timsort)
[![codecov](https://codecov.io/gh/mimol91/timsort/branch/master/graph/badge.svg)](https://codecov.io/gh/mimol91/timsort)
# Timsort

Timsort is a hybrid stable sorting algorithm, derived from merge sort and insertion sort, designed to perform well on many kinds of real-world data.
[Details](https://en.wikipedia.org/wiki/Timsort)

# Usuage
```golang
import "github.com/mimol91/timsort"

elementsToSort := int[]{1,6,2}
sorted := Timsort.Sort[int](elementsToSort)
```
