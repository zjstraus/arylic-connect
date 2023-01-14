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
	"fmt"
	"strconv"
)

// rawEndpointStatus is the truly awful mess of a  struct the far API returns
// from its only status endpoint. Everything is strings and almost nothing is
// documented.
//
// https://developer.arylic.com/httpapi/#get-device-metadata
type rawEndpointStatus struct {
	UUID                    string `json:"uuid"`
	DeviceName              string `json:"DeviceName"`
	GroupName               string `json:"GroupName"`
	SSID                    string `json:"ssid"`
	Language                string `json:"language"`
	Firmware                string `json:"firmware"`
	Hardware                string `json:"hardware"`
	Build                   string `json:"build"`
	Project                 string `json:"project"`
	PrivPrj                 string `json:"priv_prj"`
	ProjectBuildName        string `json:"project_build_name"`
	Release                 string `json:"release"`
	TempUUID                string `json:"temp_uuid"`
	HideSSID                string `json:"hideSSID"`
	SSIDStrategy            string `json:"SSIDStrategy"`
	Branch                  string `json:"branch"`
	Group                   string `json:"group"`
	WmrmVersion             string `json:"wmrm_version"`
	Internet                string `json:"internet"`
	MAC                     string `json:"MAC"`
	STAMAC                  string `json:"STA_MAC"`
	CountryCode             string `json:"CountryCode"`
	CountryRegion           string `json:"CountryRegion"`
	Netstat                 string `json:"netstat"`
	ESSID                   string `json:"ESSID"`
	Apcli0                  string `json:"apcli0"`
	Eth2                    string `json:"eth2"`
	Ra0                     string `json:"ra0"`
	EthDhcp                 string `json:"eth_dhcp"`
	VersionUpdate           string `json:"VersionUpdate"`
	NewVer                  string `json:"newVer"`
	SetDnsEnable            string `json:"set_dns_enable"`
	McuVer                  string `json:"mcu_ver"`
	McuVerNew               string `json:"mcu_ver_new"`
	DspVer                  string `json:"dsp_ver"`
	DspVerNew               string `json:"dsp_ver_new"`
	Date                    string `json:"date"`
	Time                    string `json:"time"`
	Tz                      string `json:"tz"`
	Region                  string `json:"region"`
	PromptStatus            string `json:"prompt_status"`
	IotVer                  string `json:"iot_ver"`
	UpnpVersion             string `json:"upnp_version"`
	Cap1                    string `json:"cap1"`
	Capability              string `json:"capability"`
	Languages               string `json:"languages"`
	Streams                 string `json:"streams"`
	StreamsAll              string `json:"streams_all"`
	External                string `json:"external"`
	PlmSupport              string `json:"plm_support"`
	PresetKey               string `json:"preset_key"`
	SpotifyActive           string `json:"spotify_active"`
	LbcSupport              string `json:"lbc_support"`
	PrivacyMode             string `json:"privacy_mode"`
	WifiChannel             string `json:"WifiChannel"`
	RSSI                    string `json:"RSSI"`
	BSSID                   string `json:"BSSID"`
	Battery                 string `json:"battery"`
	BatteryPercent          string `json:"battery_percent"`
	SecureMode              string `json:"securemode"`
	Auth                    string `json:"auth"`
	Encry                   string `json:"encry"`
	UpnpUUID                string `json:"upnp_uuid"`
	UartPassPort            string `json:"uart_pass_port"`
	CommunicationPort       string `json:"communication_port"`
	WebFirmwareUpdateHide   string `json:"web_firmware_update_hide"`
	IgnoreTalkStart         string `json:"ignore_talkstart"`
	WebLoginResult          string `json:"web_login_result"`
	SilenceOtaTime          string `json:"silenceOTATime"`
	IgnoreSilenceOtaTime    string `json:"ignore_silenceOTATime"`
	NewTuneinPresetAndAlarm string `json:"new_tunein_preset_and_alarm"`
	TidalVersion            string `json:"tidal_version"`
	ServiceVersion          string `json:"service_version"`
	ETHMAC                  string `json:"ETH_MAC"`
	Security                string `json:"security"`
	SecurityVersion         string `json:"security_version"`
}

