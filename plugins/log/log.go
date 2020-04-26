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
	"errors"

	"github.com/arnumina/jargon/pkg/command"
)

type cmdLog struct {
	fileName string
}

func (cl *cmdLog) Name() string {
	return "log"
}

func (cl *cmdLog) Description() string {
	return "print the log file in real time"
}

func (cl *cmdLog) Setup(args []string) error {
	fs := command.NewFlagSet(cl)

	fs.StringVar(&cl.fileName, "file", "", "the log file to be printed")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if cl.fileName == "" {
		return errors.New("no log file has been specified") ////////////////////////////////////////////////////////////
	}

	return nil
}

func (cl *cmdLog) Run() error {
	return tailLogFile(cl)
}

// Command AFAIRE
func Command() command.Command {
	return &cmdLog{}
}

func main() { _ = Command() } // avoid linter errors

/*
######################################################################################################## @(°_°)@ #######
*/
