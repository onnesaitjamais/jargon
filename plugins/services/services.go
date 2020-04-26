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

package main

import (
	"github.com/arnumina/jargon/pkg/command"
)

type cmdServices struct{}

func (cs *cmdServices) Name() string {
	return "services"
}

func (cs *cmdServices) Description() string {
	return "management of services"
}

func (cs *cmdServices) Setup(args []string) error {
	fs := command.NewFlagSet(cs)

	if err := fs.Parse(args); err != nil {
		return err
	}

	return nil
}

func (cs *cmdServices) Run() error {
	return nil
}

// Command AFAIRE
func Command() command.Command {
	return &cmdServices{}
}

func main() { _ = Command() } // avoid linter errors

/*
######################################################################################################## @(°_°)@ #######
*/
