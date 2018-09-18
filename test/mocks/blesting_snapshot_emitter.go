// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package mocks

import (
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
)

var (
	mBlestingSnapshotIn  = stats.Int64("blesting_snap_emitter/snap_in", "The number of snapshots in", "1")
	mBlestingSnapshotOut = stats.Int64("blesting_snap_emitter/snap_out", "The number of snapshots out", "1")

	blestingsnapshotInView = &view.View{
		Name:        "blesting_snap_emitter/snap_in",
		Measure:     mBlestingSnapshotIn,
		Description: "The number of snapshots updates coming in",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	blestingsnapshotOutView = &view.View{
		Name:        "blesting_snap_emitter/snap_out",
		Measure:     mBlestingSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(blestingsnapshotInView, blestingsnapshotOutView)
}

type BlestingEmitter interface {
	Register() error
	FakeResource() FakeResourceClient
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *BlestingSnapshot, <-chan error, error)
}

func NewBlestingEmitter(fakeResourceClient FakeResourceClient) BlestingEmitter {
	return NewBlestingEmitterWithEmit(fakeResourceClient, make(chan struct{}))
}

func NewBlestingEmitterWithEmit(fakeResourceClient FakeResourceClient, emit <-chan struct{}) BlestingEmitter {
	return &blestingEmitter{
		fakeResource: fakeResourceClient,
		forceEmit:    emit,
	}
}

type blestingEmitter struct {
	forceEmit    <-chan struct{}
	fakeResource FakeResourceClient
}

func (c *blestingEmitter) Register() error {
	if err := c.fakeResource.Register(); err != nil {
		return err
	}
	return nil
}

func (c *blestingEmitter) FakeResource() FakeResourceClient {
	return c.fakeResource
}

func (c *blestingEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *BlestingSnapshot, <-chan error, error) {
	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for FakeResource */
	type fakeResourceListWithNamespace struct {
		list      FakeResourceList
		namespace string
	}
	fakeResourceChan := make(chan fakeResourceListWithNamespace)

	for _, namespace := range watchNamespaces {
		/* Setup watch for FakeResource */
		fakeResourceNamespacesChan, fakeResourceErrs, err := c.fakeResource.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting FakeResource watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, fakeResourceErrs, namespace+"-fakes")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case fakeResourceList := <-fakeResourceNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case fakeResourceChan <- fakeResourceListWithNamespace{list: fakeResourceList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}

	snapshots := make(chan *BlestingSnapshot)
	go func() {
		originalSnapshot := BlestingSnapshot{}
		currentSnapshot := originalSnapshot.Clone()
		timer := time.NewTicker(time.Second * 5)
		sync := func() {
			if originalSnapshot.Hash() == currentSnapshot.Hash() {
				return
			}
			originalSnapshot = currentSnapshot.Clone()
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}

		/* TODO (yuval-k): figure out how to make this work to avoid a stale snapshot.
		   		// construct the first snapshot from all the configs that are currently there
		   		// that guarantees that the first snapshot contains all the data.
		   		for range watchNamespaces {
		      fakeResourceNamespacedList := <- fakeResourceChan:
		   	namespace := fakeResourceNamespacedList.namespace
		   	fakeResourceList := fakeResourceNamespacedList.list

		   	currentSnapshot.Fakes.Clear(namespace)
		   	currentSnapshot.Fakes.Add(fakeResourceList...)
		   		}
		*/

		for {
			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case fakeResourceNamespacedList := <-fakeResourceChan:
				namespace := fakeResourceNamespacedList.namespace
				fakeResourceList := fakeResourceNamespacedList.list

				currentSnapshot.Fakes.Clear(namespace)
				currentSnapshot.Fakes.Add(fakeResourceList...)
			}

			// if we got here its because a new entry in the channel
			stats.Record(ctx, mBlestingSnapshotIn.M(1))
		}
	}()
	return snapshots, errs, nil
}
