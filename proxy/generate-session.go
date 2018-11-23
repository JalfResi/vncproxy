package proxy

import (
	"encoding/json"
	"os/exec"
	"vncproxy/logger"
)

type response struct {
	ID             string `json:"id"`
	Target         string `json:"target"`
	TargetHostname string `json:"hostname"`
	TargetPort     string `json:"port"`
	TargetPassword string `json:"password"`
}

func GenerateSession(cmd string) (*VncSession, error) {

	out, err := exec.Command(cmd).Output()
	if err != nil {
		logger.Errorf("Failure to get a response from external command: %s", err.Error())
		return nil, err
	}

	var r response
	r.ID = "generatedSession"

	err = json.Unmarshal(out, &r)
	if err != nil {
		logger.Errorf("Failure decoding response when generating session: %s", err.Error())
		return nil, err
	}

	logger.Infof("External Session read: %s", out)

	return &VncSession{
		Target:         r.Target,
		TargetHostname: r.TargetHostname,
		TargetPort:     r.TargetPort,
		TargetPassword: r.TargetPassword, //"vncPass",
		ID:             r.ID,
		Status:         SessionStatusInit,
		Type:           SessionTypeProxyPass,
	}, nil
}
