package store

import (
	"encoding/binary"
	"errors"

	"go.etcd.io/bbolt"

	"github.com/grafana/pyroscope/pkg/iter"
)

var ErrInvalidBlockEntry = errors.New("invalid block entry")

var blockQueueBucketName = []byte("compaction_block_queue")

// BlockEntry represents a block metadata entry in the compaction block queue.
type BlockEntry struct {
	Index      uint64
	ID         string
	AppendedAt int64
	Level      uint32
	Shard      uint32
	Tenant     string
}

// BlockQueueStore provides methods to store and retrieve block queues.
// The store is optimized for two cases: load the entire queue (preserving
// the original order) and remove an entry from the queue.
//
// Compactor maintains an in-memory queue of blocks to compact, therefore
// the store never reads individual entries.
//
// NOTE(kolesnikovae): We can leverage the fact that removed entries are
// always ordered in ascending order by index and use the same cursor when
// removing entries from the database:
// DeleteEntry(*bbolt.Tx, ...store.BlockEntry) error
type BlockQueueStore struct{ bucketName []byte }

func NewBlockQueueStore() *BlockQueueStore {
	return &BlockQueueStore{bucketName: blockQueueBucketName}
}

func (s BlockQueueStore) CreateBuckets(tx *bbolt.Tx) error {
	_, err := tx.CreateBucketIfNotExists(blockQueueBucketName)
	return err
}

func (s BlockQueueStore) StoreEntry(tx *bbolt.Tx, entry BlockEntry) error {
	e := marshalBlockEntry(entry)
	return tx.Bucket(s.bucketName).Put(e.key, e.value)
}

func (s BlockQueueStore) DeleteEntry(tx *bbolt.Tx, index uint64, id string) error {
	return tx.Bucket(s.bucketName).Delete(marshalBlockEntryKey(index, id))
}

func (s BlockQueueStore) ListEntries(tx *bbolt.Tx) iter.Iterator[BlockEntry] {
	return newBlockEntriesIterator(tx.Bucket(s.bucketName))
}

type blockEntriesIterator struct {
	iter *cursorIterator
	cur  BlockEntry
	err  error
}

func newBlockEntriesIterator(bucket *bbolt.Bucket) *blockEntriesIterator {
	return &blockEntriesIterator{iter: newCursorIter(nil, bucket.Cursor())}
}

func (x *blockEntriesIterator) Next() bool {
	if x.err != nil || !x.iter.Next() {
		return false
	}
	x.err = unmarshalBlockEntry(&x.cur, x.iter.At())
	return x.err == nil
}

func (x *blockEntriesIterator) At() BlockEntry { return x.cur }

func (x *blockEntriesIterator) Close() error { return x.iter.Close() }

func (x *blockEntriesIterator) Err() error {
	if err := x.iter.Err(); err != nil {
		return err
	}
	return x.err
}

func marshalBlockEntry(e BlockEntry) kv {
	k := marshalBlockEntryKey(e.Index, e.ID)
	b := make([]byte, 8+4+4+len(e.Tenant))
	binary.BigEndian.PutUint64(b[0:8], uint64(e.AppendedAt))
	binary.BigEndian.PutUint32(b[8:12], e.Level)
	binary.BigEndian.PutUint32(b[12:16], e.Shard)
	copy(b[16:], e.Tenant)
	return kv{key: k, value: b}
}

func marshalBlockEntryKey(index uint64, id string) []byte {
	b := make([]byte, 8+len(id))
	binary.BigEndian.PutUint64(b, index)
	copy(b[8:], id)
	return b
}

func unmarshalBlockEntry(dst *BlockEntry, e kv) error {
	if len(e.key) < 8 || len(e.value) < 16 {
		return ErrInvalidBlockEntry
	}
	dst.Index = binary.BigEndian.Uint64(e.key)
	dst.ID = string(e.key[8:])
	dst.AppendedAt = int64(binary.BigEndian.Uint64(e.value[0:8]))
	dst.Level = binary.BigEndian.Uint32(e.value[8:12])
	dst.Shard = binary.BigEndian.Uint32(e.value[12:16])
	dst.Tenant = string(e.value[16:])
	return nil
}
