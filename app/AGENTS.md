# Frontend Guidelines

These rules apply to everything inside the `app/` folder (React + TypeScript PWA).

## Code style

- Follow the project `.editorconfig` for indentation (4 spaces for `.ts`/`.tsx`).
- Use functional React components and hooks.
- Keep imports using the `@/` alias for paths under `src/` when possible.
- Run `pnpm exec eslint src --ext ts,tsx` to lint the code.
- Run `pnpm exec tsc --noEmit` to check TypeScript types.
- Keep `pnpm-lock.yaml` in sync when dependencies change.

## Pull Requests

- Summarize frontend changes and include lint/type-check results in the PR description.

## Getting started

- Run `make install-deps` and make sure the packages are installed

## Developing

- Always document the work. Ensure related functions or components include meaningful\
  docstrings when relevant.
  - If more context about the product is needed, put a doc file in `docs/product/`

## Ending development

- Run `make lint` and make sure there are no lint issues.
  - If two or more lint rules conflict, disable the least relevant rule.
