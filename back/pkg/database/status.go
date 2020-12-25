package database

type QueryStatus int

const (
	SUCCESSFUL QueryStatus = iota
	FAILED
)
