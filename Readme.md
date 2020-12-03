
### ENVs

* `GITHUBDIR`

### Tips

```zsh
du -csh dir # show dir size
```

### zshrc

```zsh
export JAVA_HOME=$(/usr/libexec/java_home -v 11)

alias python3=/usr/local/bin/python3.8
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
