# AGENTS Instructions

## Copilot Instructions

### File Identification

- Always check the first line of a file for a comment in the format `# file: $(relative_path_to_file)`.
- Use this path to determine which file you're working on and where to apply generated changes.
- If this comment is present, prioritize it over any other indications of file path.
- When generating code modifications, reference this path in your response.

### Code Documentation

- Document functions with parameters, return values, and purpose.
- Document methods with parameters, return values, and purpose.
- Document classes with their responsibilities and usage patterns.
- Document tests with clear descriptions of what's being tested and expected outcomes.
- Escape triple backticks with a backslash in documentation.
- Use the existing documentation style in the repository (Go doc comments).

### Documentation Organization Policy

- **README.md**: Repository introduction, setup instructions, basic usage, critical information. Temporary breaking changes should be listed at the top.
- **TODO.md**: Project roadmap, planning, implementation status, architectural decisions, reasoning behind choices, diagrams, and technical plans.
- **CHANGELOG.md**: Version information, release notes, feature additions, bug fixes, major breaking changes, and other consolidated technical documentation.

### Code Style

- Follow the established code style in the repository.
- Use consistent naming conventions for variables, functions, and structs.
- Prefer explicit type annotations where applicable.
- Keep functions small and focused on a single responsibility.
- Use meaningful variable names.

### Testing

- Write comprehensive tests for new functionality.
- Update documentation when tests change.
- Follow existing test naming conventions.
- Include edge cases and error handling in tests.
- Maintain test coverage when modifying existing code.

### Security & Best Practices

- Avoid hardcoding sensitive information.
- Use proper error handling and input validation.
- Consider performance implications of code changes.

### Version Control

- Write clear commit messages that explain the purpose of changes.
- Keep pull requests focused on a single feature or fix.
- Reference issue numbers in commits and PRs when applicable.

### Project-Specific

- Import from project modules instead of duplicating functionality.
- Respect the established architecture patterns.
