//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>
package sites

type Sites struct {
	// Validate trough custom regular expression "resourceId"
	ProviderId      string `json:"providerId" validate:"resourceId"`
	CloudRegionCode string `json:"cloudRegionCode"  validate:"email"`
	Type            string `json:"type" validate:"required"`
}