// EndpointStatus is the collected status of the connected device. The upstream
// API only has one status command that returns all of this. We reformat it to
// be as useful as possible and pass it all along.
//
// The upstream documentation is bad, so many of these are unknown and may move
// around as more things are figured out.
type EndpointStatus struct {
	DeviceID   string `json:"deviceId"`   // Stable, unique ID for this hardware
	InstanceID string `json:"instanceId"` // ID that will change between reboots

	Model       string `json:"model"`
	PresetCount int    `json:"presetCount"`

	DeviceName string `json:"deviceName"`
	GroupName  string `json:"groupName"`

	Wifi struct {
		Status string `json:"status"`

		ExternalSSID string  `json:"externalSsid"`
		RSSI         float32 `json:"rssi"`
		BSSID        string  `json:"bssid"`

		StationMAC string `json:"stationMac"`
		Channel    int    `json:"channel"`

		LocalSSID     string `json:"localSsid"`
		HideSSID      bool   `json:"hideSsid"`
		LocalMAC      string `json:"localMac"`
		CountryCode   string `json:"countryCode"`   // Unknown
		CountryRegion string `json:"countryRegion"` // Unknown

		SSIDStrategy string `json:"ssidStrategy"` // Unknown
		Auth         string `json:"auth"`
		Encryption   string `json:"encryption"`

		ClientIp string `json:"clientIp"`
		ApIp     string `json:"apIp"`
	} `json:"wifi"`

	Updates struct {
		UpdateAvailable bool   `json:"updateAvailable"`
		Version         string `json:"version"`

		MCUAvailable string `json:"mcuAvailable"`
		DSPAvailable string `json:"dspAvailable"`
	} `json:"Updates"`

	Versions struct {
		Firmware   string `json:"firmware"`
		Build      string `json:"build"`
		Project    string `json:"project"`
		ProjectPrv string `json:"projectPrv"`
		BuildName  string `json:"buildName"`
		Release    string `json:"release"`
		Branch     string `json:"branch"`

		DSP string `json:"dsp"`
		MCU string `json:"mcu"`

		MultiroomLib string `json:"multiroomLib"`
		IOTLib       string `json:"iotLib"` // Unused?
	} `json:"versions"`

	Battery struct {
		Charging bool    `json:"charging"`
		Level    float32 `json:"level"` // 0 - 1
	} `json:"battery"`

	Slave struct {
		IsSlave  bool   `json:"isSlave"`
		MasterID string `json:"masterId"`
	} `json:"slave"`

	Locale struct {
		Language string `json:"language"`
		Date     string `json:"date"`
		Time     string `json:"time"`
		Tz       string `json:"tz"`
		Region   string `json:"region"`
	} `json:"locale"`

	NetworkServicePorts struct {
		Communication   int `json:"communication"`
		UartPassthrough int `json:"uartPassthrough"`
	} `json:"networkServicePorts"`

	Ethernet struct {
		Ip  string `json:"ip"`
		MAC string `json:"mac"`
	} `json:"ethernet"`

	Network struct {
		InternetActive bool `json:"internetActive"`
		DHCPEnabled    bool `json:"dhcpEnabled"`
		DNSEnabled     bool `json:"dnsEnabled"`
	} `json:"network"`

	UPNP struct {
		Version string `json:"version"`
		ID      string `json:"id"`
	} `json:"upnp"`

	ModuleFeatureMask [32]bool `json:"moduleFeatureMask"`
	StreamsMask       [32]bool `json:"streamsMask"`
	InputsMask        [16]bool `json:"inputsMask"`

	Settings struct {
		VoicePrompt           bool   `json:"voicePrompt"`
		WebFirmwareUpdateHide bool   `json:"webFirmwareUpdateHide"`
		IgnoreSilenceOtaTime  bool   `json:"ignoreSilenceOtaTime"`
		SilenceOtaTime        string `json:"silenceOTATime"`
		IgnoreTalkStart       bool   `json:"ignoreTalkStart"`
	} `json:"settings"`

	Unknown struct {
		Cap1          string `json:"cap1"`
		Languages     string `json:"languages"`
		StreamsAll    string `json:"streamsAll"`
		External      string `json:"external"`
		SpotifyActive string `json:"spotifyActive"`
		LbcSupport    string `json:"lbcSupport"`
		PrivacyMode   string `json:"privacyMode"`
		SecureMode    string `json:"secureMode"`

		WebLoginResult          string `json:"webLoginResult"`
		NewTuneinPresetAndAlarm string `json:"newTuneinPresetAndAlarm"`
		TidalVersion            string `json:"tidalVersion"`
		ServiceVersion          string `json:"serviceVersion"`
		Security                string `json:"security"`
		SecurityVersion         string `json:"securityVersion"`
	}
}

