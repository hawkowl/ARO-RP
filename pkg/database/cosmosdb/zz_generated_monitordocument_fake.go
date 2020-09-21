// Code generated by github.com/jim-minter/go-cosmosdb, DO NOT EDIT.

package cosmosdb

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/ugorji/go/codec"

	pkg "github.com/Azure/ARO-RP/pkg/api"
)

type FakeMonitorDocumentTrigger func(context.Context, *pkg.MonitorDocument) error
type FakeMonitorDocumentQuery func(MonitorDocumentClient, *Query, *Options) MonitorDocumentRawIterator

var _ MonitorDocumentClient = &FakeMonitorDocumentClient{}

func NewFakeMonitorDocumentClient(h *codec.JsonHandle, uniqueKeys []string) *FakeMonitorDocumentClient {
	return &FakeMonitorDocumentClient{
		docs:       make(map[string][]byte),
		triggers:   make(map[string]FakeMonitorDocumentTrigger),
		queries:    make(map[string]FakeMonitorDocumentQuery),
		uniqueKeys: uniqueKeys,
		jsonHandle: h,
		lock:       &sync.RWMutex{},
		sorter:     func(in []*pkg.MonitorDocument) {},
	}
}

type FakeMonitorDocumentClient struct {
	docs       map[string][]byte
	jsonHandle *codec.JsonHandle
	lock       *sync.RWMutex
	triggers   map[string]FakeMonitorDocumentTrigger
	queries    map[string]FakeMonitorDocumentQuery
	uniqueKeys []string
	sorter     func([]*pkg.MonitorDocument)

	// unavailable, if not nil, is an error to throw when attempting to
	// communicate with this Client
	unavailable error
}

func decodeMonitorDocument(s []byte, handle *codec.JsonHandle) (*pkg.MonitorDocument, error) {
	res := &pkg.MonitorDocument{}
	err := codec.NewDecoder(bytes.NewBuffer(s), handle).Decode(&res)
	return res, err
}

func decodeMonitorDocumentToMap(s []byte, handle *codec.JsonHandle) (map[interface{}]interface{}, error) {
	var res interface{}
	err := codec.NewDecoder(bytes.NewBuffer(s), handle).Decode(&res)
	if err != nil {
		return nil, err
	}
	ret, ok := res.(map[interface{}]interface{})
	if !ok {
		return nil, errors.New("Could not coerce")
	}
	return ret, err
}

func encodeMonitorDocument(doc *pkg.MonitorDocument, handle *codec.JsonHandle) (res []byte, err error) {
	buf := &bytes.Buffer{}
	err = codec.NewEncoder(buf, handle).Encode(doc)
	if err != nil {
		return
	}
	res = buf.Bytes()
	return
}

func (c *FakeMonitorDocumentClient) MakeUnavailable(err error) {
	c.unavailable = err
}

func (c *FakeMonitorDocumentClient) UseSorter(sorter func([]*pkg.MonitorDocument)) {
	c.sorter = sorter
}

func (c *FakeMonitorDocumentClient) encodeAndCopy(doc *pkg.MonitorDocument) (*pkg.MonitorDocument, []byte, error) {
	encoded, err := encodeMonitorDocument(doc, c.jsonHandle)
	if err != nil {
		return nil, nil, err
	}
	res, err := decodeMonitorDocument(encoded, c.jsonHandle)
	if err != nil {
		return nil, nil, err
	}
	return res, encoded, err
}

func (c *FakeMonitorDocumentClient) Create(ctx context.Context, partitionkey string, doc *pkg.MonitorDocument, options *Options) (*pkg.MonitorDocument, error) {
	if c.unavailable != nil {
		return nil, c.unavailable
	}
	c.lock.Lock()
	defer c.lock.Unlock()

	if options != nil {
		err := c.processPreTriggers(ctx, doc, options)
		if err != nil {
			return nil, err
		}
	}

	res, enc, err := c.encodeAndCopy(doc)
	if err != nil {
		return nil, err
	}
	docAsMap, err := decodeMonitorDocumentToMap(enc, c.jsonHandle)
	if err != nil {
		return nil, err
	}

	for _, ext := range c.docs {
		extDecoded, err := decodeMonitorDocumentToMap(ext, c.jsonHandle)
		if err != nil {
			return nil, err
		}

		for _, key := range c.uniqueKeys {
			var ourKeyStr string
			var theirKeyStr string
			ourKey, ourKeyOk := docAsMap[key]
			if ourKeyOk {
				ourKeyStr, ourKeyOk = ourKey.(string)
			}
			theirKey, theirKeyOk := extDecoded[key]
			if theirKeyOk {
				theirKeyStr, theirKeyOk = theirKey.(string)
			}
			if ourKeyOk && theirKeyOk && ourKeyStr != "" && ourKeyStr == theirKeyStr {
				return nil, &Error{
					StatusCode: http.StatusConflict,
					Message:    "Entity with the specified id already exists in the system",
				}
			}
		}
	}

	c.docs[doc.ID] = enc
	return res, nil
}

