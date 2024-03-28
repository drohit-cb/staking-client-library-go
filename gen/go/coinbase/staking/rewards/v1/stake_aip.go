// Code generated by protoc-gen-go-aip. DO NOT EDIT.
//
// versions:
// 	protoc-gen-go-aip development
// 	protoc (unknown)
// source: coinbase/staking/rewards/v1/stake.proto

package v1

import (
	fmt "fmt"
	resourcename "go.einride.tech/aip/resourcename"
	strings "strings"
)

type StakeResourceName struct {
	Protocol string
	Stake    string
}

func (n ProtocolResourceName) StakeResourceName(
	stake string,
) StakeResourceName {
	return StakeResourceName{
		Protocol: n.Protocol,
		Stake:    stake,
	}
}

func (n StakeResourceName) Validate() error {
	if n.Protocol == "" {
		return fmt.Errorf("protocol: empty")
	}
	if strings.IndexByte(n.Protocol, '/') != -1 {
		return fmt.Errorf("protocol: contains illegal character '/'")
	}
	if n.Stake == "" {
		return fmt.Errorf("stake: empty")
	}
	if strings.IndexByte(n.Stake, '/') != -1 {
		return fmt.Errorf("stake: contains illegal character '/'")
	}
	return nil
}

func (n StakeResourceName) ContainsWildcard() bool {
	return false || n.Protocol == "-" || n.Stake == "-"
}

func (n StakeResourceName) String() string {
	return resourcename.Sprint(
		"protocols/{protocol}/stakes/{stake}",
		n.Protocol,
		n.Stake,
	)
}

func (n StakeResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *StakeResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"protocols/{protocol}/stakes/{stake}",
		&n.Protocol,
		&n.Stake,
	)
}

func (n StakeResourceName) ProtocolResourceName() ProtocolResourceName {
	return ProtocolResourceName{
		Protocol: n.Protocol,
	}
}
