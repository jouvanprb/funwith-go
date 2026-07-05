# slice.go

## Description

Create a Go program named `slice.go` that demonstrates the use of **slices** and **dynamic expansion** using `append()`.

The program should continuously accept integer input from the user until the user enters **`X`** to terminate the program.


## Requirements

The program must:

1. Create an integer slice before entering the loop.
2. Initialize the slice with:

   * Length: `0`
   * Capacity: `3`
3. Prompt the user to enter integers one at a time.
4. After each valid integer input:

   * Add the number to the slice using `append()`.
   * Print the updated slice.
5. The slice should automatically grow when its capacity is exceeded.
6. Stop accepting input when the user enters **`X`** (uppercase).


## Example

### Input

```text
3
1
5
X
```

### Output

```text
[3]
[3 1]
[3 1 5]
```

## Learning Objectives

By completing this exercise, you will practice:

* Creating slices with `make()`
* Understanding slice **length** and **capacity**
* Using `append()` to dynamically grow a slice
* Reading user input in a loop
* Using `break` to terminate a loop
* Basic input validation

## Suggested Implementation

* Create the slice using:

```go
numbers := make([]int, 0, 3)
```

* Use an infinite `for` loop.
* Read user input as a string.
* If the input is `"X"`, exit the loop.
* Otherwise, convert the input to an integer.
* Append the integer to the slice.
* Print the current contents of the slice after every successful insertion.


## Expected Behavior

```text
Enter a number (X to quit): 3
[3]

Enter a number (X to quit): 1
[3 1]

Enter a number (X to quit): 5
[3 1 5]

Enter a number (X to quit): X
Program terminated.
```


## Solution and Output
Go to `main.go`

![output slice](../img/slice/slice-go.avif)