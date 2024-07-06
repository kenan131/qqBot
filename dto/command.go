package dto

import "regexp"

var CommandRegular = regexp.MustCompile(`/(\d{3})\s*(.*)`)
