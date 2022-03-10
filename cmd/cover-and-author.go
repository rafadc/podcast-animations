package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"podcast-animations/pkg"
)

var (
	coverAndAuthorCmd = &cobra.Command{
		Use:   "cover-and-author <background.jpg> <cover.jpg> <author.jpg> ",
		Short: "Animation for displaying the cover of a book with its author",
		Long:  `Animation for displaying the cover of a book with its author. It resizes the background and the images properly`,
		Args:  cobra.MinimumNArgs(3),
		Run:   coverAndAuthorRun,
	}
)

func coverAndAuthorRun(cmd *cobra.Command, args []string) {
	log.Info("Creating cover and author animation")

	backgroundFileName := args[0]
	coverFileName := args[1]
	authorFileName := args[2]

	pkg.CoverAndAuthor(videoConfiguration, backgroundFileName, coverFileName, authorFileName)
}
