<!-- file: .github/instructions/go.instructions.md -->
<!-- version: 1.6.0 -->
<!-- guid: 4f5a6b7c-8d9e-0f1a-2b3c-4d5e6f7a8b9c -->
<!-- DO NOT EDIT: This file is managed centrally in ghcommon repository -->
<!-- To update: Create an issue/PR in jdfalk/ghcommon -->

---
applyTo: "**/*.go"
description: |
  Go language-specific coding, documentation, and testing rules for Copilot/AI agents and VS Code Copilot customization. These rules extend the general instructions in `general-coding.instructions.md` and merge all unique content from the Google Go Style Guide.
---

------



# Go Coding Instructions# Go Coding Instructions



- Follow the [general coding instructions](general-coding.instructions.md).- Follow the [general coding instructions](general-coding.instructions.md).

- Follow the- Follow the

  [Google Go Style Guide](https://google.github.io/styleguide/go/index.html) for  [Google Go Style Guide](https://google.github.io/styleguide/go/index.html) for

  additional best practices.  additional best practices.

- All Go files must begin with the required file header (see general- All Go files must begin with the required file header (see general

  instructions for details and Go example).  instructions for details and Go example).



## Core Principles## Core Principles



- Clarity over cleverness: Code should be clear and readable- Clarity over cleverness: Code should be clear and readable

- Simplicity: Prefer simple solutions over complex ones- Simplicity: Prefer simple solutions over complex ones

- Consistency: Follow established patterns within the codebase- Consistency: Follow established patterns within the codebase

- Readability: Code is written for humans to read- Readability: Code is written for humans to read



## Version Requirements## Version Requirements



- **MANDATORY**: All Go projects must use Go 1.23.0 or higher- **MANDATORY**: All Go projects must use Go 1.23.0 or higher

- **NO EXCEPTIONS**: Do not use older Go versions in any repository- **NO EXCEPTIONS**: Do not use older Go versions in any repository

- Update `go.mod` files to specify `go 1.23` minimum version- Update `go.mod` files to specify `go 1.23` minimum version

- Update `go.work` files to specify `go 1.23` minimum version- Update `go.work` files to specify `go 1.23` minimum version

- All Go file headers must use version 1.23.0 or higher- All Go file headers must use version 1.23.0 or higher

- Use `go version` to verify your installation meets requirements- Use `go version` to verify your installation meets requirements



## Naming Conventions## Naming Conventions



- Use short, concise, evocative package names (lowercase, no underscores)- Use short, concise, evocative package names (lowercase, no underscores)

- Use camelCase for unexported names, PascalCase for exported names- Use camelCase for unexported names, PascalCase for exported names

- Use short names for short-lived variables, descriptive names for longer-lived- Use short names for short-lived variables, descriptive names for longer-lived

  variables  variables

- Use PascalCase for exported constants, camelCase for unexported constants- Use PascalCase for exported constants, camelCase for unexported constants

- Single-method interfaces should end in "-er" (e.g., Reader, Writer)- Single-method interfaces should end in "-er" (e.g., Reader, Writer)



## Code Organization## Code Organization



- Use `goimports` to format imports automatically- Use `goimports` to format imports automatically

- Group imports: standard library, third-party, local- Group imports: standard library, third-party, local

- No blank lines within groups, one blank line between groups- No blank lines within groups, one blank line between groups

- Keep functions short and focused- Keep functions short and focused

- Use blank lines to separate logical sections- Use blank lines to separate logical sections

- Order: receiver, name, parameters, return values- Order: receiver, name, parameters, return values



## Formatting## Formatting



- Use tabs for indentation, spaces for alignment- Use tabs for indentation, spaces for alignment

- Opening brace on same line as declaration, closing brace on its own line- Opening brace on same line as declaration, closing brace on its own line

- No strict line length limit, but aim for readability- No strict line length limit, but aim for readability



## Comments## Comments



- Every package should have a package comment- Every package should have a package comment

- Public functions must have comments starting with the function name- Public functions must have comments starting with the function name

- Comment exported variables, explain purpose and constraints- Comment exported variables, explain purpose and constraints



## Error Handling## Error Handling



- Use lowercase for error messages, no punctuation at end- Use lowercase for error messages, no punctuation at end

- Be specific about what failed- Be specific about what failed

- Create custom error types for specific error conditions- Create custom error types for specific error conditions

- Use `errors.Is` and `errors.As` for error checking- Use `errors.Is` and `errors.As` for error checking



## Best Practices## Best Practices



- Use short variable declarations (`:=`) when possible- Use short variable declarations (`:=`) when possible

- Use `var` for zero values or when type is important- Use `var` for zero values or when type is important

- Use `make()` for slices and maps with known capacity- Use `make()` for slices and maps with known capacity

- Accept interfaces, return concrete types- Accept interfaces, return concrete types

- Keep interfaces small and focused- Keep interfaces small and focused

- Use channels for communication between goroutines- Use channels for communication between goroutines

- Use sync primitives for protecting shared state- Use sync primitives for protecting shared state

- Test file names end with `_test.go`, test function names start with `Test`- Test file names end with `_test.go`, test function names start with `Test`

- Use table-driven tests for multiple scenarios- Use table-driven tests for multiple scenarios



## Required File Header## Required File Header



All Go files must begin with a standard header as described in theAll Go files must begin with a standard header as described in the

[general coding instructions](general-coding.instructions.md). Example for Go:[general coding instructions](general-coding.instructions.md). Example for Go:



```go```go

// file: path/to/file.go// file: path/to/file.go

// version: 1.0.0// version: 1.0.0

// guid: 123e4567-e89b-12d3-a456-426614174000// guid: 123e4567-e89b-12d3-a456-426614174000

``````



## Google Go Style Guide (Complete)## Google Go Style Guide (Complete)

    WriteFile(string, []byte) error

Follow the complete Google Go Style Guide below for all Go code:}



