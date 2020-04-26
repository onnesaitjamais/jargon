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

import "fmt"

// black:   30
// red:     31
// green:   32
// yellow:  33
// blue:    34
// magenta: 35
// cyan:    36
// white:   37

func printLine(level, line string) {
	color := 0

	switch level {
	case "{TRA}":
		color = 34
	case "{DEB}":
		color = 36
	case "{NOT}":
		color = 32
	case "{WAR}":
		color = 33
	case "{ERR}":
		color = 31
	case "{CRI}":
		color = 35
	}

	if color == 0 {
		fmt.Print(line)
	} else {
		fmt.Printf("\x1b[%dm%s\x1b[0m", color, line)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
