//go:build ignore

/* -----------------------------------------------------------------
 *					L o r d  O f   S c r i p t s (tm)
 *				  Copyright (C)2025 Dídimo Grimaldo T.
 *							   go-carousel
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * Application-related functions.
 * Moved to github.com/lordofscripts/goapp
 *-----------------------------------------------------------------*/
package app

import (
	"errors"
	"os"
)

/* ----------------------------------------------------------------
 *					F u n c t i o n s
 *-----------------------------------------------------------------*/

func DirExists(filename string) bool {
	fi, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist) && fi.IsDir()
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}