### Google Go Style Guide (Complete)// Bad

type IReader interface {  // Don't prefix with I

This style guide provides comprehensive conventions for writing clean, readable, and maintainable Go code.    Read([]byte) (int, error)

}

#### Formatting```



**gofmt:** All Go code must be formatted with `gofmt`. This is non-negotiable.**Functions and Methods:**

- Use MixedCaps

**Line Length:** No hard limit, but prefer shorter lines. Break long lines sensibly.- Exported functions start with capital letter

- Unexported functions start with lowercase letter

**Indentation:** Use tabs for indentation (handled automatically by gofmt).

```go

**Spacing:** Let gofmt handle spacing. Generally:// Good - exported

- No space inside parentheses: `f(a, b)`func CalculateTotal(price, tax float64) float64 {

- Space around binary operators: `a + b`    return price + tax

- No space around unary operators: `!condition`}



#### Naming Conventions// Good - unexported

func validateInput(input string) bool {

**Packages:**    return len(input) > 0

- Short, concise, evocative names}

- Lowercase, no underscores or mixedCaps```

- Often single words

**Variables:**

```go- Use MixedCaps

// Good- Short names for short scopes

package user- Longer descriptive names for longer scopes

package httputil

package json```go

// Good - short scope

// Badfor i, v := range items {

package userService    process(i, v)

package http_util}

```

// Good - longer scope

**Interfaces:**func processUserData(userData map[string]interface{}) error {

- Use -er suffix for single-method interfaces    userID, exists := userData["id"]

- Use MixedCaps    if !exists {

        return errors.New("user ID not found")

```go    }

// Good    // ... more processing

type Reader interface {}

    Read([]byte) (int, error)

}// Bad

func processUserData(d map[string]interface{}) error {  // 'd' too short for scope

type FileWriter interface {    userIdentificationNumber, exists := d["id"]  // Too long for simple value

    WriteFile(string, []byte) error    // ...

}}

```

// Bad

type IReader interface {  // Don't prefix with I**Constants:**

    Read([]byte) (int, error)- Use MixedCaps

}- Group related constants in blocks

```

```go

**Functions and Methods:**// Good

- Use MixedCapsconst (

- Exported functions start with capital letter    StatusOK       = 200

- Unexported functions start with lowercase letter    StatusNotFound = 404

    StatusError    = 500

```go)

// Good - exported

func CalculateTotal(price, tax float64) float64 {const DefaultTimeout = 30 * time.Second

    return price + tax

}// Bad

