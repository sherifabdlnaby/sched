package sched

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
)

type MiddleWarehandler interface {
	Handler(s *Schedule, newstate State) (stop bool, err error)
}

type eboCtxKey struct{}

type ExponentialBackoffHandler struct {
	mx             sync.RWMutex
	bo             *backoff.ExponentialBackOff
	handlePanic    bool
	handleOverlap  bool
	handleDeferred bool
}

func (ebh *ExponentialBackoffHandler) HandlePanic(val bool) {
	ebh.handlePanic = val
}

func (ebh *ExponentialBackoffHandler) HandleOverlap(val bool) {
	ebh.handleOverlap = val
}
func (ebh *ExponentialBackoffHandler) HandleDeferred(val bool) {
	ebh.handleDeferred = val
}

func (ebh *ExponentialBackoffHandler) shouldHandleState(state State) bool {
	if ebh.handlePanic && state == PANICED {
		return true
	}
	if ebh.handleOverlap && state == OVERLAPPINGJOB {
		return true
	}
	if ebh.handleDeferred && state == DEFERRED {
		return true
	}
	return false
}

func (ebh *ExponentialBackoffHandler) getBackOffCtx(s *Schedule) *backoff.ExponentialBackOff {
	bo, ok := s.ctx.Value(eboCtxKey{}).(*backoff.ExponentialBackOff)
	if !ok {
		return nil
	}
	return bo
}

func (ebh *ExponentialBackoffHandler) setBackoffCtx(s *Schedule, bo *backoff.ExponentialBackOff) {
	s.ctx = context.WithValue(s.ctx, eboCtxKey{}, bo)
}

func (ebh *ExponentialBackoffHandler) Handler(s *Schedule, newstate State) (bool, error) {
	ebh.mx.Lock()
	defer ebh.mx.Unlock()
	bo := ebh.getBackOffCtx(s)
	if newstate == NEW {
		s.logger.Infow("Creating New Exponential Backoff for Schedule")
		if bo == nil {
			bo := backoff.NewExponentialBackOff()
			bo.InitialInterval = ebh.bo.InitialInterval
			bo.MaxElapsedTime = ebh.bo.MaxElapsedTime
			bo.MaxInterval = ebh.bo.MaxInterval
			bo.Multiplier = ebh.bo.Multiplier
			bo.RandomizationFactor = ebh.bo.RandomizationFactor
			ebh.setBackoffCtx(s, bo)
		}
	}
	if newstate == COMPLETED {
		bo.Reset()
		s.RetryJob(0)
		return false, nil
	}
	s.logger.Infow("JOb State: ", "state", newstate.String())
	if ebh.shouldHandleState(newstate) {
		next := bo.NextBackOff()
		if next != bo.Stop {
			if s.RetryJob(next) {
				s.logger.Infow(fmt.Sprintf("Exponential BO Handler Retrying %s Job in %s", newstate.String(), next))
			}
		} else {
			return false, fmt.Errorf("max Elapsed Time Exceeded")
		}
	}
	return false, nil
}

func NewDefaultExponentialBackoffMW() *ExponentialBackoffHandler {
	val := NewExponentialBackoffMW(backoff.NewExponentialBackOff())
	return val
}

func NewExponentialBackoffMW(ebo *backoff.ExponentialBackOff) *ExponentialBackoffHandler {
	val := ExponentialBackoffHandler{
		bo:             ebo,
		handlePanic:    true,
		handleOverlap:  true,
		handleDeferred: true,
	}
	return &val
}

type constantCtxKey struct{}

type ConstantBackoffHandler struct {
	mx             sync.RWMutex
	interval       time.Duration
	handlePanic    bool
	handleOverlap  bool
	handleDeferred bool
}

func (ebh *ConstantBackoffHandler) HandlePanic(val bool) {
	ebh.handlePanic = val
}

func (ebh *ConstantBackoffHandler) HandleOverlap(val bool) {
	ebh.handleOverlap = val
}
func (ebh *ConstantBackoffHandler) HandleDeferred(val bool) {
	ebh.handleDeferred = val
}

func (ebh *ConstantBackoffHandler) shouldHandleState(state State) bool {
	if ebh.handlePanic && state == PANICED {
		return true
	}
	if ebh.handleOverlap && state == OVERLAPPINGJOB {
		return true
	}
	if ebh.handleDeferred && state == DEFERRED {
		return true
	}
	return false
}

func (ebh *ConstantBackoffHandler) getBackOffCtx(s *Schedule) *backoff.ConstantBackOff {
	bo, ok := s.ctx.Value(constantCtxKey{}).(*backoff.ConstantBackOff)
	if !ok {
		return nil
	}
	return bo
}

