# Vector
A very simple math vector library that does 0 memory allocation, allowing you to reuse objects.

Simply specify an `out *Vector` and it will store the result there.
I have found this to be espescially useful in iterative integration, where I can store the results
of math operations in a variable declared outside of the for loop to make allocations O(1) instead of O(n)

### With other libraries
O(n) allocations; slow.
```go
counter := &Vector2f{.1, .1}
for i := 0; i < 200; i++ {
  counter = counter.Add(counter)
}
```
Ignore that this code can be simplified, this is just demonstrating that the Add operation returns a new object, hence an allocation.

### With this library
O(1) allocations; fast.
```go
counter := &Vector2f{.1, .1}
for i := 0; i < 200; i++ {
    // Add2f(v1, v2, out *Vector2f)
    vector.Add2f(counter, counter, counter)
}
```
This writes the result to the out vector directly, not creating
any new object. The allocation is up to the user, which can reuse
objects to make code faster, and possibly even cache friendly.

## Structures and functions
Currently only `Vector2f` and `Vector3f` are implemented.

For brevity the `2f` and `3f` will be replaced with `Xf`
```go
// out = v1 + v2
func AddXf(v1, v2, out *VectorXf)

// out = v1 - v2
func SubXf(v1, v2 out *VectorXf)

// out = v * k
func MulXf(v *VectorXf, k float64, out *VectorXf)

// out = v / k
func DivXf(v *VectorXf, k float64, out *VectorXf)

// return ||v||
func (v *VectorXf) Mag() float64

// return "<x, y[, z]>|mag|"
func (v *VectorXf) String() string
```

## Possible improvements
Implementing SIMD is not something this library is aiming to do.
This library is solely to allow the user to manage memory as allocations hog down calculations, espescially in tight loops.

* [ ] Cross and dot multiplication
* [ ] Unit tests
* [ ] More types with more dimensions? This is up for discussion if it is necessary and at which point there is enough.

Pull requests and contributions welcome
