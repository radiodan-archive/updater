package utils

import (
  "os/exec"
)

func MachineHardware() (machineName string) {
  cmd := exec.Command("uname", "-m")
	out, _ := cmd.Output()
	return string(out[:])
}
