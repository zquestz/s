#!/bin/bash

# bash/zsh provider completion support for s
#
# Usage:
#
# 1. Have s installed
# 2. Add the following lines to .bash_profile or .zshrc
#
# if [ -f $GOPATH/src/github.com/zquestz/s/autocomplete/s-completion.bash ]; then
#     . $GOPATH/src/github.com/zquestz/s/autocomplete/s-completion.bash
# fi

if [[ -n ${ZSH_VERSION-} ]]; then
    autoload -U +X compinit && compinit
    autoload -U +X bashcompinit && bashcompinit
fi

_provider_completion()
{
    local cur=${COMP_WORDS[COMP_CWORD]}
    local prev=${COMP_WORDS[COMP_CWORD-1]}
    local longOpts="--binary --cert --help --key --list-providers --list-tags --output --port --provider --server --tag --verbose --version"

    # _filedir (in newer bash completions) is a huge improvement on compgen -f or compgen -G
    # because it deals correctly with spaces, ~ expansion, and .inputrc preferences.
    comp      () { COMPREPLY=( $(compgen "$@" -- "$cur") ); }
    comp_path () { if type _filedir >/dev/null; then _filedir ; else comp -G $cur\* ; fi; }

    case "$cur" in
        -) comp -W "-b -c -h -k -l -o -p -s -t -v" && return 0 ;;
       -*) comp -W "$longOpts" && return 0 ;;
    esac

    # show the long options if current word is empty and previous one
    # isn't an option which expects an argument
    case "$prev" in
                       -p|--provider) comp -W "$(s --list-providers)" ;;
                            -t|--tag) comp -W "$(s --list-tags)" ;;
      -c|--cert|-k|--key|-b|--binary) comp_path ;;
                                   *) [[ -z "$cur" ]] && comp -W "$longOpts" ;;
    esac


} && complete -o filenames -F _provider_completion s
