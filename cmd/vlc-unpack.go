package cmd

import (
	"archiver/lib/vlc"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var chihardaUnpackCmd = &cobra.Command{
	Use:   "chiharda",
	Short: "Unpack file using variable-length code",
	Run:   unpack,
}

const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := vlc.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

func unpackedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension
}

func init() {
	unpackCmd.AddCommand(chihardaUnpackCmd)
}
