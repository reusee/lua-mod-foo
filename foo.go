package main

/*
#include <lua.h>
#cgo pkg-config: lua

extern void reg(lua_State*);
*/
import "C"

func main() {}

//export luaopen_foo
func luaopen_foo(l *C.lua_State) C.int {
	C.reg(l)
	return 0
}

//export Foo
func Foo(l *C.lua_State) C.int {
	C.lua_pushstring(l, C.CString("Foo!"))
	return 1
}
