package ui

import (
	"fmt"
	"strings"
)

func (m Model) renderDetails(height int) string {
	var content string

	switch categories[m.category] {

	case "System":
		content = m.renderSystemDetails()

	case "Storage":
		content = m.renderStorageDetails()

	case "Devices":
		content = m.renderDeviceDetails()

	case "Network":
		content = m.renderNetworkDetails()
	}

	return panelStyle(m.panel == panelDetails).
		Width(45).
		Height(height).
		Render(content)
}

func (m Model) renderSystemDetails() string {
	switch m.selectedComponent() {

	case "Overview":
		return titleStyle.Render("System") +
			"\n\n" +
			normalStyle.Render(
				"Hostname        "+m.system.Hostname+"\n"+
					"OS              "+m.system.OS+"\n"+
					"Kernel          "+m.system.Kernel+"\n"+
					"Architecture    "+m.system.Architecture+"\n"+
					"Uptime          "+formatUptime(m.system.UptimeSeconds),
			)

	case "CPU":
		cpu := m.system.CPU

		return titleStyle.Render("CPU") +
			"\n\n" +
			normalStyle.Render(
				"Model           "+cpu.Model+"\n"+
					"Architecture    "+cpu.Architecture+"\n"+
					"Cores           "+fmt.Sprint(cpu.Cores)+"\n"+
					"Threads         "+fmt.Sprint(cpu.Threads)+"\n"+
					"Frequency       "+fmt.Sprint(cpu.FrequencyMHz)+" MHz",
			)

	case "Memory":
		memory := m.system.Memory

		return titleStyle.Render("Memory") +
			"\n\n" +
			normalStyle.Render(
				"Total           "+formatBytes(memory.TotalBytes)+"\n"+
					"Used            "+formatBytes(memory.UsedBytes)+"\n"+
					"Available       "+formatBytes(memory.AvailableBytes),
			)
	}

	return ""
}

func (m Model) renderStorageDetails() string {
	if len(m.system.Storage) == 0 {
		return titleStyle.Render("Storage") +
			"\n\n" +
			normalStyle.Render("No storage devices found.")
	}

	name := m.selectedComponent()

	for _, disk := range m.system.Storage {
		if disk.Name == name {
			return titleStyle.Render(disk.Name) +
				"\n\n" +
				normalStyle.Render(
					"Device          "+disk.Device+"\n"+
						"Size            "+formatBytes(disk.SizeBytes),
				)
		}
	}

	return ""
}

func (m Model) renderDeviceDetails() string {
	if len(m.system.PCI) == 0 {
		return titleStyle.Render("PCI") +
			"\n\n" +
			normalStyle.Render("No PCI devices found.")
	}

	address := m.selectedComponent()

	for _, device := range m.system.PCI {
		if device.Address == address {
			driver := device.Driver

			if driver == "" {
				driver = "None"
			}

			return titleStyle.Render("PCI Device") +
				"\n\n" +
				normalStyle.Render(
					"Address         "+device.Address+"\n"+
						"Vendor ID       "+device.VendorID+"\n"+
						"Device ID       "+device.DeviceID+"\n"+
						"Class           "+device.Class+"\n"+
						"Driver          "+driver,
				)
		}
	}

	return ""
}

func (m Model) renderNetworkDetails() string {
	if len(m.system.Network) == 0 {
		return titleStyle.Render("Network") +
			"\n\n" +
			normalStyle.Render("No network interfaces found.")
	}

	name := m.selectedComponent()

	for _, network := range m.system.Network {
		if network.Name == name {

			var addresses strings.Builder

			for _, address := range network.Addresses {
				addresses.WriteString(address)
				addresses.WriteString("\n")
			}

			return titleStyle.Render(network.Name) +
				"\n\n" +
				normalStyle.Render(
					"State           "+network.State+"\n"+
						"MAC             "+network.MACAddress+"\n"+
						"\nAddresses\n"+
						addresses.String(),
				)
		}
	}

	return ""
}
