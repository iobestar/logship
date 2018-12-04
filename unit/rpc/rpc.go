package rpc

import (
	"time"
	"github.com/iobestar/logship/unit"
)

func NewLogService(manager *unit.Manager) (LogUnitServiceServer, error) {
	return &DefaultLogService{
		unitManager: manager,
	}, nil
}

type DefaultLogService struct {
	unitManager *unit.Manager
}

func (dls *DefaultLogService) GetUnits(rq *Empty, stream LogUnitService_GetUnitsServer) error {

	for _, luId := range dls.unitManager.GetLogUnitIds() {
		stream.Send(&UnitRS{
			Unit: luId,
		})
	}
	return nil
}

func (dls *DefaultLogService) GetNLines(rq *NLineRQ, stream LogUnitService_GetNLinesServer) error {
	logUnit := dls.unitManager.GetLogUnit(rq.UnitId)
	if nil == logUnit {
		return nil
	}

	return logUnit.GetNLines(int(rq.Count), func(line string) error {
		r := NLineRS{
			Line: line,
		}
		stream.Send(&r)
		return nil
	})
}

func (dls *DefaultLogService) GetNLogs(rq *NLogRQ, stream LogUnitService_GetNLogsServer) error {

	logUnit := dls.unitManager.GetLogUnit(rq.UnitId)
	if nil == logUnit {
		return nil
	}

	return logUnit.GetNLogs(int(rq.Count), func(logEntry *unit.LogEntry) error {
		l := LogRS{
			Payload: logEntry.Log(),
		}
		stream.Send(&l)
		return nil
	})
}

func (dls *DefaultLogService) GetTLogs(rq *TLogRQ, stream LogUnitService_GetTLogsServer) error {

	unitId := rq.UnitId

	logUnit := dls.unitManager.GetLogUnit(unitId)
	if nil == logUnit {
		return nil
	}

	d, err := time.ParseDuration(rq.Duration)
	if nil != err {
		return err
	}

	offset := rq.Offset
	if offset <= 0 {
		offset = time.Now().UnixNano()
	}

	return logUnit.GetTLogs(d, offset, func(logEntry *unit.LogEntry) error {
		l := LogRS{
			Payload: logEntry.Log(),
		}
		stream.Send(&l)
		return nil
	})

	return nil
}
