
### ENVs

* `GITHUB_DIR`
* `FLINK_DIR`

### Tips

```zsh
du -csh dir # show dir size
```

### zshrc

```zsh
export JAVA_HOME=$(/usr/libexec/java_home -v 11)

alias pyfmt="yapf --in-place --style='{based_on_style: pep8, indent_width: 2}'"
```

### JDK

```zsh
brew tap homebrew/homebrew-core
brew tap AdoptOpenJDK/openjdk
brew cask install adoptopenjdk11
# .zshrc
export JAVA_HOME=$(/usr/libexec/java_home -v 11)
```

### Go Protobuf

```zsh
brew install protobuf # or, brew upgrade protobuf
protoc --version
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
```