const STATUS_OK = 200  // Don't use underscores

// Good - unexported```

func validateInput(input string) bool {

    return len(input) > 0#### Package Organization

}

```**Package Names:**

- Choose package names that are both short and clear

**Variables:**- Avoid generic names like "util", "common", "misc"

- Use MixedCaps- Package name should describe what it provides, not what it contains

- Short names for short scopes

- Longer descriptive names for longer scopes```go

// Good

```gopackage user     // for user management

// Good - short scopepackage auth     // for authentication

for i, v := range items {package httputil // for HTTP utilities

    process(i, v)

}// Bad

package utils    // Too generic

// Good - longer scopepackage stuff    // Too vague

func processUserData(userData map[string]interface{}) error {```

    userID, exists := userData["id"]

    if !exists {**Import Organization:**

        return errors.New("user ID not found")- Group imports: standard library, third-party, local

    }- Use goimports to handle this automatically

    // ... more processing

}```go

import (

// Bad    // Standard library

func processUserData(d map[string]interface{}) error {  // 'd' too short for scope    "fmt"

    userIdentificationNumber, exists := d["id"]  // Too long for simple value    "os"

    // ...    "time"

}

```    // Third-party

    "github.com/gorilla/mux"

**Constants:**    "google.golang.org/grpc"

- Use MixedCaps

- Group related constants in blocks    // Local

    "myproject/internal/auth"

```go    "myproject/pkg/utils"

// Good)

const (```

    StatusOK       = 200

    StatusNotFound = 404#### Error Handling

    StatusError    = 500

)**Error Strings:**

- Don't capitalize error messages

const DefaultTimeout = 30 * time.Second- Don't end with punctuation

- Be descriptive but concise

// Bad

const STATUS_OK = 200  // Don't use underscores```go

```// Good

return fmt.Errorf("failed to connect to database: %w", err)

#### Package Organizationreturn errors.New("invalid user ID")



**Package Names:**// Bad

- Choose package names that are both short and clearreturn errors.New("Failed to connect to database.")  // Capitalized, punctuation

- Avoid generic names like "util", "common", "misc"return errors.New("error")  // Too vague

- Package name should describe what it provides, not what it contains```



```go**Error Wrapping:**

// Good- Use fmt.Errorf with %w verb to wrap errors

package user     // for user management- Add context to errors as they bubble up

package auth     // for authentication

package httputil // for HTTP utilities```go

