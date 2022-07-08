## 소스 복사 후 환경 설정

- **git clone 소스코드**

```jsx
cd {최상위}/go/src/github.com/davidboybob/

git clone https://github.com/davidboybob/blockchainEduProject.git
```

- **go.mod 설정**

```jsx
cd $HOME/go/src/github.com/davidboybob/blockchainEduProject

go mod init # go.mod 초기화

go mod tidy # 소스 코드 패키지 동기화(다운로드)
```

<aside>
💡 Warning) go mod init {name} ⇒ name 명 달라졌을 시, 모든 소스 코드 import 문을 변경해줘야 함.

</aside>

- **Install Cobra, Cobra-cli**

```jsx
go get -u github.com/spf13/cobra@latest
go get github.com/spf13/cobra-cli@latest
```

- **Initalizing a Cobra CLI application**

```jsx
cd $HOME/go/src/github.com/davidboybob/blockchainEduProject
cobra-cli init
go run main.go
```

<aside>
💡 Path 시스템 환경 변수 설정

명령어) go env 
set GOBIN=C:\Users\david\go\bin
set GOPATH=C:\Users\david\go

</aside>

- **Go build**

```jsx
$ go build -o block.exe .

--- build options ---
-a
Force rebuilding of packages that are already up-to-date.

-n
Print the commands but do not run them.

-p n
The number of programs, such as build commands or test binaries, that can be run in parallel.
The default is the number of CPUs available.

-race
Enable data race detection.
Supported only on linux/amd64, freebsd/amd64, darwin/amd64, windows/amd64, linux/ppc64le and linux/arm64 (only for 48-bit VMA).

-msan
Enable interoperation with memory sanitizer.
Supported only on linux/amd64, linux/arm64, and only with Clang/LLVM as the host C compiler.
On linux/arm64, pie build mode will be used.

-v
Print the names of packages as they are compiled.

-work
Print the name of the temporary work directory and do not delete it when exiting.

-x
Print the commands.

-asmflags '[pattern=]arg list’
Arguments to pass on each go tool asm invocation.

-buildmode mode
Build mode to use. See ‘go help buildmode’ for more.

-compiler name
Name of compiler to use, as in runtime.Compiler (gccgo or gc).

-gccgoflags '[pattern=]arg list’
Arguments to pass on each gccgo compiler/linker invocation.

-gcflags '[pattern=]arg list’
Arguments to pass on each go tool compile invocation.

-installsuffix suffix
A suffix to use in the name of the package installation directory, in order to keep output separate from default builds.
If using the -race flag, the install suffix is automatically set to race or, if set explicitly, has _race appended to it. Likewise for the -msan flag. Using a -buildmode option that requires non-default compile flags has a similar effect.

-ldflags '[pattern=]arg list’
Arguments to pass on each go tool link invocation.

-linkshared
Build code that will be linked against shared libraries previously created with -buildmode=shared.

-mod mode
Module download mode to use: readonly, vendor, or mod.
See ‘go help modules’ for more.

-modcacherw
Leave newly-created directories in the module cache read-write instead of making them read-only.

-modfile file
In module aware mode, read (and possibly write) an alternate go.mod file instead of the one in the module root directory. A file named “go.mod” must still be present in order to determine the module root directory, but it is not accessed. When -modfile is specified, an alternate go.sum file is also used: its path is derived from the -modfile flag by trimming the “.mod” extension and appending “.sum”.

-pkgdir dir
Install and load all packages from dir instead of the usual locations.
For example, when building with a non-standard configuration, use -pkgdir to keep generated packages in a separate location.

-tags tag,list
A comma-separated list of build tags to consider satisfied during the build. For more information about build tags, see the description of build constraints in the documentation for the go/build package.

(Earlier versions of Go used a space-separated list, and that form is deprecated but still recognized.)

-trimpath
Remove all file system paths from the resulting executable.
Instead of absolute file system paths, the recorded file names will begin with either “go” (for the standard library), or a module path@version (when using modules), or a plain import path (when using GOPATH).

-toolexec 'cmd args’
A program to use to invoke toolchain programs like vet and asm.
For example, instead of running asm, the go command will run ‘cmd args /path/to/asm <arguments for asm>’.

```

