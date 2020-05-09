package simple

import "github.com/lni/dragonboat/v3/raftio"
import pb "github.com/lni/dragonboat/v3/raftpb"

type RaftLog struct {
}

func (r *RaftLog) Name() string {
	return "simple"
}

func (r *RaftLog) Close() {

}

func (r *RaftLog) BinaryFormat() uint32 {
	return 1
}

func (r *RaftLog) GetLogDBThreadContext() raftio.IContext {
	return nil
}

func (r *RaftLog) ListNodeInfo() ([]raftio.NodeInfo, error) {
	return nil, nil
}

func (r *RaftLog) SaveBootstrapInfo(clusterID uint64, nodeID uint64, bootstrap pb.Bootstrap) error {
	return nil
}

func (r *RaftLog) GetBootstrapInfo(clusterID uint64, nodeID uint64) (*pb.Bootstrap, error) {
	return nil, nil
}

func (r *RaftLog) SaveRaftState(updates []pb.Update, ctx raftio.IContext) error {
	return nil
}

func (r *RaftLog) IterateEntries(ents []pb.Entry,
	size uint64,
	clusterID uint64,
	nodeID uint64,
	low uint64,
	high uint64, maxSize uint64) ([]pb.Entry, uint64, error) {
	return nil, 0, nil
}

func (r *RaftLog) ReadRaftState(clusterID uint64, nodeID uint64, lastIndex uint64) (*raftio.RaftState, error) {
	return nil, nil
}

func (r *RaftLog) RemoveEntriesTo(clusterID uint64, nodeID uint64, index uint64) error {
	return nil
}

func (r *RaftLog) CompactEntriesTo(clusterID uint64, nodeID uint64, index uint64) (<-chan struct{}, error) {
	return nil, nil
}

func (r *RaftLog) SaveSnapshots([]pb.Update) error {
	return nil
}

func (r *RaftLog) DeleteSnapshot(clusterID uint64, nodeID uint64, index uint64) error {
	return nil
}

func (r *RaftLog) ListSnapshots(clusterID uint64, nodeID uint64, index uint64) ([]pb.Snapshot, error) {
	return nil, nil
}

func (r *RaftLog) RemoveNodeData(clusterID uint64, nodeID uint64) error {
	return nil
}

func (r *RaftLog) ImportSnapshot(snapshot pb.Snapshot, nodeID uint64) error {
	return nil
}

func NewLogDB(dirs []string, lldirs []string) (raftio.ILogDB, error) {
	return &RaftLog{}, nil
}
