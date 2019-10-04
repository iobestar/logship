package rpc

import (
	"context"
	"github.com/go-kit/kit/log/level"
	"github.com/iobestar/logship/logging"
	"github.com/iobestar/logship/unit"
	"io"
	"regexp"
	"time"
)

func NewLogService(logUnits unit.LogUnits) (LogUnitServiceServer, error) {
	return &DefaultLogUnitService{
		logUnits: logUnits,
	}, nil
}

type DefaultLogUnitService struct {
	logUnits unit.LogUnits
}

func (s *DefaultLogUnitService) NLines(rq *NLineRQ, stream LogUnitService_NLinesServer) error {
	if u, ok := s.logUnits[rq.UnitId]; ok {
		ctx, cancel := context.WithCancel(stream.Context())
		lines, errors := u.StreamLines(ctx)
		count := 0
		for {
			if count == int(rq.Count) {
				cancel()
				return nil
			}
			select {
			case l, ok := <-lines:
				if ok {
					count++
					err := stream.Send(&NLineRS{
						Line: l,
					})
					if err == io.EOF {
						return nil
					}
					if nil != err {
						level.Error(logging.Logger).Log("err", err.Error())
						return err
					}
				} else {
					return nil
				}
			case err, ok := <-errors:
				if ok {
					return err
				}
			case <-stream.Context().Done():
				return nil
			}
		}
	}
	return nil
}

func (s *DefaultLogUnitService) NLogs(rq *NLogRQ, stream LogUnitService_NLogsServer) error {
	if u, ok := s.logUnits[rq.UnitId]; ok {
		logPattern, err := regexp.Compile(rq.LogPattern)
		if nil != err {
			return err
		}
		ctx, cancel := context.WithCancel(stream.Context())
		logs, errors := u.StreamLogs(ctx, rq.DateTimeLayout, logPattern)
		count := 0
		for {
			if count == int(rq.Count) {
				cancel()
				return nil
			}
			select {
			case l, ok := <-logs:
				if ok {
					count++
					err := stream.Send(&LogRS{
						Log: l.Log(),
					})
					if err == io.EOF {
						return nil
					}
					if nil != err {
						level.Error(logging.Logger).Log("err", err.Error())
						return err
					}
				} else {
					return nil
				}
			case err, ok := <-errors:
				if ok {
					return err
				}
			case <-stream.Context().Done():
				return nil
			}
		}
	}
	return nil
}

func (s *DefaultLogUnitService) TLogs(rq *TLogRQ, stream LogUnitService_TLogsServer) error {
	if u, ok := s.logUnits[rq.UnitId]; ok {
		d, err := time.ParseDuration(rq.Duration)
		if nil != err {
			return err
		}

		logPattern, err := regexp.Compile(rq.LogPattern)
		if nil != err {
			return err
		}

		ctx, cancel := context.WithCancel(stream.Context())
		logs, errors := u.StreamLogs(ctx, rq.DateTimeLayout, logPattern)
		limit := time.Now().UnixNano() - d.Nanoseconds()
		for {
			select {
			case l, ok := <-logs:
				if ok {

					if l.Timestamp < limit {
						cancel()
						return nil
					}

					err := stream.Send(&LogRS{
						Log: l.Log(),
					})
					if err == io.EOF {
						return nil
					}
					if nil != err {
						level.Error(logging.Logger).Log("err", err.Error())
						return err
					}
				} else {
					return nil
				}
			case err, ok := <-errors:
				if ok {
					return err
				}
			case <-stream.Context().Done():
				return nil
			}
		}
	}
	return nil
}

func (s *DefaultLogUnitService) GetUnits(rq *Empty, stream LogUnitService_GetUnitsServer) error {

	for _, luId := range s.logUnits.GetLogUnitIds() {
		err := stream.Send(&UnitRS{
			Unit: luId,
		})
		if err == io.EOF {
			return nil
		}
		if nil != err {
			level.Error(logging.Logger).Log("err", err.Error())
			return err
		}
	}
	return nil
}
