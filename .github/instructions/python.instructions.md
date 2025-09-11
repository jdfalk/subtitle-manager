<!-- file: .## Python Style Guidelines

Follow the complete Google Python Style Guide below for all Python code:

### Google Python Style Guide (Complete)

This style guide is a comprehensive set of conventions for writing readable and maintainable Python code.

#### Python Language Rules

**Lint:** Use pylint to check your code. Suppress warnings with inline comments when appropriate.

```python
# Good: specific pylint disable
def connect_to_next_port(self, minimum):  # pylint: disable=invalid-name
    # Function name follows external API convention
    pass
```

**Imports:**
- Use import for packages and modules only, not individual classes/functions
- Use absolute imports when possible
- Use relative imports only within packages
- Imports should be on separate lines

```python
# Good
import os
import sys
from typing import List, Dict
from myproject.subpackage import mymodule

# Bad
from myproject.subpackage.mymodule import MyClass, my_function
import os, sys  # Multiple imports on one line
```

**Packages:** Import using full package names.

**Exceptions:**
- Use built-in exception classes when possible
- Create custom exceptions that inherit from Exception
- Never use bare except: clauses

```python
# Good
try:
    result = risky_operation()
except ValueError as e:
    logging.error('Invalid value: %s', e)
except (TypeError, AttributeError):
    logging.error('Type or attribute error occurred')

# Bad
try:
    result = risky_operation()
except:  # Too broad
    pass
```

**Global Variables:** Avoid global mutable state. Use module-level constants sparingly.

**Nested/Local/Inner Classes and Functions:** Fine to use when they close over local variables.

**Comprehensions & Generator Expressions:**
- Use for simple cases
- Multiple for clauses or filter expressions should use regular loops

```python
# Good
result = [mapping_expr for value in iterable if filter_expr]

# Bad - too complex
result = [(x, y) for x in range(10) for y in range(5) if x * y > 10]

# Better for complex logic
result = []
for x in range(10):
    for y in range(5):
        if x * y > 10:
            result.append((x, y))
```

**Default Iterators and Operators:** Use when they make code simpler.

**Lambda Functions:** Okay for one-liners. Use def for anything more complex.

**Conditional Expressions:** Okay for simple cases.

```python
# Good
x = 1 if condition else 2

# Bad - too complex
x = very_long_variable_name if some_complex_condition_that_spans_multiple_words else another_very_long_variable_name
```

**Default Argument Values:**
- Don't use mutable objects as default values
- Use None as default for mutable arguments

```python
# Good
def function(a, b=None):
    if b is None:
        b = []

# Bad
def function(a, b=[]):  # Mutable default
    pass
```

**Properties:** Use @property for accessing or setting data with simple logic.

**True/False Evaluations:** Use implicit false when possible.

```python
# Good
if not users:
if foo:

# Bad
if len(users) == 0:
if foo != []:
```

#### Python Style Rules

**Semicolons:** Never terminate statements with semicolons.

**Line Length:** Maximum 80 characters.

**Parentheses:**
- Use sparingly
- Don't use in return statements unless needed for clarity

```python
# Good
if foo:
    bar()
return value

# Bad
if (foo):
    bar()
return (value)
```

**Indentation:** Use 4 spaces, never tabs.

**Blank Lines:**
- Two blank lines between top-level definitions
- One blank line between method definitions
- Use blank lines sparingly within functions

**Whitespace:**
- No trailing whitespace
- Surround binary operators with spaces
- No space inside parentheses, brackets, or braces

```python
# Good
spam(ham[1], {eggs: 2})
if x == 4:
    print(x, y)
x, y = y, x

# Bad
spam( ham[ 1 ], { eggs: 2 } )
if x==4:
    print(x,y)
x,y = y,x
```

**Shebang Line:** Use `#!/usr/bin/env python3` for executable scripts.

**Comments and Docstrings:**

**Docstrings:** Required for all public modules, functions, classes, and methods.

```python
def calculate_distance(point1, point2):
    """Calculate the Euclidean distance between two points.

    Args:
        point1: A tuple (x, y) representing the first point.
        point2: A tuple (x, y) representing the second point.

    Returns:
        The Euclidean distance between point1 and point2.

    Raises:
        TypeError: If points are not tuples or lists.
        ValueError: If points don't have exactly 2 coordinates.
    """
    if not isinstance(point1, (tuple, list)) or not isinstance(point2, (tuple, list)):
        raise TypeError("Points must be tuples or lists")

    if len(point1) != 2 or len(point2) != 2:
        raise ValueError("Points must have exactly 2 coordinates")

    dx = point2[0] - point1[0]
    dy = point2[1] - point1[1]
    return (dx ** 2 + dy ** 2) ** 0.5
```

**Classes:**

```python
class SampleClass:
    """Summary of class here.

    Longer class information... Lorem ipsum dolor sit amet, consectetur
    adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore
    magna aliqua.

    Attributes:
        likes_spam: A boolean indicating if we like SPAM or not.
        eggs: An integer count of the eggs we have laid.
    """

    def __init__(self, likes_spam=False):
        """Initializes the instance based on spam preference.

        Args:
            likes_spam: A boolean indicating if we like SPAM or not.
        """
        self.likes_spam = likes_spam
        self.eggs = 0

    def public_method(self):
        """Performs operation blah."""
        pass
```

