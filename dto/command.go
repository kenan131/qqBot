package dto

import "regexp"

var CommandRegular = regexp.MustCompile(`/(\d{3})\s*(.*)`)

var AtRE = regexp.MustCompile(`<@!\d+>`)

const SpaceCharSet = " \u00A0"
