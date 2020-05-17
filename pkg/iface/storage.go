package iface

import (
	"context"

	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variant"
	"io"
	"time"
)

type IStorageConfiguration interface {
	/*Storage Config*/
	CacheSize() int
	BatchSize() int
	QueueSize() int

	IndexEnable() bool
	AutoIndexMeta() bool
	AutoBuildIndex() bool

	StorageTaskQueue() variant.TaskQueue

	StateStoragePath() string
	StateStorageSecretKey() []byte

	IndexStoragePath() string
	IndexStorageSecretKey() []byte

	StoragePluginState() IStateStorage
	StoragePluginIndex() IIndexStorage

	StateStorageCustomConfiguration() interface{}
	IndexStorageCustomConfiguration() interface{}

	AddTransactionProcessor(tp ITransactionProcessor)
	IsTransactionProcessorExists(family, version string) bool
	GetTransactionProcessor(family, version string) ITransactionProcessor

	ProposalReceiver(proposal *pb.Proposal, status int)
}

type IStorage interface {
	RunGC()
	Open() error
	Close() error

	SetConfiguration(configuration IStorageConfiguration) bool

	ChangeSecretKey(path string,
		oldSecretKey []byte,
		newSecretKey []byte,
		encryptionKeyRotationDuration time.Duration) error

	PrepareSnapshot() (interface{}, error)
	RecoverFromSnapshot(r io.Reader) error
	SaveSnapshot(snapshotContext interface{}, w io.Writer) error

	SaveAppliedIndex(u uint64) error
	QueryAppliedIndex() (uint64, error)

	Search(_ variant.SearchRequest) (interface{}, error)
	Lookup(request variant.LookupRequest) (interface{}, error)
	ApplyProposal(ctx context.Context, proposal *pb.Proposal, entryIndex uint64) *pb.ProposalResponse
}