func (c *FakeMonitorDocumentClient) List(*Options) MonitorDocumentIterator {
	if c.unavailable != nil {
		return NewFakeMonitorDocumentClientErroringRawIterator(c.unavailable)
	}
	c.lock.RLock()
	defer c.lock.RUnlock()

	docs := make([]*pkg.MonitorDocument, 0, len(c.docs))
	for _, d := range c.docs {
		r, err := decodeMonitorDocument(d, c.jsonHandle)
		if err != nil {
			return NewFakeMonitorDocumentClientErroringRawIterator(err)
		}
		docs = append(docs, r)
	}
	c.sorter(docs)
	return NewFakeMonitorDocumentClientRawIterator(docs, 0)
}

func (c *FakeMonitorDocumentClient) ListAll(context.Context, *Options) (*pkg.MonitorDocuments, error) {
	if c.unavailable != nil {
		return nil, c.unavailable
	}
	c.lock.RLock()
	defer c.lock.RUnlock()

	monitorDocuments := &pkg.MonitorDocuments{
		Count:            len(c.docs),
		MonitorDocuments: make([]*pkg.MonitorDocument, 0, len(c.docs)),
	}

	for _, d := range c.docs {
		dec, err := decodeMonitorDocument(d, c.jsonHandle)
		if err != nil {
			return nil, err
		}
		monitorDocuments.MonitorDocuments = append(monitorDocuments.MonitorDocuments, dec)
	}
	c.sorter(monitorDocuments.MonitorDocuments)
	return monitorDocuments, nil
}

func (c *FakeMonitorDocumentClient) Get(ctx context.Context, partitionkey string, documentId string, options *Options) (*pkg.MonitorDocument, error) {
	if c.unavailable != nil {
		return nil, c.unavailable
	}
	c.lock.RLock()
	defer c.lock.RUnlock()

	out, ext := c.docs[documentId]
	if !ext {
		return nil, &Error{StatusCode: http.StatusNotFound}
	}
	return decodeMonitorDocument(out, c.jsonHandle)
}

func (c *FakeMonitorDocumentClient) Replace(ctx context.Context, partitionkey string, doc *pkg.MonitorDocument, options *Options) (*pkg.MonitorDocument, error) {
	if c.unavailable != nil {
		return nil, c.unavailable
	}
	c.lock.Lock()
	defer c.lock.Unlock()

	_, exists := c.docs[doc.ID]
	if !exists {
		return nil, &Error{StatusCode: http.StatusNotFound}
	}

	if options != nil {
		err := c.processPreTriggers(ctx, doc, options)
		if err != nil {
			return nil, err
		}
	}

	res, enc, err := c.encodeAndCopy(doc)
	if err != nil {
		return nil, err
	}
	c.docs[doc.ID] = enc
	return res, nil
}

func (c *FakeMonitorDocumentClient) Delete(ctx context.Context, partitionKey string, doc *pkg.MonitorDocument, options *Options) error {
	if c.unavailable != nil {
		return c.unavailable
	}
	c.lock.Lock()
	defer c.lock.Unlock()

	_, ext := c.docs[doc.ID]
	if !ext {
		return &Error{StatusCode: http.StatusNotFound}
	}

	delete(c.docs, doc.ID)
	return nil
}

func (c *FakeMonitorDocumentClient) ChangeFeed(*Options) MonitorDocumentIterator {
	if c.unavailable != nil {
		return NewFakeMonitorDocumentClientErroringRawIterator(c.unavailable)
	}
	return NewFakeMonitorDocumentClientErroringRawIterator(ErrNotImplemented)
}

