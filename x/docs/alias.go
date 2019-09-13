package docs

import (
	"github.com/commercionetwork/commercionetwork/x/docs/internal/keeper"
	"github.com/commercionetwork/commercionetwork/x/docs/internal/types"
)

const (
	ModuleName   = types.ModuleName
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	TestSetup     = keeper.SetupTestInput
	RegisterCodec = types.RegisterCodec
	ModuleCdc     = types.ModuleCdc
)

type (
	Keeper                              = keeper.Keeper
	Document                            = types.Document
	DocumentReceipt                     = types.DocumentReceipt
	MsgShareDocument                    = types.MsgShareDocument
	MsgSendDocumentReceipt              = types.MsgSendDocumentReceipt
	MsgAddSupportedMetadataSchema       = types.MsgAddSupportedMetadataSchema
	MsgAddTrustedMetadataSchemaProposer = types.MsgAddTrustedMetadataSchemaProposer
)
