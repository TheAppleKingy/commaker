package internal

const PROMPT string = `
Hi! Forget the previous context and generate a response as if this is my first message.
Create a commit message based on the git diff I'm sending. The message must strictly follow this template:

<type>(<scope>): <Brief description in imperative mood>

<detailed description of changes (why, what impact)>

Added:
- what entities, functions, methods, classes or logic are added
Removed:
- what entities, functions, methods, classes or logic are removed
Modified:
- what entities, functions, methods, classes or logic are modified
Moved:
- what entities, functions, methods, classes or logic are moved

Rules:
- Type must be one of: feat, fix, docs, style, refactor, test, chore, perf, ci
- Scope is optional, use lowercase with hyphens if needed (e.g., auth, user-profile, api)
- First line must not exceed 50 characters
- Leave a blank line between header and description
- Wrap description and bullet lists at 72 characters
- Use imperative mood (e.g., "add" not "added", "fix" not "fixes")
- Focus on explaining WHY the change was made, not just WHAT changed
- If this is a breaking change, add "BREAKING CHANGE:" in the footer

Send me only the commit message strictly following the template and no extra words! Here is the diff:
`
