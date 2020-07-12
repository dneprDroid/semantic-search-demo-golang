package rpc 

type Pool struct {}

func NewPool() *Pool {
	return &Pool{}
}

func (self *Pool) Get() (*Rpc, error) {
	rpcEngine := NewRpc()
	if err := rpcEngine.Connect(); err != nil {
		return nil, err 
	}
	return rpcEngine, nil 
}

func (self *Pool) InContextSync(callback func(*Rpc) error) error {
	rpcEngine, err := self.Get()
	if err != nil {
		return err 
	}
	defer rpcEngine.Release()
	return callback(rpcEngine)
}