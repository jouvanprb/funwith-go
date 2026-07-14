# Go - Displacement Function (Closure)

## Overview

This program calculates displacement using the following physics formula:

```math
s = \frac{1}{2}at^2 + v_ot + s_o
```


Go implementation:

```go
s = 0.5*a*t*t + vo*t + so
```

Where:

- **a** = acceleration
- **vo** = initial velocity
- **so** = initial displacement
- **t** = time
- **s** = displacement

The user is prompted to enter the values of:

- Acceleration (`a`)
- Initial velocity (`vo`)
- Initial displacement (`so`)
- Time (`t`)

The program then computes and prints the displacement.

---

## Learning Objective

This exercise is **not only about the displacement formula**.

Its primary goal is to understand **higher-order functions** and **closures** in Go.

Concepts learned:

- Functions are first-class values.
- Functions can be returned from another function.
- Anonymous functions.
- Closures (returned functions remember variables from their outer scope).

---

## Program Structure

```
main()
│
├── Read user input (a, vo, so)
│
├── fn := GenDisplaceFn(a, vo, so)
│
├── Read user input (t)
│
└── Print fn(t)
```

`GenDisplaceFn()` generates a new function that remembers the values of `a`, `vo`, and `so`.

```
GenDisplaceFn(a, vo, so)
            │
            ▼
returns

func(t float64) float64
```

The returned function only requires the value of `t` because the other variables are stored inside the closure.

---

## Closure Illustration

```go
fn := GenDisplaceFn(10, 2, 1)
```

The generated function behaves as if it were:

```go
func(t float64) float64 {
    return 0.5*10*t*t + 2*t + 1
}
```

So calling:

```go
fn(3)
```

computes:

```
0.5 × 10 × 3² + 2 × 3 + 1
```

---

##  Example

**Input**

```
Acceleration: 10
Initial Velocity: 2
Initial Displacement: 1
Time: 3
```

**Output**

```
52
```

---

## 💡 Key Takeaways

- A function can return another function.
- A returned function can remember variables from its surrounding scope (closure).
- Functions in Go are values and can be assigned to variables.
- Closures are commonly used in callbacks, HTTP handlers, middleware, and other Go applications.