package main
import (
	"flag"
	"fmt"
	"os"
)
/*
   subject : Command-Line Subcommands
*/
func main() {

	//We declare a subcommand using the NewFlagSet function,
	//and proceed to define new flags specific for this subcommand.
	

	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	//For a different subcommand we can define different supported flags.

	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	//The subcommand is expected as the first argument to the program.

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	//Check which subcommand is invoked.

	switch os.Args[1] {

	//For every subcommand, we parse its own flags and have access to trailing positional arguments.

	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

// Akira@MBP practice % go build process08.go 

// Akira@MBP practice % ./process08 
// expected 'foo' or 'bar' subcommands

// Akira@MBP practice % ./process08 foo -enable -name=joe a1 a2
// subcommand 'foo'
//   enable: true
//   name: joe
//   tail: [a1 a2]

// Akira@MBP practice % ./process08 bar -level 8 a1
// subcommand 'bar'
//   level: 8
//   tail: [a1]

// Akira@MBP practice % ./process08 bar -enable a1
// flag provided but not defined: -enable
// Usage of bar:
//   -level int
//     	level
// Akira@MBP practice %
