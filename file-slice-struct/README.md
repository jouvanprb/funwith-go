## read.go

Write a Go program named read.go that performs the following steps:

### 1. Define a struct named name with two fields:
```
fname (first name)  
lname (last name)
```

### 2. Prompt the user to enter the name of a text file.

### 3. The text file will contain one first name and one last name per line, separated by a space.

Example file content:
```
John Smith
Jane Doe
Alice Brown
```

### 4. Read each line from the file.

#### For each line:

- Create a struct containing the first and last name.  
- Add the struct to a **slice of structs**.

#### After reading the entire file:

- Iterate through the slice.
- Print each first name and last name pair.


## Solution
Open `read.go`

![read.go](../img/file-slice-struct/file-struct-slice.avif)