/* -----------------------------------------------------------------
 *					L o r d  O f   S c r i p t s (tm)
 *				  Copyright (C)2025 Dídimo Grimaldo T.
 *							 go-carousel
 * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
 * Linux Desktop Wallpaper Carousel.
 *-----------------------------------------------------------------*/
package carousel

import (
	"fmt"
	//_ "embed"
	"github.com/lordofscripts/goapp/app"
)

/* ----------------------------------------------------------------
 *							G l o b a l s
 *-----------------------------------------------------------------*/

const (
	NAME           string = "Go Carousel"
	_DESC          string = "A desktop wallpaper carousel with scheduling options"
	MANUAL_VERSION string = "1.1.0"
)

const (
	// Useful Unicode Characters
	CHR_COPYRIGHT       = '\u00a9'      // ©
	CHR_REGISTERD       = '\u00ae'      // ®
	CHR_GUILLEMET_L     = '\u00ab'      // «
	CHR_GUILLEMET_R     = '\u00bb'      // »
	CHR_TRADEMARK       = '\u2122'      // ™
	CHR_SAMARITAN       = '\u214f'      // ⅏
	CHR_PLACEOFINTEREST = '\u2318'      // ⌘
	CHR_HIGHVOLTAGE     = '\u26a1'      // ⚡
	CHR_TRIDENT         = rune(0x1f531) // 🔱
	CHR_SPLATTER        = rune(0x1fadf)
	CHR_WARNING         = '\u26a0' // ⚠
	CHR_EXCLAMATION     = '\u2757'
	CHR_SKULL           = '\u2620'      // ☠
	CHR_HONEYBEE        = rune(0x1f41d) // 🐝

	CO1 = "odlamirG omidiD 6202-5202)C("
	CO2 = "stpircS fO droL 6202-5202)C("
	CO3 = "gnitirwnitsol"
)

var (
	vcsVersion  string // automatically injected with linker
	vcsCommit   string
	vcsDate     string
	vcsBuildNum string
	//NOT USEDgo:embed version.txt
)

var (
	// NOTE: Change these values accordingly
	ModuleVersion app.PackageVersion = app.NewReleaseCandidateVersion(NAME, _DESC, MANUAL_VERSION, 1)
)

/* ----------------------------------------------------------------
 *							T y p e s
 *-----------------------------------------------------------------*/
type status = string

type version struct {
	n  string // name
	v  string // version tag
	s  status // status
	sv int    // Alpha/Beta/RC-### sequence
}

/* ----------------------------------------------------------------
 *							M e t h o d s
 *-----------------------------------------------------------------*/

func BuildMeta() string {
	ver := vcsVersion
	if len(vcsVersion) == 0 {
		ver = "v" + MANUAL_VERSION
	}
	return fmt.Sprintf("\t\t%s-%s %s", ver, vcsBuildNum, vcsDate)
}

/* ----------------------------------------------------------------
 *							F u n c t i o n s
 *-----------------------------------------------------------------*/

// Funny LordOfScripts logo
func Logo() string {
	const (
		whiteStar rune = '\u269d' // ⚝
		unisex    rune = '\u26a5' // ⚥
		hotSpring rune = '\u2668' // ♨
		leftConv  rune = '\u269e' // ⚞
		rightConv rune = '\u269f' // ⚟
		eye       rune = '\u25d5' // ◕
		mouth     rune = '\u035c' // ͜	‿ \u203f
		skull     rune = '\u2620' // ☠
	)
	return fmt.Sprintf("%c%c%c %c%c", leftConv, eye, mouth, eye, rightConv)
	//fmt.Sprintf("(%c%c %c)", eye, mouth, eye)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

/*
  This commented part is used by the Makefile target that extracts it
  to a temporary file, compiles it and uses it to print out the full
  Module Version that includes attributes like Alpha,Beta,RC which the
  plain old "make version" didn't

//>>>BEGIN Versioner
package main

import (
    "os"
    "fmt"
    "strings"
    "github.com/lordofscripts/carousel"
)

func main() {
    if len(os.Args) == 2 && strings.EqualFold(os.Args[1], "short") {
        fmt.Println(carousel.ModuleVersion.Short())
    } else {
        fmt.Println(carousel.ModuleVersion)
    }
}
//>>>END Versioner
*/
