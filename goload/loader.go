package goload

import (
	"go/build"

	"golang.org/x/tools/go/loader"
)

// Program is a loaded program from a set of packages.
type Program struct {
	*loader.Program

	Pkgs []string
}

// Pkgs loads a list of package from source files into a loader program.
func Pkgs(pkgs []string) (*Program, error) {
	b := build.Default
	b.CgoEnabled = false

	conf := loader.Config{Build: &b}
	for _, p := range pkgs {
		conf.Import(p)
	}

	ret, e := conf.Load()
	if e != nil {
		return nil, e
	}

	pkgsCopied := make([]string, len(pkgs))
	copy(pkgsCopied, pkgs)

	return &Program{Program: ret, Pkgs: pkgsCopied}, e
}

// Pkg loads a set of pacakges from source files into a loader program.
// It just calls LoadPkgs with the parameter list
func Pkg(pkgs ...string) (*Program, error) {
	return Pkgs(pkgs)
}
