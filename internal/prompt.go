package internal

const PROMPT string = `
Hi! Forget the previous context and generate a response as if this is my first message.
Create commit message based on git diff i'm sending. Message must strictly match this template:

Brief description of changes:
<here is brief description>
added:
- what entities, functions, methods, classes or logic are added etc
removed:
- what entities, functions, methods, classes or logic are removed etc
modified:
- what entities, functions, methods, classes or logic are modified etc
moved:
- what entities, functions, methods, classes or logic are moved etc
Profit(or brief summary)

Send me only commit message strictly following the template and no more words!. Here is the diff:
`
