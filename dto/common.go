package dto

import "regexp"

var CommandRegular = regexp.MustCompile(`/(\d{3})\s*(.*)`)

var AtRE = regexp.MustCompile(`<@!\d+>`)

const SpaceCharSet = " \u00A0"
const ConfigStr = "config.yaml"
const GuessNumDefaultRange = 10000
const DefaultQueueSize = 10000
const (
	GameStart int = 1
	GameEnd   int = -1
)
