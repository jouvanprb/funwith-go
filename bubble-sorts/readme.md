# Bubble Sort in Go

## Description

This project implements the **Bubble Sort** algorithm in Go.

The program prompts the user to enter a sequence of **up to 10 integers**, stores them in a slice, and then sorts the numbers from **smallest to largest** using the Bubble Sort algorithm.

The implementation is divided into two functions:

- `BubbleSort()` — sorts the slice using the Bubble Sort algorithm.
- `Swap()` — swaps two adjacent elements in the slice.

---

## Learning Objectives

This project is intended to practice:

- Working with slices
- Function decomposition
- Passing slices to functions
- In-place data modification
- Nested loops
- Basic sorting algorithms

---

## Bubble Sort Algorithm

Bubble Sort repeatedly compares two adjacent elements in a slice.

If the left element is greater than the right element, their positions are swapped.

This process continues until no more swaps are needed, meaning the slice is sorted.

Example:

Input:

```
5 3 8 1 2
```

Pass 1:

```
3 5 1 2 8
```

Pass 2:

```
3 1 2 5 8
```

Pass 3:

```
1 2 3 5 8
```

Output:

```
1 2 3 5 8
```

---

## Functions

### BubbleSort()

```go
func BubbleSort(numbers []int)
```

Sorts the slice in ascending order.

Parameters:

- `numbers []int` — slice of integers to be sorted.

Returns:

- Nothing.

The function modifies the original slice directly.

---

### Swap()

```go
func Swap(numbers []int, i int)
```

Swaps two adjacent elements:

```
numbers[i]
```

and

```
numbers[i+1]
```

Parameters:

- `numbers []int`
- `i int`

Returns:

- Nothing.

---

## Example

### Input

```
Enter up to 10 integers:
5
9
2
7
1
x
```

### Output

```
Sorted:
1 2 5 7 9
```

---

## Project Structure

```
bubble-sort/
│
├── main.go
└── README.md
```

---

## Concepts Used

- Variables
- Functions
- Slices
- For loops
- Conditional statements
- Bubble Sort algorithm
- Swapping elements
- In-place sorting
