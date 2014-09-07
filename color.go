package main

type color string

// colors as escape characters
const (
	none      color = ""
	black           = "\x1b[30m"
	blue            = "\x1b[34m"
	green           = "\x1b[32m"
	red             = "\x1b[31m"
	yellow          = "\x1b[33m"
	resetCode       = "\x1b[39;49m"
)
