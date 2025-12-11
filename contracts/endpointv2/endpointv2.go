// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package endpointv2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MessagingFee is an auto generated low-level Go binding around an user-defined struct.
type MessagingFee struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}

// MessagingParams is an auto generated low-level Go binding around an user-defined struct.
type MessagingParams struct {
	DstEid       uint32
	Receiver     [32]byte
	Message      []byte
	Options      []byte
	PayInLzToken bool
}

// MessagingReceipt is an auto generated low-level Go binding around an user-defined struct.
type MessagingReceipt struct {
	Guid  [32]byte
	Nonce uint64
	Fee   MessagingFee
}

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// SetConfigParam is an auto generated low-level Go binding around an user-defined struct.
type SetConfigParam struct {
	Eid        uint32
	ConfigType uint32
	Config     []byte
}

// EndpointV2MetaData contains all meta data concerning the EndpointV2 contract.
var EndpointV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LZ_AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_ComposeExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"LZ_ComposeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_DefaultReceiveLibUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_DefaultSendLibUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredNative\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"suppliedNative\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredLzToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"suppliedLzToken\",\"type\":\"uint256\"}],\"name\":\"LZ_InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_InvalidExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"LZ_InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_InvalidPayloadHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_InvalidReceiveLibrary\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlyNonDefaultLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlyReceiveLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlyRegisteredLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlyRegisteredOrDefaultLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlySendLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_PathNotInitializable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_PathNotVerifiable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"LZ_PayloadHashNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_SameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_SendReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_UnsupportedEid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_UnsupportedInterface\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_ZeroLzTokenFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer_NativeFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Transfer_ToAddressIsZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"}],\"name\":\"ComposeDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"ComposeSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"DefaultReceiveLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldLib\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"DefaultReceiveLibraryTimeoutSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"DefaultSendLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"DelegateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"InboundNonceSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"LibraryRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"LzComposeAlert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"LzReceiveAlert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"LzTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"name\":\"PacketBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"PacketDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"name\":\"PacketNilified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendLibrary\",\"type\":\"address\"}],\"name\":\"PacketSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"name\":\"PacketVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"ReceiveLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldLib\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"ReceiveLibraryTimeoutSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"SendLibrarySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EMPTY_PAYLOAD_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NIL_PAYLOAD_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockedLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"clear\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"}],\"name\":\"composeQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"}],\"name\":\"defaultReceiveLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"}],\"name\":\"defaultReceiveLibraryTimeout\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"}],\"name\":\"defaultSendLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oapp\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eid\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_configType\",\"type\":\"uint32\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"}],\"name\":\"getReceiveLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredLibraries\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSendContext\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"}],\"name\":\"getSendLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"}],\"name\":\"inboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"inboundNonce\",\"type\":\"uint64\"}],\"name\":\"inboundPayloadHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"initializable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"}],\"name\":\"isDefaultSendLibrary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"name\":\"isRegisteredLibrary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSendingMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"}],\"name\":\"isSupportedEid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_actualReceiveLib\",\"type\":\"address\"}],\"name\":\"isValidReceiveLibrary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"lazyInboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"_index\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzCompose\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"_index\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"lzComposeAlert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"lzReceiveAlert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lzToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_receiver\",\"type\":\"bytes32\"}],\"name\":\"nextGuid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"nilify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"}],\"name\":\"outboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"payInLzToken\",\"type\":\"bool\"}],\"internalType\":\"structMessagingParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"quote\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"}],\"name\":\"receiveLibraryTimeout\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"}],\"name\":\"registerLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"payInLzToken\",\"type\":\"bool\"}],\"internalType\":\"structMessagingParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"}],\"internalType\":\"structMessagingReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"_index\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendCompose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"configType\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"internalType\":\"structSetConfigParam[]\",\"name\":\"_params\",\"type\":\"tuple[]\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newLib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gracePeriod\",\"type\":\"uint256\"}],\"name\":\"setDefaultReceiveLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expiry\",\"type\":\"uint256\"}],\"name\":\"setDefaultReceiveLibraryTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newLib\",\"type\":\"address\"}],\"name\":\"setDefaultSendLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lzToken\",\"type\":\"address\"}],\"name\":\"setLzToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newLib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gracePeriod\",\"type\":\"uint256\"}],\"name\":\"setReceiveLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expiry\",\"type\":\"uint256\"}],\"name\":\"setReceiveLibraryTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newLib\",\"type\":\"address\"}],\"name\":\"setSendLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"name\":\"skip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"verifiable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EndpointV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use EndpointV2MetaData.ABI instead.
var EndpointV2ABI = EndpointV2MetaData.ABI

// EndpointV2 is an auto generated Go binding around an Ethereum contract.
type EndpointV2 struct {
	EndpointV2Caller     // Read-only binding to the contract
	EndpointV2Transactor // Write-only binding to the contract
	EndpointV2Filterer   // Log filterer for contract events
}

// EndpointV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type EndpointV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndpointV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type EndpointV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndpointV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EndpointV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndpointV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EndpointV2Session struct {
	Contract     *EndpointV2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EndpointV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EndpointV2CallerSession struct {
	Contract *EndpointV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EndpointV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EndpointV2TransactorSession struct {
	Contract     *EndpointV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EndpointV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type EndpointV2Raw struct {
	Contract *EndpointV2 // Generic contract binding to access the raw methods on
}

// EndpointV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EndpointV2CallerRaw struct {
	Contract *EndpointV2Caller // Generic read-only contract binding to access the raw methods on
}

// EndpointV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EndpointV2TransactorRaw struct {
	Contract *EndpointV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEndpointV2 creates a new instance of EndpointV2, bound to a specific deployed contract.
func NewEndpointV2(address common.Address, backend bind.ContractBackend) (*EndpointV2, error) {
	contract, err := bindEndpointV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EndpointV2{EndpointV2Caller: EndpointV2Caller{contract: contract}, EndpointV2Transactor: EndpointV2Transactor{contract: contract}, EndpointV2Filterer: EndpointV2Filterer{contract: contract}}, nil
}

// NewEndpointV2Caller creates a new read-only instance of EndpointV2, bound to a specific deployed contract.
func NewEndpointV2Caller(address common.Address, caller bind.ContractCaller) (*EndpointV2Caller, error) {
	contract, err := bindEndpointV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EndpointV2Caller{contract: contract}, nil
}

// NewEndpointV2Transactor creates a new write-only instance of EndpointV2, bound to a specific deployed contract.
func NewEndpointV2Transactor(address common.Address, transactor bind.ContractTransactor) (*EndpointV2Transactor, error) {
	contract, err := bindEndpointV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EndpointV2Transactor{contract: contract}, nil
}

// NewEndpointV2Filterer creates a new log filterer instance of EndpointV2, bound to a specific deployed contract.
func NewEndpointV2Filterer(address common.Address, filterer bind.ContractFilterer) (*EndpointV2Filterer, error) {
	contract, err := bindEndpointV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EndpointV2Filterer{contract: contract}, nil
}

// bindEndpointV2 binds a generic wrapper to an already deployed contract.
func bindEndpointV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EndpointV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EndpointV2 *EndpointV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EndpointV2.Contract.EndpointV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EndpointV2 *EndpointV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EndpointV2.Contract.EndpointV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EndpointV2 *EndpointV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EndpointV2.Contract.EndpointV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EndpointV2 *EndpointV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EndpointV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EndpointV2 *EndpointV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EndpointV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EndpointV2 *EndpointV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EndpointV2.Contract.contract.Transact(opts, method, params...)
}

// EMPTYPAYLOADHASH is a free data retrieval call binding the contract method 0xcb5026b9.
//
// Solidity: function EMPTY_PAYLOAD_HASH() view returns(bytes32)
func (_EndpointV2 *EndpointV2Caller) EMPTYPAYLOADHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "EMPTY_PAYLOAD_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EMPTYPAYLOADHASH is a free data retrieval call binding the contract method 0xcb5026b9.
//
// Solidity: function EMPTY_PAYLOAD_HASH() view returns(bytes32)
func (_EndpointV2 *EndpointV2Session) EMPTYPAYLOADHASH() ([32]byte, error) {
	return _EndpointV2.Contract.EMPTYPAYLOADHASH(&_EndpointV2.CallOpts)
}

// EMPTYPAYLOADHASH is a free data retrieval call binding the contract method 0xcb5026b9.
//
// Solidity: function EMPTY_PAYLOAD_HASH() view returns(bytes32)
func (_EndpointV2 *EndpointV2CallerSession) EMPTYPAYLOADHASH() ([32]byte, error) {
	return _EndpointV2.Contract.EMPTYPAYLOADHASH(&_EndpointV2.CallOpts)
}

// NILPAYLOADHASH is a free data retrieval call binding the contract method 0x2baf0be7.
//
// Solidity: function NIL_PAYLOAD_HASH() view returns(bytes32)
func (_EndpointV2 *EndpointV2Caller) NILPAYLOADHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "NIL_PAYLOAD_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NILPAYLOADHASH is a free data retrieval call binding the contract method 0x2baf0be7.
//
// Solidity: function NIL_PAYLOAD_HASH() view returns(bytes32)
func (_EndpointV2 *EndpointV2Session) NILPAYLOADHASH() ([32]byte, error) {
	return _EndpointV2.Contract.NILPAYLOADHASH(&_EndpointV2.CallOpts)
}

// NILPAYLOADHASH is a free data retrieval call binding the contract method 0x2baf0be7.
//
// Solidity: function NIL_PAYLOAD_HASH() view returns(bytes32)
func (_EndpointV2 *EndpointV2CallerSession) NILPAYLOADHASH() ([32]byte, error) {
	return _EndpointV2.Contract.NILPAYLOADHASH(&_EndpointV2.CallOpts)
}

// BlockedLibrary is a free data retrieval call binding the contract method 0x73318091.
//
// Solidity: function blockedLibrary() view returns(address)
func (_EndpointV2 *EndpointV2Caller) BlockedLibrary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "blockedLibrary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockedLibrary is a free data retrieval call binding the contract method 0x73318091.
//
// Solidity: function blockedLibrary() view returns(address)
func (_EndpointV2 *EndpointV2Session) BlockedLibrary() (common.Address, error) {
	return _EndpointV2.Contract.BlockedLibrary(&_EndpointV2.CallOpts)
}

// BlockedLibrary is a free data retrieval call binding the contract method 0x73318091.
//
// Solidity: function blockedLibrary() view returns(address)
func (_EndpointV2 *EndpointV2CallerSession) BlockedLibrary() (common.Address, error) {
	return _EndpointV2.Contract.BlockedLibrary(&_EndpointV2.CallOpts)
}

