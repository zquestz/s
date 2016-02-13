function __fish_s_needs_option_argument
    set -l cmd (commandline -opc)
    set -l opt $argv[1]
    if test $cmd[-1] = $opt
        return 0
    else if [ (count $cmd) -ge 2 ]
        return (test $cmd[-2] = $opt)
    else
        return 1
    end
end

function __fish_s_provider_list
    s -l
end

complete -f -c s -s p
complete -f -c s -n '__fish_s_needs_option_argument -p' -a '(__fish_s_provider_list)'
