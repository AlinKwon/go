package internal

import (
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

func StartService(name string) error {
	var logger = GetLogger()
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		logger.Errorf("error opening service: %v", err)
		return err
	}
	defer s.Close()
	err = s.Start("is", "manual-started")
	if err != nil {
		logger.Errorf("could not starting service: %v", err)
		return err
	}
	return nil
}

func ControlService(name string, c svc.Cmd, to svc.State) error {
	var logger = GetLogger()
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		logger.Errorf("could not access service: %v", err)
		return err
	}
	status, err := s.Control(c)
	if err != nil {
		logger.Errorf("could not control service: %v", err)
		return err
	}
	timeout := time.Now().Add(10 * time.Second)
	for status.State != to {
		if timeout.Before(time.Now()) {
			logger.Errorf("timeout waiting for service to go to state=%d", to)
			return err
		}
		time.Sleep(300 * time.Millisecond)
		status, err = s.Query()
		if err != nil {
			logger.Errorf("could not retrieve service status: %v", err)
			return err
		}
	}
	return nil
}
