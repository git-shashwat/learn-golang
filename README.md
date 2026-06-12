# Learn Go

A hands-on Go learning repository with small, runnable examples organized by topic. Each lesson lives in its own directory with a standalone `main.go` file you can run independently.

## Prerequisites

- [Go](https://go.dev/dl/) 1.21 or later

Verify your installation:

```bash
go version
```

## Getting Started

Clone the repository and run any lesson from its directory:

```bash
git clone <repository-url>
cd learn-golang
go run introduction/1.1_hello_world/main.go
```

Or run from the project root by passing the full path:

```bash
go run variables/2.1_declaration/main.go
```

## Project Structure

```
learn-golang/
├── introduction/          # Getting started with Go
│   ├── 1.1_hello_world/
│   ├── 1.2_bugfix/
│   ├── 1.3_compile_error/
│   └── 1.4_typing/
└── variables/             # Variables, types, and control flow
    ├── 2.1_declaration/
    ├── 2.2_short_declaration/
    ├── 2.3_multiple_declaration/
    ├── 2.4_type_casting/
    ├── 2.5_constants/
    ├── 2.6_computed_constant/
    ├── 2.7_formatted_string/
    └── 2.8_conditionals/
```

## Lessons

### Introduction

| Lesson | Topic | What you'll learn |
|--------|-------|-------------------|
| `1.1_hello_world` | Hello World | Your first Go program with `fmt.Println` |
| `1.2_bugfix` | Bug Fix | Using constants and `fmt.Printf` for formatted output |
| `1.3_compile_error` | Compile Errors | How syntax mistakes cause compile-time failures |
| `1.4_typing` | Static Typing | Go's type system and fixing type mismatches |

### Variables

| Lesson | Topic | What you'll learn |
|--------|-------|-------------------|
| `2.1_declaration` | Variable Declaration | Declaring variables with explicit types |
| `2.2_short_declaration` | Short Declaration | The `:=` shorthand for variable assignment |
| `2.3_multiple_declaration` | Multiple Declaration | Declaring multiple variables in one statement |
| `2.4_type_casting` | Type Casting | Converting between numeric types |
| `2.5_constants` | Constants | Declaring immutable values with `const` |
| `2.6_computed_constant` | Computed Constants | Deriving constants from other constants |
| `2.7_formatted_string` | Formatted Strings | Building strings with `fmt.Sprintf` |
| `2.8_conditionals` | Conditionals | Branching logic with `if` / `else` |

## Running All Lessons

Run every lesson in order:

```bash
for dir in introduction/*/ variables/*/; do
  echo "Running $dir"
  go run "$dir/main.go"
  echo
done
```

## License

This project is for personal learning. Feel free to use and adapt the examples.
