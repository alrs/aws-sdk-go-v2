// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

type AcceleratorStatus string

// Enum values for AcceleratorStatus
const (
	AcceleratorStatusDeployed   AcceleratorStatus = "DEPLOYED"
	AcceleratorStatusInProgress AcceleratorStatus = "IN_PROGRESS"
)

func (enum AcceleratorStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AcceleratorStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Affinity string

// Enum values for Affinity
const (
	AffinityNone     Affinity = "NONE"
	AffinitySourceIp Affinity = "SOURCE_IP"
)

func (enum Affinity) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Affinity) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ByoipCidrState string

// Enum values for ByoipCidrState
const (
	ByoipCidrStatePendingProvisioning   ByoipCidrState = "PENDING_PROVISIONING"
	ByoipCidrStateReady                 ByoipCidrState = "READY"
	ByoipCidrStatePendingAdvertising    ByoipCidrState = "PENDING_ADVERTISING"
	ByoipCidrStateAdvertising           ByoipCidrState = "ADVERTISING"
	ByoipCidrStatePendingWithdrawing    ByoipCidrState = "PENDING_WITHDRAWING"
	ByoipCidrStatePendingDeprovisioning ByoipCidrState = "PENDING_DEPROVISIONING"
	ByoipCidrStateDeprovisioned         ByoipCidrState = "DEPROVISIONED"
	ByoipCidrStateFailedProvision       ByoipCidrState = "FAILED_PROVISION"
	ByoipCidrStateFailedAdvertising     ByoipCidrState = "FAILED_ADVERTISING"
	ByoipCidrStateFailedWithdraw        ByoipCidrState = "FAILED_WITHDRAW"
	ByoipCidrStateFailedDeprovision     ByoipCidrState = "FAILED_DEPROVISION"
)

func (enum ByoipCidrState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ByoipCidrState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HealthCheckProtocol string

// Enum values for HealthCheckProtocol
const (
	HealthCheckProtocolTcp   HealthCheckProtocol = "TCP"
	HealthCheckProtocolHttp  HealthCheckProtocol = "HTTP"
	HealthCheckProtocolHttps HealthCheckProtocol = "HTTPS"
)

func (enum HealthCheckProtocol) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HealthCheckProtocol) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HealthState string

// Enum values for HealthState
const (
	HealthStateInitial   HealthState = "INITIAL"
	HealthStateHealthy   HealthState = "HEALTHY"
	HealthStateUnhealthy HealthState = "UNHEALTHY"
)

func (enum HealthState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HealthState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type IpAddressType string

// Enum values for IpAddressType
const (
	IpAddressTypeIpv4 IpAddressType = "IPV4"
)

func (enum IpAddressType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum IpAddressType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Protocol string

// Enum values for Protocol
const (
	ProtocolTcp Protocol = "TCP"
	ProtocolUdp Protocol = "UDP"
)

func (enum Protocol) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Protocol) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
