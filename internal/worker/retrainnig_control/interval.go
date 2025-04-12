package worker

func (rcw *retrainingControlWorker) Interval() float64 {
	return rcw.interval.Minutes()
}
