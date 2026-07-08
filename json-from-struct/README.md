# Make JSON from User Input Using Struct

## Objective

Write a Go program that:

1. Prompts the user to enter their **name**, **age**, and **address**.
2. Stores the input in a **struct**.
3. Converts the struct into JSON format.
4. Prints the JSON string to the console.

---

## Requirements

- Use `bufio.NewReader` to read user input.
- Convert the age from `string` to `int` using `strconv.Atoi`.
- Define a struct to store the user data.
- Convert the struct into JSON using `json.Marshal`.
- Handle errors when converting the age and when marshaling JSON.

---

## Example Input

```text
Enter name: Jouvan Augusto Purba
Enter Age: 22
Enter Address: Sukabirus, Citeureup 08
```

## Expected Output

```json
{"Name":"Jouvan Augusto Purba","Age":22,"Address":"Sukabirus, Citeureup 08"}
```

---

## Packages to Use

- `fmt`
- `bufio`
- `os`
- `strings`
- `strconv`
- `encoding/json`

---

## Expected Struct

```go
type User struct {
    Name    string
    Age     int
    Address string
}
```

---

## Learning Outcomes

After completing this assignment, you should be able to:

- Read user input using `bufio.Reader`.
- Clean input using `strings.TrimSpace`.
- Convert strings to integers using `strconv.Atoi`.
- Define and use a struct.
- Initialize a struct with user input.
- Convert a struct into JSON using `json.Marshal`.
- Handle basic errors in Go.