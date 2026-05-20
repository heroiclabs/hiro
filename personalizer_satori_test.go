// Copyright 2026 Heroic Labs & Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hiro

import (
	"context"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
)

type stubLogger struct{}

func (l *stubLogger) Debug(format string, v ...any)                   {}
func (l *stubLogger) Info(format string, v ...any)                    {}
func (l *stubLogger) Warn(format string, v ...any)                    {}
func (l *stubLogger) Error(format string, v ...any)                   {}
func (l *stubLogger) WithField(key string, v any) runtime.Logger      { return l }
func (l *stubLogger) WithFields(fields map[string]any) runtime.Logger { return l }
func (l *stubLogger) Fields() map[string]any                          { return nil }

type stubSatori struct {
	runtime.Satori
	eventsPublishCalled bool
	flagsListCalled     bool
}

func (s *stubSatori) EventsPublish(_ context.Context, _ string, _ []*runtime.Event, _ ...string) error {
	s.eventsPublishCalled = true
	return nil
}

func (s *stubSatori) FlagsList(_ context.Context, _ string, _ []string, _ []string) (*runtime.FlagList, error) {
	s.flagsListCalled = true
	return &runtime.FlagList{}, nil
}

type stubNakamaModule struct {
	runtime.NakamaModule
	satori *stubSatori
}

func (m *stubNakamaModule) GetSatori() runtime.Satori { return m.satori }

type stubSystem struct{}

func (s *stubSystem) GetType() SystemType { return SystemTypeBase }
func (s *stubSystem) GetConfig() any      { return nil }

//
// Send Tests

func TestSend_PublishesEventsWhenNotLimited(t *testing.T) {
	ctx := t.Context()
	p := NewSatoriPersonalizer(ctx, SatoriPersonalizerPublishAllEvents())

	satori := &stubSatori{}
	nk := &stubNakamaModule{satori: satori}

	p.Send(context.Background(), &stubLogger{}, nk, "user", []*PublisherEvent{{Name: "e", System: &stubSystem{}}})

	if !satori.eventsPublishCalled {
		t.Error("expected EventsPublish to be called, but it was not")
	}
}

func TestSend_SkipsEventsWhenLimitedForAnalytics(t *testing.T) {
	ctx := t.Context()
	p := NewSatoriPersonalizer(ctx, SatoriPersonalizerPublishAllEvents())

	limited := context.WithValue(context.Background(), runtime.RUNTIME_CTX_VARS, map[string]string{"LimitedForAnalytics": "true"})

	// nk is nil: if Send does not return early it will panic, failing the test.
	p.Send(limited, &stubLogger{}, nil, "user", []*PublisherEvent{{Name: "e", System: &stubSystem{}}})
}

func TestSend_PublishesEventsWhenNotLimitedKeySet(t *testing.T) {
	ctx := t.Context()
	p := NewSatoriPersonalizer(ctx, SatoriPersonalizerPublishAllEvents())

	satori := &stubSatori{}
	nk := &stubNakamaModule{satori: satori}

	limited := context.WithValue(context.Background(), runtime.RUNTIME_CTX_VARS, map[string]string{"LimitedForAnalytics": "false"})
	p.Send(limited, &stubLogger{}, nk, "user", []*PublisherEvent{{Name: "e", System: &stubSystem{}}})

	if !satori.eventsPublishCalled {
		t.Error("expected EventsPublish to be called, but it was not")
	}
}

//
// GetValue Tests

func TestGetValue_CallsSatoriWhenNotLimited(t *testing.T) {
	p := NewSatoriPersonalizer(context.Background(), SatoriPersonalizerNoCache())

	satori := &stubSatori{}
	nk := &stubNakamaModule{satori: satori}

	if _, err := p.GetValue(context.Background(), &stubLogger{}, nk, &stubSystem{}, "user"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !satori.flagsListCalled {
		t.Error("expected FlagsList to be called, but it was not")
	}
}

func TestGetValue_ReturnsNilWhenLimitedForPersonalize(t *testing.T) {
	p := NewSatoriPersonalizer(context.Background(), SatoriPersonalizerNoCache())

	// GetValue should return nil, nil here.
	limited := context.WithValue(context.Background(), runtime.RUNTIME_CTX_VARS, map[string]string{"LimitedForPersonalize": "true"})
	val, err := p.GetValue(limited, &stubLogger{}, nil, nil, "user")

	if val != nil {
		t.Errorf("expected nil value, got %v", val)
	}
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGetValue_CallsSatoriWhenNotLimitedKeySet(t *testing.T) {
	p := NewSatoriPersonalizer(context.Background(), SatoriPersonalizerNoCache())

	satori := &stubSatori{}
	nk := &stubNakamaModule{satori: satori}

	limited := context.WithValue(context.Background(), runtime.RUNTIME_CTX_VARS, map[string]string{"LimitedForPersonalize": "false"})
	if _, err := p.GetValue(limited, &stubLogger{}, nk, &stubSystem{}, "user"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !satori.flagsListCalled {
		t.Error("expected FlagsList to be called, but it was not")
	}
}