func (ebh *ConstantBackoffHandler) setBackoffCtx(s *Schedule, bo *backoff.ConstantBackOff) {
	s.ctx = context.WithValue(s.ctx, constantCtxKey{}, bo)
}

func (ebh *ConstantBackoffHandler) Handler(s *Schedule, newstate State) (bool, error) {
	ebh.mx.Lock()
	defer ebh.mx.Unlock()
	bo := ebh.getBackOffCtx(s)
	if newstate == NEW {
		s.logger.Infow("Creating New Constant Backoff for Schedule")
		if bo == nil {
			bo := backoff.NewConstantBackOff(ebh.interval)
			ebh.setBackoffCtx(s, bo)
		}
	}
	if newstate == COMPLETED {
		bo.Reset()
		s.RetryJob(0)
		return false, nil
	}
	if ebh.shouldHandleState(newstate) {
		next := bo.NextBackOff()
		if next != backoff.Stop {
			if s.RetryJob(next) {
				s.logger.Infow(fmt.Sprintf("Constant BO Handler Retrying %s Job in %s", newstate.String(), next))
			}
		} else {
			return false, fmt.Errorf("max Elapsed Time Exceeded")
		}
	}
	return false, nil
}

func NewDefaultConstandBackoffMW() *ConstantBackoffHandler {
	val := NewConstandBackoffMW(1 * time.Second)
	return val
}

func NewConstandBackoffMW(interval time.Duration) *ConstantBackoffHandler {
	val := ConstantBackoffHandler{
		interval:       interval,
		handlePanic:    true,
		handleOverlap:  true,
		handleDeferred: true,
	}
	return &val
}

type hasTagCtxKey struct{}

type HasTagHandler struct {
	mx             sync.RWMutex
	wantTags       map[string]interface{}
	haveTags       map[string]interface{}
	handlePanic    bool
	handleOverlap  bool
	handleDeferred bool
}

func (hth *HasTagHandler) HandlePanic(val bool) {
	hth.handlePanic = val
}

func (hth *HasTagHandler) HandleOverlap(val bool) {
	hth.handleOverlap = val
}
func (hth *HasTagHandler) HandleDeferred(val bool) {
	hth.handleDeferred = val
}

func (hth *HasTagHandler) SetHaveTags(tag string) {
	hth.mx.Lock()
	defer hth.mx.Unlock()
	hth.haveTags[tag] = nil
}

func (hth *HasTagHandler) DelHaveTags(tag string) {
	hth.mx.Lock()
	defer hth.mx.Unlock()
	delete(hth.haveTags, tag)
}

func (hth *HasTagHandler) IsHaveTag(tag string) bool {
	hth.mx.RLock()
	defer hth.mx.RUnlock()
	_, ok := hth.haveTags[tag]
	return ok
}

func (hth *HasTagHandler) SetWantTags(tag string) {
	hth.mx.Lock()
	defer hth.mx.Unlock()
	hth.wantTags[tag] = nil
}

func (hth *HasTagHandler) DelWantTags(tag string) {
	hth.mx.Lock()
	defer hth.mx.Unlock()
	delete(hth.wantTags, tag)
}

func (hth *HasTagHandler) IsWantTag(tag string) bool {
	hth.mx.RLock()
	defer hth.mx.RUnlock()
	_, ok := hth.wantTags[tag]
	return ok
}

func (hth *HasTagHandler) shouldHandleState(state State) bool {
	if hth.handlePanic && state == PANICED {
		return true
	}
	if hth.handleOverlap && state == OVERLAPPINGJOB {
		return true
	}
	if hth.handleDeferred && state == DEFERRED {
		return true
	}
	return false
}

func (hth *HasTagHandler) getTagCtx(s *Schedule) *backoff.ConstantBackOff {
	bo, ok := s.ctx.Value(hasTagCtxKey{}).(*backoff.ConstantBackOff)
	if !ok {
		return nil
	}
	return bo
}

func (hth *HasTagHandler) setTagCtx(s *Schedule, bo *backoff.ConstantBackOff) {
	s.ctx = context.WithValue(s.ctx, hasTagCtxKey{}, bo)
}

func (hth *HasTagHandler) Handler(s *Schedule, newstate State) (bool, error) {
	if newstate == DISPATCHED  {
		s.logger.Infow("Checking Tags", "have", hth.haveTags, "Want", hth.wantTags)
		for k := range hth.wantTags {
			if !(hth.IsHaveTag(k)) {
				s.logger.Infow("Missing Tag", "tag", k)
				return true, nil
			}
		}
	}
	return false, nil
}

func NewTagHandlerMW() *HasTagHandler {
	val := HasTagHandler{
		wantTags:       make(map[string]interface{}),
		haveTags:       make(map[string]interface{}),
		handlePanic:    true,
		handleOverlap:  true,
		handleDeferred: true,
	}
	return &val
}
