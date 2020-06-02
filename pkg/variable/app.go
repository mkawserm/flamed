package variable

import (
	"github.com/lni/dragonboat/v3/config"
	"github.com/lni/dragonboat/v3/raftio"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/iface"
)

var Name = "flamed"
var ShortDescription = "mFlamed is an open-source distributed embeddable NoSQL database"
var LongDescription = "mFlamed is an open-source distributed embeddable NoSQL database"
var DefaultPasswordHashAlgorithm = "argon2"
var DefaultPasswordHashAlgorithmFactory iface.IPasswordHashAlgorithmFactory = crypto.DefaultPasswordHashAlgorithmFactory()

var DefaultLogDbFactory config.LogDBFactoryFunc
var DefaultRaftRPCFactory config.RaftRPCFactoryFunc
var DefaultRaftEventListener raftio.IRaftEventListener
var DefaultSystemEventListener raftio.ISystemEventListener
