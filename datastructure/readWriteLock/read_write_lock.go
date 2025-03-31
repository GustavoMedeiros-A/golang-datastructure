package readwritelock

import "sync"

// ReadWriteLock struct {
// 	mutex        → tranca o acesso aos campos abaixo
// 	readers      → conta quantos leitores estão ativos agora
// 	writerActive → diz se um escritor está ativo
// 	readWait     → fila de espera dos leitores
// 	writeWait    → fila de espera dos escritores
//   }

// type ReadWriteLock struct {
// 	mutex       sync.Mutex // garante que só uma goroutine por vez acesse/modifique as variáveis internas. Protege o acesso ao estado interno do lock
// 	readers     int        // serve pra saber se há leitores ativos (e, se tiver, bloquear o escritor)
// 	writeActive bool       // serve pra bloquear leitores enquanto o escritor estiver escrevendo.
// 	readWait    *sync.Cond // se um leitor tentar ler enquanto um escritor estiver ativo, ele espera aqui. Quando o escritor termina, ativa os leitores
// 	writeWait   *sync.Cond // é uma condição de espera para os escritores. Garante que o escritor só entra quando estiver seguro pra escrever
// }

type ReadWriteLock struct {
	mutex       sync.Mutex
	readers     int
	writeActive bool
	readWait    *sync.Cond
	writeWait   *sync.Cond
}

func NewReadWriteLock() *ReadWriteLock {
	rw := &ReadWriteLock{}
	rw.readWait = sync.NewCond(&rw.mutex)
	rw.writeWait = sync.NewCond(&rw.mutex)
	return rw
}

func (rw *ReadWriteLock) ReadLock() {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()

	for rw.writeActive {
		rw.readWait.Wait()
	}

	rw.readers++
}

func (rw *ReadWriteLock) ReadUnlock() {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()

	rw.readers--
	if rw.readers == 0 && rw.writeActive {
		rw.writeWait.Signal()
	}
}

func (rw *ReadWriteLock) WriteLock() {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()

	for rw.readers > 0 || rw.writeActive {
		rw.writeWait.Wait()
	}

	rw.writeActive = true
}

func (rw *ReadWriteLock) WriteUnlock() {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()

	rw.writeActive = false
	rw.readWait.Broadcast()
	rw.writeWait.Signal()
}
