//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>
package utils

import (
	"fmt"
)

type ServiceError struct {
	Status         int    `json:"-"`
	Code           string `json:"code,omitempty"`
	Message        string `json:"message,omitempty"`
	Details        string `json:"details,omitempty"`
	TraceId        string `json:"traceId,omitempty"`
	FlightRecorder string `json:"flightRecorder,omitempty"`
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
