package types

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	proto "github.com/gogo/protobuf/proto"
)

var _ codectypes.UnpackInterfacesMessage = &Grant{}

// Grantee represents a generic grant grantee
type Grantee interface {
	proto.Message

	isGrantee()
	Validate() error
}

// NewUserGrantee is a constructor for the UserGrantee type
func NewUserGrantee(user string) *UserGrantee {
	return &UserGrantee{
		User: user,
	}
}

// isGrantee implements Grantee
func (t *UserGrantee) isGrantee() {}

// isGrantee implements Grantee
func (t *UserGrantee) Validate() error {
	_, err := sdk.AccAddressFromBech32(t.User)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid grantee address")
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// NewGroupGrantee is a constructor for the GroupGrantee type
func NewGroupGrantee(groupID uint32) *GroupGrantee {
	return &GroupGrantee{
		GroupID: groupID,
	}
}

// isGrantee implements Grantee
func (t *GroupGrantee) isGrantee() {}

// isGrantee implements Grantee
func (t *GroupGrantee) Validate() error {
	if t.GroupID == 0 {
		return fmt.Errorf("invalid group id: %d", t.GroupID)
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// NewGrant is a constructor for the Grant type
func NewGrant(subspaceID uint64, granter string, grantee Grantee, feeAllowance feegranttypes.FeeAllowanceI) Grant {
	allowanceProto, isProto := feeAllowance.(proto.Message)
	if !isProto {
		panic(sdkerrors.Wrapf(sdkerrors.ErrPackAny, "cannot proto marshal %T", feeAllowance))
	}

	allowanceAny, err := codectypes.NewAnyWithValue(allowanceProto)
	if err != nil {
		panic("failed to pack allowance to any type")
	}

	granteeAny, err := codectypes.NewAnyWithValue(grantee)
	if err != nil {
		panic("failed to pack grantee to any type")
	}

	return Grant{
		SubspaceID: subspaceID,
		Granter:    granter,
		Grantee:    granteeAny,
		Allowance:  allowanceAny,
	}
}

// Validate implements fmt.Validator
func (g Grant) Validate() error {
	if g.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", g.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(g.Granter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid granter address")
	}

	grantee, isGrantee := g.Grantee.GetCachedValue().(Grantee)
	if !isGrantee {
		return fmt.Errorf("invalid grantee type: %T", grantee)
	}

	err = grantee.Validate()
	if err != nil {
		return err
	}

	if u, isUserGrantee := grantee.(*UserGrantee); isUserGrantee {
		if u.User == g.Granter {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "cannot self-grant fee authorization")
		}
	}

	f, err := g.GetUnpackedAllowance()
	if err != nil {
		return err
	}

	return f.ValidateBasic()
}

// GetUnpackedAllowance returns unpacked allowance
func (u Grant) GetUnpackedAllowance() (feegranttypes.FeeAllowanceI, error) {
	allowance, isAllowance := u.Allowance.GetCachedValue().(feegranttypes.FeeAllowanceI)
	if !isAllowance {
		return nil, fmt.Errorf("failed to unpack allowance")
	}

	return allowance, nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (u Grant) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var grantee Grantee
	err := unpacker.UnpackAny(u.Grantee, &grantee)
	if err != nil {
		return err
	}

	var allowance feegranttypes.FeeAllowanceI
	return unpacker.UnpackAny(u.Allowance, &allowance)
}
