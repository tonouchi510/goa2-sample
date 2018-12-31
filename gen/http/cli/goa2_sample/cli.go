// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// goa2-sample HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	usersc "github.com/tonouchi510/goa2-sample/gen/http/users/client"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `users (list|show|add|update|remove)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` users list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		usersFlags = flag.NewFlagSet("users", flag.ContinueOnError)

		usersListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		usersShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		usersShowIDFlag = usersShowFlags.String("id", "REQUIRED", "ID of user to show")

		usersAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		usersAddBodyFlag = usersAddFlags.String("body", "REQUIRED", "")

		usersUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		usersUpdateBodyFlag = usersUpdateFlags.String("body", "REQUIRED", "")
		usersUpdateIDFlag   = usersUpdateFlags.String("id", "REQUIRED", "ID of user to show")

		usersRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		usersRemoveIDFlag = usersRemoveFlags.String("id", "REQUIRED", "ID of user to remove")
	)
	usersFlags.Usage = usersUsage
	usersListFlags.Usage = usersListUsage
	usersShowFlags.Usage = usersShowUsage
	usersAddFlags.Usage = usersAddUsage
	usersUpdateFlags.Usage = usersUpdateUsage
	usersRemoveFlags.Usage = usersRemoveUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if len(os.Args) < flag.NFlag()+3 {
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = os.Args[1+flag.NFlag()]
		switch svcn {
		case "users":
			svcf = usersFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(os.Args[2+flag.NFlag():]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = os.Args[2+flag.NFlag()+svcf.NFlag()]
		switch svcn {
		case "users":
			switch epn {
			case "list":
				epf = usersListFlags

			case "show":
				epf = usersShowFlags

			case "add":
				epf = usersAddFlags

			case "update":
				epf = usersUpdateFlags

			case "remove":
				epf = usersRemoveFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if len(os.Args) > 2+flag.NFlag()+svcf.NFlag() {
		if err := epf.Parse(os.Args[3+flag.NFlag()+svcf.NFlag():]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "users":
			c := usersc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = usersc.BuildShowPayload(*usersShowIDFlag)
			case "add":
				endpoint = c.Add()
				data, err = usersc.BuildAddPayload(*usersAddBodyFlag)
			case "update":
				endpoint = c.Update()
				data, err = usersc.BuildUpdatePayload(*usersUpdateBodyFlag, *usersUpdateIDFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = usersc.BuildRemovePayload(*usersRemoveIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// usersUsage displays the usage of the users command and its subcommands.
func usersUsage() {
	fmt.Fprintf(os.Stderr, `users serves user relative information.
Usage:
    %s [globalflags] users COMMAND [flags]

COMMAND:
    list: List all stored users
    show: Show user by ID
    add: Add new user and return its ID.
    update: Update user item.
    remove: Remove user from storage

Additional help:
    %s users COMMAND --help
`, os.Args[0], os.Args[0])
}
func usersListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] users list

List all stored users

Example:
    `+os.Args[0]+` users list
`, os.Args[0])
}

func usersShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] users show -id INT64

Show user by ID
    -id INT64: ID of user to show

Example:
    `+os.Args[0]+` users show --id 8943934585442295493
`, os.Args[0])
}

func usersAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] users add -body JSON

Add new user and return its ID.
    -body JSON: 

Example:
    `+os.Args[0]+` users add --body '{
      "email": "Et voluptas neque voluptas doloribus.",
      "name": "hoge fuga"
   }'
`, os.Args[0])
}

func usersUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] users update -body JSON -id INT64

Update user item.
    -body JSON: 
    -id INT64: ID of user to show

Example:
    `+os.Args[0]+` users update --body '{
      "email": "Magni nostrum doloribus accusantium enim.",
      "name": "Sapiente recusandae."
   }' --id 5746158754770524014
`, os.Args[0])
}

func usersRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] users remove -id INT64

Remove user from storage
    -id INT64: ID of user to remove

Example:
    `+os.Args[0]+` users remove --id 6383285931504119743
`, os.Args[0])
}
