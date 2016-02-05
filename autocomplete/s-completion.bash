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

    if [[ "$prev" == "-p" ]] || [[ "$prev" == "--provider" ]]; then
        # Get all providers using `s -l`
        COMPREPLY=( $(compgen -W "$(s -l)" -- $cur) )
    fi
}
complete -F _provider_completion s
