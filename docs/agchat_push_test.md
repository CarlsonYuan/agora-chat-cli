## agchat push test

Test whether the push notification credentials and notification services work properly

```
agchat push test [flags]
```

### Examples

```
# Send a test push notification to a specific user
$ agchat push test --user <user-id>

```

### Options

```
  -h, --help             help for test
  -m, --message string   JSON string for the push message (default "{\"title\": \"Admin sent you a message\", \"content\": \"For push notification testing\", \"sub_title\": \"Test message is sent\"}")
  -u, --user string      the user ID of the target user
```

### Options inherited from parent commands

```
  -v, --verbose   enable verbose output
```

### SEE ALSO

* [agchat push](agchat_push.md)	 - Manage push notifications

