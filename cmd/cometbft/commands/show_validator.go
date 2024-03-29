package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	cmtjson "github.com/tendermint/tendermint/libs/json"
	cmtos "github.com/tendermint/tendermint/libs/os"
	"github.com/tendermint/tendermint/privval"
)

// ShowValidatorCmd adds capabilities for showing the validator info.
var ShowValidatorCmd = &cobra.Command{
	Use:     "show-validator",
	Aliases: []string{"show_validator"},
	Short:   "Show this node's validator info",
	RunE:    showValidator,
}

func showValidator(cmd *cobra.Command, args []string) error {
	keyFilePath := config.PrivValidatorKeyFile()
	if !cmtos.FileExists(keyFilePath) {
		return fmt.Errorf("private validator file %s does not exist", keyFilePath)
	}

	pv := privval.LoadFilePV(keyFilePath, config.PrivValidatorStateFile())

	pubKey, err := pv.GetPubKey()
	if err != nil {
		return fmt.Errorf("can't get pubkey: %w", err)
	}

	bz, err := cmtjson.Marshal(pubKey)
	if err != nil {
		return fmt.Errorf("failed to marshal private validator pubkey: %w", err)
	}

	fmt.Println(string(bz))
	return nil
}
