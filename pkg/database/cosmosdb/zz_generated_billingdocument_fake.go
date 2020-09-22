// Code generated by github.com/jim-minter/go-cosmosdb, DO NOT EDIT.

package cosmosdb

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"sync"

	"github.com/ugorji/go/codec"

	pkg "github.com/Azure/ARO-RP/pkg/api"
)

type fakeBillingDocumentTrigger func(context.Context, *pkg.BillingDocument) error
type fakeBillingDocumentQuery func(BillingDocumentClient, *Query, *Options) BillingDocumentRawIterator

var _ BillingDocumentClient = &FakeBillingDocumentClient{}

func NewFakeBillingDocumentClient(h *codec.JsonHandle) *FakeBillingDocumentClient {
	return &FakeBillingDocumentClient{
		docs:              make(map[string][]byte),
		triggers:          make(map[string]fakeBillingDocumentTrigger),
		queries:           make(map[string]fakeBillingDocumentQuery),
		jsonHandle:        h,
		lock:              &sync.RWMutex{},
		sorter:            func(in []*pkg.BillingDocument) {},
		checkDocsConflict: func(*pkg.BillingDocument, *pkg.BillingDocument) bool { return false },
	}
}

type FakeBillingDocumentClient struct {
	docs       map[string][]byte
	jsonHandle *codec.JsonHandle
	lock       *sync.RWMutex
	triggers   map[string]fakeBillingDocumentTrigger
	queries    map[string]fakeBillingDocumentQuery
	sorter     func([]*pkg.BillingDocument)

	// returns true if documents conflict
	checkDocsConflict func(*pkg.BillingDocument, *pkg.BillingDocument) bool

	// unavailable, if not nil, is an error to throw when attempting to
	// communicate with this Client
	unavailable error
}

func (c *FakeBillingDocumentClient) decodeBillingDocument(s []byte) (*pkg.BillingDocument, error) {
	res := &pkg.BillingDocument{}
	err := codec.NewDecoderBytes(s, c.jsonHandle).Decode(&res)
	return res, err
}

