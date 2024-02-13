//SPDX-FileCopyrightText: © 2023 3nets, Inc. <it@3nets.io>
package templates

type Templates struct {
	Name        string `json:"name" validate:"name"`
	Description string `json:"description"  validate:"description"`
}
