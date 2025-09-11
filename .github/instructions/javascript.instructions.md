<!-- file: .github/instructions/javascript.instructions.md -->
## JavaScript Style Guidelines

Follow the complete Google JavaScript Style Guide below for all JavaScript code:

### Google JavaScript Style Guide (Complete)

This style guide is a comprehensive set of conventions for writing readable and maintainable JavaScript code.

#### Source File Basics

**File encoding:** UTF-8 only. Use Unicode escape sequences (\u221e) for non-ASCII characters with explanatory comments.

**Whitespace:** Use spaces, not tabs. Indent by 2 spaces.

```javascript
// Good
function example() {
  const value = 42;
  return value;
}

// Bad
function example() {
    const value = 42;    // 4-space indent
    return value;
}
```

#### Source File Structure

Files consist of the following in order:
1. License or copyright information (if present)
2. @fileoverview JSDoc (if present)
3. goog.module statement
4. goog.require statements
5. The file's implementation

#### Language Features

**Variable Declarations:**
- Use `const` and `let`, never `var`
- Use `const` by default, `let` only when reassigning
- Declare variables as close to first use as possible
- One variable per declaration

```javascript
// Good
const items = getItems();
let count = 0;

// Bad
var items = getItems();
var count = 0, total = 0;  // Multiple declarations
```

**Array Literals:**
- Use trailing commas
- Don't use the Array constructor
- Use spread syntax for copying

```javascript
// Good
const items = [
  'apple',
  'banana',
  'cherry',
];

const copy = [...items];

// Bad
const items = new Array('apple', 'banana');
```

**Object Literals:**
- Use trailing commas
- Don't quote keys unless necessary
- Use computed property names sparingly
- Use object shorthand when possible

```javascript
// Good
const obj = {
  a: 1,
  b: 2,
  method() {
    return this.a;
  },
};

// Bad
const obj = {
  'a': 1,
  b: 2,
};
```

**Classes:**
- Use ES6 class syntax
- Constructor comes first, then instance methods, then static methods
- Don't manipulate prototypes directly

```javascript
// Good
class MyClass {
  constructor(value) {
    this.value_ = value;
  }

  getValue() {
    return this.value_;
  }

  static createDefault() {
    return new MyClass(0);
  }
}
```

**Functions:**
- Prefer function declarations for top-level functions
- Use arrow functions for callbacks and inline functions
- Don't use `function` keyword in function expressions

```javascript
// Good: function declaration
function topLevelFunction() {
  return 42;
}

// Good: arrow function for callback
items.map(item => item.value);

// Good: arrow function preserves `this`
class Handler {
  constructor() {
    element.addEventListener('click', () => this.handle());
  }
}
```

**Arrow Functions:**
- Use parentheses around parameters unless single untyped parameter
- Use block body when statements are needed, expression body for single expressions
- Don't use arrow functions for object methods or prototype methods

```javascript
// Good
const square = x => x * x;
const sum = (a, b) => a + b;
const process = items => {
  const results = [];
  for (const item of items) {
    results.push(transform(item));
  }
  return results;
};
```

**Destructuring:**
- Use object destructuring for accessing multiple properties
- Use array destructuring for unpacking arrays
- Use default values in destructuring

```javascript
// Good
const {name, age} = person;
const [first, second] = array;
const {title = 'Default Title'} = config;
```

#### Statements

**Loops:**
- Use for-of for iteration when possible
- Use for-in only for objects (not arrays)
- Use Object.keys(), Object.values(), Object.entries() for objects

```javascript
// Good
for (const item of items) {
  process(item);
}

for (const key of Object.keys(obj)) {
  console.log(key, obj[key]);
}

// Bad
for (let i = 0; i < items.length; i++) {
  process(items[i]);  // Prefer for-of when index not needed
}
```

**Conditionals:**
- Use strict equality (=== and !==)
- Use truthiness carefully - prefer explicit comparisons for clarity

```javascript
// Good
if (array.length > 0) { ... }
if (string !== '') { ... }
if (object !== null) { ... }

// Acceptable for truthiness
if (array.length) { ... }
if (object) { ... }
```

**Exception Handling:**
- Always use Error objects when throwing
- Use try/catch for expected exceptions
- Don't use empty catch blocks

```javascript
// Good
if (condition) {
  throw new Error('Something went wrong');
}

try {
  riskyOperation();
} catch (err) {
  console.error('Operation failed:', err.message);
  return fallbackValue;
}
```

#### Naming

**Conventions:**
- `lowerCamelCase` for variables, functions, methods, properties
- `UpperCamelCase` for constructors and classes
- `CONSTANT_CASE` for module-level constants
- `UPPER_SNAKE_CASE` for global constants

