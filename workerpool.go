 package workerpool

 type WorkerPool struct {
  messages chan interface{}
  end chan bool
 }

 func New(buffer int) WorkerPool {
  return WorkerPool{
    messages: make(chan interface{}, buffer),
    end: make(chan bool),
  }
 }

 type Task func(interface{})

 func (wp WorkerPool) Delegate(t Task, width int) {
  for i := 0; i < width; i++ {
    go func () {
      for {
        select {
          case message := <-wp.messages:
            t(message)
          case <-wp.end:
            return
        }
      }
    } ()
  }
 }

 func (wp WorkerPool) Throw(data interface{}) {
  wp.messages <-data
 }

 func (wp WorkerPool) End() {
  wp.end <-true
 }
