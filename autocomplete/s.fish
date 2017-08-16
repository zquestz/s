function __fish_s_needs_option_argument
    set -l cmd (commandline -opc)
    set -l opt $argv[1]
    if test $cmd[-1] = $opt
        return 0
    else if [ (count $cmd) -ge 2 ]
        if test $cmd[-2] = $opt
            if not contains -- $cmd[-1] (s -l)
                return 0
            end
        end
    end
    return 1
end

function __fish_s_find_pem_files_matching_prefix
    set -l cmd (commandline -opc)
    set -l opt $argv[1]
    if test $cmd[-1] = $opt
        # user has not started typing, look for any .pem in (pwd)
        ls *.pem
    else
        # user has started typing a path. look for any .pem starting from the
        # path
        set path $cmd[-1]
        set prefix ''
        if test -f path
            set prefix "$path*.pem"
        else
            set path_as_list (echo $path | tr '/' '\n')
            if test (count $path_as_list) = 1
                set prefix "$path*.pem"
            else
                set path_prefix (echo $path_as_list[1..-2] | tr ' ' '/')
                set prefix "$path_prefix/$path_as_list[-1]*.pem"
            end
        end
        ls $prefix
    end
end

function __fish_s_find_executable_matching_prefix
    set -l cmd (commandline -opc)
    set -l opt $argv[1]
    set completions
    if test $cmd[-1] = $opt
        for path in $PATH
            ls $path
        end
    else
        for path in $PATH
            ls "$path/$cmd[-1]"
        end
    end
end

# s options
complete -f -c s -o b -l binary         -d 'binary to launch search URI'
complete -f -c s -o c -l cert           -d 'path to cert.pem for TLS'
complete -f -c s -o k -l key            -d 'path to key.pem for TLS'
complete -f -c s -o l -l list-providers -d 'list supported providers'
complete -f -c s      -l list-tags      -d 'list available tags'
complete -f -c s -o o -l output         -d 'output only mode'
complete -f -c s      -l port           -d 'server port (default 8080)'
complete -f -c s -o p -l provider       -d 'search provider (default "google")'
complete -f -c s -o s -l server         -d 'launch web server'
complete -f -c s -o t -l tag            -d 'search tag'
complete -f -c s -o v -l verbose        -d 'display URL when opening'
complete -f -c s      -l version        -d 'display version'

# s {-b|--binary} options
complete -c s -n '__fish_s_needs_option_argument -b' -a '(__fish_s_find_executable_matching_prefix -b)'
complete -c s -n '__fish_s_needs_option_argument --binary' -a '(__fish_s_find_executable_matching_prefix --binary)'

# s {-c|--cert} options
complete -c s -n '__fish_s_needs_option_argument -c' -a '(__fish_s_find_pem_files_matching_prefix -c)'
complete -c s -n '__fish_s_needs_option_argument --cert' -a '(__fish_s_find_pem_files_matching_prefix --cert)'

# s {-k|--key} options
complete -c s -n '__fish_s_needs_option_argument -k' -a '(__fish_s_find_pem_files_matching_prefix -k)'
complete -c s -n '__fish_s_needs_option_argument --key' -a '(__fish_s_find_pem_files_matching_prefix --key)'

# s {-p|--provider} options
complete -f -c s -n '__fish_s_needs_option_argument -p' -a '(s -l)'
complete -f -c s -n '__fish_s_needs_option_argument --provider' -a '(s -l)'

# s {-t|--tag} options
complete -f -c s -n '__fish_s_needs_option_argument -t' -a '(s --list-tags)'
complete -f -c s -n '__fish_s_needs_option_argument --tag' -a '(s --list-tags)'
