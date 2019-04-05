/*
 * This file is part of the IoT Agent
 * Copyright 2019 Canonical Ltd.
 *
 * This program is free software: you can redistribute it and/or modify it
 * under the terms of the GNU General Public License version 3, as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranties of MERCHANTABILITY,
 * SATISFACTORY QUALITY, or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package mqtt

import (
	"encoding/json"
	"github.com/CanonicalLtd/iot-agent/snapdapi"
)

// SubscribeAct is the interface for actions
type SubscribeAct interface {
	SnapInstall()
}

// SubscribeAction is the message format for the action topic
type SubscribeAction struct {
	ID     string `json:"id"`
	Action string `json:"action"`
	Snap   string `json:"snap"`
	Data   string `json:"data"`
}

var snapd snapdapi.SnapdClient = snapdapi.NewClientAdapter()

// SnapInstall installs a new snap
func (act *SubscribeAction) SnapInstall() PublishResponse {
	if len(act.Snap) == 0 {
		return PublishResponse{ID: act.ID, Success: false, Message: "No snap name provided for install"}
	}

	// Call the snapd API
	result, err := snapd.Install(act.Snap, nil)
	if err != nil {
		return PublishResponse{ID: act.ID, Success: false, Message: err.Error()}
	}
	return PublishResponse{ID: act.ID, Success: true, Result: result}
}

// SnapRemove removes an existing snap
func (act *SubscribeAction) SnapRemove() PublishResponse {
	if len(act.Snap) == 0 {
		return PublishResponse{ID: act.ID, Success: false, Message: "No snap name provided for remove"}
	}

	// Call the snapd API
	result, err := snapd.Remove(act.Snap, nil)
	if err != nil {
		return PublishResponse{ID: act.ID, Success: false, Message: err.Error()}
	}
	return PublishResponse{ID: act.ID, Success: true, Result: result}
}

// SnapList lists installed snaps
func (act *SubscribeAction) SnapList() PublishResponse {
	// Call the snapd API
	snaps, err := snapd.List([]string{}, nil)
	if err != nil {
		return PublishResponse{ID: act.ID, Success: false, Message: err.Error()}
	}
	return PublishResponse{ID: act.ID, Success: true, Result: snaps}
}

// SnapRefresh refreshes an existing snap
func (act *SubscribeAction) SnapRefresh() PublishResponse {
	if len(act.Snap) == 0 {
		return PublishResponse{ID: act.ID, Success: false, Message: "No snap name provided for refresh"}
	}

	// Call the snapd API
	result, err := snapd.Refresh(act.Snap, nil)
	if err != nil {
		return PublishResponse{ID: act.ID, Success: false, Message: err.Error()}
	}
	return PublishResponse{ID: act.ID, Success: true, Result: result}
}

// SnapRevert reverts an existing snap
func (act *SubscribeAction) SnapRevert() PublishResponse {
	if len(act.Snap) == 0 {
		return PublishResponse{ID: act.ID, Success: false, Message: "No snap name provided for revert"}
	}

	// Call the snapd API
	result, err := snapd.Revert(act.Snap, nil)
	if err != nil {
		return PublishResponse{ID: act.ID, Success: false, Message: err.Error()}
	}
	return PublishResponse{ID: act.ID, Success: true, Result: result}
}

// SnapEnable enables an existing snap
func (act *SubscribeAction) SnapEnable() PublishResponse {
	if len(act.Snap) == 0 {
		return PublishResponse{ID: act.ID, Success: false, Message: "No snap name provided for enable"}
	}

	// Call the snapd API
	result, err := snapd.Enable(act.Snap, nil)
	if err != nil {
		return PublishResponse{ID: act.ID, Success: false, Message: err.Error()}
	}
	return PublishResponse{ID: act.ID, Success: true, Result: result}
}

// SnapDisable disables an existing snap
func (act *SubscribeAction) SnapDisable() PublishResponse {
	if len(act.Snap) == 0 {
		return PublishResponse{ID: act.ID, Success: false, Message: "No snap name provided for disable"}
	}

	// Call the snapd API
	result, err := snapd.Disable(act.Snap, nil)
	if err != nil {
		return PublishResponse{ID: act.ID, Success: false, Message: err.Error()}
	}
	return PublishResponse{ID: act.ID, Success: true, Result: result}
}

// SnapConf gets the config for a snap
func (act *SubscribeAction) SnapConf() PublishResponse {
	if len(act.Snap) == 0 {
		return PublishResponse{ID: act.ID, Success: false, Message: "No snap name provided for config"}
	}

	// Call the snapd API
	result, err := snapd.Conf(act.Snap)
	if err != nil {
		return PublishResponse{ID: act.ID, Success: false, Message: err.Error()}
	}

	return PublishResponse{ID: act.ID, Success: true, Result: result}
}

// SnapSetConf sets the config for a snap
func (act *SubscribeAction) SnapSetConf() PublishResponse {
	if len(act.Snap) == 0 {
		return PublishResponse{ID: act.ID, Success: false, Message: "No snap name provided for set config"}
	}

	// Deserialize the settings
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(act.Data), &data); err != nil {
		return PublishResponse{ID: act.ID, Success: false, Message: err.Error()}
	}

	// Call the snapd API
	result, err := snapd.SetConf(act.Snap, data)
	if err != nil {
		return PublishResponse{ID: act.ID, Success: false, Message: err.Error()}
	}

	return PublishResponse{ID: act.ID, Success: true, Result: result}
}