**Block and Inline Comments:**
- Use complete sentences
- Keep comments up to date
- Use # for line comments

```python
# We use a weighted dictionary search to find out where i is in
# the array. We extrapolate position based on the largest num
# in the array and the array size and then do binary search to
# get the exact number.

x = x + 1  # Compensate for border
```

**Classes:** Use CapWords convention.

**Functions:** Use lowercase_with_underscores.

**Variables:** Use lowercase_with_underscores.

**Constants:** Use CAPS_WITH_UNDERSCORES.

```python
# Good
class MyClass:
    CONSTANT_VALUE = 42

    def __init__(self):
        self.instance_variable = None

    def my_method(self):
        local_variable = "hello"
        return local_variable

# Bad
class myClass:  # Should be CapWords
    constantValue = 42  # Should be CAPS_WITH_UNDERSCORES

    def MyMethod(self):  # Should be lowercase_with_underscores
        LocalVariable = "hello"  # Should be lowercase_with_underscores
        return LocalVariable
```

**Files and Directories:** Use lowercase_with_underscores.py

**Type Annotations:**
- Use type hints for function signatures
- Use from __future__ import annotations for forward references

```python
from typing import List, Dict, Optional, Union

def process_items(items: List[str],
                  config: Dict[str, int]) -> Optional[str]:
    """Process a list of items with given configuration.

    Args:
        items: List of strings to process.
        config: Configuration dictionary.

    Returns:
        Processed result or None if processing fails.
    """
    if not items:
        return None

    # Process items according to config
    result = "processed"
    return result
```

**String Formatting:**
- Use f-strings for simple formatting
- Use .format() for more complex cases
- Avoid % formatting

```python
# Good
name = "John"
age = 30
message = f"Hello {name}, you are {age} years old"

# Good for complex formatting
template = "User {user} has {count} items"
message = template.format(user=username, count=item_count)

# Bad
message = "Hello %s, you are %d years old" % (name, age)
```

**Main Guard:**

```python
def main():
    """Main function."""
    pass

if __name__ == '__main__':
    main()
```

This covers the essential Python style guidelines including:/instructions/python.instructions.md -->
<!-- version: 1.5.0 -->
<!-- guid: 2a5b7c8d-9e1f-4a2b-8c3d-6e9f1a5b7c8d -->
<!-- DO NOT EDIT: This file is managed centrally in ghcommon repository -->
<!-- To update: Create an issue/PR in jdfalk/ghcommon -->

---
applyTo: "**/*.py"
description: |
  Python language-specific coding, documentation, and testing rules for Copilot/AI agents and VS Code Copilot customization. These rules extend the general instructions in `general-coding.instructions.md` and merge all unique content from the Google Python Style Guide.
---

# Python Coding Instructions

- Follow the [general coding instructions](general-coding.instructions.md).
- Follow the
  [Google Python Style Guide](https://google.github.io/styleguide/pyguide.html)
  for additional best practices.
- All Python files must begin with the required file header (see general
  instructions for details and Python example).

## Version Requirements

- **MANDATORY**: All Python projects must use Python 3.13.0 or higher
- **NO EXCEPTIONS**: Do not use older Python versions in any repository
- Update `requirements.txt` files to specify `python>=3.13.0`
- Update `pyproject.toml` files to specify `requires-python = ">=3.13.0"`
- Update `setup.py` files to specify `python_requires=">=3.13.0"`
- Update `.python-version` files to specify `3.13.0` or higher
- Use `python --version` to verify your installation meets requirements
- All Python file headers must use version 1.0.0 or higher (following semantic versioning)

## Core Principles

- Be consistent: Follow the established patterns in your codebase
- Readability counts: Code is read more often than it is written
- Simple is better than complex: Prefer clarity over cleverness
- Use tools: Leverage formatters like `ruff`, `black`, and `isort` for
  consistency

## Language Rules

- Use `pylint` with Google's pylintrc configuration
- Use absolute imports, never relative
- Avoid mutable global state
- Use nested functions only when closing over a local value
- Use comprehensions for simple cases
- Use default iterators and operators
- Use generators as needed
- Use lambda for one-liners, prefer generator expressions
- Use properties for simple computations
- Use implicit false when possible, `if foo is None:` for None checks
- Do not use mutable objects as default argument values
- Use 4 spaces for indentation, never tabs
- Maximum line length is 80 characters
- Use parentheses for implied line continuation
- Two blank lines between top-level definitions
- One blank line between method definitions
- Use triple-double-quotes for docstrings, Google-style docstring format
- Use f-strings, `%` operator, or `format()` for formatting
- Use `with` statements for resource management
- Use `# TODO:` comments with context
- Imports on separate lines, grouped stdlib/third-party/local
- Use descriptive names, avoid single character names
- Use type annotations where applicable
- Put main functionality in `main()` and check `if __name__ == '__main__':`
- Prefer small and focused functions

## Required File Header

All Python files must begin with a standard header as described in the
[general coding instructions](general-coding.instructions.md). Example for
Python:

```python
#!/usr/bin/env python3
# file: path/to/file.py
# version: 1.0.0
# guid: 123e4567-e89b-12d3-a456-426614174000
```
