package mlog

type Deferred struct {
	Log Logger
}

// NewDeferred constructs a new deferred logging handler
func NewDeferred(log Logger) *Deferred {
	return &Deferred{Log: log}
}

// IfErr logs an error IFF err != nil
func (d *Deferred) IfErr(fn func() error) {
	if err := fn(); err != nil {
		d.Log.Println(err.Error())
	}
}

// IfErrF logs an error IFF err != nil
func (d *Deferred) IfErrF(format string, fn func() error) {
	if err := fn(); err != nil {
		d.Log.Printf(format, err)
	}
}
