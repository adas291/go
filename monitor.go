package main

import (
	"fmt"
	"sync"
)

type Monitor struct {
	MonitorLock   sync.Mutex
	Container     []Payment
	Capacity      int
	CurrentLength int
	cond          *sync.Cond
}

func CreateMonitor(capacity int) *Monitor {
	m := &Monitor{
		Capacity: capacity,
		cond:     sync.NewCond(&sync.Mutex{}),
	}
	m.Container = make([]Payment, capacity)
	return m
}

func (m *Monitor) GetCurrentLength() int {
	m.MonitorLock.Lock()
	defer m.MonitorLock.Unlock()
	return m.CurrentLength
}

func (m *Monitor) Add(p Payment) {
	m.MonitorLock.Lock()
	
	for m.CurrentLength == m.Capacity {
		m.cond.Wait()
	}

	fmt.Printf("added %d/%d\n", m.CurrentLength, m.Capacity)
	m.Container[m.CurrentLength] = p
	m.CurrentLength++
	m.cond.Broadcast()
	m.MonitorLock.Unlock()
}

func (m *Monitor) Remove() Payment {
	m.MonitorLock.Lock()
	defer m.MonitorLock.Unlock()

	for m.CurrentLength == 0 {
		m.cond.Wait()
	}

	item := m.Container[m.CurrentLength]
	m.CurrentLength--

	return item
}

type Payment struct {
	Name  string
	Price float32
	Count int16
}

func (p Payment) ToString() string {
	return fmt.Sprintf("this is custom tostring (%s)", p.Name)
}