// ComposeQueue is a free data retrieval call binding the contract method 0x35d330b0.
//
// Solidity: function composeQueue(address from, address to, bytes32 guid, uint16 index) view returns(bytes32 messageHash)
func (_EndpointV2 *EndpointV2Caller) ComposeQueue(opts *bind.CallOpts, from common.Address, to common.Address, guid [32]byte, index uint16) ([32]byte, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "composeQueue", from, to, guid, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComposeQueue is a free data retrieval call binding the contract method 0x35d330b0.
//
// Solidity: function composeQueue(address from, address to, bytes32 guid, uint16 index) view returns(bytes32 messageHash)
func (_EndpointV2 *EndpointV2Session) ComposeQueue(from common.Address, to common.Address, guid [32]byte, index uint16) ([32]byte, error) {
	return _EndpointV2.Contract.ComposeQueue(&_EndpointV2.CallOpts, from, to, guid, index)
}

// ComposeQueue is a free data retrieval call binding the contract method 0x35d330b0.
//
// Solidity: function composeQueue(address from, address to, bytes32 guid, uint16 index) view returns(bytes32 messageHash)
func (_EndpointV2 *EndpointV2CallerSession) ComposeQueue(from common.Address, to common.Address, guid [32]byte, index uint16) ([32]byte, error) {
	return _EndpointV2.Contract.ComposeQueue(&_EndpointV2.CallOpts, from, to, guid, index)
}

// DefaultReceiveLibrary is a free data retrieval call binding the contract method 0x6f50a803.
//
// Solidity: function defaultReceiveLibrary(uint32 srcEid) view returns(address lib)
func (_EndpointV2 *EndpointV2Caller) DefaultReceiveLibrary(opts *bind.CallOpts, srcEid uint32) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "defaultReceiveLibrary", srcEid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultReceiveLibrary is a free data retrieval call binding the contract method 0x6f50a803.
//
// Solidity: function defaultReceiveLibrary(uint32 srcEid) view returns(address lib)
func (_EndpointV2 *EndpointV2Session) DefaultReceiveLibrary(srcEid uint32) (common.Address, error) {
	return _EndpointV2.Contract.DefaultReceiveLibrary(&_EndpointV2.CallOpts, srcEid)
}

// DefaultReceiveLibrary is a free data retrieval call binding the contract method 0x6f50a803.
//
// Solidity: function defaultReceiveLibrary(uint32 srcEid) view returns(address lib)
func (_EndpointV2 *EndpointV2CallerSession) DefaultReceiveLibrary(srcEid uint32) (common.Address, error) {
	return _EndpointV2.Contract.DefaultReceiveLibrary(&_EndpointV2.CallOpts, srcEid)
}

// DefaultReceiveLibraryTimeout is a free data retrieval call binding the contract method 0x6e83f5bb.
//
// Solidity: function defaultReceiveLibraryTimeout(uint32 srcEid) view returns(address lib, uint256 expiry)
func (_EndpointV2 *EndpointV2Caller) DefaultReceiveLibraryTimeout(opts *bind.CallOpts, srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "defaultReceiveLibraryTimeout", srcEid)

	outstruct := new(struct {
		Lib    common.Address
		Expiry *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lib = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Expiry = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DefaultReceiveLibraryTimeout is a free data retrieval call binding the contract method 0x6e83f5bb.
//
// Solidity: function defaultReceiveLibraryTimeout(uint32 srcEid) view returns(address lib, uint256 expiry)
func (_EndpointV2 *EndpointV2Session) DefaultReceiveLibraryTimeout(srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	return _EndpointV2.Contract.DefaultReceiveLibraryTimeout(&_EndpointV2.CallOpts, srcEid)
}

// DefaultReceiveLibraryTimeout is a free data retrieval call binding the contract method 0x6e83f5bb.
//
// Solidity: function defaultReceiveLibraryTimeout(uint32 srcEid) view returns(address lib, uint256 expiry)
func (_EndpointV2 *EndpointV2CallerSession) DefaultReceiveLibraryTimeout(srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	return _EndpointV2.Contract.DefaultReceiveLibraryTimeout(&_EndpointV2.CallOpts, srcEid)
}

// DefaultSendLibrary is a free data retrieval call binding the contract method 0xf64be4c7.
//
// Solidity: function defaultSendLibrary(uint32 dstEid) view returns(address lib)
func (_EndpointV2 *EndpointV2Caller) DefaultSendLibrary(opts *bind.CallOpts, dstEid uint32) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "defaultSendLibrary", dstEid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultSendLibrary is a free data retrieval call binding the contract method 0xf64be4c7.
//
// Solidity: function defaultSendLibrary(uint32 dstEid) view returns(address lib)
func (_EndpointV2 *EndpointV2Session) DefaultSendLibrary(dstEid uint32) (common.Address, error) {
	return _EndpointV2.Contract.DefaultSendLibrary(&_EndpointV2.CallOpts, dstEid)
}

// DefaultSendLibrary is a free data retrieval call binding the contract method 0xf64be4c7.
//
// Solidity: function defaultSendLibrary(uint32 dstEid) view returns(address lib)
func (_EndpointV2 *EndpointV2CallerSession) DefaultSendLibrary(dstEid uint32) (common.Address, error) {
	return _EndpointV2.Contract.DefaultSendLibrary(&_EndpointV2.CallOpts, dstEid)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address oapp) view returns(address delegate)
func (_EndpointV2 *EndpointV2Caller) Delegates(opts *bind.CallOpts, oapp common.Address) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "delegates", oapp)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address oapp) view returns(address delegate)
func (_EndpointV2 *EndpointV2Session) Delegates(oapp common.Address) (common.Address, error) {
	return _EndpointV2.Contract.Delegates(&_EndpointV2.CallOpts, oapp)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address oapp) view returns(address delegate)
func (_EndpointV2 *EndpointV2CallerSession) Delegates(oapp common.Address) (common.Address, error) {
	return _EndpointV2.Contract.Delegates(&_EndpointV2.CallOpts, oapp)
}

// Eid is a free data retrieval call binding the contract method 0x416ecebf.
//
// Solidity: function eid() view returns(uint32)
func (_EndpointV2 *EndpointV2Caller) Eid(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "eid")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Eid is a free data retrieval call binding the contract method 0x416ecebf.
//
// Solidity: function eid() view returns(uint32)
func (_EndpointV2 *EndpointV2Session) Eid() (uint32, error) {
	return _EndpointV2.Contract.Eid(&_EndpointV2.CallOpts)
}

// Eid is a free data retrieval call binding the contract method 0x416ecebf.
//
// Solidity: function eid() view returns(uint32)
func (_EndpointV2 *EndpointV2CallerSession) Eid() (uint32, error) {
	return _EndpointV2.Contract.Eid(&_EndpointV2.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0x2b3197b9.
//
// Solidity: function getConfig(address _oapp, address _lib, uint32 _eid, uint32 _configType) view returns(bytes config)
func (_EndpointV2 *EndpointV2Caller) GetConfig(opts *bind.CallOpts, _oapp common.Address, _lib common.Address, _eid uint32, _configType uint32) ([]byte, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "getConfig", _oapp, _lib, _eid, _configType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0x2b3197b9.
//
// Solidity: function getConfig(address _oapp, address _lib, uint32 _eid, uint32 _configType) view returns(bytes config)
func (_EndpointV2 *EndpointV2Session) GetConfig(_oapp common.Address, _lib common.Address, _eid uint32, _configType uint32) ([]byte, error) {
	return _EndpointV2.Contract.GetConfig(&_EndpointV2.CallOpts, _oapp, _lib, _eid, _configType)
}

// GetConfig is a free data retrieval call binding the contract method 0x2b3197b9.
//
// Solidity: function getConfig(address _oapp, address _lib, uint32 _eid, uint32 _configType) view returns(bytes config)
func (_EndpointV2 *EndpointV2CallerSession) GetConfig(_oapp common.Address, _lib common.Address, _eid uint32, _configType uint32) ([]byte, error) {
	return _EndpointV2.Contract.GetConfig(&_EndpointV2.CallOpts, _oapp, _lib, _eid, _configType)
}

// GetReceiveLibrary is a free data retrieval call binding the contract method 0x402f8468.
//
// Solidity: function getReceiveLibrary(address _receiver, uint32 _srcEid) view returns(address lib, bool isDefault)
func (_EndpointV2 *EndpointV2Caller) GetReceiveLibrary(opts *bind.CallOpts, _receiver common.Address, _srcEid uint32) (struct {
	Lib       common.Address
	IsDefault bool
}, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "getReceiveLibrary", _receiver, _srcEid)

	outstruct := new(struct {
		Lib       common.Address
		IsDefault bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lib = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsDefault = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetReceiveLibrary is a free data retrieval call binding the contract method 0x402f8468.
//
// Solidity: function getReceiveLibrary(address _receiver, uint32 _srcEid) view returns(address lib, bool isDefault)
func (_EndpointV2 *EndpointV2Session) GetReceiveLibrary(_receiver common.Address, _srcEid uint32) (struct {
	Lib       common.Address
	IsDefault bool
}, error) {
	return _EndpointV2.Contract.GetReceiveLibrary(&_EndpointV2.CallOpts, _receiver, _srcEid)
}

// GetReceiveLibrary is a free data retrieval call binding the contract method 0x402f8468.
//
// Solidity: function getReceiveLibrary(address _receiver, uint32 _srcEid) view returns(address lib, bool isDefault)
func (_EndpointV2 *EndpointV2CallerSession) GetReceiveLibrary(_receiver common.Address, _srcEid uint32) (struct {
	Lib       common.Address
	IsDefault bool
}, error) {
	return _EndpointV2.Contract.GetReceiveLibrary(&_EndpointV2.CallOpts, _receiver, _srcEid)
}

// GetRegisteredLibraries is a free data retrieval call binding the contract method 0x9132e5c3.
//
// Solidity: function getRegisteredLibraries() view returns(address[])
func (_EndpointV2 *EndpointV2Caller) GetRegisteredLibraries(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "getRegisteredLibraries")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredLibraries is a free data retrieval call binding the contract method 0x9132e5c3.
//
// Solidity: function getRegisteredLibraries() view returns(address[])
func (_EndpointV2 *EndpointV2Session) GetRegisteredLibraries() ([]common.Address, error) {
	return _EndpointV2.Contract.GetRegisteredLibraries(&_EndpointV2.CallOpts)
}

// GetRegisteredLibraries is a free data retrieval call binding the contract method 0x9132e5c3.
//
// Solidity: function getRegisteredLibraries() view returns(address[])
func (_EndpointV2 *EndpointV2CallerSession) GetRegisteredLibraries() ([]common.Address, error) {
	return _EndpointV2.Contract.GetRegisteredLibraries(&_EndpointV2.CallOpts)
}

// GetSendContext is a free data retrieval call binding the contract method 0x14f651a9.
//
// Solidity: function getSendContext() view returns(uint32, address)
func (_EndpointV2 *EndpointV2Caller) GetSendContext(opts *bind.CallOpts) (uint32, common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "getSendContext")

	if err != nil {
		return *new(uint32), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetSendContext is a free data retrieval call binding the contract method 0x14f651a9.
//
// Solidity: function getSendContext() view returns(uint32, address)
func (_EndpointV2 *EndpointV2Session) GetSendContext() (uint32, common.Address, error) {
	return _EndpointV2.Contract.GetSendContext(&_EndpointV2.CallOpts)
}

// GetSendContext is a free data retrieval call binding the contract method 0x14f651a9.
//
// Solidity: function getSendContext() view returns(uint32, address)
func (_EndpointV2 *EndpointV2CallerSession) GetSendContext() (uint32, common.Address, error) {
	return _EndpointV2.Contract.GetSendContext(&_EndpointV2.CallOpts)
}

// GetSendLibrary is a free data retrieval call binding the contract method 0xb96a277f.
//
// Solidity: function getSendLibrary(address _sender, uint32 _dstEid) view returns(address lib)
func (_EndpointV2 *EndpointV2Caller) GetSendLibrary(opts *bind.CallOpts, _sender common.Address, _dstEid uint32) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "getSendLibrary", _sender, _dstEid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSendLibrary is a free data retrieval call binding the contract method 0xb96a277f.
//
// Solidity: function getSendLibrary(address _sender, uint32 _dstEid) view returns(address lib)
func (_EndpointV2 *EndpointV2Session) GetSendLibrary(_sender common.Address, _dstEid uint32) (common.Address, error) {
	return _EndpointV2.Contract.GetSendLibrary(&_EndpointV2.CallOpts, _sender, _dstEid)
}

// GetSendLibrary is a free data retrieval call binding the contract method 0xb96a277f.
//
// Solidity: function getSendLibrary(address _sender, uint32 _dstEid) view returns(address lib)
func (_EndpointV2 *EndpointV2CallerSession) GetSendLibrary(_sender common.Address, _dstEid uint32) (common.Address, error) {
	return _EndpointV2.Contract.GetSendLibrary(&_EndpointV2.CallOpts, _sender, _dstEid)
}

// InboundNonce is a free data retrieval call binding the contract method 0xa0dd43fc.
//
// Solidity: function inboundNonce(address _receiver, uint32 _srcEid, bytes32 _sender) view returns(uint64)
func (_EndpointV2 *EndpointV2Caller) InboundNonce(opts *bind.CallOpts, _receiver common.Address, _srcEid uint32, _sender [32]byte) (uint64, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "inboundNonce", _receiver, _srcEid, _sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InboundNonce is a free data retrieval call binding the contract method 0xa0dd43fc.
//
// Solidity: function inboundNonce(address _receiver, uint32 _srcEid, bytes32 _sender) view returns(uint64)
func (_EndpointV2 *EndpointV2Session) InboundNonce(_receiver common.Address, _srcEid uint32, _sender [32]byte) (uint64, error) {
	return _EndpointV2.Contract.InboundNonce(&_EndpointV2.CallOpts, _receiver, _srcEid, _sender)
}

// InboundNonce is a free data retrieval call binding the contract method 0xa0dd43fc.
//
// Solidity: function inboundNonce(address _receiver, uint32 _srcEid, bytes32 _sender) view returns(uint64)
func (_EndpointV2 *EndpointV2CallerSession) InboundNonce(_receiver common.Address, _srcEid uint32, _sender [32]byte) (uint64, error) {
	return _EndpointV2.Contract.InboundNonce(&_EndpointV2.CallOpts, _receiver, _srcEid, _sender)
}

// InboundPayloadHash is a free data retrieval call binding the contract method 0xc9fc7bcd.
//
// Solidity: function inboundPayloadHash(address receiver, uint32 srcEid, bytes32 sender, uint64 inboundNonce) view returns(bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Caller) InboundPayloadHash(opts *bind.CallOpts, receiver common.Address, srcEid uint32, sender [32]byte, inboundNonce uint64) ([32]byte, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "inboundPayloadHash", receiver, srcEid, sender, inboundNonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboundPayloadHash is a free data retrieval call binding the contract method 0xc9fc7bcd.
//
// Solidity: function inboundPayloadHash(address receiver, uint32 srcEid, bytes32 sender, uint64 inboundNonce) view returns(bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Session) InboundPayloadHash(receiver common.Address, srcEid uint32, sender [32]byte, inboundNonce uint64) ([32]byte, error) {
	return _EndpointV2.Contract.InboundPayloadHash(&_EndpointV2.CallOpts, receiver, srcEid, sender, inboundNonce)
}

// InboundPayloadHash is a free data retrieval call binding the contract method 0xc9fc7bcd.
//
// Solidity: function inboundPayloadHash(address receiver, uint32 srcEid, bytes32 sender, uint64 inboundNonce) view returns(bytes32 payloadHash)
func (_EndpointV2 *EndpointV2CallerSession) InboundPayloadHash(receiver common.Address, srcEid uint32, sender [32]byte, inboundNonce uint64) ([32]byte, error) {
	return _EndpointV2.Contract.InboundPayloadHash(&_EndpointV2.CallOpts, receiver, srcEid, sender, inboundNonce)
}

// Initializable is a free data retrieval call binding the contract method 0x861e1ca5.
//
// Solidity: function initializable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_EndpointV2 *EndpointV2Caller) Initializable(opts *bind.CallOpts, _origin Origin, _receiver common.Address) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "initializable", _origin, _receiver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initializable is a free data retrieval call binding the contract method 0x861e1ca5.
//
// Solidity: function initializable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_EndpointV2 *EndpointV2Session) Initializable(_origin Origin, _receiver common.Address) (bool, error) {
	return _EndpointV2.Contract.Initializable(&_EndpointV2.CallOpts, _origin, _receiver)
}

// Initializable is a free data retrieval call binding the contract method 0x861e1ca5.
//
// Solidity: function initializable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) Initializable(_origin Origin, _receiver common.Address) (bool, error) {
	return _EndpointV2.Contract.Initializable(&_EndpointV2.CallOpts, _origin, _receiver)
}

// IsDefaultSendLibrary is a free data retrieval call binding the contract method 0xdc93c8a2.
//
// Solidity: function isDefaultSendLibrary(address _sender, uint32 _dstEid) view returns(bool)
func (_EndpointV2 *EndpointV2Caller) IsDefaultSendLibrary(opts *bind.CallOpts, _sender common.Address, _dstEid uint32) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "isDefaultSendLibrary", _sender, _dstEid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDefaultSendLibrary is a free data retrieval call binding the contract method 0xdc93c8a2.
//
// Solidity: function isDefaultSendLibrary(address _sender, uint32 _dstEid) view returns(bool)
func (_EndpointV2 *EndpointV2Session) IsDefaultSendLibrary(_sender common.Address, _dstEid uint32) (bool, error) {
	return _EndpointV2.Contract.IsDefaultSendLibrary(&_EndpointV2.CallOpts, _sender, _dstEid)
}

// IsDefaultSendLibrary is a free data retrieval call binding the contract method 0xdc93c8a2.
//
// Solidity: function isDefaultSendLibrary(address _sender, uint32 _dstEid) view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) IsDefaultSendLibrary(_sender common.Address, _dstEid uint32) (bool, error) {
	return _EndpointV2.Contract.IsDefaultSendLibrary(&_EndpointV2.CallOpts, _sender, _dstEid)
}

// IsRegisteredLibrary is a free data retrieval call binding the contract method 0xdc706a62.
//
// Solidity: function isRegisteredLibrary(address lib) view returns(bool)
func (_EndpointV2 *EndpointV2Caller) IsRegisteredLibrary(opts *bind.CallOpts, lib common.Address) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "isRegisteredLibrary", lib)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredLibrary is a free data retrieval call binding the contract method 0xdc706a62.
//
// Solidity: function isRegisteredLibrary(address lib) view returns(bool)
func (_EndpointV2 *EndpointV2Session) IsRegisteredLibrary(lib common.Address) (bool, error) {
	return _EndpointV2.Contract.IsRegisteredLibrary(&_EndpointV2.CallOpts, lib)
}

// IsRegisteredLibrary is a free data retrieval call binding the contract method 0xdc706a62.
//
// Solidity: function isRegisteredLibrary(address lib) view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) IsRegisteredLibrary(lib common.Address) (bool, error) {
	return _EndpointV2.Contract.IsRegisteredLibrary(&_EndpointV2.CallOpts, lib)
}

// IsSendingMessage is a free data retrieval call binding the contract method 0x79624ca9.
//
// Solidity: function isSendingMessage() view returns(bool)
func (_EndpointV2 *EndpointV2Caller) IsSendingMessage(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "isSendingMessage")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSendingMessage is a free data retrieval call binding the contract method 0x79624ca9.
//
// Solidity: function isSendingMessage() view returns(bool)
func (_EndpointV2 *EndpointV2Session) IsSendingMessage() (bool, error) {
	return _EndpointV2.Contract.IsSendingMessage(&_EndpointV2.CallOpts)
}

// IsSendingMessage is a free data retrieval call binding the contract method 0x79624ca9.
//
// Solidity: function isSendingMessage() view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) IsSendingMessage() (bool, error) {
	return _EndpointV2.Contract.IsSendingMessage(&_EndpointV2.CallOpts)
}

// IsSupportedEid is a free data retrieval call binding the contract method 0x6750cd4c.
//
// Solidity: function isSupportedEid(uint32 _eid) view returns(bool)
func (_EndpointV2 *EndpointV2Caller) IsSupportedEid(opts *bind.CallOpts, _eid uint32) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "isSupportedEid", _eid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedEid is a free data retrieval call binding the contract method 0x6750cd4c.
//
// Solidity: function isSupportedEid(uint32 _eid) view returns(bool)
func (_EndpointV2 *EndpointV2Session) IsSupportedEid(_eid uint32) (bool, error) {
	return _EndpointV2.Contract.IsSupportedEid(&_EndpointV2.CallOpts, _eid)
}

// IsSupportedEid is a free data retrieval call binding the contract method 0x6750cd4c.
//
// Solidity: function isSupportedEid(uint32 _eid) view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) IsSupportedEid(_eid uint32) (bool, error) {
	return _EndpointV2.Contract.IsSupportedEid(&_EndpointV2.CallOpts, _eid)
}

