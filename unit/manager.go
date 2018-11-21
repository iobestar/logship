package unit

import (
	"github.com/iobestar/logship/config"
	"sync"
)

type Manager struct {
	logUnits logUnits
	mutex    *sync.RWMutex
}

type logUnits map[string]*LogUnit

func NewManager(config config.Config) (*Manager, error) {

	var lum = &Manager{
		logUnits: logUnits{},
		mutex:    &sync.RWMutex{},
	}
	err := lum.Reload(config)
	if nil != err {
		return nil, err
	}
	return lum, err
}

func (lum *Manager) Reload(config config.Config) error {

	result := make(logUnits)
	for _, cLogUnit := range config.LogUnits {
		lUnit, err := NewLogUnit(cLogUnit)
		if nil != err {
			return err
		}
		result[lUnit.Id] = lUnit
	}

	lum.mutex.Lock()
	defer lum.mutex.Unlock()
	lum.logUnits = result
	return nil
}

func (lum *Manager) GetLogUnit(unitId string) *LogUnit {

	lum.mutex.RLock()
	defer lum.mutex.RUnlock()

	if u, ok := lum.logUnits[unitId]; ok {
		return u
	}
	return nil
}

func (lum *Manager) GetLogUnitIds() []string {

	lum.mutex.RLock()
	defer lum.mutex.RUnlock()
	result := []string{}
	for _, v := range lum.logUnits {
		result = append(result, v.Id)
	}
	return result
}
