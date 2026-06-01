# Agent Instructions

## Critical Rules

- **NEVER guess or hallucinate answers.** If unsure, say "I don't know" and research the answer first. A confident wrong answer is worse than admitting uncertainty. Always verify before stating something as fact.

## Before Every Commit

Always run `make verify` before committing. It runs fmt, vet, modernize, and lint. Do not commit if any check fails.
