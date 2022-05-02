 package workerpool

 type WorkerPool struct {
  messages chan []byte
 }

 func New(buffer int) WorkerPool {
  return WorkerPool{
    messages: make(chan []byte, buffer)
  }
 }

 type Task func([]byte)

 func (wp WorkerPool) Delegate(t Task, width int) {
  for i := 0; i < width; i++ {
    go func () {
      for {
        select {
          case message := <-wp.messages:
            t(message)
        }
      }
    } ()
  }
 }

 func (wp WorkerPool) Throw(data []byte) {
  wp.messages <-data
 }
