/*
arylic-connect, an API broker for Arylic Audio devices
Copyright (C) 2023  Zach Strauss

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package websocketControl

import (
	"context"
	"encoding/json"
)

// GetStatus queries the device for its current status summary.
func (rpc *RPC) GetStatus(ctx context.Context) (StatusChangeMessage, error) {
	status := StatusChangeMessage{}

	data, reqErr := requestWithResponse(ctx, rpc.transport, "#CMD:STATUS", "STATUS")
	if reqErr != nil {
		return status, reqErr
	}
	jsonErr := json.Unmarshal(data, &status)
	return status, jsonErr
}
