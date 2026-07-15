# Go - Animal Information Program

## Description

This program allows the user to retrieve information about three predefined animals:

- Cow
- Bird
- Snake

Each animal stores three pieces of information:

- Food
- Method of locomotion
- Spoken sound

The user enters a command in the following format:

```text
> <animal> <action>
```

Example:

```text
> cow eat
grass

> bird move
fly

> snake speak
hsss
```

The program continues accepting requests until it is terminated manually.

---

## Predefined Animals

| Animal | Food | Locomotion | Sound |
|---------|------|------------|-------|
| Cow | grass | walk | moo |
| Bird | worms | fly | peep |
| Snake | mice | slither | hsss |

---

## Learning Objectives

This exercise focuses on understanding Object-Oriented Programming concepts in Go using structs and methods.

Concepts learned:

- Structs
- Methods
- Method Receivers
- Creating Objects
- User Input
- String Processing
- Loops
- Nested Switch Statements

---

## 🏗 Program Structure

```
main()
│
├── Create Animal objects
│      ├── cow
│      ├── bird
│      └── snake
│
├── Read user input
│
├── Split input into:
│      ├── animal
│      └── action
│
├── Determine selected animal
│
├── Call corresponding method
│      ├── Eat()
│      ├── Move()
│      └── Speak()
│
└── Repeat 
```

---

## Struct Design

Each animal is represented by an `Animal` struct.

```go
type Animal struct {
    food       string
    locomotion string
    noise      string
}
```

Each object stores its own values.

Example:

```go
cow := Animal{
    food:       "grass",
    locomotion: "walk",
    noise:      "moo",
}
```

---

## ⚙ Methods

Each `Animal` has three methods attached through method receivers.

```go
Eat()
Move()
Speak()
```

These methods print information stored inside the struct.

For example:

```text
cow.Eat()   → grass
cow.Move()  → walk
cow.Speak() → moo
```

---

## Input Format

The program expects two words separated by a space.

```
<animal> <action>
```

Valid animals:

- cow
- bird
- snake

Valid actions:

- eat
- move
- speak

---


## 💡 Key Takeaways

- Structs group related data together.
- Methods allow behaviors to be attached to a struct.
- Method receivers provide access to an object's data.
- Objects can share the same methods while storing different values.
- User commands can be parsed using `strings.Split()`.
- `switch` statements are useful for handling multiple user choices.

