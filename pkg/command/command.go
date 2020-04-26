/*
#######
##            _
##           (_)__ ________ ____  ___
##          / / _ `/ __/ _ `/ _ \/ _ \
##       __/ /\_,_/_/  \_, /\___/_//_/
##      |___/         /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package command

import (
	"errors"
	"flag"
	"fmt"
)

// ErrNoError AFAIRE
var ErrNoError = errors.New("no error")

// Command AFAIRE
type Command interface {
	Name() string
	Description() string
	Setup(args []string) error
	Run() error
}

// FlagSet AFAIRE
type FlagSet struct {
	*flag.FlagSet
	command Command
}

func (fs *FlagSet) printUsage() {
	fmt.Println()
	fmt.Println(fs.Name(), "-", fs.command.Description())
	fmt.Println("=====================================================================================================")
	fmt.Println("Options:")
	fs.PrintDefaults()
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println()
}

// NewFlagSet AFAIRE
func NewFlagSet(cmd Command) *FlagSet {
	fs := &FlagSet{
		FlagSet: flag.NewFlagSet(cmd.Name(), flag.ContinueOnError),
		command: cmd,
	}

	fs.FlagSet.Usage = fs.printUsage

	return fs
}

// Parse AFAIRE
func (fs *FlagSet) Parse(args []string) error {
	if err := fs.FlagSet.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return ErrNoError
		}

		return err
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
