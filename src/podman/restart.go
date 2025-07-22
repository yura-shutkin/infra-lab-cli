package podman

// RestartMachine restarts the specified podman machine by stopping and then starting it
func RestartMachine(binaryName, machineName string) error {
	// Stop the machine
	err := StopMachine(binaryName, machineName)
	if err != nil {
		return err
	}

	// Start the machine
	err = StartMachine(binaryName, machineName)
	if err != nil {
		return err
	}

	return nil
}
