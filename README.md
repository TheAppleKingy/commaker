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
|  COMMAKER_AI_API_KEY  | Your API key.                                    
|  COMMAKER_CONFIG_PATH | Path to the config file.                                                 


---

## 🚀 Quickstart

1. Get API key to access the model.
2. Set API key at the environment variable `COMMAKER_AI_API_KEY`.
3. Create config file '.commaker.yaml' and set it at the environment variable `COMMAKER_CONFIG_PATH`. See config pattern at the project files.

---

## 📦 Usage

Commands:  
 - `commit` - Create a commit with an AI-generated message based on staged changes.

Options:  
 - `-p` - Push the commit to the remote repository after committing.

Examples:  
 - `commaker commit` - Create a commit using AI-generated message from current staged diff.

 - `commaker commit -p` - Commit and push to the current branch on origin.


