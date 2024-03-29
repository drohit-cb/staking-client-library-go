// Code generated by protoc-gen-go-aip. DO NOT EDIT.
//
// versions:
// 	protoc-gen-go-aip development
// 	protoc (unknown)
// source: coinbase/staking/v1alpha1/action.proto

package v1alpha1

import (
	fmt "fmt"
	resourcename "go.einride.tech/aip/resourcename"
	strings "strings"
)

type ActionResourceName struct {
	Protocol string
	Network  string
	Action   string
}

func (n ProtocolResourceName) ActionResourceName(
	network string,
	action string,
) ActionResourceName {
	return ActionResourceName{
		Protocol: n.Protocol,
		Network:  network,
		Action:   action,
	}
}

func (n NetworkResourceName) ActionResourceName(
	action string,
) ActionResourceName {
	return ActionResourceName{
		Protocol: n.Protocol,
		Network:  n.Network,
		Action:   action,
	}
}

func (n ActionResourceName) Validate() error {
	if n.Protocol == "" {
		return fmt.Errorf("protocol: empty")
	}
	if strings.IndexByte(n.Protocol, '/') != -1 {
		return fmt.Errorf("protocol: contains illegal character '/'")
	}
	if n.Network == "" {
		return fmt.Errorf("network: empty")
	}
	if strings.IndexByte(n.Network, '/') != -1 {
		return fmt.Errorf("network: contains illegal character '/'")
	}
	if n.Action == "" {
		return fmt.Errorf("action: empty")
	}
	if strings.IndexByte(n.Action, '/') != -1 {
		return fmt.Errorf("action: contains illegal character '/'")
	}
	return nil
}

func (n ActionResourceName) ContainsWildcard() bool {
	return false || n.Protocol == "-" || n.Network == "-" || n.Action == "-"
}

func (n ActionResourceName) String() string {
	return resourcename.Sprint(
		"protocols/{protocol}/networks/{network}/actions/{action}",
		n.Protocol,
		n.Network,
		n.Action,
	)
}

func (n ActionResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *ActionResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"protocols/{protocol}/networks/{network}/actions/{action}",
		&n.Protocol,
		&n.Network,
		&n.Action,
	)
}

func (n ActionResourceName) ProtocolResourceName() ProtocolResourceName {
	return ProtocolResourceName{
		Protocol: n.Protocol,
	}
}

func (n ActionResourceName) NetworkResourceName() NetworkResourceName {
	return NetworkResourceName{
		Protocol: n.Protocol,
		Network:  n.Network,
	}
}