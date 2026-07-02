# Task responsibility-boundary validator prompts

## Purpose

Provide the prompt fragments used to evaluate one Task record against the canonical Task responsibility contract.

These files are derived evaluator assets.
They are not canonical Task authoring authority.

## Canonical authority

Use `spec:product.design_records.authoring_standards.task_authoring` as the normative source.

When a prompt fragment conflicts with that Specification, the Specification wins.
Do not load this skill as normal Task authoring guidance.

## Prompt composition

For one Task evaluation, compose these files in order:

1. `prompts/evaluator-instructions.md`;
2. `prompts/common.md`;
3. `prompts/task-types/<task_type>.md` for the declared canonical `task_type`.

Evaluate every composed criterion.
Do not add or remove criteria at invocation time.

## Boundary

This skill owns:

- evaluator instruction text;
- common responsibility criteria;
- one type-specific criterion file for each canonical Task type;
- deterministic prompt-fragment selection and order.

This skill does not own:

- canonical Task authoring rules;
- structural Task validation;
- model, provider, runtime, retry, timeout, or decode policy;
- external validator response field names;
- validator implementation or DRMCP integration;
- Task correction, release, or human exception decisions.