func wifiStatusMap(rawVal string) string {
	switch rawVal {
	case "0":
		return "No Connection"
	case "1":
		return "Connecting"
	case "2":
		return "Connected"
	default:
		return fmt.Sprintf("Unknown (%s)", rawVal)
	}
}

func (status *EndpointStatus) UnmarshalJSON(input []byte) error {
	rawStruct := rawEndpointStatus{}
	initialParseErr := json.Unmarshal(input, &rawStruct)
	if initialParseErr != nil {
		return initialParseErr
	}

	status.DeviceID = rawStruct.UUID
	status.InstanceID = rawStruct.TempUUID
	status.Model = rawStruct.Hardware

	presetCount, _ := strconv.Atoi(rawStruct.PresetKey)
	status.PresetCount = presetCount

	status.DeviceName = rawStruct.DeviceName
	status.GroupName = rawStruct.GroupName

	status.Wifi.Status = wifiStatusMap(rawStruct.Netstat)
	essid, _ := hex.DecodeString(rawStruct.ESSID)
	status.Wifi.ExternalSSID = string(essid)
	rssi, _ := strconv.ParseFloat(rawStruct.RSSI, 32)
	status.Wifi.RSSI = float32(rssi)
	status.Wifi.BSSID = rawStruct.BSSID
	status.Wifi.StationMAC = rawStruct.STAMAC
	channel, _ := strconv.Atoi(rawStruct.WifiChannel)
	status.Wifi.Channel = channel
	status.Wifi.LocalSSID = rawStruct.SSID
	status.Wifi.HideSSID = rawStruct.HideSSID == "1"
	status.Wifi.LocalMAC = rawStruct.MAC
	status.Wifi.CountryCode = rawStruct.CountryCode
	status.Wifi.CountryRegion = rawStruct.CountryRegion
	status.Wifi.SSIDStrategy = rawStruct.SSIDStrategy
	status.Wifi.Auth = rawStruct.Auth
	status.Wifi.Encryption = rawStruct.Encry
	status.Wifi.ClientIp = rawStruct.Apcli0
	status.Wifi.ApIp = rawStruct.Ra0

	status.Updates.UpdateAvailable = rawStruct.VersionUpdate == "1"
	status.Updates.Version = rawStruct.NewVer
	status.Updates.MCUAvailable = rawStruct.McuVerNew
	status.Updates.DSPAvailable = rawStruct.DspVerNew

	status.Versions.Firmware = rawStruct.Firmware
	status.Versions.Build = rawStruct.Build
	status.Versions.Project = rawStruct.Project
	status.Versions.ProjectPrv = rawStruct.PrivPrj
	status.Versions.BuildName = rawStruct.ProjectBuildName
	status.Versions.Release = rawStruct.Release
	status.Versions.Branch = rawStruct.Branch
	status.Versions.DSP = rawStruct.DspVer
	status.Versions.MCU = rawStruct.McuVer
	status.Versions.MultiroomLib = rawStruct.WmrmVersion
	status.Versions.IOTLib = rawStruct.IotVer

	status.Battery.Charging = rawStruct.Battery == "1"
	batteryLevel, _ := strconv.Atoi(rawStruct.BatteryPercent)
	status.Battery.Level = float32(batteryLevel) / 100

	status.Slave.IsSlave = rawStruct.Group == "1"
	status.Slave.MasterID = ""

	status.Locale.Language = rawStruct.Language
	status.Locale.Date = rawStruct.Date
	status.Locale.Time = rawStruct.Time
	status.Locale.Tz = rawStruct.Tz
	status.Locale.Region = rawStruct.Region

	comPort, _ := strconv.Atoi(rawStruct.CommunicationPort)
	status.NetworkServicePorts.Communication = comPort
	uartPort, _ := strconv.Atoi(rawStruct.UartPassPort)
	status.NetworkServicePorts.UartPassthrough = uartPort

	status.Ethernet.MAC = rawStruct.ETHMAC
	status.Ethernet.Ip = rawStruct.Eth2

	status.Network.InternetActive = rawStruct.Internet == "1"
	status.Network.DHCPEnabled = rawStruct.EthDhcp == "1"
	status.Network.DNSEnabled = rawStruct.SetDnsEnable == "1"

	status.UPNP.Version = rawStruct.UpnpVersion
	status.UPNP.ID = rawStruct.UpnpUUID

	moduleMask, _ := strconv.Atoi(rawStruct.Capability)
	for i, _ := range status.ModuleFeatureMask {
		status.ModuleFeatureMask[i] = (moduleMask & (2 ^ i)) > 0
	}

	streamsMask, _ := strconv.Atoi(rawStruct.Streams)
	for i, _ := range status.StreamsMask {
		status.StreamsMask[i] = (streamsMask & (2 ^ i)) > 0
	}

	inputMask, _ := strconv.Atoi(rawStruct.PlmSupport)
	for i, _ := range status.InputsMask {
		status.InputsMask[i] = (inputMask & (2 ^ i)) > 0
	}

	status.Settings.VoicePrompt = rawStruct.PromptStatus == "1"
	status.Settings.WebFirmwareUpdateHide = rawStruct.WebFirmwareUpdateHide == "1"
	status.Settings.IgnoreSilenceOtaTime = rawStruct.IgnoreSilenceOtaTime == "1"
	status.Settings.SilenceOtaTime = rawStruct.SilenceOtaTime
	status.Settings.IgnoreTalkStart = rawStruct.IgnoreTalkStart == "1"

	status.Unknown.Cap1 = rawStruct.Cap1
	status.Unknown.Languages = rawStruct.Languages
	status.Unknown.StreamsAll = rawStruct.StreamsAll
	status.Unknown.External = rawStruct.External
	status.Unknown.SpotifyActive = rawStruct.SpotifyActive
	status.Unknown.LbcSupport = rawStruct.LbcSupport
	status.Unknown.PrivacyMode = rawStruct.PrivacyMode
	status.Unknown.SecureMode = rawStruct.SecureMode
	status.Unknown.WebLoginResult = rawStruct.WebLoginResult
	status.Unknown.NewTuneinPresetAndAlarm = rawStruct.NewTuneinPresetAndAlarm
	status.Unknown.TidalVersion = rawStruct.TidalVersion
	status.Unknown.ServiceVersion = rawStruct.ServiceVersion
	status.Unknown.Security = rawStruct.Security
	status.Unknown.SecurityVersion = rawStruct.SecurityVersion

	return nil
}

// GetStatus queries the device for its current status summary.
func (rpc *RPC) GetStatus(ctx context.Context) (EndpointStatus, error) {
	status := EndpointStatus{}

	reply, reqErr := rpc.transport.MakeRequest(ctx, "getStatusEx")
	if reqErr != nil {
		return status, reqErr
	}
	parseErr := json.Unmarshal(reply, &status)

	return status, parseErr
}
