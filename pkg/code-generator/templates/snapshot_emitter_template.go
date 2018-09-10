package templates

import (
	"text/template"
)

var ResourceGroupEmitterTemplate = template.Must(template.New("resource_group_emitter").Funcs(funcs).Parse(`package {{ .Project.PackageName }}

{{- $client_declarations := new_str_slice }}
{{- $clients := new_str_slice }}
{{- range .Resources}}
{{- $client_declarations := (append_str_slice $client_declarations (printf "%vClient %vClient"  (lower_camel .Name) .Name)) }}
{{- $clients := (append_str_slice $clients (printf "%vClient"  (lower_camel .Name))) }}
{{- end}}
{{- $client_declarations := (join_str_slice $client_declarations ", ") }}
{{- $clients := (join_str_slice $clients ", ") }}

import (
	"sync"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
)

type {{ .GoName }}Emitter interface {
	Register() error
{{- range .Resources}}
	{{ .Name }}() {{ .Name }}Client
{{- end}}
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *{{ .GoName }}Snapshot, <-chan error, error)
}

func New{{ .GoName }}Emitter({{ $client_declarations }}) {{ .GoName }}Emitter {
	return New{{ .GoName }}EmitterWithEmit({{ $clients }}, make(chan struct{}))
}

func New{{ .GoName }}EmitterWithEmit({{ $client_declarations }}, emit <-chan struct{}) {{ .GoName }}Emitter {
	return &{{ lower_camel .GoName }}Emitter{
{{- range .Resources}}
		{{ lower_camel .Name }}: {{ lower_camel .Name }}Client,
{{- end}}
		forceEmit: emit,
	}
}

type {{ lower_camel .GoName }}Emitter struct {
	forceEmit <- chan struct{}
{{- range .Resources}}
	{{ lower_camel .Name }} {{ .Name }}Client
{{- end}}
}

func (c *{{ lower_camel .GoName }}Emitter) Register() error {
{{- range .Resources}}
	if err := c.{{ lower_camel .Name }}.Register(); err != nil {
		return err
	}
{{- end}}
	return nil
}

{{- range .Resources}}

func (c *{{ lower_camel $.GoName }}Emitter) {{ .Name }}() {{ .Name }}Client {
	return c.{{ lower_camel .Name }}
}
{{- end}}

func (c *{{ lower_camel .GoName }}Emitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *{{ .GoName }}Snapshot, <-chan error, error) {
	errs := make(chan error)
	var done sync.WaitGroup


{{- range .Resources}}
	/* Create channel for {{ .Name }} */
	type {{ lower_camel .Name }}ListWithNamespace struct {
		list {{ .Name }}List
		namespace string
	}
	{{ lower_camel .Name }}Chan := make(chan {{ lower_camel .Name }}ListWithNamespace)
{{- end}}

	for _, namespace := range watchNamespaces {
{{- range .Resources}}
		/* Setup watch for {{ .Name }} */
		{{ lower_camel .Name }}NamespacesChan, {{ lower_camel .Name }}Errs, err := c.{{ lower_camel .Name }}.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting {{ .Name }} watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(opts.Ctx, errs, {{ lower_camel .Name }}Errs, namespace+"-{{ lower_camel .PluralName }}")
		}(namespace)

{{- end}}


		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-opts.Ctx.Done():
					return
{{- range .Resources}}
				case {{ lower_camel .Name }}List := <- {{ lower_camel .Name }}NamespacesChan:
					select {
					case <-opts.Ctx.Done():
						return
					case {{ lower_camel .Name }}Chan <- {{ lower_camel .Name }}ListWithNamespace{list:{{ lower_camel .Name }}List, namespace:namespace}:
					}
{{- end}}
				}
			}
		}(namespace)
	}

	
	snapshots := make(chan *{{ .GoName }}Snapshot)
	go func() {
		currentSnapshot := {{ .GoName }}Snapshot{}
		sync := func(newSnapshot {{ .GoName }}Snapshot) {
			if currentSnapshot.Hash() == newSnapshot.Hash() {
				return
			}
			currentSnapshot = newSnapshot
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}
		for {
			select {
			case <-opts.Ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
{{- range .Resources}}
			case {{ lower_camel .Name }}NamespacedList := <- {{ lower_camel .Name }}Chan:
				namespace := {{ lower_camel .Name }}NamespacedList.namespace
				{{ lower_camel .Name }}List := {{ lower_camel .Name }}NamespacedList.list

				newSnapshot := currentSnapshot.Clone()
				newSnapshot.{{ .PluralName }}.Clear(namespace)
				newSnapshot.{{ .PluralName }}.Add({{ lower_camel .Name }}List...)
				sync(newSnapshot)
{{- end}}
			}
		}
	}()
	return snapshots, errs, nil
}
`))
