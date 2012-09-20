package options

import (
	"fmt"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	_, err := Parse(`
    usage: haraway <flags>... <command> <args>...
    --
    root=     -r,--root=,HARAWAY_ROOT     Path to the haraway data root
    prefix=   -p,--prefix,HARAWAY_PREFIX  Path to the haraway install prefix.
    verbose   -v,--verbose                Show more info
    debug     -d,--debug,HARAWAY_DEBUG    Show debug info
    --
    --
    exec      exec                        Execute a command within the haraway sanbox
    shell     sh,shell                    Open a shell within the haraway sanbox
    --
    `)

	if err != nil {
		t.Error(err)
	}
}

func TestParse(t *testing.T) {
	spec, err := Parse(`
    usage: haraway <flags>... <command> <args>...
    --
    root=     -r,--root=,HARAWAY_ROOT     Path to the haraway data root
    prefix=   -p,--prefix,HARAWAY_PREFIX  Path to the haraway install prefix.
    verbose   -v,--verbose                Show more info
    debug     -d,--debug,HARAWAY_DEBUG    Show debug info
    --
    --
    exec      c,exec                      Execute a command within the haraway sanbox
    shell     sh,shell                    Open a shell within the haraway sanbox
    --
    `)
	if err != nil {
		t.Error(err)
	}

	opts, err := spec.Interpret([]string{"haraway", "-p", "/usr/local", "-r=hello", "-v", "c", "ls"}, []string{})

	if err != nil {
		t.Fatal(err)
	}

	if opts.Get("root") != "hello" {
		t.Error("--root != hello")
	}

	if opts.Get("verbose") != "true" {
		t.Error("--verbose != true")
	}

	if opts.GetBool("verbose") != true {
		t.Error("--verbose != true (bool)")
	}

	if strings.Join(opts.Args, " ") != "exec ls" {
		t.Errorf(".Args != [`exec`, `ls`] (was: %+v)", opts.Args)
	}
}

func ExampleParse() {
	spec, err := Parse(`
    usage: example-tool
    A short description of the command
    --
    flag        --flag,-f,FLAG           A description for this flag
    option=     --option=,-o=,OPTION=    A description for this option
                                         the description continues here
    !required=  --required,-r=,REQUIRED= A required option
    --
    env_var=    ENV_VAR=                 An environment variable
    --
    help        help,h                   Show this help message
    run         run                      Run some function
    --
    More freestyle text
    `)
	if err != nil {
		spec.PrintUsageWithError(err)
	}

	opts, err := spec.Interpret([]string{"example-tool", "--required", "hello world"}, []string{})
	if err != nil {
		spec.PrintUsageWithError(err)
	}

	fmt.Printf("required: %s", opts.Get("required"))

	// Output:
	// required: hello world
}
