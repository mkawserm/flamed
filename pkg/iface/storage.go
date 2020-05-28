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

	ProposalReceiver(proposal *pb.Proposal, status pb.Status)
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

	GetAppliedIndex(ctx context.Context) (interface{}, error)

	Search(ctx context.Context, search *pb.SearchInput) (interface{}, error)
	Iterate(ctx context.Context, iterate *pb.IterateInput) (interface{}, error)
	Retrieve(ctx context.Context, retrieve *pb.RetrieveInput) (interface{}, error)
	GlobalSearch(ctx context.Context, globalSearch *pb.GlobalSearchInput) (interface{}, error)
	GlobalIterate(ctx context.Context, globalIterate *pb.GlobalIterateInput) (interface{}, error)

	ApplyProposal(ctx context.Context, proposal *pb.Proposal, entryIndex uint64) *pb.ProposalResponse
}
