package embed

import _ "embed"

//go:embed note.txt
var Note []byte

//go:embed hello-world.txt
var HW string