func (c *FakeMonitorDocumentClient) processPreTriggers(ctx context.Context, doc *pkg.MonitorDocument, options *Options) error {
	for _, trigger := range options.PreTriggers {
		trig, ok := c.triggers[trigger]
		if ok {
			err := trig(ctx, doc)
			if err != nil {
				return err
			}
		} else {
			return ErrNotImplemented
		}
	}
	return nil
}

func (c *FakeMonitorDocumentClient) Query(name string, query *Query, options *Options) MonitorDocumentRawIterator {
	if c.unavailable != nil {
		return NewFakeMonitorDocumentClientErroringRawIterator(c.unavailable)
	}
	c.lock.RLock()
	defer c.lock.RUnlock()

	quer, ok := c.queries[query.Query]
	if ok {
		return quer(c, query, options)
	} else {
		return NewFakeMonitorDocumentClientErroringRawIterator(ErrNotImplemented)
	}
}

func (c *FakeMonitorDocumentClient) QueryAll(ctx context.Context, partitionkey string, query *Query, options *Options) (*pkg.MonitorDocuments, error) {
	if c.unavailable != nil {
		return nil, c.unavailable
	}
	c.lock.RLock()
	defer c.lock.RUnlock()

	quer, ok := c.queries[query.Query]
	if ok {
		items := quer(c, query, options)
		res := &pkg.MonitorDocuments{}
		err := items.NextRaw(ctx, -1, res)
		return res, err
	} else {
		return nil, ErrNotImplemented
	}
}

func (c *FakeMonitorDocumentClient) InjectTrigger(trigger string, impl FakeMonitorDocumentTrigger) {
	c.triggers[trigger] = impl
}

func (c *FakeMonitorDocumentClient) InjectQuery(query string, impl FakeMonitorDocumentQuery) {
	c.queries[query] = impl
}

// NewFakeMonitorDocumentClientRawIterator creates a RawIterator that will produce only
// MonitorDocuments from Next() and NextRaw().
func NewFakeMonitorDocumentClientRawIterator(docs []*pkg.MonitorDocument, continuation int) MonitorDocumentRawIterator {
	return &fakeMonitorDocumentClientRawIterator{docs: docs, continuation: continuation}
}

type fakeMonitorDocumentClientRawIterator struct {
	docs         []*pkg.MonitorDocument
	continuation int
}

func (i *fakeMonitorDocumentClientRawIterator) Next(ctx context.Context, maxItemCount int) (*pkg.MonitorDocuments, error) {
	out := &pkg.MonitorDocuments{}
	err := i.NextRaw(ctx, maxItemCount, out)

	if out.Count == 0 {
		return nil, nil
	}
	return out, err
}

func (i *fakeMonitorDocumentClientRawIterator) NextRaw(ctx context.Context, maxItemCount int, out interface{}) error {
	if i.continuation >= len(i.docs) {
		return nil
	}

	var docs []*pkg.MonitorDocument
	if maxItemCount == -1 {
		docs = i.docs[i.continuation:]
		i.continuation = len(i.docs)
	} else {
		max := i.continuation + maxItemCount
		if max > len(i.docs) {
			max = len(i.docs)
		}
		docs = i.docs[i.continuation:max]
		i.continuation += max
	}

	d := out.(*pkg.MonitorDocuments)
	d.MonitorDocuments = docs
	d.Count = len(d.MonitorDocuments)
	return nil
}

func (i *fakeMonitorDocumentClientRawIterator) Continuation() string {
	if i.continuation >= len(i.docs) {
		return ""
	}
	return fmt.Sprintf("%d", i.continuation)
}

// fakeMonitorDocumentErroringRawIterator is a RawIterator that will return an error on use.
func NewFakeMonitorDocumentClientErroringRawIterator(err error) *fakeMonitorDocumentErroringRawIterator {
	return &fakeMonitorDocumentErroringRawIterator{err: err}
}

type fakeMonitorDocumentErroringRawIterator struct {
	err error
}

func (i *fakeMonitorDocumentErroringRawIterator) Next(ctx context.Context, maxItemCount int) (*pkg.MonitorDocuments, error) {
	return nil, i.err
}

func (i *fakeMonitorDocumentErroringRawIterator) NextRaw(context.Context, int, interface{}) error {
	return i.err
}

func (i *fakeMonitorDocumentErroringRawIterator) Continuation() string {
	return ""
}
