你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
