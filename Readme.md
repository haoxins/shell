
### ENVs

* `GITHUB_DIR`

### zshrc

```zsh
export JAVA_HOME=$(/usr/libexec/java_home -v 11)

alias pyfmt="yapf --in-place --style='{based_on_style: pep8, indent_width: 2}'"
```

### JDK

```zsh
brew tap homebrew/homebrew-core
brew tap AdoptOpenJDK/openjdk
brew install adoptopenjdk11 --cask
# .zshrc
export JAVA_HOME=$(/usr/libexec/java_home -v 11)
```

### Online tools

* [Excalidraw](https://excalidraw.com)
  - [GitHub](https://github.com/excalidraw/excalidraw)
* [AsyncAPI](https://www.asyncapi.com)
