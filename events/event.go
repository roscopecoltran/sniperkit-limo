package events

import (
	"fmt"

	pb "github.com/roscopecoltran/sniperkit-limo/model/generate/plugin"
	"github.com/roscopecoltran/sniperkit-limo/plugin"
)

const (
	InternalNamespace = "keds"
	PluginNamespace   = "plugin"
)

type Event struct {
	*pb.PluginEvent
	SourcePlugin *plugin.Plugin
}

type EventOption func(e *Event)

func NewEvent(event *pb.PluginEvent, src *plugin.Plugin) *Event {
	e := &Event{SourcePlugin: src}
	if event != nil {
		e.PluginEvent = event
	} else {
		e.PluginEvent = &pb.PluginEvent{}
	}
	e.Args = make([]string, 0)
	e.Data = make(map[string]string, 0)
	if e.Source == "" && e.SourcePlugin != nil {
		e.Source = e.SourcePlugin.GetName()
	}
	return e
}

func NewEventWithOptions(opts ...EventOption) *Event {
	e := NewEvent(nil, nil)
	e.ApplyOptions(opts...)
	if e.Source == "" && e.SourcePlugin != nil {
		e.Source = e.SourcePlugin.GetName()
	}
	return e
}

func (e *Event) ApplyOptions(opts ...EventOption) {
	for _, option := range opts {
		option(e)
	}
}

func (e *Event) Proto() *pb.PluginEvent {
	return e.PluginEvent
}

func CreateEventServerQuit(src *plugin.Plugin, exitCode int) (event *Event) {
	return NewEventWithOptions(WithExitCode(exitCode), WithName("keds.exit"), WithSourcePlugin(src))
}

func CreateEventCommandInvoked(target *plugin.Plugin, args []string) (event *Event) {
	return NewEventWithOptions(WithArgs(args), WithName("keds.command_invoked"), WithSource("keds"), WithTarget(target.GetName()))
}

func WithSourcePlugin(plug *plugin.Plugin) EventOption {
	return func(e *Event) {
		e.SourcePlugin = plug
	}
}

func WithName(name string) EventOption {
	return func(e *Event) {
		e.Name = name
	}
}

func WithTarget(target string) EventOption {
	return func(e *Event) {
		e.Target = target
	}
}

func WithSource(source string) EventOption {
	return func(e *Event) {
		e.Source = source
	}
}

func WithArgs(args []string) EventOption {
	return func(e *Event) {
		e.Args = args
	}
}

func WithData(data map[string]string) EventOption {
	return func(e *Event) {
		e.Data = data
	}
}

func WithExitCode(code int) EventOption {
	return WithKeyPair("exit_code", fmt.Sprintf("%d", code))
}

func WithKeyPair(key string, val string) EventOption {
	return func(e *Event) {
		e.Data[key] = val
	}
}

func WithArg(val string) EventOption {
	return func(e *Event) {
		e.Args = append(e.Args, val)
	}
}