func (c *FakeBillingDocumentClient) encodeBillingDocument(doc *pkg.BillingDocument) ([]byte, error) {
	res := make([]byte, 0)
	err := codec.NewEncoderBytes(&res, c.jsonHandle).Encode(doc)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (c *FakeBillingDocumentClient) MakeUnavailable(err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.unavailable = err
}

func (c *FakeBillingDocumentClient) UseSorter(sorter func([]*pkg.BillingDocument)) {
	c.sorter = sorter
}

func (c *FakeBillingDocumentClient) UseDocumentConflictChecker(checker func(*pkg.BillingDocument, *pkg.BillingDocument) bool) {
	c.checkDocsConflict = checker
}

func (c *FakeBillingDocumentClient) InjectTrigger(trigger string, impl fakeBillingDocumentTrigger) {
	c.triggers[trigger] = impl
}

func (c *FakeBillingDocumentClient) InjectQuery(query string, impl fakeBillingDocumentQuery) {
	c.queries[query] = impl
}

func (c *FakeBillingDocumentClient) encodeAndCopy(doc *pkg.BillingDocument) (*pkg.BillingDocument, []byte, error) {
	encoded, err := c.encodeBillingDocument(doc)
	if err != nil {
		return nil, nil, err
	}
	res, err := c.decodeBillingDocument(encoded)
	if err != nil {
		return nil, nil, err
	}
	return res, encoded, err
}

func (c *FakeBillingDocumentClient) apply(ctx context.Context, partitionkey string, doc *pkg.BillingDocument, options *Options, isNew bool) (*pkg.BillingDocument, error) {
	var docExists bool
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

	for _, ext := range c.docs {
		dec, err := c.decodeBillingDocument(ext)
		if err != nil {
			return nil, err
		}

		if dec.ID == res.ID {
			// If the document exists in the database, we want to error out in a
			// create but mark the document as extant so it can be replaced if
			// it is an update
			if isNew {
				return nil, &Error{
					StatusCode: http.StatusConflict,
					Message:    "Entity with the specified id already exists in the system",
				}
			} else {
				docExists = true
			}
		} else {
			if c.checkDocsConflict(dec, res) {
				return nil, &Error{
					StatusCode: http.StatusConflict,
					Message:    "Entity with the specified id already exists in the system",
				}
			}
		}
	}

	if !isNew && !docExists {
		return nil, &Error{StatusCode: http.StatusNotFound}
	}

	c.docs[doc.ID] = enc
	return res, nil
}

func (c *FakeBillingDocumentClient) Create(ctx context.Context, partitionkey string, doc *pkg.BillingDocument, options *Options) (*pkg.BillingDocument, error) {
	if c.unavailable != nil {
		return nil, c.unavailable
	}
	return c.apply(ctx, partitionkey, doc, options, true)
}

func (c *FakeBillingDocumentClient) Replace(ctx context.Context, partitionkey string, doc *pkg.BillingDocument, options *Options) (*pkg.BillingDocument, error) {
	if c.unavailable != nil {
		return nil, c.unavailable
	}
	return c.apply(ctx, partitionkey, doc, options, false)
}

func (c *FakeBillingDocumentClient) List(*Options) BillingDocumentIterator {
	if c.unavailable != nil {
		return NewFakeBillingDocumentClientErroringRawIterator(c.unavailable)
	}
	c.lock.RLock()
	defer c.lock.RUnlock()

	docs := make([]*pkg.BillingDocument, 0, len(c.docs))
	for _, d := range c.docs {
		r, err := c.decodeBillingDocument(d)
		if err != nil {
			return NewFakeBillingDocumentClientErroringRawIterator(err)
		}
		docs = append(docs, r)
	}
	c.sorter(docs)
	return NewFakeBillingDocumentClientRawIterator(docs, 0)
}

func (c *FakeBillingDocumentClient) ListAll(ctx context.Context, opts *Options) (*pkg.BillingDocuments, error) {
	if c.unavailable != nil {
		return nil, c.unavailable
	}
	iter := c.List(opts)
	billingDocuments, err := iter.Next(ctx, -1)
	if err != nil {
		return nil, err
	}
	return billingDocuments, nil
}

func (c *FakeBillingDocumentClient) Get(ctx context.Context, partitionkey string, documentId string, options *Options) (*pkg.BillingDocument, error) {
	if c.unavailable != nil {
		return nil, c.unavailable
	}
	c.lock.RLock()
	defer c.lock.RUnlock()

	out, ext := c.docs[documentId]
	if !ext {
		return nil, &Error{StatusCode: http.StatusNotFound}
	}
	return c.decodeBillingDocument(out)
}

func (c *FakeBillingDocumentClient) Delete(ctx context.Context, partitionKey string, doc *pkg.BillingDocument, options *Options) error {
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

func (c *FakeBillingDocumentClient) ChangeFeed(*Options) BillingDocumentIterator {
	if c.unavailable != nil {
		return NewFakeBillingDocumentClientErroringRawIterator(c.unavailable)
	}
	return NewFakeBillingDocumentClientErroringRawIterator(ErrNotImplemented)
}

func (c *FakeBillingDocumentClient) processPreTriggers(ctx context.Context, doc *pkg.BillingDocument, options *Options) error {
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

func (c *FakeBillingDocumentClient) Query(name string, query *Query, options *Options) BillingDocumentRawIterator {
	if c.unavailable != nil {
		return NewFakeBillingDocumentClientErroringRawIterator(c.unavailable)
	}
	c.lock.RLock()
	defer c.lock.RUnlock()

	quer, ok := c.queries[query.Query]
	if ok {
		return quer(c, query, options)
	} else {
		return NewFakeBillingDocumentClientErroringRawIterator(ErrNotImplemented)
	}
}

func (c *FakeBillingDocumentClient) QueryAll(ctx context.Context, partitionkey string, query *Query, options *Options) (*pkg.BillingDocuments, error) {
	iter := c.Query("", query, options)
	return iter.Next(ctx, -1)
}

// NewFakeBillingDocumentClientRawIterator creates a RawIterator that will produce only
// BillingDocuments from Next() and NextRaw().
func NewFakeBillingDocumentClientRawIterator(docs []*pkg.BillingDocument, continuation int) BillingDocumentRawIterator {
	return &fakeBillingDocumentClientRawIterator{docs: docs, continuation: continuation}
}

type fakeBillingDocumentClientRawIterator struct {
	docs         []*pkg.BillingDocument
	continuation int
	done         bool
}

func (i *fakeBillingDocumentClientRawIterator) Next(ctx context.Context, maxItemCount int) (out *pkg.BillingDocuments, err error) {
	err = i.NextRaw(ctx, maxItemCount, &out)
	return
}

func (i *fakeBillingDocumentClientRawIterator) NextRaw(ctx context.Context, maxItemCount int, out interface{}) error {
	if i.done {
		return nil
	}

	var docs []*pkg.BillingDocument
	if maxItemCount == -1 {
		docs = i.docs[i.continuation:]
		i.continuation = len(i.docs)
		i.done = true
	} else {
		max := i.continuation + maxItemCount
		if max > len(i.docs) {
			max = len(i.docs)
		}
		docs = i.docs[i.continuation:max]
		i.continuation += max
		i.done = i.Continuation() == ""
	}

	y := reflect.ValueOf(out)
	d := &pkg.BillingDocuments{}
	d.BillingDocuments = docs
	d.Count = len(d.BillingDocuments)
	y.Elem().Set(reflect.ValueOf(d))
	return nil
}

func (i *fakeBillingDocumentClientRawIterator) Continuation() string {
	if i.continuation >= len(i.docs) {
		return ""
	}
	return fmt.Sprintf("%d", i.continuation)
}

// fakeBillingDocumentErroringRawIterator is a RawIterator that will return an error on use.
func NewFakeBillingDocumentClientErroringRawIterator(err error) *fakeBillingDocumentErroringRawIterator {
	return &fakeBillingDocumentErroringRawIterator{err: err}
}

type fakeBillingDocumentErroringRawIterator struct {
	err error
}

func (i *fakeBillingDocumentErroringRawIterator) Next(ctx context.Context, maxItemCount int) (*pkg.BillingDocuments, error) {
	return nil, i.err
}

func (i *fakeBillingDocumentErroringRawIterator) NextRaw(context.Context, int, interface{}) error {
	return i.err
}

func (i *fakeBillingDocumentErroringRawIterator) Continuation() string {
	return ""
}
