# author
author is a tool to get the history of commits of a specific author

<!-- ロゴ -->
<img width="200" alt="Point Ameba" src="https://github.com/Pianoopera/author/assets/42969626/4a2276d8-6101-442d-b589-60224560fd1e">

## Description
You can use it to get the history of commits of a specific author in all the repositories in a directory.

<!-- install -->
## Install

### for Mac
```bash
brew install Pianoopera/tap/author
```

### for Ubuntu
Please download author_Linux_x86_64.tar.gz from the [release page](https://github.com/Pianoopera/author/releases)

<!-- 使い方 -->
## Way to Use Something

### Usage
```bash
author you/dir/path account_name
```

### Information
```bash
author is a tool to get the history of commits of a specific author
You can use it to get the history of commits of a specific author in all the repositories in a directory.

Usage:
  author [flags]
  author [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     author version

Flags:
  -a, --accounts strings   accounts to search for
  -h, --help               help for author
  -s, --since string       since how many months ago (default "3")

Use "author [command] --help" for more information about a command.
```

### subCommand
```bash
author version
author help
```

### Option

-h, --help | author command help
```bash
author --help
```
---

-a, --accounts strings | accounts to search for
```bash
author you/dir/path account_name --accounts another_account_name1,another_account_name2
```

---

-s, --since string | since how many months ago (default "3")
```bash
author author you/dir/path account_name --since 1
```