```javascript
// Good
const userName = 'john';
const MAX_COUNT = 100;

function getUserData() { ... }

class UserManager { ... }

// Bad
const user_name = 'john';  // Wrong case
const maxcount = 100;      // Should be CONSTANT_CASE
```

**Descriptive Names:**
- Use clear, descriptive names
- Avoid abbreviations unless widely understood
- Method names should be verbs, variables should be nouns

```javascript
// Good
const userAccountList = getActiveUsers();
function calculateTotal() { ... }

// Bad
const usrAcctLst = getActvUsrs();  // Too abbreviated
function calc() { ... }            // Unclear
```

#### Formatting

**Line Length:** Maximum 80 characters. Break long lines logically.

**Indentation:** 2 spaces for each level. No tabs.

**Semicolons:** Required at the end of every statement.

**Commas:** Trailing commas in arrays and objects.

```javascript
// Good formatting
const config = {
  apiUrl: 'https://api.example.com',
  timeout: 5000,
  retries: 3,
};

const longFunctionCall = someVeryLongFunctionName(
    firstParameter,
    secondParameter,
    thirdParameter);
```

#### Comments and Documentation

**JSDoc:**
- Document all public APIs
- Use proper JSDoc tags (@param, @return, @throws)
- Include examples for complex functions

```javascript
/**
 * Calculates the distance between two points.
 * @param {number} x1 X coordinate of first point
 * @param {number} y1 Y coordinate of first point
 * @param {number} x2 X coordinate of second point
 * @param {number} y2 Y coordinate of second point
 * @return {number} The distance between the points
 */
function calculateDistance(x1, y1, x2, y2) {
  const dx = x2 - x1;
  const dy = y2 - y1;
  return Math.sqrt(dx * dx + dy * dy);
}
```

**Inline Comments:**
- Use // for line comments
- Explain why, not what
- Keep comments up to date with code changes

#### Type Annotations (JSDoc)

When not using TypeScript, use JSDoc for type information:

```javascript
/**
 * @param {string} name
 * @param {number} age
 * @param {Array<string>} hobbies
 * @return {Object} User object
 */
function createUser(name, age, hobbies) {
  return {name, age, hobbies};
}
```

This covers the essential JavaScript style guidelines including:nstructions/javascript.instructions.md -->
<!-- version: 1.2.0 -->
<!-- guid: 8e7d6c5b-4a3c-2d1e-0f9a-8b7c6d5e4f3a -->
<!-- DO NOT EDIT: This file is managed centrally in ghcommon repository -->
<!-- To update: Create an issue/PR in jdfalk/ghcommon -->

---
applyTo: "**/*.{js,jsx}"
description: |
  JavaScript language-specific coding, documentation, and testing rules for Copilot/AI agents and VS Code Copilot customization. These rules extend the general instructions in `general-coding.instructions.md` and merge all unique content from the Google
JavaScript Style Guide.

---

# JavaScript Coding Instructions

- Follow the [general coding instructions](general-coding.instructions.md).
- Follow the
  [Google JavaScript Style Guide](https://google.github.io/styleguide/jsguide.html)
  for additional best practices.
- All JavaScript files must begin with the required file header (see general
  instructions for details and JavaScript example).

## File Structure

- Use `camelCase` for file names
- Each file should contain exactly one ES module
- Prefer ES6 modules (`import`/`export`) over other module systems

## Formatting

- Use 2 spaces for indentation
- Line length maximum of 80 characters
- Use semicolons to terminate statements
- Use single quotes for string literals
- Add trailing commas for multi-line object/array literals
- Use parentheses only where required for clarity or priority

## Naming Conventions

- `functionNamesLikeThis`, `variableNamesLikeThis`, `ClassNamesLikeThis`,
  `EnumNamesLikeThis`, `methodNamesLikeThis`, `CONSTANT_VALUES_LIKE_THIS`,
  `private_values_with_underscore`, `package-names-like-this`

## Comments

- Use JSDoc for documentation
- Comment all non-obvious code sections

## Language Features

- Use `const` and `let` instead of `var`
- Use arrow functions for anonymous functions, especially callbacks
- Use template literals instead of string concatenation
- Use object/array destructuring where it improves readability
- Use default parameters instead of conditional statements
- Use spread operator and rest parameters where appropriate

## Best Practices

- Avoid using the global scope
- Avoid deeply nested code blocks
- Use early returns to reduce nesting
- Limit line length to improve readability
- Separate logic and display concerns

## Error Handling

- Always handle Promise rejections
- Use try/catch blocks appropriately
- Provide useful error messages

## Testing

- Write unit tests for all code
- Use descriptive test names
- Follow AAA pattern (Arrange, Act, Assert)

## Required File Header

All JavaScript files must begin with a standard header as described in the
[general coding instructions](general-coding.instructions.md). Example for
JavaScript:

```js
// file: path/to/file.js
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174000
```
