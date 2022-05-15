/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yeldisbayev/crypto-cli/internal/delivery"
	"github.com/yeldisbayev/crypto-cli/internal/repository/storage"
	"github.com/yeldisbayev/crypto-cli/internal/usecase/interactor"
	"github.com/yeldisbayev/crypto-cli/pkg/crypto"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "File decryption",
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()
		secret := cmd.Flag("secret").Value.String()
		usecase := interactor.NewCryptoUsecase(
			extension,
			crypto.NewCrypto([]byte(secret)),
			storage.NewReadWriter(),
		)

		handler := delivery.NewHandler(usecase)
		err := handler.Decrypt(path)

		if err != nil {
			panic(err)
		}

		fmt.Println("Completed!")

	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	decryptCmd.PersistentFlags().String("path", "", "Path to file")
	decryptCmd.PersistentFlags().String("secret", "", "Secret for encryption")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
