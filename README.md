# Golang GitHub App example

To-do: Update doc with step-by-step instructions applicable only to this Go app.
To-do: Submit to upstream as Go example app, alongside ruby example.

See [Building your first GitHub app](https://developer.github.com/apps/building-your-first-github-app/#one-time-setup)

```sh
export GITHUB_WEBHOOK_SECRET=<see lastpass>
export GITHUB_APP_IDENTIFIER=<see app registration>
export GITHUB_INSTALLATION_IDENTIFIER=<see installation URL>
export GITHUB_PRIVATE_KEY_FILE=path/to/your/private-key.pem
```

Development ngrok (or smee.io)

- Ensure ngrok is installed `which ngrok || brew cask install ngrok`
- Optional: [log in with github acct](https://dashboard.ngrok.com) and "connect your account"
- Start a tunnel: `ngrok http 3000`
