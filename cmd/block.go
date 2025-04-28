package cmd

import (
	"fmt"
	"github.com/txn2/txeh"

	"blocky/pkg"
	"github.com/spf13/cobra"
)

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "The block subcommand will be creating the blocks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Block certain URLs in /etc/hosts")
		fmt.Println("Or where the users configured!")
		hostsFile, _ := cmd.Flags().GetString("host-file")
		sites, _ := cmd.Flags().GetStringSlice("sites")
		fmt.Println(hostsFile)
		fmt.Println(sites)
		hosts, err := txeh.NewHostsDefault()
		if err != nil {
			panic(err)
		}
		block.BlockSites(hosts, sites)
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	blockCmd.PersistentFlags().StringP("host-file", "f", "/etc/hosts", "a custom hostfile path other than /etc/hosts")
	blockCmd.PersistentFlags().StringSliceP("sites", "s", []string{}, "sites to block")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//blockCmd.Flags().BoolP("host-file", "", false, "a custom hostfile path other than /etc/hosts")
}
