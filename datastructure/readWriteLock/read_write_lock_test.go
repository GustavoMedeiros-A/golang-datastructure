package readwritelock

import (
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSingleReader(t *testing.T) {
	readWrite := NewReadWriteLock()

	var readCount int32

	go func() {
		readWrite.ReadLock()
		atomic.AddInt32(&readCount, 1)
		log.Println("Start reading", readCount)
		time.Sleep(20 * time.Millisecond)
		readWrite.ReadUnlock()
		log.Println("Finishing reading", readCount)
	}()
	time.Sleep(10 * time.Millisecond)

	if atomic.LoadInt32(&readCount) != 1 {
		t.Error("Expected read count to be 1, got", atomic.LoadInt32(&readCount))
	}
}

func TestMultipleReaders(t *testing.T) {
	readWrite := NewReadWriteLock()

	var readCount int32
	var waitGroup sync.WaitGroup

	numReaders := 5
	waitGroup.Add(numReaders)

	for i := 0; i < numReaders; i++ {
		go func(id int) {
			defer waitGroup.Done()
			readWrite.ReadLock()
			atomic.AddInt32(&readCount, 1)
			log.Println("Reader", id, "started")
			time.Sleep(20 * time.Millisecond)
			log.Println("Reader", id, "finishing")
			readWrite.ReadUnlock()
		}(i)
	}

	waitGroup.Wait()
	if atomic.LoadInt32(&readCount) != int32(numReaders) {
		t.Error("Expected read count to be", numReaders, "got", atomic.LoadInt32(&readCount))
	}

}

func TestWriterBlocksReaders(t *testing.T) {
	readWrite := NewReadWriteLock()
	var readStarted int32
	var writeDone int32
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	// Writers
	go func() {
		defer waitGroup.Done()
		readWrite.WriteLock()
		log.Println("Write started")
		time.Sleep(30 * time.Millisecond)
		atomic.StoreInt32(&writeDone, 1)
		log.Println("Write finished")
		readWrite.WriteUnlock()
	}()

	// Readers
	go func() {
		defer waitGroup.Done()
		log.Println("Write count when readers is started", atomic.LoadInt32(&writeDone))
		time.Sleep(5 * time.Millisecond)
		readWrite.ReadLock()
		if atomic.LoadInt32(&writeDone) == 0 {
			t.Error("Reader should not start while write is active")
		}
		log.Println("Reader started after write")
		atomic.AddInt32(&readStarted, 1)
		readWrite.ReadUnlock()
	}()

	waitGroup.Wait()

	if atomic.LoadInt32(&readStarted) != 1 {
		t.Error("Expected read to start exactly once, got", atomic.LoadInt32(&readStarted))
	}
}
