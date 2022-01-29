# boring

A bunch of boring utilities

Requires Go >= 1.18 (for generics)

## Iterators

Now that generics are a thing it makes it plausible to come up with a way to create something like iterators/generators
from other languages. This repo contains some experiments toward creating that.

The approach used here is to use goroutines feeding a channel which is simple and easy to use. It does have some
downsides, however. Anytime an iterator is not completely consumed, we leak a channel and goroutine. It's also easy to
make infinite iterators, which are very useful but will always leak (since they cannot be fully consumed)

To fix this we'd need a way to proactively cancel an iterator, but to do so seems likely to require a much less
clean interface.
