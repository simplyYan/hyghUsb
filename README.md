# hyghUsb Library

[![License: BSD-3-Clause](https://img.shields.io/badge/license-BSD--3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

## Overview

`hyghUsb` is a high-level USB interaction library for Go, built on top of the `gousb` package. It aims to simplify USB device management, providing an easy-to-use interface for connecting, listing, obtaining information, sending and receiving data, and disabling USB devices.

## Features

- **Device Connection:** Easily connect to USB devices using the `NewUsbManager` function.

- **Device Listing:** List all connected USB devices with the `ListDevices` function.

- **Device Information:** Retrieve information about a USB device using the `GetDeviceInformation` function.

- **Data Transmission:** Send and receive data to/from a USB device with the `SendData` and `ReceiveData` functions.

- **Device Disabling:** Disable a USB device using the `DisableDevice` function.

- **BSD-3-Clause License:** `hyghUsb` is released under the BSD-3-Clause license, providing flexibility for both open source and commercial projects.

## Installation

Install `hyghUsb` using `go get`:

```bash
go get github.com/simplyYan/hyghUsb
```
## Usage

Here's a simple example demonstrating the basic functionality of hyghUsb:
```go
package main

import (
    "fmt"
    "github.com/simplyYan/hyghUsb"
    "log"
)

func main() {
    // Create a new USB manager
    usbManager, err := hyghUsb.NewUsbManager()
    if err != nil {
        log.Fatalf("Error creating USB manager: %v", err)
    }
    defer usbManager.Close()

    // List all connected USB devices
    devices, err := usbManager.ListDevices()
    if err != nil {
        log.Fatalf("Error listing USB devices: %v", err)
    }

    for _, device := range devices {
        // Get information about each device
        info, err := usbManager.GetDeviceInformation(device)
        if err != nil {
            log.Printf("Error getting device information: %v", err)
            continue
        }
        fmt.Println("Device Information:", info)

        // Add additional operations as needed, such as sending or receiving data.
    }

    // Example: Disable the first device (if it exists)
    if len(devices) > 0 {
        err := usbManager.DisableDevice(devices[0])
        if err != nil {
            log.Printf("Error disabling device: %v", err)
        } else {
            fmt.Println("Device disabled successfully.")
        }
    }
}

```
## Contributing
Feel free to contribute to hyghUsb by opening issues or pull requests on the GitHub repository.
License

#### hyghUsb is licensed under the BSD-3-Clause license. See the LICENSE file for details.

For more information, visit: https://opensource.org/licenses/BSD-3-Clause

> A project founded by Wesley Yan Soares Brehmer
