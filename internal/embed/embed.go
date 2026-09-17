package embed

import "embed"

//go:embed note.txt
var Note []byte

//go:embed hello-world.txt
var HW string

//go:embed 00001-auth.sql
var UserScript embed.FS
