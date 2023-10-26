package builtInFunctions

import (
	"math/big"
	"testing"

	"github.com/Dharitri-org/sme-core/core"
	vmcommon "github.com/Dharitri-org/sme-vm-common"
	"github.com/stretchr/testify/require"
)

func TestNewEntryForNFT(t *testing.T) {
	t.Parallel()

	vmOutput := &vmcommon.VMOutput{}
	addDCTEntryInVMOutput(vmOutput, []byte(core.BuiltInFunctionDCTNFTCreate), []byte("my-token"), 5, big.NewInt(1), []byte("caller"), []byte("receiver"))
	require.Equal(t, &vmcommon.LogEntry{
		Identifier: []byte(core.BuiltInFunctionDCTNFTCreate),
		Address:    []byte("caller"),
		Topics:     [][]byte{[]byte("my-token"), big.NewInt(0).SetUint64(5).Bytes(), big.NewInt(1).Bytes(), []byte("receiver")},
		Data:       nil,
	}, vmOutput.Logs[0])
}
