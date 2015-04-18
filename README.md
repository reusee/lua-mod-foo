# Build
```bash
go build -buildmode=c-shared -o foo.so
```

# Test
```
~> lua
Lua 5.2.4  Copyright (C) 1994-2015 Lua.org, PUC-Rio
> require 'foo'
> print(foo())
Foo!
> 
```
