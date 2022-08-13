package filecmds

import (
	"dfis-utils/pkg/fileutils"
	"errors"

	"github.com/ibraimgm/libcmd"
)

func FileRaw(cmd *libcmd.Cmd) error {
	numBytes := cmd.GetInt("bytes")
	file := cmd.Operand("FILE")

	if *numBytes <= 0 {
		return errors.New("you must specify a valid number of bytes to read")
	}

	if file == "" {
		return errors.New("you need to specify the FILE")
	}

	return fileutils.Raw(file, *numBytes)
}

func FileInfo(cmd *libcmd.Cmd) error {
	file := cmd.Operand("FILE")

	if file == "" {
		return errors.New("you need to specify the FILE")
	}

	return fileutils.Info(file)
}

func MagicCopy(cmd *libcmd.Cmd) error {
	src := cmd.Operand("SOURCE")
	dst := cmd.Operand("DEST")

	if src == "" {
		return errors.New("you need to specify the SOURCE")
	}

	if dst == "" {
		return errors.New("you need to specify the DEST")
	}

	ftype, err := fileutils.ParseFileType(*cmd.GetChoice("type"))
	if err != nil {
		return err
	}

	return fileutils.Extract(src, dst, ftype)
}

func Large(cmd *libcmd.Cmd) error {
	max := cmd.GetInt("max")
	path := cmd.Operand("PATH")

	if *max <= 0 {
		return errors.New("you must specify a positive number the '--max' parameter")
	}

	if path == "" {
		return errors.New("you need to specify the PATH")
	}

	return fileutils.Large(path, *max)
}

func Recent(cmd *libcmd.Cmd) error {
	max := cmd.GetInt("max")
	path := cmd.Operand("PATH")

	if *max <= 0 {
		return errors.New("you must specify a positive number the '--max' parameter")
	}

	if path == "" {
		return errors.New("you need to specify the PATH")
	}

	return fileutils.Recent(path, *max)
}

func Shred(cmd *libcmd.Cmd) error {
	file := cmd.Operand("FILE")

	if file == "" {
		return errors.New("you need to specify the FILE")
	}

	return fileutils.Del(file)
}

func FileCompare(cmd *libcmd.Cmd) error {
	src := cmd.Operand("SOURCE")
	dst := cmd.Operand("DEST")
	bs := cmd.GetInt64("buffersize")

	if src == "" {
		return errors.New("you need to specify the SOURCE")
	}

	if dst == "" {
		return errors.New("you need to specify the DEST")
	}

	if *bs <= int64(0) {
		return errors.New("buffer must be a positive number")
	}

	return fileutils.Comp(src, dst, *bs)
}
