// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package genericattributeprofile

type GattWriteOption int32

const SignatureGattWriteOption string = "enum(Windows.Devices.Bluetooth.GenericAttributeProfile.GattWriteOption;i4)"

const (
	GattWriteOptionWriteWithResponse    GattWriteOption = 0
	GattWriteOptionWriteWithoutResponse GattWriteOption = 1
)