- **Error**

```jsx

#1) go build error
$ go build
# github.com/mattn/go-sqlite3
cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%

Commnet) windows
gcc (the GNU Compiler Collection) provides a C compiler. 
On Windows, install TDM-GCC. The github.com/miekg/pkcs11 package uses cgo. 
Cgo enables the creation of Go packages that call C code.
다운로드 주소 : https://jmeubank.github.io/tdm-gcc/download/

해당 사이트 들어가서 다운로드 받음 (minimal)

Ubuntu

If you are running Ubuntu do:

apt-get install build-essential
This solved the problem. It installs the gcc/g++ compilers and libraries.

import error ## 삽질 엄청함...

settings.json 로 이동.
f1 > setting.json 입력 후 클릭
아래 내용 추가...
"gopls": {
        "experimentalWorkspaceModule": true,
},

---
#2 database create error
$ ./blockchainEduProject.exe init
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x1 addr=0x20 pc=0x7ff7cdcdee93]

goroutine 1 [running]:
database/sql.(*DB).conn(0x0, {0x7ff7cdef6fc8, 0xc0000160d8}, 0x1)
        C:/Go/src/database/sql/sql.go:1290 +0x53
database/sql.(*DB).prepare(0x0?, {0x7ff7cdef6fc8, 0xc0000160d8}, {0x7ff7cdeb07ee, 0x97}, 0x0?)
        C:/Go/src/database/sql/sql.go:1588 +0x45
database/sql.(*DB).PrepareContext(0x0?, {0x7ff7cdef6fc8, 0xc0000160d8}, {0x7ff7cdeb07ee, 0x97})
        C:/Go/src/database/sql/sql.go:1557 +0x98
database/sql.(*DB).Prepare(...)
        C:/Go/src/database/sql/sql.go:1578
github.com/davidboybob/blockchainEduProject/data.CreateTable()
        C:/Users/david/go/src/github.com/davidboybob/blockchainEduProject/data/data.go:31 +0x3e
github.com/davidboybob/blockchainEduProject/cmd.glob..func2(0x7ff7cde40e40?, {0x7ff7cdea0b1a?, 0x0?, 0x0?})
        C:/Users/david/go/src/github.com/davidboybob/blockchainEduProject/cmd/init.go:19 +0x17
github.com/spf13/cobra.(*Command).execute(0x7ff7cde40e40, {0x7ff7ce08b800, 0x0, 0x0})
        C:/Users/david/go/pkg/mod/github.com/spf13/cobra@v1.5.0/command.go:876 +0x67b
github.com/spf13/cobra.(*Command).ExecuteC(0x7ff7cde415c0)
        C:/Users/david/go/pkg/mod/github.com/spf13/cobra@v1.5.0/command.go:990 +0x3b4
github.com/spf13/cobra.(*Command).Execute(...)
        C:/Users/david/go/pkg/mod/github.com/spf13/cobra@v1.5.0/command.go:918
github.com/davidboybob/blockchainEduProject/cmd.Execute()
        C:/Users/david/go/src/github.com/davidboybob/blockchainEduProject/cmd/root.go:33 +0x25
main.main()
        C:/Users/david/go/src/github.com/davidboybob/blockchainEduProject/main.go:12 +0x17

해결)
main.go 파일에서 data.OpenDataBase() 코드를 입력 안함. (connection 오류)

package main

import (
	"github.com/davidboybob/blockchainEduProject/cmd"
	"github.com/davidboybob/blockchainEduProject/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
---
```

- CLI Install
```
go install {name}

*{name} : same as go mod init {name}

=> bin/{name}.exe 파일 생김

{name} 을 cli 명령어로 사용 로 가능.
```