package gid

func NewMaxPeakIdGenerator(epoch int64, deviceId, mode, version uint64, opts ...OptionFunc) (IGenerator, error) {
	g, err := NewIdGenerator(&Config{
		Epoch:  epoch,
		WorkID: deviceId,
		SegmentBits: &SegmentBits{
			DurationBits: 31,
			WorkerIDBits: 10,
			SequenceBits: 22,
		},
	}, opts...)
	if err != nil {
		return nil, err
	}
	return g, nil
}
