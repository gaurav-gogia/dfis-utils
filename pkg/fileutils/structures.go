package fileutils

import "time"

type filenode struct {
	fpath string
	fname string
	ftime time.Time
	fsize int64
}
