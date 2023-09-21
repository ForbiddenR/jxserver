package manage

type Interface interface {
	SwitchLogging(name string, instructment uint8) error
	GetConnections(mode uint8) (uint64, string, error)
	CloseConnection(sn string) error
	GetConnectionStatus(sn string) (ct string, lht string, local string, remote string, err error)
	GetConnectionAlarmRule() (rule string, limit uint, err error)
	SetConnectionAlarmRules(rule string, limit uint) error
}

type NoopInterface struct{}

func (*NoopInterface) SwitchLogging(name string, instructment uint8) error {
	return nil
}

func (*NoopInterface) GetConnections(mode uint8) (uint64, string, error) {
	return 0, "", nil
}

func (*NoopInterface) CloseConnection(sn string) error {
	return nil
}

func (*NoopInterface) GetConnectionStatus(sn string) (ct string, lht string, local string, remote string, err error) {
	return "", "", "127.0.0.1", "127.0.0.1", nil
}

func (*NoopInterface) GetConnectionAlarmRule() (rule string, limit uint, err error) {
	return "", 0, nil
}

func (*NoopInterface) SetConnectionAlarmRules(rule string, limit uint) error {
	return nil
}
