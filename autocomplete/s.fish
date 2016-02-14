function __fish_s_needs_option_argument
    set -l cmd (commandline -opc)
    set -l opt $argv[1]
    if test $cmd[-1] = $opt
        return 0
    else if [ (count $cmd) -ge 2 ]
        if test $cmd[-2] = $opt
            if not contains $cmd[-1] (s -l)
                return 0
            end
        end
    end
    return 1
end

complete -f -c s -s p
complete -f -c s -l provider
complete -f -c s -n '__fish_s_needs_option_argument -p' -a '(s -l)'
complete -f -c s -n '__fish_s_needs_option_argument --provider' -a '(s -l)'
