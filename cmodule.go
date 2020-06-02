package main

// gcc --version
//
// Configured with: --prefix=/Applications/Xcode.app/Contents/Developer/usr --with-gxx-include-dir=/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/include/c++/4.2.1
// Apple clang version 11.0.0 (clang-1100.0.33.16)
// Target: x86_64-apple-darwin18.7.0
// Thread model: posix
// InstalledDir: /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin

/*
typedef struct {
	int a;
} Foo;
#include <stdio.h>
void ppint(int a) {
	printf("%d\n", a);
}
*/
import "C"

func foo() interface{} {
	f := &C.Foo{}
	f.a = 3
	return f
}

func printInt(a int) {
	C.ppint(C.int(a))
}
