package internal

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/svc/mgr"
)

func exePath() (string, error) {
	var logger = GetLogger()
	prog := os.Args[0]
	p, err := filepath.Abs(prog)
	if err != nil {
		return "", err
	}

	fi, err := os.Stat(p)
	if err == nil {
		if !fi.Mode().IsDir() {
			return p, nil
		}

		logger.Errorf("%s is directory", p)
		err = errors.New(fmt.Sprintf("%s is not a directory", p))
	}

	if filepath.Ext(p) == "" {
		p += ".exe"
		fi, err := os.Stat(p)
		if err == nil {
			if !fi.Mode().IsDir() {
				return p, nil
			}
		}
		logger.Errorf("%s is directory", p)
		err = errors.New(fmt.Sprintf("%s is not a directory", p))
	}

	return "", err
}

func InstallService(name, desc string) error {
	var logger = GetLogger()
	exepath, err := exePath()
	if err != nil {
		return err
	}
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	_, err = m.OpenService(name)
	if err == nil {
		logger.Errorf("service %s already exists", name)
		return errors.New(fmt.Sprintf("service %s already exists", name))
	}

	s, err := m.CreateService(name, exepath, mgr.Config{DisplayName: desc}, "is", "auto-started")
	if err != nil {
		return err
	}
	defer s.Close()
	/*
		err = eventlog.InstallAsEventCreate(name, eventlog.Error|eventlog.Warning|eventlog.Info)
		if err != nil {
			s.Delete()
			logger.Errorf("SetupEventLogSource() failed: %s", err)
			return errors.New(fmt.Sprintf("SetupEventLogSource() failed: %s", err))
		}
	*/
	logger.Info(fmt.Sprintf("%s install completed %s", name, desc))
	return nil
}

func RemoveService(name string) error {
	var logger = GetLogger()
	logger.Info(fmt.Sprintf("%s uninstall called", name))
	m, err := mgr.Connect()
	if err != nil {
		logger.Error(fmt.Sprintf("%v", err))
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		logger.Error(fmt.Sprintf("%v", err))
		return err
	}
	defer s.Close()
	err = s.Delete()
	if err != nil {
		logger.Error(fmt.Sprintf("%v", err))
		return err
	}
	/*
		err = eventlog.Remove(name)
		if err != nil {
			return fmt.Errorf("RemoveEventLogSource() failed %s", err)
		}
	*/
	logger.Info(fmt.Sprintf("%s uninstall completed", name))
	return nil
}
