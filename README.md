
# antibote

The Antidote for GitHub botnets to uncover botnets with shared operators.



## What

A lot of GitHub accounts are bots that are parts of the same
shared botnet. Turns out, different botnet operators have
different PGP keys that they use to push commits to those
accounts.

This tool scrapes a user's social connections on GitHub and
maps the keys back to the fake accounts, so that you can trace
which botnet operator (co-)controls what fake accounts.

## Why

Multiple accounts use the same PGP key, and this tool
tries to scrape related accounts of a botnet user's followers,
repositories, and contributors (pull requests) to uncover
the botnet on a larger scale.

Usually lots of botnets work in a way that they create other
fake accounts and follow each other once they go online, so
that these accounts get pushed into "not being a bot" by
GitHub's very flawed bot detector.

## How

```bash
# Generate Personal Access Token
echo "My-Personal-Access-Token" > github/Token.env;

# Start tracing the botnet behind fake user account
go run cmds/antibote/main.go xiexinch;

# GPG key for this botnet operator was: B5690EEEBB952194
cat ~/Antibote/github/xiexinch.json;
cat ~/Antibote/keymap.json;
```

