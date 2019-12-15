package cmd

import (
	"bytes"

	"github.com/spf13/cobra"
	"helm.sh/helm/v3/cmd/helm/require"
)

func init() {
	completionCmd.AddCommand(zshCmd)
	completionCmd.AddCommand(bashCmd)
}

var zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: `Output zsh completion script for gobc.`,
	Long: `
Generates a zsh autocompletion script for gobc.

To load completion run

	. <(gobc completion zsh)
`,
	Args: require.NoArgs,
	RunE: func(command *cobra.Command, args []string) error {
		return runCompletionZsh(rootCmd)
	},
}

var bashCmd = &cobra.Command{
	Use:   "bash",
	Short: `Output bash completion script for gobc.`,
	Long: `
Generates a bash autocompletion script for gobc.

To load completion run

	. <(gobc completion bash)
`,
	Args: require.NoArgs,
	RunE: func(command *cobra.Command, args []string) error {
		return rootCmd.GenBashCompletion(command.OutOrStdout())
	},
}

func runCompletionZsh(cmd *cobra.Command) error {
	out := cmd.OutOrStdout()

	zshInitialization := `#compdef helm

__helm_bash_source() {
	alias shopt=':'
	alias _expand=_bash_expand
	alias _complete=_bash_comp
	emulate -L sh
	setopt kshglob noshglob braceexpand
	source "$@"
}
__helm_type() {
	# -t is not supported by zsh
	if [ "$1" == "-t" ]; then
		shift
		# fake Bash 4 to disable "complete -o nospace". Instead
		# "compopt +-o nospace" is used in the code to toggle trailing
		# spaces. We don't support that, but leave trailing spaces on
		# all the time
		if [ "$1" = "__helm_compopt" ]; then
			echo builtin
			return 0
		fi
	fi
	type "$@"
}
__helm_compgen() {
	local completions w
	completions=( $(compgen "$@") ) || return $?
	# filter by given word as prefix
	while [[ "$1" = -* && "$1" != -- ]]; do
		shift
		shift
	done
	if [[ "$1" == -- ]]; then
		shift
	fi
	for w in "${completions[@]}"; do
		if [[ "${w}" = "$1"* ]]; then
			echo "${w}"
		fi
	done
}
__helm_compopt() {
	true # don't do anything. Not supported by bashcompinit in zsh
}
__helm_ltrim_colon_completions()
{
	if [[ "$1" == *:* && "$COMP_WORDBREAKS" == *:* ]]; then
		# Remove colon-word prefix from COMPREPLY items
		local colon_word=${1%${1##*:}}
		local i=${#COMPREPLY[*]}
		while [[ $((--i)) -ge 0 ]]; do
			COMPREPLY[$i]=${COMPREPLY[$i]#"$colon_word"}
		done
	fi
}
__helm_get_comp_words_by_ref() {
	cur="${COMP_WORDS[COMP_CWORD]}"
	prev="${COMP_WORDS[${COMP_CWORD}-1]}"
	words=("${COMP_WORDS[@]}")
	cword=("${COMP_CWORD[@]}")
}
__helm_filedir() {
	local RET OLD_IFS w qw
	__debug "_filedir $@ cur=$cur"
	if [[ "$1" = \~* ]]; then
		# somehow does not work. Maybe, zsh does not call this at all
		eval echo "$1"
		return 0
	fi
	OLD_IFS="$IFS"
	IFS=$'\n'
	if [ "$1" = "-d" ]; then
		shift
		RET=( $(compgen -d) )
	else
		RET=( $(compgen -f) )
	fi
	IFS="$OLD_IFS"
	IFS="," __debug "RET=${RET[@]} len=${#RET[@]}"
	for w in ${RET[@]}; do
		if [[ ! "${w}" = "${cur}"* ]]; then
			continue
		fi
		if eval "[[ \"\${w}\" = *.$1 || -d \"\${w}\" ]]"; then
			qw="$(__helm_quote "${w}")"
			if [ -d "${w}" ]; then
				COMPREPLY+=("${qw}/")
			else
				COMPREPLY+=("${qw}")
			fi
		fi
	done
}
__helm_quote() {
	if [[ $1 == \'* || $1 == \"* ]]; then
		# Leave out first character
		printf %q "${1:1}"
	else
		printf %q "$1"
	fi
}
autoload -U +X bashcompinit && bashcompinit
# use word boundary patterns for BSD or GNU sed
LWORD='[[:<:]]'
RWORD='[[:>:]]'
if sed --help 2>&1 | grep -q 'GNU\|BusyBox'; then
	LWORD='\<'
	RWORD='\>'
fi
__helm_convert_bash_to_zsh() {
	sed \
	-e 's/declare -F/whence -w/' \
	-e 's/_get_comp_words_by_ref "\$@"/_get_comp_words_by_ref "\$*"/' \
	-e 's/local \([a-zA-Z0-9_]*\)=/local \1; \1=/' \
	-e 's/flags+=("\(--.*\)=")/flags+=("\1"); two_word_flags+=("\1")/' \
	-e 's/must_have_one_flag+=("\(--.*\)=")/must_have_one_flag+=("\1")/' \
	-e "s/${LWORD}_filedir${RWORD}/__helm_filedir/g" \
	-e "s/${LWORD}_get_comp_words_by_ref${RWORD}/__helm_get_comp_words_by_ref/g" \
	-e "s/${LWORD}__ltrim_colon_completions${RWORD}/__helm_ltrim_colon_completions/g" \
	-e "s/${LWORD}compgen${RWORD}/__helm_compgen/g" \
	-e "s/${LWORD}compopt${RWORD}/__helm_compopt/g" \
	-e "s/${LWORD}declare${RWORD}/builtin declare/g" \
	-e "s/\\\$(type${RWORD}/\$(__helm_type/g" \
	-e 's/aliashash\["\(.\{1,\}\)"\]/aliashash[\1]/g' \
	-e 's/FUNCNAME/funcstack/g' \
	<<'BASH_COMPLETION_EOF'
`
	if _, err := out.Write([]byte(zshInitialization)); err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	if err := cmd.Root().GenBashCompletion(buf); err != nil {
		return err
	}

	if _, err := out.Write(buf.Bytes()); err != nil {
		return err
	}

	zshTail := `
BASH_COMPLETION_EOF
}
__helm_bash_source <(__helm_convert_bash_to_zsh)
`
	_, err := out.Write([]byte(zshTail))
	return err
}
