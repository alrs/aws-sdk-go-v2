// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

type BulkPublishStatus string

// Enum values for BulkPublishStatus
const (
	BulkPublishStatusNotStarted BulkPublishStatus = "NOT_STARTED"
	BulkPublishStatusInProgress BulkPublishStatus = "IN_PROGRESS"
	BulkPublishStatusFailed     BulkPublishStatus = "FAILED"
	BulkPublishStatusSucceeded  BulkPublishStatus = "SUCCEEDED"
)

func (enum BulkPublishStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BulkPublishStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Operation string

// Enum values for Operation
const (
	OperationReplace Operation = "replace"
	OperationRemove  Operation = "remove"
)

func (enum Operation) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Operation) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Platform string

// Enum values for Platform
const (
	PlatformApns        Platform = "APNS"
	PlatformApnsSandbox Platform = "APNS_SANDBOX"
	PlatformGcm         Platform = "GCM"
	PlatformAdm         Platform = "ADM"
)

func (enum Platform) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Platform) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StreamingStatus string

// Enum values for StreamingStatus
const (
	StreamingStatusEnabled  StreamingStatus = "ENABLED"
	StreamingStatusDisabled StreamingStatus = "DISABLED"
)

func (enum StreamingStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StreamingStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
