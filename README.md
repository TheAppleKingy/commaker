# commaker

**commaker** is a lightweight CLI tool that helps you automatically generate meaningful Git commit messages using an AI model. Optionally, it can also push your commit to a remote repository.

- [Install](#install)
- [Quickstart](#quickstart)
- [Usage](#usage)

---

## ⚙️ Install

Before using tool, define the following environment variables:

| Variable              | Description                                                                  |
|-----------------------|------------------------------------------------------------------------------|
|  OPEN_AI_KEY          | Your API key from the OpenRouter                                        |
|  MODEL			    | The model identifier (e.g., qwen/qwen3-coder:free)
|  AI_MODEL_HOST        |The full API URL to the AI chat endpoint                                                    


---

## 🚀 Quickstart

1. Create account on OpenRouter - [https://openrouter.ai/](https://openrouter.ai/)
2. Get API key and set as `OPEN_AI_KEY` env variable.
3. Set 'https://openrouter.ai/api/v1/chat/completions' at `AI_MODEL_HOST` in env.
4. Pick a supported model and assign it to the `MODEL` env variable
(e.g., qwen/qwen3-coder:free)
5. Download the appropriate binary from the [Releases](https://github.com/TheAppleKingy/commaker/releases) page.
6. Make the binary executable and add it to your system's `PATH`
(e.g., move it to /usr/local/bin/ on Linux/macOS).

---

## 📦 Usage

Commands:  
 - `commit` - Create a commit with an AI-generated message based on staged changes.

Options:  
 - `-p` - Push the commit to the remote repository after committing.

Examples:  
 - `commaker commit` - Create a commit using AI-generated message from current staged diff.

 - `commaker commit -p` - Commit and push to the current branch on origin.


