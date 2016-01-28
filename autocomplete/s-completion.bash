#!/bin/bash

# bash provider completion support for s
# 
# Usage:
#
# 1. Have s installed
# 2. Add the following lines to .bash_profile
#       if [ -f /path/to/s-completion.bash ]; then
#           . /path/to/s-completion.bash
#       fi

_provider_completion()
{
    local cur=${COMP_WORDS[COMP_CWORD]}
    local prev=${COMP_WORDS[COMP_CWORD-1]}

    if [[ "$prev" == "-p" ]]; then
        # Get all providers using `s -l`
        COMPREPLY=( $( compgen -W "$(s -l)" -- $cur ) )
    fi
}
complete -F _provider_completion s