// IsValidReceiveLibrary is a free data retrieval call binding the contract method 0x9d7f9775.
//
// Solidity: function isValidReceiveLibrary(address _receiver, uint32 _srcEid, address _actualReceiveLib) view returns(bool)
func (_EndpointV2 *EndpointV2Caller) IsValidReceiveLibrary(opts *bind.CallOpts, _receiver common.Address, _srcEid uint32, _actualReceiveLib common.Address) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "isValidReceiveLibrary", _receiver, _srcEid, _actualReceiveLib)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidReceiveLibrary is a free data retrieval call binding the contract method 0x9d7f9775.
//
// Solidity: function isValidReceiveLibrary(address _receiver, uint32 _srcEid, address _actualReceiveLib) view returns(bool)
func (_EndpointV2 *EndpointV2Session) IsValidReceiveLibrary(_receiver common.Address, _srcEid uint32, _actualReceiveLib common.Address) (bool, error) {
	return _EndpointV2.Contract.IsValidReceiveLibrary(&_EndpointV2.CallOpts, _receiver, _srcEid, _actualReceiveLib)
}

// IsValidReceiveLibrary is a free data retrieval call binding the contract method 0x9d7f9775.
//
// Solidity: function isValidReceiveLibrary(address _receiver, uint32 _srcEid, address _actualReceiveLib) view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) IsValidReceiveLibrary(_receiver common.Address, _srcEid uint32, _actualReceiveLib common.Address) (bool, error) {
	return _EndpointV2.Contract.IsValidReceiveLibrary(&_EndpointV2.CallOpts, _receiver, _srcEid, _actualReceiveLib)
}

// LazyInboundNonce is a free data retrieval call binding the contract method 0x5b17bb70.
//
// Solidity: function lazyInboundNonce(address receiver, uint32 srcEid, bytes32 sender) view returns(uint64 nonce)
func (_EndpointV2 *EndpointV2Caller) LazyInboundNonce(opts *bind.CallOpts, receiver common.Address, srcEid uint32, sender [32]byte) (uint64, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "lazyInboundNonce", receiver, srcEid, sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LazyInboundNonce is a free data retrieval call binding the contract method 0x5b17bb70.
//
// Solidity: function lazyInboundNonce(address receiver, uint32 srcEid, bytes32 sender) view returns(uint64 nonce)
func (_EndpointV2 *EndpointV2Session) LazyInboundNonce(receiver common.Address, srcEid uint32, sender [32]byte) (uint64, error) {
	return _EndpointV2.Contract.LazyInboundNonce(&_EndpointV2.CallOpts, receiver, srcEid, sender)
}

// LazyInboundNonce is a free data retrieval call binding the contract method 0x5b17bb70.
//
// Solidity: function lazyInboundNonce(address receiver, uint32 srcEid, bytes32 sender) view returns(uint64 nonce)
func (_EndpointV2 *EndpointV2CallerSession) LazyInboundNonce(receiver common.Address, srcEid uint32, sender [32]byte) (uint64, error) {
	return _EndpointV2.Contract.LazyInboundNonce(&_EndpointV2.CallOpts, receiver, srcEid, sender)
}

// LzToken is a free data retrieval call binding the contract method 0xe4fe1d94.
//
// Solidity: function lzToken() view returns(address)
func (_EndpointV2 *EndpointV2Caller) LzToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "lzToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LzToken is a free data retrieval call binding the contract method 0xe4fe1d94.
//
// Solidity: function lzToken() view returns(address)
func (_EndpointV2 *EndpointV2Session) LzToken() (common.Address, error) {
	return _EndpointV2.Contract.LzToken(&_EndpointV2.CallOpts)
}

// LzToken is a free data retrieval call binding the contract method 0xe4fe1d94.
//
// Solidity: function lzToken() view returns(address)
func (_EndpointV2 *EndpointV2CallerSession) LzToken() (common.Address, error) {
	return _EndpointV2.Contract.LzToken(&_EndpointV2.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_EndpointV2 *EndpointV2Caller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_EndpointV2 *EndpointV2Session) NativeToken() (common.Address, error) {
	return _EndpointV2.Contract.NativeToken(&_EndpointV2.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_EndpointV2 *EndpointV2CallerSession) NativeToken() (common.Address, error) {
	return _EndpointV2.Contract.NativeToken(&_EndpointV2.CallOpts)
}

// NextGuid is a free data retrieval call binding the contract method 0xaafe5e07.
//
// Solidity: function nextGuid(address _sender, uint32 _dstEid, bytes32 _receiver) view returns(bytes32)
func (_EndpointV2 *EndpointV2Caller) NextGuid(opts *bind.CallOpts, _sender common.Address, _dstEid uint32, _receiver [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "nextGuid", _sender, _dstEid, _receiver)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextGuid is a free data retrieval call binding the contract method 0xaafe5e07.
//
// Solidity: function nextGuid(address _sender, uint32 _dstEid, bytes32 _receiver) view returns(bytes32)
func (_EndpointV2 *EndpointV2Session) NextGuid(_sender common.Address, _dstEid uint32, _receiver [32]byte) ([32]byte, error) {
	return _EndpointV2.Contract.NextGuid(&_EndpointV2.CallOpts, _sender, _dstEid, _receiver)
}

// NextGuid is a free data retrieval call binding the contract method 0xaafe5e07.
//
// Solidity: function nextGuid(address _sender, uint32 _dstEid, bytes32 _receiver) view returns(bytes32)
func (_EndpointV2 *EndpointV2CallerSession) NextGuid(_sender common.Address, _dstEid uint32, _receiver [32]byte) ([32]byte, error) {
	return _EndpointV2.Contract.NextGuid(&_EndpointV2.CallOpts, _sender, _dstEid, _receiver)
}

// OutboundNonce is a free data retrieval call binding the contract method 0x9c6d7340.
//
// Solidity: function outboundNonce(address sender, uint32 dstEid, bytes32 receiver) view returns(uint64 nonce)
func (_EndpointV2 *EndpointV2Caller) OutboundNonce(opts *bind.CallOpts, sender common.Address, dstEid uint32, receiver [32]byte) (uint64, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "outboundNonce", sender, dstEid, receiver)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OutboundNonce is a free data retrieval call binding the contract method 0x9c6d7340.
//
// Solidity: function outboundNonce(address sender, uint32 dstEid, bytes32 receiver) view returns(uint64 nonce)
func (_EndpointV2 *EndpointV2Session) OutboundNonce(sender common.Address, dstEid uint32, receiver [32]byte) (uint64, error) {
	return _EndpointV2.Contract.OutboundNonce(&_EndpointV2.CallOpts, sender, dstEid, receiver)
}

// OutboundNonce is a free data retrieval call binding the contract method 0x9c6d7340.
//
// Solidity: function outboundNonce(address sender, uint32 dstEid, bytes32 receiver) view returns(uint64 nonce)
func (_EndpointV2 *EndpointV2CallerSession) OutboundNonce(sender common.Address, dstEid uint32, receiver [32]byte) (uint64, error) {
	return _EndpointV2.Contract.OutboundNonce(&_EndpointV2.CallOpts, sender, dstEid, receiver)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EndpointV2 *EndpointV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EndpointV2 *EndpointV2Session) Owner() (common.Address, error) {
	return _EndpointV2.Contract.Owner(&_EndpointV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EndpointV2 *EndpointV2CallerSession) Owner() (common.Address, error) {
	return _EndpointV2.Contract.Owner(&_EndpointV2.CallOpts)
}

// Quote is a free data retrieval call binding the contract method 0xddc28c58.
//
// Solidity: function quote((uint32,bytes32,bytes,bytes,bool) _params, address _sender) view returns((uint256,uint256))
func (_EndpointV2 *EndpointV2Caller) Quote(opts *bind.CallOpts, _params MessagingParams, _sender common.Address) (MessagingFee, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "quote", _params, _sender)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xddc28c58.
//
// Solidity: function quote((uint32,bytes32,bytes,bytes,bool) _params, address _sender) view returns((uint256,uint256))
func (_EndpointV2 *EndpointV2Session) Quote(_params MessagingParams, _sender common.Address) (MessagingFee, error) {
	return _EndpointV2.Contract.Quote(&_EndpointV2.CallOpts, _params, _sender)
}

// Quote is a free data retrieval call binding the contract method 0xddc28c58.
//
// Solidity: function quote((uint32,bytes32,bytes,bytes,bool) _params, address _sender) view returns((uint256,uint256))
func (_EndpointV2 *EndpointV2CallerSession) Quote(_params MessagingParams, _sender common.Address) (MessagingFee, error) {
	return _EndpointV2.Contract.Quote(&_EndpointV2.CallOpts, _params, _sender)
}

// ReceiveLibraryTimeout is a free data retrieval call binding the contract method 0xef667aa1.
//
// Solidity: function receiveLibraryTimeout(address receiver, uint32 srcEid) view returns(address lib, uint256 expiry)
func (_EndpointV2 *EndpointV2Caller) ReceiveLibraryTimeout(opts *bind.CallOpts, receiver common.Address, srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "receiveLibraryTimeout", receiver, srcEid)

	outstruct := new(struct {
		Lib    common.Address
		Expiry *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lib = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Expiry = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReceiveLibraryTimeout is a free data retrieval call binding the contract method 0xef667aa1.
//
// Solidity: function receiveLibraryTimeout(address receiver, uint32 srcEid) view returns(address lib, uint256 expiry)
func (_EndpointV2 *EndpointV2Session) ReceiveLibraryTimeout(receiver common.Address, srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	return _EndpointV2.Contract.ReceiveLibraryTimeout(&_EndpointV2.CallOpts, receiver, srcEid)
}

// ReceiveLibraryTimeout is a free data retrieval call binding the contract method 0xef667aa1.
//
// Solidity: function receiveLibraryTimeout(address receiver, uint32 srcEid) view returns(address lib, uint256 expiry)
func (_EndpointV2 *EndpointV2CallerSession) ReceiveLibraryTimeout(receiver common.Address, srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	return _EndpointV2.Contract.ReceiveLibraryTimeout(&_EndpointV2.CallOpts, receiver, srcEid)
}

// Verifiable is a free data retrieval call binding the contract method 0xc9a54a99.
//
// Solidity: function verifiable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_EndpointV2 *EndpointV2Caller) Verifiable(opts *bind.CallOpts, _origin Origin, _receiver common.Address) (bool, error) {
	var out []interface{}
	err := _EndpointV2.contract.Call(opts, &out, "verifiable", _origin, _receiver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verifiable is a free data retrieval call binding the contract method 0xc9a54a99.
//
// Solidity: function verifiable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_EndpointV2 *EndpointV2Session) Verifiable(_origin Origin, _receiver common.Address) (bool, error) {
	return _EndpointV2.Contract.Verifiable(&_EndpointV2.CallOpts, _origin, _receiver)
}

// Verifiable is a free data retrieval call binding the contract method 0xc9a54a99.
//
// Solidity: function verifiable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_EndpointV2 *EndpointV2CallerSession) Verifiable(_origin Origin, _receiver common.Address) (bool, error) {
	return _EndpointV2.Contract.Verifiable(&_EndpointV2.CallOpts, _origin, _receiver)
}

// Burn is a paid mutator transaction binding the contract method 0x40f80683.
//
// Solidity: function burn(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_EndpointV2 *EndpointV2Transactor) Burn(opts *bind.TransactOpts, _oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "burn", _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Burn is a paid mutator transaction binding the contract method 0x40f80683.
//
// Solidity: function burn(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_EndpointV2 *EndpointV2Session) Burn(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Burn(&_EndpointV2.TransactOpts, _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Burn is a paid mutator transaction binding the contract method 0x40f80683.
//
// Solidity: function burn(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_EndpointV2 *EndpointV2TransactorSession) Burn(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Burn(&_EndpointV2.TransactOpts, _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Clear is a paid mutator transaction binding the contract method 0x2a56c1b0.
//
// Solidity: function clear(address _oapp, (uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message) returns()
func (_EndpointV2 *EndpointV2Transactor) Clear(opts *bind.TransactOpts, _oapp common.Address, _origin Origin, _guid [32]byte, _message []byte) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "clear", _oapp, _origin, _guid, _message)
}

// Clear is a paid mutator transaction binding the contract method 0x2a56c1b0.
//
// Solidity: function clear(address _oapp, (uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message) returns()
func (_EndpointV2 *EndpointV2Session) Clear(_oapp common.Address, _origin Origin, _guid [32]byte, _message []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Clear(&_EndpointV2.TransactOpts, _oapp, _origin, _guid, _message)
}

// Clear is a paid mutator transaction binding the contract method 0x2a56c1b0.
//
// Solidity: function clear(address _oapp, (uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message) returns()
func (_EndpointV2 *EndpointV2TransactorSession) Clear(_oapp common.Address, _origin Origin, _guid [32]byte, _message []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Clear(&_EndpointV2.TransactOpts, _oapp, _origin, _guid, _message)
}

// LzCompose is a paid mutator transaction binding the contract method 0x91d20fa1.
//
// Solidity: function lzCompose(address _from, address _to, bytes32 _guid, uint16 _index, bytes _message, bytes _extraData) payable returns()
func (_EndpointV2 *EndpointV2Transactor) LzCompose(opts *bind.TransactOpts, _from common.Address, _to common.Address, _guid [32]byte, _index uint16, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "lzCompose", _from, _to, _guid, _index, _message, _extraData)
}

// LzCompose is a paid mutator transaction binding the contract method 0x91d20fa1.
//
// Solidity: function lzCompose(address _from, address _to, bytes32 _guid, uint16 _index, bytes _message, bytes _extraData) payable returns()
func (_EndpointV2 *EndpointV2Session) LzCompose(_from common.Address, _to common.Address, _guid [32]byte, _index uint16, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.LzCompose(&_EndpointV2.TransactOpts, _from, _to, _guid, _index, _message, _extraData)
}

// LzCompose is a paid mutator transaction binding the contract method 0x91d20fa1.
//
// Solidity: function lzCompose(address _from, address _to, bytes32 _guid, uint16 _index, bytes _message, bytes _extraData) payable returns()
func (_EndpointV2 *EndpointV2TransactorSession) LzCompose(_from common.Address, _to common.Address, _guid [32]byte, _index uint16, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.LzCompose(&_EndpointV2.TransactOpts, _from, _to, _guid, _index, _message, _extraData)
}

// LzComposeAlert is a paid mutator transaction binding the contract method 0x697fe6b6.
//
// Solidity: function lzComposeAlert(address _from, address _to, bytes32 _guid, uint16 _index, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_EndpointV2 *EndpointV2Transactor) LzComposeAlert(opts *bind.TransactOpts, _from common.Address, _to common.Address, _guid [32]byte, _index uint16, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "lzComposeAlert", _from, _to, _guid, _index, _gas, _value, _message, _extraData, _reason)
}

// LzComposeAlert is a paid mutator transaction binding the contract method 0x697fe6b6.
//
// Solidity: function lzComposeAlert(address _from, address _to, bytes32 _guid, uint16 _index, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_EndpointV2 *EndpointV2Session) LzComposeAlert(_from common.Address, _to common.Address, _guid [32]byte, _index uint16, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.LzComposeAlert(&_EndpointV2.TransactOpts, _from, _to, _guid, _index, _gas, _value, _message, _extraData, _reason)
}

// LzComposeAlert is a paid mutator transaction binding the contract method 0x697fe6b6.
//
// Solidity: function lzComposeAlert(address _from, address _to, bytes32 _guid, uint16 _index, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_EndpointV2 *EndpointV2TransactorSession) LzComposeAlert(_from common.Address, _to common.Address, _guid [32]byte, _index uint16, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.LzComposeAlert(&_EndpointV2.TransactOpts, _from, _to, _guid, _index, _gas, _value, _message, _extraData, _reason)
}

