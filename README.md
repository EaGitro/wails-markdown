# README

## About

This is the official Wails Vanilla-TS template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: [https://wails.io/docs/reference/project-config](https://wails.io/docs/reference/project-config)

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.


> aaaaa
> bbbbb


## Code

```py
a = 10
def func():
    return "hoge"

```

```go
package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}
```

```riscv
 	.file	"call-func.c"
	.intel_syntax noprefix
	.text
	.globl	plus
	.type	plus, @function
plus:
	endbr64
	push	rbp
	mov	rbp, rsp
	mov	DWORD PTR -4[rbp], edi
	mov	DWORD PTR -8[rbp], esi
	mov	edx, DWORD PTR -4[rbp]
	mov	eax, DWORD PTR -8[rbp]
	add	eax, edx
	pop	rbp
	ret
;; 省略


main:
	endbr64
	push	rbp
	mov	rbp, rsp
	mov	esi, 4
	mov	edi, 3
	call	plus
	pop	rbp
	ret
;; 省略
 
```

