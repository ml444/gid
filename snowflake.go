package gid

func NewSnowflakeIdGenerator(epoch int64, deviceId, version uint64, opts ...OptionFunc) (IGenerator, error) {
	g, err := NewIdGenerator(&Config{
		Epoch:  epoch,
		WorkID: deviceId,
		SegmentBits: &SegmentBits{
			DurationBits: 41,
			WorkerIDBits: 10,
			SequenceBits: 12,
		},
	}, opts...)
	if err != nil {
		return nil, err
	}
	return g, nil
}
