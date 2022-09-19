package models

type WorkerReturn struct {
	Error      error
	Stats      *Timings
	PaymentIDs *[]uint64
}
