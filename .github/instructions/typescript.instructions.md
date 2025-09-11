<!-- file: .github/instructions/typescript.instructions.md -->
## TypeScript Style Guidelines

Follow the complete Google TypeScript Style Guide below for all TypeScript code:

### Google TypeScript Style Guide (Complete)

This guide is based on the internal Google TypeScript style guide and provides comprehensive standards for TypeScript development.

#### Key Principles
- Use structural typing over nominal typing
- Prefer `const` and `let` over `var`
- Use interfaces for object types, not type aliases
- Prefer named exports over default exports
- Use strict TypeScript compiler settings

#### Source File Structure
Files consist of the following, in order:
1. Copyright information, if present
2. JSDoc with `@fileoverview`, if present
3. Imports, if present
4. The file's implementation

Exactly one blank line separates each section that is present.

#### Imports and Exports
- Use paths to import other TypeScript code
- Prefer relative imports (`./foo`) for same-project files
- Use named imports for frequently used symbols
- Use namespace imports for large APIs
- **Never use default exports** - use named exports exclusively
- Export individual constants and functions, not container classes

```typescript
// Good: named exports
export class Foo { ... }
export const BAR = 1;

// Bad: default exports
export default class Foo { ... }
```

#### Variable Declarations
- Always use `const` or `let`, never `var`
- Use `const` by default, `let` only when reassigning
- One variable per declaration

```typescript
const foo = otherValue;  // Use if "foo" never changes
let bar = someValue;     // Use if "bar" is reassigned later
```

#### Array and Object Literals
- Use bracket notation for arrays: `[1, 2, 3]`
- Don't use the `Array()` constructor
- Use object literals `{}` instead of `Object` constructor
- Use spread syntax `{...obj}` for shallow copying

#### Classes
- Class declarations must not be terminated with semicolons
- Separate method declarations with single blank lines
- Use parameter properties for constructor parameters
- Initialize fields where declared when possible
- Use TypeScript visibility modifiers (`private`, `protected`, `public`)
- Don't use `#private` fields

```typescript
class Foo {
  private readonly userList: string[] = [];

  constructor(private readonly barService: BarService) {}

  doThing() {
    console.log("A");
  }
}
```

#### Functions
- Prefer function declarations for named functions
- Use arrow functions for callbacks and short expressions
- Don't use function expressions (use arrow functions instead)
- Use concise bodies for single expressions, block bodies otherwise

```typescript
// Good: function declaration
function foo() {
  return 42;
}

// Good: arrow function for callback
const items = list.map(item => item.name);

// Good: block body when return value unused
myPromise.then(v => {
  console.log(v);
});
```

#### Type System
- Rely on type inference for simple cases
- Use explicit types for complex expressions
- Prefer `unknown` over `any` when type is truly unknown
- Use structural types and interfaces
- Use `Array<T>` for complex types, `T[]` for simple types

```typescript
// Good: type inference
const x = 15;
const users = new Set<string>();

// Good: explicit type for complex expression
const value: string[] = await rpc.getSomeValue().transform();
```

#### Control Structures
- Always use braced blocks, even for single statements
- Use triple equals (`===`) and not equals (`!==`)
- Handle errors with proper Error instances
- Use proper exception handling with try/catch

```typescript
// Good: braced blocks
if (x) {
  doSomething();
}

// Good: strict equality
if (foo === 'bar') {
  // ...
}
```

#### Naming Conventions
- `lowerCamelCase` for variables, parameters, functions, methods, properties
- `UpperCamelCase` for classes, interfaces, types, enums, decorators
- `CONSTANT_CASE` for global constant values
- Use descriptive names, avoid abbreviations
- Don't use `_` prefix/suffix for private members

#### Comments and Documentation
- Use `/** JSDoc */` for documentation comments
- Use `//` for implementation comments
- Document all top-level exports
- Parameter and return descriptions may be omitted if obvious
- Write JSDoc before decorators

```typescript
/**
 * Calculates the total price including tax.
 * @param basePrice The price before tax
 * @param taxRate The tax rate as a decimal (e.g., 0.1 for 10%)
 */
function calculateTotal(basePrice: number, taxRate: number): number {
  return basePrice * (1 + taxRate);
}
```

#### Disallowed Features
- Don't use `const enum` (use regular `enum`)
- Don't use wrapper objects (`String`, `Boolean`, `Number`)
- Don't use `eval` or `Function` constructor
- Don't use `with` statement
- Don't use `@ts-ignore` (prefer proper typing)
- Don't rely on Automatic Semicolon Insertion

This covers the essential TypeScript style guidelines including:/instructions/typescript.instructions.md -->
<!-- version: 1.2.0 -->
<!-- guid: ts123456-e89b-12d3-a456-426614174000 -->
<!-- DO NOT EDIT: This file is managed centrally in ghcommon repository -->
<!-- To update: Create an issue/PR in jdfalk/ghcommon -->

---
applyTo: "**/*.ts"
description: |
  TypeScript language-specific coding, documentation, and testing rules for Copilot/AI agents and VS Code Copilot customization. These rules extend the general instructions in `general-coding.instructions.md` and merge all unique content from the Google TypeScript Style Guide.
---

# TypeScript Coding Instructions

- Follow the [general coding instructions](general-coding.instructions.md).
- Follow the
  [Google TypeScript Style Guide](https://google.github.io/styleguide/tsguide.html)
  for additional best practices.
- All TypeScript files must begin with the required file header (see general
  instructions for details and TypeScript example).

## Core Principles

- Readability: Code should be clear and understandable
- Consistency: Follow established patterns and conventions
- Type Safety: Use TypeScript features to catch errors at compile time
- Simplicity: Prefer simple, straightforward solutions

## File Organization

- Use `.ts` for TypeScript files, `.tsx` for TypeScript with JSX
- Use ES6 import/export style
- License header (if required), then imports, then main export, then
  implementation

## Naming Conventions

- camelCase for variables and functions
- PascalCase for classes and interfaces
- SCREAMING_SNAKE_CASE for module-level constants
- Use descriptive names, avoid abbreviations
- Use PascalCase for enum names and members

## Type Annotations

- Always annotate function parameters and return types
- Use arrow functions for short functions
- Use explicit types for complex objects
- Use interfaces for object shapes that might be extended
- Use type aliases for unions, primitives, or computed types
- Use extends for generic constraints
- Use built-in utility types

## Code Formatting

- Maximum 80 characters per line
- Use 2 spaces for indentation, no tabs
- Always use semicolons
- Use single quotes for strings, template literals for interpolation
- Use trailing commas in multiline structures

## Best Practices

- Use strict null checks
- Use array methods instead of loops
- Use object spread for copying
- Use destructuring for extraction
- Prefer async/await over promises
- Keep functions small and focused
- Use pure functions when possible
- Use function overloads for different signatures

## Testing

- Use descriptive test names
- Follow AAA pattern (Arrange, Act, Assert)
- Use table-driven tests for multiple scenarios

## Required File Header

All TypeScript files must begin with a standard header as described in the
[general coding instructions](general-coding.instructions.md). Example for
TypeScript:

```typescript
// file: path/to/file.ts
// version: 1.0.0
// guid: 123e4567-e89b-12d3-a456-426614174000
```
