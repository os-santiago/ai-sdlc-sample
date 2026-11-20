# ai-sdlc-sample

A small Go module used as an external sample repository for
[ai-sdlc](https://github.com/os-santiago/ai-sdlc): issues on this repository
are delivered end to end by the runtime (verification with the repository's
own tests, independent review, pull request and merge).

```bash
go test ./...
```

The module has no dependencies, so its tests run offline.