func processUser(id string) error {

// Bad    user, err := getUserFromDB(id)

package utils    // Too generic    if err != nil {

package stuff    // Too vague        return fmt.Errorf("failed to get user %s: %w", id, err)

```    }



**Import Organization:**    if err := validateUser(user); err != nil {

- Group imports: standard library, third-party, local        return fmt.Errorf("user validation failed: %w", err)

- Use goimports to handle this automatically    }



```go    return nil

import (}

    // Standard library```

    "fmt"

    "os"**Error Checking:**

    "time"- Check errors immediately after operations

- Don't ignore errors (use _ only when truly appropriate)

    // Third-party

    "github.com/gorilla/mux"```go

    "google.golang.org/grpc"// Good

file, err := os.Open(filename)

    // Localif err != nil {

    "myproject/internal/auth"    return fmt.Errorf("failed to open file: %w", err)

    "myproject/pkg/utils"}

)defer file.Close()

```

// Bad

#### Error Handlingfile, _ := os.Open(filename)  // Ignoring error

// ... later in code ...

**Error Strings:**if file == nil {  // Too late to handle properly

- Don't capitalize error messages    return errors.New("file is nil")

- Don't end with punctuation}

- Be descriptive but concise```



```go#### Function Design

// Good

return fmt.Errorf("failed to connect to database: %w", err)**Function Length:** Keep functions short and focused. If a function is very long, consider breaking it up.

return errors.New("invalid user ID")

**Function Signature:**

// Bad- Related parameters should be grouped

return errors.New("Failed to connect to database.")  // Capitalized, punctuation- Use meaningful parameter names

return errors.New("error")  // Too vague

``````go

// Good

**Error Wrapping:**func CreateUser(firstName, lastName, email string, age int) *User {

- Use fmt.Errorf with %w verb to wrap errors    return &User{

- Add context to errors as they bubble up        FirstName: firstName,

        LastName:  lastName,

```go        Email:     email,

func processUser(id string) error {        Age:       age,

    user, err := getUserFromDB(id)    }

    if err != nil {}

        return fmt.Errorf("failed to get user %s: %w", id, err)

    }// Bad

func CreateUser(a, b, c string, d int) *User {  // Unclear parameter names

    if err := validateUser(user); err != nil {    return &User{

        return fmt.Errorf("user validation failed: %w", err)        FirstName: a,

    }        LastName:  b,

        Email:     c,

    return nil        Age:       d,

}    }

```}

```

**Error Checking:**

- Check errors immediately after operations**Return Values:**

- Don't ignore errors (use _ only when truly appropriate)- Return errors as the last value

- Use named return parameters sparingly

```go

// Good```go

file, err := os.Open(filename)// Good

if err != nil {func divide(a, b float64) (float64, error) {

    return fmt.Errorf("failed to open file: %w", err)    if b == 0 {

}        return 0, errors.New("division by zero")

defer file.Close()    }

    return a / b, nil

// Bad}

file, _ := os.Open(filename)  // Ignoring error

// ... later in code ...// Acceptable for short, clear functions

if file == nil {  // Too late to handle properlyfunc split(path string) (dir, file string) {

    return errors.New("file is nil")    // ... implementation

}    return

```}

```

#### Function Design

#### Struct Design

**Function Length:** Keep functions short and focused. If a function is very long, consider breaking it up.

**Field Organization:**

**Function Signature:**- Group related fields together

- Related parameters should be grouped- Consider field alignment for memory efficiency

- Use meaningful parameter names

```go

```gotype User struct {

// Good    // Identity fields

func CreateUser(firstName, lastName, email string, age int) *User {    ID       int64

    return &User{    Username string

        FirstName: firstName,    Email    string

        LastName:  lastName,

        Email:     email,    // Personal information

        Age:       age,    FirstName string

    }    LastName  string

}    Age       int



// Bad    // Metadata

func CreateUser(a, b, c string, d int) *User {  // Unclear parameter names    CreatedAt time.Time

    return &User{    UpdatedAt time.Time

        FirstName: a,    Active    bool

        LastName:  b,}

        Email:     c,```

        Age:       d,

    }**Constructor Functions:**

}- Use New prefix for constructor functions

```- Return pointers for structs that will be modified



**Return Values:**```go

- Return errors as the last valuefunc NewUser(username, email string) *User {

- Use named return parameters sparingly    return &User{

        Username:  username,

```go        Email:     email,

// Good        CreatedAt: time.Now(),

func divide(a, b float64) (float64, error) {        Active:    true,

    if b == 0 {    }

        return 0, errors.New("division by zero")}

    }```

    return a / b, nil

}#### Concurrency



// Acceptable for short, clear functions**Goroutines:**

func split(path string) (dir, file string) {- Use goroutines for independent tasks

    // ... implementation- Always consider how goroutines will exit

    return

}```go

```// Good

func processItems(items []Item) {

#### Struct Design    var wg sync.WaitGroup



**Field Organization:**    for _, item := range items {

- Group related fields together        wg.Add(1)

- Consider field alignment for memory efficiency        go func(item Item) {

            defer wg.Done()

```go            process(item)

type User struct {        }(item)

    // Identity fields    }

    ID       int64

    Username string    wg.Wait()

    Email    string}

```

    // Personal information

    FirstName string**Channels:**

    LastName  string- Use channels for communication between goroutines

    Age       int- Close channels when done sending



    // Metadata```go

    CreatedAt time.Timefunc producer(ch chan<- int) {

    UpdatedAt time.Time    defer close(ch)

    Active    bool    for i := 0; i < 10; i++ {

}        ch <- i

```    }

}

**Constructor Functions:**

- Use New prefix for constructor functionsfunc consumer(ch <-chan int) {

- Return pointers for structs that will be modified    for value := range ch {

        fmt.Println(value)

```go    }

func NewUser(username, email string) *User {}

    return &User{```

        Username:  username,

        Email:     email,#### Comments and Documentation

        CreatedAt: time.Now(),

        Active:    true,**Package Comments:**

    }- Every package should have a package comment

}- Use complete sentences

```

```go

#### Concurrency// Package user provides functionality for user management,

// including authentication, authorization, and user data operations.

**Goroutines:**package user

- Use goroutines for independent tasks```

- Always consider how goroutines will exit

**Function Comments:**

```go- Document all exported functions

// Good- Start with the function name

func processItems(items []Item) {- Explain what the function does, not how

    var wg sync.WaitGroup

```go

    for _, item := range items {// CalculateTotal computes the total price including tax.

        wg.Add(1)// It returns an error if the tax rate is negative.

        go func(item Item) {func CalculateTotal(price, taxRate float64) (float64, error) {

            defer wg.Done()    if taxRate < 0 {

            process(item)        return 0, errors.New("tax rate cannot be negative")

        }(item)    }

    }    return price * (1 + taxRate), nil

}

    wg.Wait()```

}

```**Inline Comments:**

- Use for complex logic or non-obvious code

**Channels:**- Explain why, not what

- Use channels for communication between goroutines

- Close channels when done sending```go

// Sort items by priority to ensure high-priority items are processed first

```gosort.Slice(items, func(i, j int) bool {

func producer(ch chan<- int) {    return items[i].Priority > items[j].Priority

    defer close(ch)})

    for i := 0; i < 10; i++ {```

        ch <- i

    }#### Testing

}

**Test Functions:**

func consumer(ch <-chan int) {- Use TestXxx naming convention

    for value := range ch {- Use t.Run for subtests

        fmt.Println(value)

    }```go

}func TestCalculateTotal(t *testing.T) {

```    tests := []struct {

        name     string

#### Comments and Documentation        price    float64

        taxRate  float64

**Package Comments:**        expected float64

- Every package should have a package comment        hasError bool

- Use complete sentences    }{

        {

```go            name:     "positive values",

// Package user provides functionality for user management,            price:    100.0,

// including authentication, authorization, and user data operations.            taxRate:  0.1,

package user            expected: 110.0,

```            hasError: false,

        },

**Function Comments:**        {

- Document all exported functions            name:     "negative tax rate",

- Start with the function name            price:    100.0,

- Explain what the function does, not how            taxRate:  -0.1,

            expected: 0.0,

```go            hasError: true,

// CalculateTotal computes the total price including tax.        },

// It returns an error if the tax rate is negative.    }

func CalculateTotal(price, taxRate float64) (float64, error) {

    if taxRate < 0 {    for _, tt := range tests {

        return 0, errors.New("tax rate cannot be negative")        t.Run(tt.name, func(t *testing.T) {

    }            result, err := CalculateTotal(tt.price, tt.taxRate)

    return price * (1 + taxRate), nil

}            if tt.hasError {

```                if err == nil {

                    t.Errorf("expected error, got none")

**Inline Comments:**                }

- Use for complex logic or non-obvious code                return

- Explain why, not what            }



```go            if err != nil {

// Sort items by priority to ensure high-priority items are processed first                t.Errorf("unexpected error: %v", err)

sort.Slice(items, func(i, j int) bool {                return

    return items[i].Priority > items[j].Priority            }

})

```            if result != tt.expected {

                t.Errorf("expected %f, got %f", tt.expected, result)

#### Testing            }

        })

**Test Functions:**    }

- Use TestXxx naming convention}

- Use t.Run for subtests```



```go**Benchmark Functions:**

func TestCalculateTotal(t *testing.T) {

    tests := []struct {```go

        name     stringfunc BenchmarkCalculateTotal(b *testing.B) {

        price    float64    for i := 0; i < b.N; i++ {

        taxRate  float64        CalculateTotal(100.0, 0.1)

        expected float64    }

        hasError bool}

    }{```

        {

            name:     "positive values",This covers the essential Go style guidelines including:/instructions/go.instructions.md -->

            price:    100.0,<!-- version: 1.6.0 -->

            taxRate:  0.1,<!-- guid: 4f5a6b7c-8d9e-0f1a-2b3c-4d5e6f7a8b9c -->

            expected: 110.0,<!-- DO NOT EDIT: This file is managed centrally in ghcommon repository -->

            hasError: false,<!-- To update: Create an issue/PR in jdfalk/ghcommon -->

        },

        {---

            name:     "negative tax rate",applyTo: "**/*.go"

            price:    100.0,description: |

            taxRate:  -0.1,  Go language-specific coding, documentation, and testing rules for Copilot/AI agents and VS Code Copilot customization. These rules extend the general instructions in `general-coding.instructions.md` and merge all unique content from the Google Go Style Guide.

            expected: 0.0,---

            hasError: true,

        },# Go Coding Instructions

    }

- Follow the [general coding instructions](general-coding.instructions.md).

    for _, tt := range tests {- Follow the

        t.Run(tt.name, func(t *testing.T) {  [Google Go Style Guide](https://google.github.io/styleguide/go/index.html) for

            result, err := CalculateTotal(tt.price, tt.taxRate)  additional best practices.

- All Go files must begin with the required file header (see general

            if tt.hasError {  instructions for details and Go example).

                if err == nil {

                    t.Errorf("expected error, got none")## Core Principles

                }

                return- Clarity over cleverness: Code should be clear and readable

            }- Simplicity: Prefer simple solutions over complex ones

- Consistency: Follow established patterns within the codebase

            if err != nil {- Readability: Code is written for humans to read

                t.Errorf("unexpected error: %v", err)

                return## Version Requirements

            }

- **MANDATORY**: All Go projects must use Go 1.23.0 or higher

            if result != tt.expected {- **NO EXCEPTIONS**: Do not use older Go versions in any repository

                t.Errorf("expected %f, got %f", tt.expected, result)- Update `go.mod` files to specify `go 1.23` minimum version

            }- Update `go.work` files to specify `go 1.23` minimum version

        })- All Go file headers must use version 1.23.0 or higher

    }- Use `go version` to verify your installation meets requirements

}

```## Version Requirements



**Benchmark Functions:**- **MANDATORY**: All Go projects must use Go 1.23.0 or higher

- **NO EXCEPTIONS**: Do not use older Go versions in any repository

```go- Update `go.mod` files to specify `go 1.23` minimum version

func BenchmarkCalculateTotal(b *testing.B) {- Update `go.work` files to specify `go 1.23` minimum version

    for i := 0; i < b.N; i++ {- All Go file headers must use version 1.23.0 or higher

        CalculateTotal(100.0, 0.1)- Use `go version` to verify your installation meets requirements

    }

}## Naming Conventions

```

- Use short, concise, evocative package names (lowercase, no underscores)

This covers the essential Go style guidelines including formatting, naming conventions, package organization, error handling, function design, struct design, concurrency, comments, and testing best practices.- Use camelCase for unexported names, PascalCase for exported names
- Use short names for short-lived variables, descriptive names for longer-lived
  variables
- Use PascalCase for exported constants, camelCase for unexported constants
- Single-method interfaces should end in "-er" (e.g., Reader, Writer)

## Code Organization

- Use `goimports` to format imports automatically
- Group imports: standard library, third-party, local
- No blank lines within groups, one blank line between groups
- Keep functions short and focused
- Use blank lines to separate logical sections
- Order: receiver, name, parameters, return values

## Formatting

- Use tabs for indentation, spaces for alignment
- Opening brace on same line as declaration, closing brace on its own line
- No strict line length limit, but aim for readability

## Comments

- Every package should have a package comment
- Public functions must have comments starting with the function name
- Comment exported variables, explain purpose and constraints

## Error Handling

- Use lowercase for error messages, no punctuation at end
- Be specific about what failed
- Create custom error types for specific error conditions
- Use `errors.Is` and `errors.As` for error checking

## Best Practices

- Use short variable declarations (`:=`) when possible
- Use `var` for zero values or when type is important
- Use `make()` for slices and maps with known capacity
- Accept interfaces, return concrete types
- Keep interfaces small and focused
- Use channels for communication between goroutines
- Use sync primitives for protecting shared state
- Test file names end with `_test.go`, test function names start with `Test`
- Use table-driven tests for multiple scenarios

## Required File Header

All Go files must begin with a standard header as described in the
[general coding instructions](general-coding.instructions.md). Example for Go:

```go
// file: path/to/file.go
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174000
```