// LzReceive is a paid mutator transaction binding the contract method 0x0c0c389e.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, bytes _message, bytes _extraData) payable returns()
func (_EndpointV2 *EndpointV2Transactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _receiver common.Address, _guid [32]byte, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "lzReceive", _origin, _receiver, _guid, _message, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x0c0c389e.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, bytes _message, bytes _extraData) payable returns()
func (_EndpointV2 *EndpointV2Session) LzReceive(_origin Origin, _receiver common.Address, _guid [32]byte, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.LzReceive(&_EndpointV2.TransactOpts, _origin, _receiver, _guid, _message, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x0c0c389e.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, bytes _message, bytes _extraData) payable returns()
func (_EndpointV2 *EndpointV2TransactorSession) LzReceive(_origin Origin, _receiver common.Address, _guid [32]byte, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.LzReceive(&_EndpointV2.TransactOpts, _origin, _receiver, _guid, _message, _extraData)
}

// LzReceiveAlert is a paid mutator transaction binding the contract method 0x6bf73fa3.
//
// Solidity: function lzReceiveAlert((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_EndpointV2 *EndpointV2Transactor) LzReceiveAlert(opts *bind.TransactOpts, _origin Origin, _receiver common.Address, _guid [32]byte, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "lzReceiveAlert", _origin, _receiver, _guid, _gas, _value, _message, _extraData, _reason)
}

// LzReceiveAlert is a paid mutator transaction binding the contract method 0x6bf73fa3.
//
// Solidity: function lzReceiveAlert((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_EndpointV2 *EndpointV2Session) LzReceiveAlert(_origin Origin, _receiver common.Address, _guid [32]byte, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.LzReceiveAlert(&_EndpointV2.TransactOpts, _origin, _receiver, _guid, _gas, _value, _message, _extraData, _reason)
}

// LzReceiveAlert is a paid mutator transaction binding the contract method 0x6bf73fa3.
//
// Solidity: function lzReceiveAlert((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_EndpointV2 *EndpointV2TransactorSession) LzReceiveAlert(_origin Origin, _receiver common.Address, _guid [32]byte, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.LzReceiveAlert(&_EndpointV2.TransactOpts, _origin, _receiver, _guid, _gas, _value, _message, _extraData, _reason)
}

// Nilify is a paid mutator transaction binding the contract method 0x2e80fbf3.
//
// Solidity: function nilify(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_EndpointV2 *EndpointV2Transactor) Nilify(opts *bind.TransactOpts, _oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "nilify", _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Nilify is a paid mutator transaction binding the contract method 0x2e80fbf3.
//
// Solidity: function nilify(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_EndpointV2 *EndpointV2Session) Nilify(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Nilify(&_EndpointV2.TransactOpts, _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Nilify is a paid mutator transaction binding the contract method 0x2e80fbf3.
//
// Solidity: function nilify(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_EndpointV2 *EndpointV2TransactorSession) Nilify(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Nilify(&_EndpointV2.TransactOpts, _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xa7229fd9.
//
// Solidity: function recoverToken(address _token, address _to, uint256 _amount) returns()
func (_EndpointV2 *EndpointV2Transactor) RecoverToken(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "recoverToken", _token, _to, _amount)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xa7229fd9.
//
// Solidity: function recoverToken(address _token, address _to, uint256 _amount) returns()
func (_EndpointV2 *EndpointV2Session) RecoverToken(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _EndpointV2.Contract.RecoverToken(&_EndpointV2.TransactOpts, _token, _to, _amount)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xa7229fd9.
//
// Solidity: function recoverToken(address _token, address _to, uint256 _amount) returns()
func (_EndpointV2 *EndpointV2TransactorSession) RecoverToken(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _EndpointV2.Contract.RecoverToken(&_EndpointV2.TransactOpts, _token, _to, _amount)
}

// RegisterLibrary is a paid mutator transaction binding the contract method 0xe8964e81.
//
// Solidity: function registerLibrary(address _lib) returns()
func (_EndpointV2 *EndpointV2Transactor) RegisterLibrary(opts *bind.TransactOpts, _lib common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "registerLibrary", _lib)
}

// RegisterLibrary is a paid mutator transaction binding the contract method 0xe8964e81.
//
// Solidity: function registerLibrary(address _lib) returns()
func (_EndpointV2 *EndpointV2Session) RegisterLibrary(_lib common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.RegisterLibrary(&_EndpointV2.TransactOpts, _lib)
}

// RegisterLibrary is a paid mutator transaction binding the contract method 0xe8964e81.
//
// Solidity: function registerLibrary(address _lib) returns()
func (_EndpointV2 *EndpointV2TransactorSession) RegisterLibrary(_lib common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.RegisterLibrary(&_EndpointV2.TransactOpts, _lib)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EndpointV2 *EndpointV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EndpointV2 *EndpointV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _EndpointV2.Contract.RenounceOwnership(&_EndpointV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EndpointV2 *EndpointV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EndpointV2.Contract.RenounceOwnership(&_EndpointV2.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x2637a450.
//
// Solidity: function send((uint32,bytes32,bytes,bytes,bool) _params, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)))
func (_EndpointV2 *EndpointV2Transactor) Send(opts *bind.TransactOpts, _params MessagingParams, _refundAddress common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "send", _params, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0x2637a450.
//
// Solidity: function send((uint32,bytes32,bytes,bytes,bool) _params, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)))
func (_EndpointV2 *EndpointV2Session) Send(_params MessagingParams, _refundAddress common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.Send(&_EndpointV2.TransactOpts, _params, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0x2637a450.
//
// Solidity: function send((uint32,bytes32,bytes,bytes,bool) _params, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)))
func (_EndpointV2 *EndpointV2TransactorSession) Send(_params MessagingParams, _refundAddress common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.Send(&_EndpointV2.TransactOpts, _params, _refundAddress)
}

// SendCompose is a paid mutator transaction binding the contract method 0x7cb59012.
//
// Solidity: function sendCompose(address _to, bytes32 _guid, uint16 _index, bytes _message) returns()
func (_EndpointV2 *EndpointV2Transactor) SendCompose(opts *bind.TransactOpts, _to common.Address, _guid [32]byte, _index uint16, _message []byte) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "sendCompose", _to, _guid, _index, _message)
}

// SendCompose is a paid mutator transaction binding the contract method 0x7cb59012.
//
// Solidity: function sendCompose(address _to, bytes32 _guid, uint16 _index, bytes _message) returns()
func (_EndpointV2 *EndpointV2Session) SendCompose(_to common.Address, _guid [32]byte, _index uint16, _message []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.SendCompose(&_EndpointV2.TransactOpts, _to, _guid, _index, _message)
}

// SendCompose is a paid mutator transaction binding the contract method 0x7cb59012.
//
// Solidity: function sendCompose(address _to, bytes32 _guid, uint16 _index, bytes _message) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SendCompose(_to common.Address, _guid [32]byte, _index uint16, _message []byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.SendCompose(&_EndpointV2.TransactOpts, _to, _guid, _index, _message)
}

// SetConfig is a paid mutator transaction binding the contract method 0x6dbd9f90.
//
// Solidity: function setConfig(address _oapp, address _lib, (uint32,uint32,bytes)[] _params) returns()
func (_EndpointV2 *EndpointV2Transactor) SetConfig(opts *bind.TransactOpts, _oapp common.Address, _lib common.Address, _params []SetConfigParam) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setConfig", _oapp, _lib, _params)
}

// SetConfig is a paid mutator transaction binding the contract method 0x6dbd9f90.
//
// Solidity: function setConfig(address _oapp, address _lib, (uint32,uint32,bytes)[] _params) returns()
func (_EndpointV2 *EndpointV2Session) SetConfig(_oapp common.Address, _lib common.Address, _params []SetConfigParam) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetConfig(&_EndpointV2.TransactOpts, _oapp, _lib, _params)
}

// SetConfig is a paid mutator transaction binding the contract method 0x6dbd9f90.
//
// Solidity: function setConfig(address _oapp, address _lib, (uint32,uint32,bytes)[] _params) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetConfig(_oapp common.Address, _lib common.Address, _params []SetConfigParam) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetConfig(&_EndpointV2.TransactOpts, _oapp, _lib, _params)
}

// SetDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0xa718531b.
//
// Solidity: function setDefaultReceiveLibrary(uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_EndpointV2 *EndpointV2Transactor) SetDefaultReceiveLibrary(opts *bind.TransactOpts, _eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setDefaultReceiveLibrary", _eid, _newLib, _gracePeriod)
}

// SetDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0xa718531b.
//
// Solidity: function setDefaultReceiveLibrary(uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_EndpointV2 *EndpointV2Session) SetDefaultReceiveLibrary(_eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetDefaultReceiveLibrary(&_EndpointV2.TransactOpts, _eid, _newLib, _gracePeriod)
}

// SetDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0xa718531b.
//
// Solidity: function setDefaultReceiveLibrary(uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetDefaultReceiveLibrary(_eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetDefaultReceiveLibrary(&_EndpointV2.TransactOpts, _eid, _newLib, _gracePeriod)
}

// SetDefaultReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0xd4b4ec8f.
//
// Solidity: function setDefaultReceiveLibraryTimeout(uint32 _eid, address _lib, uint256 _expiry) returns()
func (_EndpointV2 *EndpointV2Transactor) SetDefaultReceiveLibraryTimeout(opts *bind.TransactOpts, _eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setDefaultReceiveLibraryTimeout", _eid, _lib, _expiry)
}

// SetDefaultReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0xd4b4ec8f.
//
// Solidity: function setDefaultReceiveLibraryTimeout(uint32 _eid, address _lib, uint256 _expiry) returns()
func (_EndpointV2 *EndpointV2Session) SetDefaultReceiveLibraryTimeout(_eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetDefaultReceiveLibraryTimeout(&_EndpointV2.TransactOpts, _eid, _lib, _expiry)
}

// SetDefaultReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0xd4b4ec8f.
//
// Solidity: function setDefaultReceiveLibraryTimeout(uint32 _eid, address _lib, uint256 _expiry) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetDefaultReceiveLibraryTimeout(_eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetDefaultReceiveLibraryTimeout(&_EndpointV2.TransactOpts, _eid, _lib, _expiry)
}

// SetDefaultSendLibrary is a paid mutator transaction binding the contract method 0xaafea312.
//
// Solidity: function setDefaultSendLibrary(uint32 _eid, address _newLib) returns()
func (_EndpointV2 *EndpointV2Transactor) SetDefaultSendLibrary(opts *bind.TransactOpts, _eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setDefaultSendLibrary", _eid, _newLib)
}

// SetDefaultSendLibrary is a paid mutator transaction binding the contract method 0xaafea312.
//
// Solidity: function setDefaultSendLibrary(uint32 _eid, address _newLib) returns()
func (_EndpointV2 *EndpointV2Session) SetDefaultSendLibrary(_eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetDefaultSendLibrary(&_EndpointV2.TransactOpts, _eid, _newLib)
}

// SetDefaultSendLibrary is a paid mutator transaction binding the contract method 0xaafea312.
//
// Solidity: function setDefaultSendLibrary(uint32 _eid, address _newLib) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetDefaultSendLibrary(_eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetDefaultSendLibrary(&_EndpointV2.TransactOpts, _eid, _newLib)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_EndpointV2 *EndpointV2Transactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_EndpointV2 *EndpointV2Session) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetDelegate(&_EndpointV2.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetDelegate(&_EndpointV2.TransactOpts, _delegate)
}

// SetLzToken is a paid mutator transaction binding the contract method 0xc28e0eed.
//
// Solidity: function setLzToken(address _lzToken) returns()
func (_EndpointV2 *EndpointV2Transactor) SetLzToken(opts *bind.TransactOpts, _lzToken common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setLzToken", _lzToken)
}

// SetLzToken is a paid mutator transaction binding the contract method 0xc28e0eed.
//
// Solidity: function setLzToken(address _lzToken) returns()
func (_EndpointV2 *EndpointV2Session) SetLzToken(_lzToken common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetLzToken(&_EndpointV2.TransactOpts, _lzToken)
}

// SetLzToken is a paid mutator transaction binding the contract method 0xc28e0eed.
//
// Solidity: function setLzToken(address _lzToken) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetLzToken(_lzToken common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetLzToken(&_EndpointV2.TransactOpts, _lzToken)
}

// SetReceiveLibrary is a paid mutator transaction binding the contract method 0x6a14d715.
//
// Solidity: function setReceiveLibrary(address _oapp, uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_EndpointV2 *EndpointV2Transactor) SetReceiveLibrary(opts *bind.TransactOpts, _oapp common.Address, _eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setReceiveLibrary", _oapp, _eid, _newLib, _gracePeriod)
}

// SetReceiveLibrary is a paid mutator transaction binding the contract method 0x6a14d715.
//
// Solidity: function setReceiveLibrary(address _oapp, uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_EndpointV2 *EndpointV2Session) SetReceiveLibrary(_oapp common.Address, _eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetReceiveLibrary(&_EndpointV2.TransactOpts, _oapp, _eid, _newLib, _gracePeriod)
}

// SetReceiveLibrary is a paid mutator transaction binding the contract method 0x6a14d715.
//
// Solidity: function setReceiveLibrary(address _oapp, uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetReceiveLibrary(_oapp common.Address, _eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetReceiveLibrary(&_EndpointV2.TransactOpts, _oapp, _eid, _newLib, _gracePeriod)
}

// SetReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0x183c834f.
//
// Solidity: function setReceiveLibraryTimeout(address _oapp, uint32 _eid, address _lib, uint256 _expiry) returns()
func (_EndpointV2 *EndpointV2Transactor) SetReceiveLibraryTimeout(opts *bind.TransactOpts, _oapp common.Address, _eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setReceiveLibraryTimeout", _oapp, _eid, _lib, _expiry)
}

// SetReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0x183c834f.
//
// Solidity: function setReceiveLibraryTimeout(address _oapp, uint32 _eid, address _lib, uint256 _expiry) returns()
func (_EndpointV2 *EndpointV2Session) SetReceiveLibraryTimeout(_oapp common.Address, _eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetReceiveLibraryTimeout(&_EndpointV2.TransactOpts, _oapp, _eid, _lib, _expiry)
}

// SetReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0x183c834f.
//
// Solidity: function setReceiveLibraryTimeout(address _oapp, uint32 _eid, address _lib, uint256 _expiry) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetReceiveLibraryTimeout(_oapp common.Address, _eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetReceiveLibraryTimeout(&_EndpointV2.TransactOpts, _oapp, _eid, _lib, _expiry)
}

// SetSendLibrary is a paid mutator transaction binding the contract method 0x9535ff30.
//
// Solidity: function setSendLibrary(address _oapp, uint32 _eid, address _newLib) returns()
func (_EndpointV2 *EndpointV2Transactor) SetSendLibrary(opts *bind.TransactOpts, _oapp common.Address, _eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "setSendLibrary", _oapp, _eid, _newLib)
}

// SetSendLibrary is a paid mutator transaction binding the contract method 0x9535ff30.
//
// Solidity: function setSendLibrary(address _oapp, uint32 _eid, address _newLib) returns()
func (_EndpointV2 *EndpointV2Session) SetSendLibrary(_oapp common.Address, _eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetSendLibrary(&_EndpointV2.TransactOpts, _oapp, _eid, _newLib)
}

// SetSendLibrary is a paid mutator transaction binding the contract method 0x9535ff30.
//
// Solidity: function setSendLibrary(address _oapp, uint32 _eid, address _newLib) returns()
func (_EndpointV2 *EndpointV2TransactorSession) SetSendLibrary(_oapp common.Address, _eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.SetSendLibrary(&_EndpointV2.TransactOpts, _oapp, _eid, _newLib)
}

// Skip is a paid mutator transaction binding the contract method 0xd70b8902.
//
// Solidity: function skip(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce) returns()
func (_EndpointV2 *EndpointV2Transactor) Skip(opts *bind.TransactOpts, _oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "skip", _oapp, _srcEid, _sender, _nonce)
}

// Skip is a paid mutator transaction binding the contract method 0xd70b8902.
//
// Solidity: function skip(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce) returns()
func (_EndpointV2 *EndpointV2Session) Skip(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64) (*types.Transaction, error) {
	return _EndpointV2.Contract.Skip(&_EndpointV2.TransactOpts, _oapp, _srcEid, _sender, _nonce)
}

