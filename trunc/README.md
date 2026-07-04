## Truncate Floating Point Number - Go Program

## Description
This Go program accepts a floating point number from the user and prints the truncated integer value. Truncation means removing all digits to the right of the decimal point without rounding.

## Program Details
- **File name:** `trunc.go`
- **Language:** Go

## How It Works
1. The program prompts the user to enter a floating point number.
2. It reads the input using `fmt.Scan()`.
3. It converts the floating point number to an integer using `int()` type conversion, which automatically truncates the decimal part.
4. It prints the truncated integer value.

## Example Output
```
Enter a floating point number:
3.7
Truncated integer: 3
```

```
Enter a floating point number:
-2.9
Truncated integer: -2
```
