package advancedError

// subject interface define set of methods for Observers subject
type subject interface {
	RegisterObserver(obs observer)
	RemoveObserver(obs observer)
	Notify()
}

// RegisterObserver for subject
func (e *ExpertError) RegisterObserver(obs observer) {
	e.observers = append(e.observers, obs)
}

// RemoveObserver from subject
func (e *ExpertError) RemoveObserver(obs observer) {
	for i, observe := range e.observers {
		if observe == obs {
			e.observers = append(e.observers[:i], e.observers[i+1:]...)
			break
		}
	}
}

// Notify call update method for all registered observers
func (e *ExpertError) Notify() {
	for _, i := range e.observers {
		i.Update(e.GetMessage())
	}
}
