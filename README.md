# reverse-hash
A simple program that generates random MD5 hashes until a particular string is found in the hash.

This program was inspired by Christian von Kleist's wonderful solution to SQL injections through hashed queries:
https://cvk.posthaven.com/sql-injection-with-raw-md5-hashes

This program is not the most efficient way to check for hashes containing the shortest possible SQL injection that can bypass a rudimentary login, but instead focuses on leveraging the concurrency of multiple CPU cores to speed up computation (why code good when more cpu do the trick?).

This is however a decent way to find specific strings in MD5 hashes should you have the need.

Below is an example output searching for "'or'4" with 8 goroutines on a Ryzen 7 3700X

```
input: 52668037446787535327598257308165296518
hex:   9d881c276f722734f474d63ee9437240
raw:   'or'4t>Cr@
time:  10m24.0585467s
```

Input is the string that was hashed, producing an output in hexadecimal and raw ASCII.