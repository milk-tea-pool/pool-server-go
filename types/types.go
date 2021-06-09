package types

type Config struct {
	Hostname string `yaml:"host"`
	Port     string `yaml:"port"`
}

type SingletonState struct {
	poolURL            string
	poolPuzzleHash     [32]byte
	relativeLockHeight uint32
	minimumDifficulty  uint64
	escaping           bool
	blockchainHeight   uint32
	singletonCoinID    [32]byte
}

type RPCClient struct {
	name string
}

type BlockchainState struct {
	sync []string
}
