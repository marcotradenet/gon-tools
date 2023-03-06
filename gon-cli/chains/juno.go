package chains

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
)

type junoChain struct {
	ChainData
}

var JunoChain = junoChain{
	ChainData{
		name:         "Juno GoN Testnet",
		chainID:      "uni-6",
		bech32Prefix: "juno",
		denom:        "ujunox",
		rpc:          "https://rpc.uni.junonetwork.io:443",
		grpc:         "http://juno-testnet-grpc.polkachu.com:12690",
	},
}

func (c junoChain) CreateIssueCreditClassMsg(denomID, denomName, schema, sender, symbol string, mintRestricted, updateRestricted bool, description, uri, uriHash, data string) sdk.Msg {
	panic("implement me")
}

func (c junoChain) CreateTransferNFTMsg(connection NFTChannel, nft NFT, fromAddress string, toAddress string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64) sdk.Msg {
	panic("implement me")
}

func (c junoChain) CreateMintNFTMsg(tokenID, classID, tokenName, tokenURI, tokenURIHash, tokenData, minterAddress string) sdk.Msg {
	panic("implement me")
}

func (c junoChain) ListNFTClassesThatHasNFTs(ctx context.Context, clientCtx client.Context, query ListNFTsQuery) []NFTClass {
	panic("implement me")
}
