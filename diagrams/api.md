# API Design

This file will be the brainstorming of how our core API will work with our plug-ins. I think I've discovered a good flow on how things can work out

## Plug-in dependencies

If plug-ins can be able to list dependencies, the core will be able to work out the order in which plug-ins should load in. Meaning things like:

- event subscriptions
- event emitting

will be able to properly work because they will already exist

For example, if A emits EventA, and b emits EventB and depends on EventA, A should load in first before B. B has a dependency on A ( []string { "A" } assuming A is the id 