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

package httpControl

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type rawDetectedWLAN struct {
	Auth    string `json:"auth"`
	BSSID   string `json:"bssid"`
	Channel string `json:"channel"`
	RSSI    string `json:"rssi"`
	SSID    string `json:"ssid"`
	Extch   string `json:"extch"`
}

type wlanGetApListReturn struct {
	Res    string            `json:"res"`
	Aplist []rawDetectedWLAN `json:"aplist"`
}

type DetectedWLAN struct {
	Auth    string  `json:"auth"`
	BSSID   string  `json:"bssid"`
	Channel int     `json:"channel"`
	RSSI    float32 `json:"rssi"`
	SSID    string  `json:"ssid"`
	Extch   string  `json:"extch"`
}

// GetApList queries the device for its currently visible wifi access points.
func (rpc *RPC) GetApList(ctx context.Context) ([]DetectedWLAN, error) {
	rawReq := wlanGetApListReturn{}

	reply, reqErr := rpc.transport.MakeRequest(ctx, "wlanGetApListEx")
	if reqErr != nil {
		return nil, reqErr
	}
	parseErr := json.Unmarshal(reply, &wlanGetApListReturn{})

	formatted := make([]DetectedWLAN, len(rawReq.Aplist))
	for i, ap := range rawReq.Aplist {
		channel, _ := strconv.Atoi(ap.Channel)
		rssRaw, _ := strconv.ParseFloat(ap.RSSI, 32)
		ssid, _ := hex.DecodeString(ap.SSID)
		formatted[i] = DetectedWLAN{
			Auth:    ap.Auth,
			BSSID:   ap.BSSID,
			Channel: channel,
			RSSI:    float32(rssRaw) / 100,
			SSID:    string(ssid),
			Extch:   ap.Extch,
		}
	}

	return formatted, parseErr
}

// ConnectToWifi requests the device to connect to a discovered AP.
func (rpc *RPC) ConnectToWifi(ctx context.Context, ssid string, channel int, auth string, encryption string, password string) error {
	flags := make([]string, 6)

	flags[0] = fmt.Sprintf("ssid=%s", strings.ToUpper(hex.EncodeToString([]byte(ssid))))
	flags[1] = fmt.Sprintf("ch=%d", channel)
	flags[2] = fmt.Sprintf("auth=%s", auth)
	flags[3] = fmt.Sprintf("encry=%s", encryption)
	flags[4] = fmt.Sprintf("pwd=%s", strings.ToUpper(hex.EncodeToString([]byte(password))))
	flags[5] = "chext=1"

	// This apparently has no reply body
	_, reqErr := rpc.transport.MakeRequest(ctx, "wlanConnectApEx", flags...)

	return reqErr
}

// ConnectToHiddenWifi requests the device to connect to an undiscovered AP.
func (rpc *RPC) ConnectToHiddenWifi(ctx context.Context, ssid string, password string) error {
	var flags []string

	flags = append(flags, fmt.Sprintf("ssid=%s", strings.ToUpper(hex.EncodeToString([]byte(ssid)))))

	if password != "" {
		flags = append(flags, fmt.Sprintf("pwd=%s", strings.ToUpper(hex.EncodeToString([]byte(password)))))
	}

	// This apparently has no reply body
	_, reqErr := rpc.transport.MakeRequest(ctx, "wlanConnectHideApEx", flags...)

	return reqErr
}

type WlanState int

const (
	WLAN_FAIL WlanState = iota
	WLAN_PAIRING_FAIL
	WLAN_IN_PROGRESS
	WLAN_OK
)

func (state WlanState) MarshalText() ([]byte, error) {
	switch state {
	case WLAN_FAIL:
		return []byte("Disconnected"), nil
	case WLAN_PAIRING_FAIL:
		return []byte("Pairing Failed"), nil
	case WLAN_IN_PROGRESS:
		return []byte("In Progress"), nil
	case WLAN_OK:
		return []byte("Connected"), nil
	default:
		return []byte(fmt.Sprintf("Unknown :%d", state)), errors.New("unknown WlanState value")
	}
}

func (state *WlanState) unmarshalApiText(input []byte) error {
	switch string(input) {
	case "FAIL":
		*state = WLAN_FAIL
	case "PAIRFAIL":
		*state = WLAN_PAIRING_FAIL
	case "PROCESS":
		*state = WLAN_IN_PROGRESS
	case "OK":
		*state = WLAN_OK
	default:
		*state = WLAN_FAIL
		return errors.New("Unknown API string: " + string(input))
	}
	return nil
}

// GetWlanState queries the device for its current WLAN connection.
func (rpc *RPC) GetWlanState(ctx context.Context) (WlanState, error) {
	var state WlanState

	reply, reqErr := rpc.transport.MakeRequest(ctx, "wlanGetConnectState")
	if reqErr != nil {
		return state, reqErr
	}
	parseErr := state.unmarshalApiText(reply)

	return state, parseErr
}
