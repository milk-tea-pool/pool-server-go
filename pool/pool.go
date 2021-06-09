package pool

import (
	redis "github.com/go-redis/redis"
	t "milkteapool.com/pool-server/types"
	"milkteapool.com/pool-server/utils"
)

var DEFAULT_ROOT_PATH string

func main() {

}

type Pool struct {
	PoolFee                 float32
	PoolURL                 string
	MinDifficulty           uint64
	Config                  t.Config
	Store                   *PoolStore
	RelativeLockHeight      uint32
	BlockchainState         t.BlockchainState
	NodeRPCClient           *t.RPCClient
	WalletRPCClient         *t.RPCClient
	PendingPointPartials    *redis.Client
	DefaultTargetPuzzleHash []byte
	PoolFeePuzzleHash       []byte
	WalletFingerprint       int
	WalletID                string
}

type PoolStore struct {
	Name string
	Num  int
}

func NewPool(config *t.Config) *Pool {
	p := new(Pool)

	p.NodeRPCClient = nil
	p.WalletRPCClient = nil

	p.Store = nil
	p.PoolFee = 0.01

	p.RelativeLockHeight = 100
	p.PoolURL = "www.mypool.com"
	p.MinDifficulty = 10

	p.DefaultTargetPuzzleHash = utils.DecodePuzzleHash("txch16mhvz4cpzwq9dmds5lkjk22h34wftcc0xx728zmqwnla2wy2gl0qancuzs")
	p.PoolFeePuzzleHash = utils.DecodePuzzleHash("txch1ww8m7ttxuc2ng6qlthk609hkrt25mklcuqn882rkngnygeuu8fks36ckvs")

	p.WalletFingerprint = 2164248527
	p.WalletID = "1"

	return p
}

func NewPoolStore() *PoolStore {
	p := new(PoolStore)
	p.Name = ""
	return p
}

func NewRPCClient(url string, port uint16, rootPath string, netconfig t.Config) *t.RPCClient {
	n := new(t.RPCClient)
	return n
}

func (p *Pool) Start() {

	p.Store = NewPoolStore()
	// p.pendingPointPartials = redisClient
	hostname := p.Config.Hostname

	p.NodeRPCClient = NewRPCClient(hostname, uint16(8555), DEFAULT_ROOT_PATH, p.Config)

	go p.confirmPartialsLoop()
	go p.collectPoolRewardsLoop()
	go p.createPaymentLoop()
	go p.submitPaymentLoop()
	go p.getPeakLoop()

}

func (p *Pool) confirmPartialsLoop()    {}
func (p *Pool) collectPoolRewardsLoop() {}

func (p *Pool) createPaymentLoop() {}

func (p *Pool) submitPaymentLoop() {}
func (p *Pool) getPeakLoop()       {}
