Golang sync
Go1.9.2 sync库里包含下面几类：Mutex/RWMutex/Cond/WaitGroup/Once/Map/Pool
sync模块下定义的所有struct及其方法如下：

    type Cond struct {
    	// L is held while observing or changing the condition
    	L Locker
    	// contains filtered or unexported fields
    }
    func NewCond(l Locker) *Cond
    func (c *Cond) Broadcast()
    func (c *Cond) Signal()
    func (c *Cond) Wait()
    
    type Locker interface {
    	Lock()
    	Unlock()
    }
    
    type Map
    func (m *Map) Delete(key interface{})
    func (m *Map) Load(key interface{}) (value interface{}, ok bool)
    func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
    func (m *Map) Range(f func(key, value interface{}) bool)
    func (m *Map) Store(key, value interface{})

    type Mutex
    func (m *Mutex) Lock()
    func (m *Mutex) Unlock()

    type Once
    func (o *Once) Do(f func())

    type Pool
    func (p *Pool) Get() interface{}
    func (p *Pool) Put(x interface{})

    type RWMutex
    func (rw *RWMutex) Lock()
    func (rw *RWMutex) RLock()
    func (rw *RWMutex) RLocker() Locker
    func (rw *RWMutex) RUnlock()
    func (rw *RWMutex) Unlock()

    type WaitGroup
    func (wg *WaitGroup) Add(delta int)
    func (wg *WaitGroup) Done()
    func (wg *WaitGroup) Wait()

