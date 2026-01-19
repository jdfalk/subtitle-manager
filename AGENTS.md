<!-- file: AGENTS.md -->
<!-- version: 2.3.1 -->
<!-- guid: 2e7c1a4b-5d3f-4b8c-9e1f-7a6b2c3d4e5f -->

# AGENTS.md

This file is a **local index** of the canonical Copilot/AI agent instructions.
The actual rules live in the shared, reusable instruction system synced from
`ghcommon` under `.github/instructions/`. Do **not** duplicate those rules
locally unless they are explicitly present in this repository.

## Canonical Instruction Sources (Use These)

These files are the canonical rules and **must** be followed as-is:

- `.github/copilot-instructions.md`
- `.github/instructions/general-coding.instructions.md`
- `.github/instructions/commit-messages.instructions.md`
- `.github/instructions/pull-request-descriptions.instructions.md`
- `.github/instructions/security.instructions.md`
- `.github/instructions/test-generation.instructions.md`
- `.github/instructions/github-actions.instructions.md`
- `.github/instructions/protobuf.instructions.md`
- Language-specific rules in `.github/instructions/*.instructions.md`
  (Go, Python, JavaScript, TypeScript, Rust, Shell, R, HTML/CSS, JSON, Markdown)

## Local Notes (Only If Present Here)

The only local overrides or additions should be repository-specific guidance
that is **not** available in the shared instruction system. If a rule is already
covered by the shared `ghcommon` instruction files, do not restate it here.

## Documentation Update Protocol (Local Reminder)

- Edit documentation directly in target files.
- Keep the file header and bump the version on changes.
- Do not use legacy doc-update scripts:
  `create-doc-update.sh`, `doc_update_manager.py`, or `.github/doc-updates/`.
