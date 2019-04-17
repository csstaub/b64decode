A base64 decoding CLI that decodes standard encoding, url-safe encoding, with
or without padding. If you're dealing with a lot of url-safe base64-encoded data,
or data that sometimes doesn't have the padding, this will be more robust
than `base64 --decode`.

# Install

```
go get github.com/csstaub/b64decode
```

# Examples

With padding
```
❯❯❯ echo 'SGVsbG8gV29ybGQhCg==' | b64decode
Hello World!
```

Without padding
```
❯❯❯ echo 'SGVsbG8gV29ybGQhCg' | b64decode
Hello World!
```

Understands standard and url-safe encoding
```
❯❯❯ echo '8J-Ygwo=' | b64decode
😃
```
