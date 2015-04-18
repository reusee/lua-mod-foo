#include <lua.h>
#include "_cgo_export.h"

void reg(lua_State *l) {
  lua_pushcclosure(l, Foo, 0);
  lua_setglobal(l, "foo");
}