// Skip is a paid mutator transaction binding the contract method 0xd70b8902.
//
// Solidity: function skip(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce) returns()
func (_EndpointV2 *EndpointV2TransactorSession) Skip(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64) (*types.Transaction, error) {
	return _EndpointV2.Contract.Skip(&_EndpointV2.TransactOpts, _oapp, _srcEid, _sender, _nonce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EndpointV2 *EndpointV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EndpointV2 *EndpointV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.TransferOwnership(&_EndpointV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EndpointV2 *EndpointV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EndpointV2.Contract.TransferOwnership(&_EndpointV2.TransactOpts, newOwner)
}

// Verify is a paid mutator transaction binding the contract method 0xa825d747.
//
// Solidity: function verify((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) returns()
func (_EndpointV2 *EndpointV2Transactor) Verify(opts *bind.TransactOpts, _origin Origin, _receiver common.Address, _payloadHash [32]byte) (*types.Transaction, error) {
	return _EndpointV2.contract.Transact(opts, "verify", _origin, _receiver, _payloadHash)
}

// Verify is a paid mutator transaction binding the contract method 0xa825d747.
//
// Solidity: function verify((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) returns()
func (_EndpointV2 *EndpointV2Session) Verify(_origin Origin, _receiver common.Address, _payloadHash [32]byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Verify(&_EndpointV2.TransactOpts, _origin, _receiver, _payloadHash)
}

// Verify is a paid mutator transaction binding the contract method 0xa825d747.
//
// Solidity: function verify((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) returns()
func (_EndpointV2 *EndpointV2TransactorSession) Verify(_origin Origin, _receiver common.Address, _payloadHash [32]byte) (*types.Transaction, error) {
	return _EndpointV2.Contract.Verify(&_EndpointV2.TransactOpts, _origin, _receiver, _payloadHash)
}

// EndpointV2ComposeDeliveredIterator is returned from FilterComposeDelivered and is used to iterate over the raw logs and unpacked data for ComposeDelivered events raised by the EndpointV2 contract.
type EndpointV2ComposeDeliveredIterator struct {
	Event *EndpointV2ComposeDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2ComposeDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2ComposeDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2ComposeDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2ComposeDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2ComposeDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2ComposeDelivered represents a ComposeDelivered event raised by the EndpointV2 contract.
type EndpointV2ComposeDelivered struct {
	From  common.Address
	To    common.Address
	Guid  [32]byte
	Index uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterComposeDelivered is a free log retrieval operation binding the contract event 0x0036c98efcf9e6641dfbc9051f66f405253e8e0c2ab4a24dccda15595b7378c8.
//
// Solidity: event ComposeDelivered(address from, address to, bytes32 guid, uint16 index)
func (_EndpointV2 *EndpointV2Filterer) FilterComposeDelivered(opts *bind.FilterOpts) (*EndpointV2ComposeDeliveredIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "ComposeDelivered")
	if err != nil {
		return nil, err
	}
	return &EndpointV2ComposeDeliveredIterator{contract: _EndpointV2.contract, event: "ComposeDelivered", logs: logs, sub: sub}, nil
}

// WatchComposeDelivered is a free log subscription operation binding the contract event 0x0036c98efcf9e6641dfbc9051f66f405253e8e0c2ab4a24dccda15595b7378c8.
//
// Solidity: event ComposeDelivered(address from, address to, bytes32 guid, uint16 index)
func (_EndpointV2 *EndpointV2Filterer) WatchComposeDelivered(opts *bind.WatchOpts, sink chan<- *EndpointV2ComposeDelivered) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "ComposeDelivered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2ComposeDelivered)
				if err := _EndpointV2.contract.UnpackLog(event, "ComposeDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseComposeDelivered is a log parse operation binding the contract event 0x0036c98efcf9e6641dfbc9051f66f405253e8e0c2ab4a24dccda15595b7378c8.
//
// Solidity: event ComposeDelivered(address from, address to, bytes32 guid, uint16 index)
func (_EndpointV2 *EndpointV2Filterer) ParseComposeDelivered(log types.Log) (*EndpointV2ComposeDelivered, error) {
	event := new(EndpointV2ComposeDelivered)
	if err := _EndpointV2.contract.UnpackLog(event, "ComposeDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2ComposeSentIterator is returned from FilterComposeSent and is used to iterate over the raw logs and unpacked data for ComposeSent events raised by the EndpointV2 contract.
type EndpointV2ComposeSentIterator struct {
	Event *EndpointV2ComposeSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2ComposeSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2ComposeSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2ComposeSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2ComposeSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2ComposeSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2ComposeSent represents a ComposeSent event raised by the EndpointV2 contract.
type EndpointV2ComposeSent struct {
	From    common.Address
	To      common.Address
	Guid    [32]byte
	Index   uint16
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterComposeSent is a free log retrieval operation binding the contract event 0x3d52ff888d033fd3dd1d8057da59e850c91d91a72c41dfa445b247dfedeb6dc1.
//
// Solidity: event ComposeSent(address from, address to, bytes32 guid, uint16 index, bytes message)
func (_EndpointV2 *EndpointV2Filterer) FilterComposeSent(opts *bind.FilterOpts) (*EndpointV2ComposeSentIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "ComposeSent")
	if err != nil {
		return nil, err
	}
	return &EndpointV2ComposeSentIterator{contract: _EndpointV2.contract, event: "ComposeSent", logs: logs, sub: sub}, nil
}

// WatchComposeSent is a free log subscription operation binding the contract event 0x3d52ff888d033fd3dd1d8057da59e850c91d91a72c41dfa445b247dfedeb6dc1.
//
// Solidity: event ComposeSent(address from, address to, bytes32 guid, uint16 index, bytes message)
func (_EndpointV2 *EndpointV2Filterer) WatchComposeSent(opts *bind.WatchOpts, sink chan<- *EndpointV2ComposeSent) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "ComposeSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2ComposeSent)
				if err := _EndpointV2.contract.UnpackLog(event, "ComposeSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseComposeSent is a log parse operation binding the contract event 0x3d52ff888d033fd3dd1d8057da59e850c91d91a72c41dfa445b247dfedeb6dc1.
//
// Solidity: event ComposeSent(address from, address to, bytes32 guid, uint16 index, bytes message)
func (_EndpointV2 *EndpointV2Filterer) ParseComposeSent(log types.Log) (*EndpointV2ComposeSent, error) {
	event := new(EndpointV2ComposeSent)
	if err := _EndpointV2.contract.UnpackLog(event, "ComposeSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2DefaultReceiveLibrarySetIterator is returned from FilterDefaultReceiveLibrarySet and is used to iterate over the raw logs and unpacked data for DefaultReceiveLibrarySet events raised by the EndpointV2 contract.
type EndpointV2DefaultReceiveLibrarySetIterator struct {
	Event *EndpointV2DefaultReceiveLibrarySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2DefaultReceiveLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2DefaultReceiveLibrarySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2DefaultReceiveLibrarySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2DefaultReceiveLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2DefaultReceiveLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2DefaultReceiveLibrarySet represents a DefaultReceiveLibrarySet event raised by the EndpointV2 contract.
type EndpointV2DefaultReceiveLibrarySet struct {
	Eid    uint32
	NewLib common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultReceiveLibrarySet is a free log retrieval operation binding the contract event 0xc16891855cffb4a5ac51ac11864a3f3c96ba816cc45fe686c987ae36277de5ec.
//
// Solidity: event DefaultReceiveLibrarySet(uint32 eid, address newLib)
func (_EndpointV2 *EndpointV2Filterer) FilterDefaultReceiveLibrarySet(opts *bind.FilterOpts) (*EndpointV2DefaultReceiveLibrarySetIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "DefaultReceiveLibrarySet")
	if err != nil {
		return nil, err
	}
	return &EndpointV2DefaultReceiveLibrarySetIterator{contract: _EndpointV2.contract, event: "DefaultReceiveLibrarySet", logs: logs, sub: sub}, nil
}

// WatchDefaultReceiveLibrarySet is a free log subscription operation binding the contract event 0xc16891855cffb4a5ac51ac11864a3f3c96ba816cc45fe686c987ae36277de5ec.
//
// Solidity: event DefaultReceiveLibrarySet(uint32 eid, address newLib)
func (_EndpointV2 *EndpointV2Filterer) WatchDefaultReceiveLibrarySet(opts *bind.WatchOpts, sink chan<- *EndpointV2DefaultReceiveLibrarySet) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "DefaultReceiveLibrarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2DefaultReceiveLibrarySet)
				if err := _EndpointV2.contract.UnpackLog(event, "DefaultReceiveLibrarySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultReceiveLibrarySet is a log parse operation binding the contract event 0xc16891855cffb4a5ac51ac11864a3f3c96ba816cc45fe686c987ae36277de5ec.
//
// Solidity: event DefaultReceiveLibrarySet(uint32 eid, address newLib)
func (_EndpointV2 *EndpointV2Filterer) ParseDefaultReceiveLibrarySet(log types.Log) (*EndpointV2DefaultReceiveLibrarySet, error) {
	event := new(EndpointV2DefaultReceiveLibrarySet)
	if err := _EndpointV2.contract.UnpackLog(event, "DefaultReceiveLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2DefaultReceiveLibraryTimeoutSetIterator is returned from FilterDefaultReceiveLibraryTimeoutSet and is used to iterate over the raw logs and unpacked data for DefaultReceiveLibraryTimeoutSet events raised by the EndpointV2 contract.
type EndpointV2DefaultReceiveLibraryTimeoutSetIterator struct {
	Event *EndpointV2DefaultReceiveLibraryTimeoutSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2DefaultReceiveLibraryTimeoutSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2DefaultReceiveLibraryTimeoutSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2DefaultReceiveLibraryTimeoutSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2DefaultReceiveLibraryTimeoutSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2DefaultReceiveLibraryTimeoutSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2DefaultReceiveLibraryTimeoutSet represents a DefaultReceiveLibraryTimeoutSet event raised by the EndpointV2 contract.
type EndpointV2DefaultReceiveLibraryTimeoutSet struct {
	Eid    uint32
	OldLib common.Address
	Expiry *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultReceiveLibraryTimeoutSet is a free log retrieval operation binding the contract event 0x55b28633cdb29709386f555dfc54418592ad475ce7a65a78ac5928af60ffb8f8.
//
// Solidity: event DefaultReceiveLibraryTimeoutSet(uint32 eid, address oldLib, uint256 expiry)
func (_EndpointV2 *EndpointV2Filterer) FilterDefaultReceiveLibraryTimeoutSet(opts *bind.FilterOpts) (*EndpointV2DefaultReceiveLibraryTimeoutSetIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "DefaultReceiveLibraryTimeoutSet")
	if err != nil {
		return nil, err
	}
	return &EndpointV2DefaultReceiveLibraryTimeoutSetIterator{contract: _EndpointV2.contract, event: "DefaultReceiveLibraryTimeoutSet", logs: logs, sub: sub}, nil
}

// WatchDefaultReceiveLibraryTimeoutSet is a free log subscription operation binding the contract event 0x55b28633cdb29709386f555dfc54418592ad475ce7a65a78ac5928af60ffb8f8.
//
// Solidity: event DefaultReceiveLibraryTimeoutSet(uint32 eid, address oldLib, uint256 expiry)
func (_EndpointV2 *EndpointV2Filterer) WatchDefaultReceiveLibraryTimeoutSet(opts *bind.WatchOpts, sink chan<- *EndpointV2DefaultReceiveLibraryTimeoutSet) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "DefaultReceiveLibraryTimeoutSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2DefaultReceiveLibraryTimeoutSet)
				if err := _EndpointV2.contract.UnpackLog(event, "DefaultReceiveLibraryTimeoutSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultReceiveLibraryTimeoutSet is a log parse operation binding the contract event 0x55b28633cdb29709386f555dfc54418592ad475ce7a65a78ac5928af60ffb8f8.
//
// Solidity: event DefaultReceiveLibraryTimeoutSet(uint32 eid, address oldLib, uint256 expiry)
func (_EndpointV2 *EndpointV2Filterer) ParseDefaultReceiveLibraryTimeoutSet(log types.Log) (*EndpointV2DefaultReceiveLibraryTimeoutSet, error) {
	event := new(EndpointV2DefaultReceiveLibraryTimeoutSet)
	if err := _EndpointV2.contract.UnpackLog(event, "DefaultReceiveLibraryTimeoutSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2DefaultSendLibrarySetIterator is returned from FilterDefaultSendLibrarySet and is used to iterate over the raw logs and unpacked data for DefaultSendLibrarySet events raised by the EndpointV2 contract.
type EndpointV2DefaultSendLibrarySetIterator struct {
	Event *EndpointV2DefaultSendLibrarySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2DefaultSendLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2DefaultSendLibrarySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2DefaultSendLibrarySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2DefaultSendLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2DefaultSendLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2DefaultSendLibrarySet represents a DefaultSendLibrarySet event raised by the EndpointV2 contract.
type EndpointV2DefaultSendLibrarySet struct {
	Eid    uint32
	NewLib common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultSendLibrarySet is a free log retrieval operation binding the contract event 0x16aa0f528038ab41019e95bae5b418a50ba8532c5800e3b7ea2f517d3fa625f5.
//
// Solidity: event DefaultSendLibrarySet(uint32 eid, address newLib)
func (_EndpointV2 *EndpointV2Filterer) FilterDefaultSendLibrarySet(opts *bind.FilterOpts) (*EndpointV2DefaultSendLibrarySetIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "DefaultSendLibrarySet")
	if err != nil {
		return nil, err
	}
	return &EndpointV2DefaultSendLibrarySetIterator{contract: _EndpointV2.contract, event: "DefaultSendLibrarySet", logs: logs, sub: sub}, nil
}

// WatchDefaultSendLibrarySet is a free log subscription operation binding the contract event 0x16aa0f528038ab41019e95bae5b418a50ba8532c5800e3b7ea2f517d3fa625f5.
//
// Solidity: event DefaultSendLibrarySet(uint32 eid, address newLib)
func (_EndpointV2 *EndpointV2Filterer) WatchDefaultSendLibrarySet(opts *bind.WatchOpts, sink chan<- *EndpointV2DefaultSendLibrarySet) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "DefaultSendLibrarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2DefaultSendLibrarySet)
				if err := _EndpointV2.contract.UnpackLog(event, "DefaultSendLibrarySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultSendLibrarySet is a log parse operation binding the contract event 0x16aa0f528038ab41019e95bae5b418a50ba8532c5800e3b7ea2f517d3fa625f5.
//
// Solidity: event DefaultSendLibrarySet(uint32 eid, address newLib)
func (_EndpointV2 *EndpointV2Filterer) ParseDefaultSendLibrarySet(log types.Log) (*EndpointV2DefaultSendLibrarySet, error) {
	event := new(EndpointV2DefaultSendLibrarySet)
	if err := _EndpointV2.contract.UnpackLog(event, "DefaultSendLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2DelegateSetIterator is returned from FilterDelegateSet and is used to iterate over the raw logs and unpacked data for DelegateSet events raised by the EndpointV2 contract.
type EndpointV2DelegateSetIterator struct {
	Event *EndpointV2DelegateSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2DelegateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2DelegateSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2DelegateSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2DelegateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2DelegateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2DelegateSet represents a DelegateSet event raised by the EndpointV2 contract.
type EndpointV2DelegateSet struct {
	Sender   common.Address
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateSet is a free log retrieval operation binding the contract event 0x6ee10e9ed4d6ce9742703a498707862f4b00f1396a87195eb93267b3d7983981.
//
// Solidity: event DelegateSet(address sender, address delegate)
func (_EndpointV2 *EndpointV2Filterer) FilterDelegateSet(opts *bind.FilterOpts) (*EndpointV2DelegateSetIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "DelegateSet")
	if err != nil {
		return nil, err
	}
	return &EndpointV2DelegateSetIterator{contract: _EndpointV2.contract, event: "DelegateSet", logs: logs, sub: sub}, nil
}

// WatchDelegateSet is a free log subscription operation binding the contract event 0x6ee10e9ed4d6ce9742703a498707862f4b00f1396a87195eb93267b3d7983981.
//
// Solidity: event DelegateSet(address sender, address delegate)
func (_EndpointV2 *EndpointV2Filterer) WatchDelegateSet(opts *bind.WatchOpts, sink chan<- *EndpointV2DelegateSet) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "DelegateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2DelegateSet)
				if err := _EndpointV2.contract.UnpackLog(event, "DelegateSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateSet is a log parse operation binding the contract event 0x6ee10e9ed4d6ce9742703a498707862f4b00f1396a87195eb93267b3d7983981.
//
// Solidity: event DelegateSet(address sender, address delegate)
func (_EndpointV2 *EndpointV2Filterer) ParseDelegateSet(log types.Log) (*EndpointV2DelegateSet, error) {
	event := new(EndpointV2DelegateSet)
	if err := _EndpointV2.contract.UnpackLog(event, "DelegateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2InboundNonceSkippedIterator is returned from FilterInboundNonceSkipped and is used to iterate over the raw logs and unpacked data for InboundNonceSkipped events raised by the EndpointV2 contract.
type EndpointV2InboundNonceSkippedIterator struct {
	Event *EndpointV2InboundNonceSkipped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2InboundNonceSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2InboundNonceSkipped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2InboundNonceSkipped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2InboundNonceSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2InboundNonceSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2InboundNonceSkipped represents a InboundNonceSkipped event raised by the EndpointV2 contract.
type EndpointV2InboundNonceSkipped struct {
	SrcEid   uint32
	Sender   [32]byte
	Receiver common.Address
	Nonce    uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInboundNonceSkipped is a free log retrieval operation binding the contract event 0x28f40053783033ef755556a0c3315379141f51a33aed8334174ffbadd90bde48.
//
// Solidity: event InboundNonceSkipped(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce)
func (_EndpointV2 *EndpointV2Filterer) FilterInboundNonceSkipped(opts *bind.FilterOpts) (*EndpointV2InboundNonceSkippedIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "InboundNonceSkipped")
	if err != nil {
		return nil, err
	}
	return &EndpointV2InboundNonceSkippedIterator{contract: _EndpointV2.contract, event: "InboundNonceSkipped", logs: logs, sub: sub}, nil
}

// WatchInboundNonceSkipped is a free log subscription operation binding the contract event 0x28f40053783033ef755556a0c3315379141f51a33aed8334174ffbadd90bde48.
//
// Solidity: event InboundNonceSkipped(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce)
func (_EndpointV2 *EndpointV2Filterer) WatchInboundNonceSkipped(opts *bind.WatchOpts, sink chan<- *EndpointV2InboundNonceSkipped) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "InboundNonceSkipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2InboundNonceSkipped)
				if err := _EndpointV2.contract.UnpackLog(event, "InboundNonceSkipped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInboundNonceSkipped is a log parse operation binding the contract event 0x28f40053783033ef755556a0c3315379141f51a33aed8334174ffbadd90bde48.
//
// Solidity: event InboundNonceSkipped(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce)
func (_EndpointV2 *EndpointV2Filterer) ParseInboundNonceSkipped(log types.Log) (*EndpointV2InboundNonceSkipped, error) {
	event := new(EndpointV2InboundNonceSkipped)
	if err := _EndpointV2.contract.UnpackLog(event, "InboundNonceSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2LibraryRegisteredIterator is returned from FilterLibraryRegistered and is used to iterate over the raw logs and unpacked data for LibraryRegistered events raised by the EndpointV2 contract.
type EndpointV2LibraryRegisteredIterator struct {
	Event *EndpointV2LibraryRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2LibraryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2LibraryRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2LibraryRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2LibraryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2LibraryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2LibraryRegistered represents a LibraryRegistered event raised by the EndpointV2 contract.
type EndpointV2LibraryRegistered struct {
	NewLib common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLibraryRegistered is a free log retrieval operation binding the contract event 0x6b374d56679ca9463f27c85c6311e2bb7fde69bf201d3da39d53f10bd9d78af5.
//
// Solidity: event LibraryRegistered(address newLib)
func (_EndpointV2 *EndpointV2Filterer) FilterLibraryRegistered(opts *bind.FilterOpts) (*EndpointV2LibraryRegisteredIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "LibraryRegistered")
	if err != nil {
		return nil, err
	}
	return &EndpointV2LibraryRegisteredIterator{contract: _EndpointV2.contract, event: "LibraryRegistered", logs: logs, sub: sub}, nil
}

// WatchLibraryRegistered is a free log subscription operation binding the contract event 0x6b374d56679ca9463f27c85c6311e2bb7fde69bf201d3da39d53f10bd9d78af5.
//
// Solidity: event LibraryRegistered(address newLib)
func (_EndpointV2 *EndpointV2Filterer) WatchLibraryRegistered(opts *bind.WatchOpts, sink chan<- *EndpointV2LibraryRegistered) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "LibraryRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2LibraryRegistered)
				if err := _EndpointV2.contract.UnpackLog(event, "LibraryRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLibraryRegistered is a log parse operation binding the contract event 0x6b374d56679ca9463f27c85c6311e2bb7fde69bf201d3da39d53f10bd9d78af5.
//
// Solidity: event LibraryRegistered(address newLib)
func (_EndpointV2 *EndpointV2Filterer) ParseLibraryRegistered(log types.Log) (*EndpointV2LibraryRegistered, error) {
	event := new(EndpointV2LibraryRegistered)
	if err := _EndpointV2.contract.UnpackLog(event, "LibraryRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2LzComposeAlertIterator is returned from FilterLzComposeAlert and is used to iterate over the raw logs and unpacked data for LzComposeAlert events raised by the EndpointV2 contract.
type EndpointV2LzComposeAlertIterator struct {
	Event *EndpointV2LzComposeAlert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2LzComposeAlertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2LzComposeAlert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2LzComposeAlert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2LzComposeAlertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2LzComposeAlertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2LzComposeAlert represents a LzComposeAlert event raised by the EndpointV2 contract.
type EndpointV2LzComposeAlert struct {
	From      common.Address
	To        common.Address
	Executor  common.Address
	Guid      [32]byte
	Index     uint16
	Gas       *big.Int
	Value     *big.Int
	Message   []byte
	ExtraData []byte
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLzComposeAlert is a free log retrieval operation binding the contract event 0x8a0b1dce321c5c5fb42349bce46d18087c04140de520917661fb923e44a904b9.
//
// Solidity: event LzComposeAlert(address indexed from, address indexed to, address indexed executor, bytes32 guid, uint16 index, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_EndpointV2 *EndpointV2Filterer) FilterLzComposeAlert(opts *bind.FilterOpts, from []common.Address, to []common.Address, executor []common.Address) (*EndpointV2LzComposeAlertIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "LzComposeAlert", fromRule, toRule, executorRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2LzComposeAlertIterator{contract: _EndpointV2.contract, event: "LzComposeAlert", logs: logs, sub: sub}, nil
}

// WatchLzComposeAlert is a free log subscription operation binding the contract event 0x8a0b1dce321c5c5fb42349bce46d18087c04140de520917661fb923e44a904b9.
//
// Solidity: event LzComposeAlert(address indexed from, address indexed to, address indexed executor, bytes32 guid, uint16 index, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_EndpointV2 *EndpointV2Filterer) WatchLzComposeAlert(opts *bind.WatchOpts, sink chan<- *EndpointV2LzComposeAlert, from []common.Address, to []common.Address, executor []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "LzComposeAlert", fromRule, toRule, executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2LzComposeAlert)
				if err := _EndpointV2.contract.UnpackLog(event, "LzComposeAlert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLzComposeAlert is a log parse operation binding the contract event 0x8a0b1dce321c5c5fb42349bce46d18087c04140de520917661fb923e44a904b9.
//
// Solidity: event LzComposeAlert(address indexed from, address indexed to, address indexed executor, bytes32 guid, uint16 index, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_EndpointV2 *EndpointV2Filterer) ParseLzComposeAlert(log types.Log) (*EndpointV2LzComposeAlert, error) {
	event := new(EndpointV2LzComposeAlert)
	if err := _EndpointV2.contract.UnpackLog(event, "LzComposeAlert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2LzReceiveAlertIterator is returned from FilterLzReceiveAlert and is used to iterate over the raw logs and unpacked data for LzReceiveAlert events raised by the EndpointV2 contract.
type EndpointV2LzReceiveAlertIterator struct {
	Event *EndpointV2LzReceiveAlert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2LzReceiveAlertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2LzReceiveAlert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2LzReceiveAlert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2LzReceiveAlertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2LzReceiveAlertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2LzReceiveAlert represents a LzReceiveAlert event raised by the EndpointV2 contract.
type EndpointV2LzReceiveAlert struct {
	Receiver  common.Address
	Executor  common.Address
	Origin    Origin
	Guid      [32]byte
	Gas       *big.Int
	Value     *big.Int
	Message   []byte
	ExtraData []byte
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLzReceiveAlert is a free log retrieval operation binding the contract event 0x7edfa10fe10193301ad8a8bea7e968c7bcabcc64981f368e3aeada40ce26ae2c.
//
// Solidity: event LzReceiveAlert(address indexed receiver, address indexed executor, (uint32,bytes32,uint64) origin, bytes32 guid, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_EndpointV2 *EndpointV2Filterer) FilterLzReceiveAlert(opts *bind.FilterOpts, receiver []common.Address, executor []common.Address) (*EndpointV2LzReceiveAlertIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "LzReceiveAlert", receiverRule, executorRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2LzReceiveAlertIterator{contract: _EndpointV2.contract, event: "LzReceiveAlert", logs: logs, sub: sub}, nil
}

// WatchLzReceiveAlert is a free log subscription operation binding the contract event 0x7edfa10fe10193301ad8a8bea7e968c7bcabcc64981f368e3aeada40ce26ae2c.
//
// Solidity: event LzReceiveAlert(address indexed receiver, address indexed executor, (uint32,bytes32,uint64) origin, bytes32 guid, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_EndpointV2 *EndpointV2Filterer) WatchLzReceiveAlert(opts *bind.WatchOpts, sink chan<- *EndpointV2LzReceiveAlert, receiver []common.Address, executor []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "LzReceiveAlert", receiverRule, executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2LzReceiveAlert)
				if err := _EndpointV2.contract.UnpackLog(event, "LzReceiveAlert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLzReceiveAlert is a log parse operation binding the contract event 0x7edfa10fe10193301ad8a8bea7e968c7bcabcc64981f368e3aeada40ce26ae2c.
//
// Solidity: event LzReceiveAlert(address indexed receiver, address indexed executor, (uint32,bytes32,uint64) origin, bytes32 guid, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_EndpointV2 *EndpointV2Filterer) ParseLzReceiveAlert(log types.Log) (*EndpointV2LzReceiveAlert, error) {
	event := new(EndpointV2LzReceiveAlert)
	if err := _EndpointV2.contract.UnpackLog(event, "LzReceiveAlert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2LzTokenSetIterator is returned from FilterLzTokenSet and is used to iterate over the raw logs and unpacked data for LzTokenSet events raised by the EndpointV2 contract.
type EndpointV2LzTokenSetIterator struct {
	Event *EndpointV2LzTokenSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2LzTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2LzTokenSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2LzTokenSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2LzTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2LzTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2LzTokenSet represents a LzTokenSet event raised by the EndpointV2 contract.
type EndpointV2LzTokenSet struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLzTokenSet is a free log retrieval operation binding the contract event 0xd476ec5ec1ac11cec3714d41e7ea49419471aceb9bd0dff1becfc3e363a62396.
//
// Solidity: event LzTokenSet(address token)
func (_EndpointV2 *EndpointV2Filterer) FilterLzTokenSet(opts *bind.FilterOpts) (*EndpointV2LzTokenSetIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "LzTokenSet")
	if err != nil {
		return nil, err
	}
	return &EndpointV2LzTokenSetIterator{contract: _EndpointV2.contract, event: "LzTokenSet", logs: logs, sub: sub}, nil
}

// WatchLzTokenSet is a free log subscription operation binding the contract event 0xd476ec5ec1ac11cec3714d41e7ea49419471aceb9bd0dff1becfc3e363a62396.
//
// Solidity: event LzTokenSet(address token)
func (_EndpointV2 *EndpointV2Filterer) WatchLzTokenSet(opts *bind.WatchOpts, sink chan<- *EndpointV2LzTokenSet) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "LzTokenSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2LzTokenSet)
				if err := _EndpointV2.contract.UnpackLog(event, "LzTokenSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLzTokenSet is a log parse operation binding the contract event 0xd476ec5ec1ac11cec3714d41e7ea49419471aceb9bd0dff1becfc3e363a62396.
//
// Solidity: event LzTokenSet(address token)
func (_EndpointV2 *EndpointV2Filterer) ParseLzTokenSet(log types.Log) (*EndpointV2LzTokenSet, error) {
	event := new(EndpointV2LzTokenSet)
	if err := _EndpointV2.contract.UnpackLog(event, "LzTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EndpointV2 contract.
type EndpointV2OwnershipTransferredIterator struct {
	Event *EndpointV2OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2OwnershipTransferred represents a OwnershipTransferred event raised by the EndpointV2 contract.
type EndpointV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EndpointV2 *EndpointV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EndpointV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2OwnershipTransferredIterator{contract: _EndpointV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EndpointV2 *EndpointV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EndpointV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2OwnershipTransferred)
				if err := _EndpointV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EndpointV2 *EndpointV2Filterer) ParseOwnershipTransferred(log types.Log) (*EndpointV2OwnershipTransferred, error) {
	event := new(EndpointV2OwnershipTransferred)
	if err := _EndpointV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2PacketBurntIterator is returned from FilterPacketBurnt and is used to iterate over the raw logs and unpacked data for PacketBurnt events raised by the EndpointV2 contract.
type EndpointV2PacketBurntIterator struct {
	Event *EndpointV2PacketBurnt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2PacketBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2PacketBurnt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2PacketBurnt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2PacketBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2PacketBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2PacketBurnt represents a PacketBurnt event raised by the EndpointV2 contract.
type EndpointV2PacketBurnt struct {
	SrcEid      uint32
	Sender      [32]byte
	Receiver    common.Address
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPacketBurnt is a free log retrieval operation binding the contract event 0x7f68a37a6e69a0de35024a234558f9efe4b33b58657753d21eaaa82d51c3510e.
//
// Solidity: event PacketBurnt(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Filterer) FilterPacketBurnt(opts *bind.FilterOpts) (*EndpointV2PacketBurntIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "PacketBurnt")
	if err != nil {
		return nil, err
	}
	return &EndpointV2PacketBurntIterator{contract: _EndpointV2.contract, event: "PacketBurnt", logs: logs, sub: sub}, nil
}

// WatchPacketBurnt is a free log subscription operation binding the contract event 0x7f68a37a6e69a0de35024a234558f9efe4b33b58657753d21eaaa82d51c3510e.
//
// Solidity: event PacketBurnt(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Filterer) WatchPacketBurnt(opts *bind.WatchOpts, sink chan<- *EndpointV2PacketBurnt) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "PacketBurnt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2PacketBurnt)
				if err := _EndpointV2.contract.UnpackLog(event, "PacketBurnt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketBurnt is a log parse operation binding the contract event 0x7f68a37a6e69a0de35024a234558f9efe4b33b58657753d21eaaa82d51c3510e.
//
// Solidity: event PacketBurnt(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Filterer) ParsePacketBurnt(log types.Log) (*EndpointV2PacketBurnt, error) {
	event := new(EndpointV2PacketBurnt)
	if err := _EndpointV2.contract.UnpackLog(event, "PacketBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2PacketDeliveredIterator is returned from FilterPacketDelivered and is used to iterate over the raw logs and unpacked data for PacketDelivered events raised by the EndpointV2 contract.
type EndpointV2PacketDeliveredIterator struct {
	Event *EndpointV2PacketDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2PacketDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2PacketDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2PacketDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2PacketDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2PacketDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2PacketDelivered represents a PacketDelivered event raised by the EndpointV2 contract.
type EndpointV2PacketDelivered struct {
	Origin   Origin
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPacketDelivered is a free log retrieval operation binding the contract event 0x3cd5e48f9730b129dc7550f0fcea9c767b7be37837cd10e55eb35f734f4bca04.
//
// Solidity: event PacketDelivered((uint32,bytes32,uint64) origin, address receiver)
func (_EndpointV2 *EndpointV2Filterer) FilterPacketDelivered(opts *bind.FilterOpts) (*EndpointV2PacketDeliveredIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "PacketDelivered")
	if err != nil {
		return nil, err
	}
	return &EndpointV2PacketDeliveredIterator{contract: _EndpointV2.contract, event: "PacketDelivered", logs: logs, sub: sub}, nil
}

// WatchPacketDelivered is a free log subscription operation binding the contract event 0x3cd5e48f9730b129dc7550f0fcea9c767b7be37837cd10e55eb35f734f4bca04.
//
// Solidity: event PacketDelivered((uint32,bytes32,uint64) origin, address receiver)
func (_EndpointV2 *EndpointV2Filterer) WatchPacketDelivered(opts *bind.WatchOpts, sink chan<- *EndpointV2PacketDelivered) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "PacketDelivered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2PacketDelivered)
				if err := _EndpointV2.contract.UnpackLog(event, "PacketDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketDelivered is a log parse operation binding the contract event 0x3cd5e48f9730b129dc7550f0fcea9c767b7be37837cd10e55eb35f734f4bca04.
//
// Solidity: event PacketDelivered((uint32,bytes32,uint64) origin, address receiver)
func (_EndpointV2 *EndpointV2Filterer) ParsePacketDelivered(log types.Log) (*EndpointV2PacketDelivered, error) {
	event := new(EndpointV2PacketDelivered)
	if err := _EndpointV2.contract.UnpackLog(event, "PacketDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2PacketNilifiedIterator is returned from FilterPacketNilified and is used to iterate over the raw logs and unpacked data for PacketNilified events raised by the EndpointV2 contract.
type EndpointV2PacketNilifiedIterator struct {
	Event *EndpointV2PacketNilified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2PacketNilifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2PacketNilified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2PacketNilified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2PacketNilifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2PacketNilifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2PacketNilified represents a PacketNilified event raised by the EndpointV2 contract.
type EndpointV2PacketNilified struct {
	SrcEid      uint32
	Sender      [32]byte
	Receiver    common.Address
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPacketNilified is a free log retrieval operation binding the contract event 0xaf0450c392c4f702515a457a362328c8aa21916048ca6d0419e248b30cb55292.
//
// Solidity: event PacketNilified(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Filterer) FilterPacketNilified(opts *bind.FilterOpts) (*EndpointV2PacketNilifiedIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "PacketNilified")
	if err != nil {
		return nil, err
	}
	return &EndpointV2PacketNilifiedIterator{contract: _EndpointV2.contract, event: "PacketNilified", logs: logs, sub: sub}, nil
}

// WatchPacketNilified is a free log subscription operation binding the contract event 0xaf0450c392c4f702515a457a362328c8aa21916048ca6d0419e248b30cb55292.
//
// Solidity: event PacketNilified(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Filterer) WatchPacketNilified(opts *bind.WatchOpts, sink chan<- *EndpointV2PacketNilified) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "PacketNilified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2PacketNilified)
				if err := _EndpointV2.contract.UnpackLog(event, "PacketNilified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketNilified is a log parse operation binding the contract event 0xaf0450c392c4f702515a457a362328c8aa21916048ca6d0419e248b30cb55292.
//
// Solidity: event PacketNilified(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Filterer) ParsePacketNilified(log types.Log) (*EndpointV2PacketNilified, error) {
	event := new(EndpointV2PacketNilified)
	if err := _EndpointV2.contract.UnpackLog(event, "PacketNilified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2PacketSentIterator is returned from FilterPacketSent and is used to iterate over the raw logs and unpacked data for PacketSent events raised by the EndpointV2 contract.
type EndpointV2PacketSentIterator struct {
	Event *EndpointV2PacketSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2PacketSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2PacketSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2PacketSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2PacketSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2PacketSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2PacketSent represents a PacketSent event raised by the EndpointV2 contract.
type EndpointV2PacketSent struct {
	EncodedPayload []byte
	Options        []byte
	SendLibrary    common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPacketSent is a free log retrieval operation binding the contract event 0x1ab700d4ced0c005b164c0f789fd09fcbb0156d4c2041b8a3bfbcd961cd1567f.
//
// Solidity: event PacketSent(bytes encodedPayload, bytes options, address sendLibrary)
func (_EndpointV2 *EndpointV2Filterer) FilterPacketSent(opts *bind.FilterOpts) (*EndpointV2PacketSentIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return &EndpointV2PacketSentIterator{contract: _EndpointV2.contract, event: "PacketSent", logs: logs, sub: sub}, nil
}

// WatchPacketSent is a free log subscription operation binding the contract event 0x1ab700d4ced0c005b164c0f789fd09fcbb0156d4c2041b8a3bfbcd961cd1567f.
//
// Solidity: event PacketSent(bytes encodedPayload, bytes options, address sendLibrary)
func (_EndpointV2 *EndpointV2Filterer) WatchPacketSent(opts *bind.WatchOpts, sink chan<- *EndpointV2PacketSent) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2PacketSent)
				if err := _EndpointV2.contract.UnpackLog(event, "PacketSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketSent is a log parse operation binding the contract event 0x1ab700d4ced0c005b164c0f789fd09fcbb0156d4c2041b8a3bfbcd961cd1567f.
//
// Solidity: event PacketSent(bytes encodedPayload, bytes options, address sendLibrary)
func (_EndpointV2 *EndpointV2Filterer) ParsePacketSent(log types.Log) (*EndpointV2PacketSent, error) {
	event := new(EndpointV2PacketSent)
	if err := _EndpointV2.contract.UnpackLog(event, "PacketSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2PacketVerifiedIterator is returned from FilterPacketVerified and is used to iterate over the raw logs and unpacked data for PacketVerified events raised by the EndpointV2 contract.
type EndpointV2PacketVerifiedIterator struct {
	Event *EndpointV2PacketVerified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2PacketVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2PacketVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2PacketVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2PacketVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2PacketVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2PacketVerified represents a PacketVerified event raised by the EndpointV2 contract.
type EndpointV2PacketVerified struct {
	Origin      Origin
	Receiver    common.Address
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPacketVerified is a free log retrieval operation binding the contract event 0x0d87345f3d1c929caba93e1c3821b54ff3512e12b66aa3cfe54b6bcbc17e59b4.
//
// Solidity: event PacketVerified((uint32,bytes32,uint64) origin, address receiver, bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Filterer) FilterPacketVerified(opts *bind.FilterOpts) (*EndpointV2PacketVerifiedIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "PacketVerified")
	if err != nil {
		return nil, err
	}
	return &EndpointV2PacketVerifiedIterator{contract: _EndpointV2.contract, event: "PacketVerified", logs: logs, sub: sub}, nil
}

// WatchPacketVerified is a free log subscription operation binding the contract event 0x0d87345f3d1c929caba93e1c3821b54ff3512e12b66aa3cfe54b6bcbc17e59b4.
//
// Solidity: event PacketVerified((uint32,bytes32,uint64) origin, address receiver, bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Filterer) WatchPacketVerified(opts *bind.WatchOpts, sink chan<- *EndpointV2PacketVerified) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "PacketVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2PacketVerified)
				if err := _EndpointV2.contract.UnpackLog(event, "PacketVerified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketVerified is a log parse operation binding the contract event 0x0d87345f3d1c929caba93e1c3821b54ff3512e12b66aa3cfe54b6bcbc17e59b4.
//
// Solidity: event PacketVerified((uint32,bytes32,uint64) origin, address receiver, bytes32 payloadHash)
func (_EndpointV2 *EndpointV2Filterer) ParsePacketVerified(log types.Log) (*EndpointV2PacketVerified, error) {
	event := new(EndpointV2PacketVerified)
	if err := _EndpointV2.contract.UnpackLog(event, "PacketVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2ReceiveLibrarySetIterator is returned from FilterReceiveLibrarySet and is used to iterate over the raw logs and unpacked data for ReceiveLibrarySet events raised by the EndpointV2 contract.
type EndpointV2ReceiveLibrarySetIterator struct {
	Event *EndpointV2ReceiveLibrarySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2ReceiveLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2ReceiveLibrarySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2ReceiveLibrarySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2ReceiveLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2ReceiveLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2ReceiveLibrarySet represents a ReceiveLibrarySet event raised by the EndpointV2 contract.
type EndpointV2ReceiveLibrarySet struct {
	Receiver common.Address
	Eid      uint32
	NewLib   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceiveLibrarySet is a free log retrieval operation binding the contract event 0xcd6f92f5ac6185a5acfa02c92090746cec64d777269cbcd0ed031e396657a1c2.
//
// Solidity: event ReceiveLibrarySet(address receiver, uint32 eid, address newLib)
func (_EndpointV2 *EndpointV2Filterer) FilterReceiveLibrarySet(opts *bind.FilterOpts) (*EndpointV2ReceiveLibrarySetIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "ReceiveLibrarySet")
	if err != nil {
		return nil, err
	}
	return &EndpointV2ReceiveLibrarySetIterator{contract: _EndpointV2.contract, event: "ReceiveLibrarySet", logs: logs, sub: sub}, nil
}

// WatchReceiveLibrarySet is a free log subscription operation binding the contract event 0xcd6f92f5ac6185a5acfa02c92090746cec64d777269cbcd0ed031e396657a1c2.
//
// Solidity: event ReceiveLibrarySet(address receiver, uint32 eid, address newLib)
func (_EndpointV2 *EndpointV2Filterer) WatchReceiveLibrarySet(opts *bind.WatchOpts, sink chan<- *EndpointV2ReceiveLibrarySet) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "ReceiveLibrarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2ReceiveLibrarySet)
				if err := _EndpointV2.contract.UnpackLog(event, "ReceiveLibrarySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceiveLibrarySet is a log parse operation binding the contract event 0xcd6f92f5ac6185a5acfa02c92090746cec64d777269cbcd0ed031e396657a1c2.
//
// Solidity: event ReceiveLibrarySet(address receiver, uint32 eid, address newLib)
func (_EndpointV2 *EndpointV2Filterer) ParseReceiveLibrarySet(log types.Log) (*EndpointV2ReceiveLibrarySet, error) {
	event := new(EndpointV2ReceiveLibrarySet)
	if err := _EndpointV2.contract.UnpackLog(event, "ReceiveLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2ReceiveLibraryTimeoutSetIterator is returned from FilterReceiveLibraryTimeoutSet and is used to iterate over the raw logs and unpacked data for ReceiveLibraryTimeoutSet events raised by the EndpointV2 contract.
type EndpointV2ReceiveLibraryTimeoutSetIterator struct {
	Event *EndpointV2ReceiveLibraryTimeoutSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2ReceiveLibraryTimeoutSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2ReceiveLibraryTimeoutSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2ReceiveLibraryTimeoutSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2ReceiveLibraryTimeoutSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2ReceiveLibraryTimeoutSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2ReceiveLibraryTimeoutSet represents a ReceiveLibraryTimeoutSet event raised by the EndpointV2 contract.
type EndpointV2ReceiveLibraryTimeoutSet struct {
	Receiver common.Address
	Eid      uint32
	OldLib   common.Address
	Timeout  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceiveLibraryTimeoutSet is a free log retrieval operation binding the contract event 0x4e0a5bbfa0c11a64effb1ada324b5437a17272e1aed9320398715ef71bb20928.
//
// Solidity: event ReceiveLibraryTimeoutSet(address receiver, uint32 eid, address oldLib, uint256 timeout)
func (_EndpointV2 *EndpointV2Filterer) FilterReceiveLibraryTimeoutSet(opts *bind.FilterOpts) (*EndpointV2ReceiveLibraryTimeoutSetIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "ReceiveLibraryTimeoutSet")
	if err != nil {
		return nil, err
	}
	return &EndpointV2ReceiveLibraryTimeoutSetIterator{contract: _EndpointV2.contract, event: "ReceiveLibraryTimeoutSet", logs: logs, sub: sub}, nil
}

// WatchReceiveLibraryTimeoutSet is a free log subscription operation binding the contract event 0x4e0a5bbfa0c11a64effb1ada324b5437a17272e1aed9320398715ef71bb20928.
//
// Solidity: event ReceiveLibraryTimeoutSet(address receiver, uint32 eid, address oldLib, uint256 timeout)
func (_EndpointV2 *EndpointV2Filterer) WatchReceiveLibraryTimeoutSet(opts *bind.WatchOpts, sink chan<- *EndpointV2ReceiveLibraryTimeoutSet) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "ReceiveLibraryTimeoutSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2ReceiveLibraryTimeoutSet)
				if err := _EndpointV2.contract.UnpackLog(event, "ReceiveLibraryTimeoutSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceiveLibraryTimeoutSet is a log parse operation binding the contract event 0x4e0a5bbfa0c11a64effb1ada324b5437a17272e1aed9320398715ef71bb20928.
//
// Solidity: event ReceiveLibraryTimeoutSet(address receiver, uint32 eid, address oldLib, uint256 timeout)
func (_EndpointV2 *EndpointV2Filterer) ParseReceiveLibraryTimeoutSet(log types.Log) (*EndpointV2ReceiveLibraryTimeoutSet, error) {
	event := new(EndpointV2ReceiveLibraryTimeoutSet)
	if err := _EndpointV2.contract.UnpackLog(event, "ReceiveLibraryTimeoutSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2SendLibrarySetIterator is returned from FilterSendLibrarySet and is used to iterate over the raw logs and unpacked data for SendLibrarySet events raised by the EndpointV2 contract.
type EndpointV2SendLibrarySetIterator struct {
	Event *EndpointV2SendLibrarySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EndpointV2SendLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2SendLibrarySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EndpointV2SendLibrarySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EndpointV2SendLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2SendLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2SendLibrarySet represents a SendLibrarySet event raised by the EndpointV2 contract.
type EndpointV2SendLibrarySet struct {
	Sender common.Address
	Eid    uint32
	NewLib common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSendLibrarySet is a free log retrieval operation binding the contract event 0x4cff966ebee29a156dcb34cf72c1d06231fb1777f6bdf6e8089819232f002b1c.
//
// Solidity: event SendLibrarySet(address sender, uint32 eid, address newLib)
func (_EndpointV2 *EndpointV2Filterer) FilterSendLibrarySet(opts *bind.FilterOpts) (*EndpointV2SendLibrarySetIterator, error) {

	logs, sub, err := _EndpointV2.contract.FilterLogs(opts, "SendLibrarySet")
	if err != nil {
		return nil, err
	}
	return &EndpointV2SendLibrarySetIterator{contract: _EndpointV2.contract, event: "SendLibrarySet", logs: logs, sub: sub}, nil
}

// WatchSendLibrarySet is a free log subscription operation binding the contract event 0x4cff966ebee29a156dcb34cf72c1d06231fb1777f6bdf6e8089819232f002b1c.
//
// Solidity: event SendLibrarySet(address sender, uint32 eid, address newLib)
func (_EndpointV2 *EndpointV2Filterer) WatchSendLibrarySet(opts *bind.WatchOpts, sink chan<- *EndpointV2SendLibrarySet) (event.Subscription, error) {

	logs, sub, err := _EndpointV2.contract.WatchLogs(opts, "SendLibrarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2SendLibrarySet)
				if err := _EndpointV2.contract.UnpackLog(event, "SendLibrarySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSendLibrarySet is a log parse operation binding the contract event 0x4cff966ebee29a156dcb34cf72c1d06231fb1777f6bdf6e8089819232f002b1c.
//
// Solidity: event SendLibrarySet(address sender, uint32 eid, address newLib)
func (_EndpointV2 *EndpointV2Filterer) ParseSendLibrarySet(log types.Log) (*EndpointV2SendLibrarySet, error) {
	event := new(EndpointV2SendLibrarySet)
	if err := _EndpointV2.contract.UnpackLog(event, "SendLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
