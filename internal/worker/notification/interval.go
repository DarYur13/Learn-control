package worker

func (nw *notificationWorker) Interval() float64 {
	return nw.interval.Minutes()
}
