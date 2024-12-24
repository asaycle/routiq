package repository

import "golang.org/x/xerrors"

var (
	NotfoundError = xerrors.New("Notfound Error")
)
