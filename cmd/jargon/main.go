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
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/arnumina/jargon/pkg/command"
)

const _exportedFunction = "Command"

var (
	_version string
	_builtAt string
)

func findPlugins() (map[string]string, error) {
	exe, err := os.Executable()
	if err != nil {
		return nil, err
	}

	app := filepath.Base(os.Args[0])

	files, err := filepath.Glob(filepath.Join(filepath.Dir(exe), app+".*.so"))
	if err != nil {
		return nil, err
	}

	plugins := make(map[string]string)

	for _, file := range files {
		plugins[strings.TrimSuffix(strings.TrimPrefix(filepath.Base(file), app+"."), ".so")] = file
	}

	return plugins, nil
}

func cmdHelp() error {
	plugins, err := findPlugins()
	if err != nil {
		return err
	}

	app := filepath.Base(os.Args[0])

	fmt.Println()
	fmt.Println("The command line client")
	fmt.Println("================================================================================")
	fmt.Println("Usage:")
	fmt.Printf("  %s [command [options]]\n", app)
	fmt.Println()
	fmt.Println("Available commands:")

	commands := make([]string, len(plugins)+1)

	i := 0

	for cmd := range plugins {
		commands[i] = cmd
		i++
	}

	commands[i] = "version"

	sort.Strings(commands)

	for _, cmd := range commands {
		fmt.Println("  " + cmd)
	}

	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("  Use '%s [command] --help' for more information about a command.\n", app)
	fmt.Println("================================================================================")
	fmt.Println()

	return nil
}

func cmdVersion() error {
	builtAt, err := strconv.ParseInt(_builtAt, 10, 64)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("  jargon")
	fmt.Println("-----------------------------------------------")
	fmt.Println("  version  :", _version)
	fmt.Println("  built at :", time.Unix(builtAt, 0).String())
	fmt.Println("  by       : Archivage Numérique © INA", time.Now().Year())
	fmt.Println("-----------------------------------------------")
	fmt.Println()

	return nil
}

func runCommand(file string) error {
	plugin, err := plugin.Open(file)
	if err != nil {
		return err
	}

	ef, err := plugin.Lookup(_exportedFunction)
	if err != nil {
		return err
	}

	fn, ok := ef.(func() command.Command)
	if !ok {
		return fmt.Errorf( /////////////////////////////////////////////////////////////////////////////////////////////
			"the function 'exported' by this plugin is not valid: plugin=%s, function=%s",
			file,
			_exportedFunction,
		)
	}

	cmd := fn()

	if err := cmd.Setup(os.Args[2:]); err != nil {
		return err
	}

	return cmd.Run()
}

func run() error {
	if len(os.Args) == 1 {
		return cmdHelp()
	}

	switch os.Args[1] {
	case "--help", "-help", "help":
		return cmdHelp()
	case "--version", "-version", "version":
		return cmdVersion()
	}

	plugins, err := findPlugins()
	if err != nil {
		return err
	}

	for name, file := range plugins {
		if name == os.Args[1] {
			return runCommand(file)
		}
	}

	return errors.New("this command does not exist") ///////////////////////////////////////////////////////////////////
}

func main() {
	if err := run(); err != nil {
		if errors.Is(err, command.ErrNoError) {
			return
		}

		fmt.Fprintf( ///////////////////////////////////////////////////////////////////////////////////////////////////
			os.Stderr,
			"Error: cmd=%s >>> %s\n",
			os.Args[1],
			err,
		)

		os.Exit(1)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
