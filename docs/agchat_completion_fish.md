## agchat completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	agchat completion fish | source

To load completions for every new session, execute once:

	agchat completion fish > ~/.config/fish/completions/agchat.fish

You will need to start a new shell for this setup to take effect.


```
agchat completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -v, --verbose   enable verbose output
```

### SEE ALSO

* [agchat completion](agchat_completion.md)	 - Generate the autocompletion script for the specified shell

