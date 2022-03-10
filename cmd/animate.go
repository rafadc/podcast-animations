package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"podcast-animations/internal"
)

var (
	videoConfiguration = internal.VideoConfiguration{}

	rootCmd = &cobra.Command{
		Use:   "animate",
		Short: "Different animations to embed in my podcasts",
	}
)

func Execute() error {
	rootCmd.PersistentFlags().Uint8Var(&videoConfiguration.Fps, "fps", 25, "frames per second of the animation")
	viper.SetDefault("fps", 25)

	rootCmd.PersistentFlags().Uint16Var(&videoConfiguration.Width, "width", 1920, "width of the output video")
	viper.SetDefault("width", 1920)

	rootCmd.PersistentFlags().Uint16Var(&videoConfiguration.Height, "height", 1080, "height of the output video")
	viper.SetDefault("width", 1080)

	rootCmd.PersistentFlags().Uint8Var(&videoConfiguration.DurationInSeconds, "duration_in_seconds", 10, "duration of the video")
	viper.SetDefault("width", 10)

	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(coverAndAuthorCmd)
}
