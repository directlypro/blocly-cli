package cmd

import (
	block "blocky/pkg"
	"github.com/spf13/cobra"
	"github.com/txn2/txeh"
)

// unblockCmd represents the unblock command
var unblockCmd = &cobra.Command{
	Use:   "unblock",
	Short: "Unblocking a blocked host or site",
	Long:  `Unblocking a blocked host or site.`,
	Run: func(cmd *cobra.Command, args []string) {
		hosts, err := txeh.NewHostsDefault()
		if err != nil {
			panic(err)
		}
		sites, _ := cmd.Flags().GetStringSlice("sites")
		block.CleanBlock(hosts, sites)
	},
}

func init() {
	rootCmd.AddCommand(unblockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	unblockCmd.PersistentFlags().StringSliceP("sites", "s", []string{}, "Unblock a blocked host or site")
	unblockCmd.MarkFlagRequired("sites")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unblockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
