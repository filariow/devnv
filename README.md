# devnv

> For this early stage the project is thought for working on ZSH + OH-MY-ZSH.

`devnv` saves project locations and helps you cd into projects folders and configure the shell with respect to the selected project.
When asked for changing directory, devnv will prompt bash code that you have to source to configure your current shell.

The prompted code simply changes the directory (cd) and executes a script file (.devnv) if present.

```console
# to add a new project
$ devnv add -f PROJ_FOLDER -p PROJ_NAME
# to cd into PROJ_FOLDER and configure the shell
$ devnv cd PROJ_NAME # inspect the code
$ source <(devnv cd PROJ_NAME)
```

you can use the "pm" companion for a simpler notation


## PM companion

It is a simple facility script that wraps the `source <(devnv cd PROJ_NAME)` command for an easier notation and a better completion support.

Dependencies:
- jq: JSON query tool
- fzf: fuzzy finder


## Install

The installation process compiles and install the devnv binary into `${GOPATH}/bin/devnv`.

> To be able to run `devnv` ensure to have `${GOPATH}/bin` in your `${PATH}`.


To install the project clone the repository and use the provided [Makefile](./Makefile)'s `install` rule:

```
$ git clone https://github.com/filariow/devnv
$ cd devnv
$ make install
```

As suggested by the make install process, add the following lines into your `~/.zshrc` file:

```
#####################################
## devnv
#####################################
source <(devnv completion oh-my-zsh)
source ${HOME}/.devnv/scripts/pm.sh
```


## Uninstall

To uninstall devnv you can use the [Makefile](./Makefile)'s `uninstall` rule:

```console
$ make uninstall
```

or remove the binary and script manually

```console
$ rm ${HOME}/.local/bin/pm.sh
$ rm ${GOPATH}/bin/devnv
```


## Contribute

Please refer to the [TODO](./TODO.md) file for a list of feature to implement.

