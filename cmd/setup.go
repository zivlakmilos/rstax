/*
Copyright © 2025 Milos Zivlak

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zivlakmilos/rstax/models"
	"github.com/zivlakmilos/rstax/tui"
	"github.com/zivlakmilos/rstax/utils"
)

var setupParams models.SetupParams

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup config data",
	Run: func(cmd *cobra.Command, args []string) {
		utils.InitParams(&setupParams)
		tui.RunSetup(setupParams)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	setupCmd.Flags().StringVarP(&setupParams.Name, "name", "n", "", "Full Name")
	setupCmd.Flags().StringVarP(&setupParams.Jmbg, "jmbg", "j", "", "JMBG")
	setupCmd.Flags().StringVarP(&setupParams.Phone, "phone", "p", "", "Phone")
	setupCmd.Flags().StringVarP(&setupParams.Email, "email", "e", "", "Email Address")
	setupCmd.Flags().StringVarP(&setupParams.Address, "address", "a", "", "Address")
	setupCmd.Flags().StringVarP(&setupParams.Place, "place", "c", "", "Place")
	setupCmd.Flags().StringVarP(&setupParams.BankAccount, "bank-account", "b", "", "Bank Account Number")
}
