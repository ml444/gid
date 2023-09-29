package gid

func NewSnowflakeIdGenerator(epoch int64, deviceId, version uint64, opts ...OptionFunc) (IdServer, error) {
	g, err := NewIdGenerator(2, 0, map[int]uint64{
		0: 0,        // sequence
		1: deviceId, // deviceId
		2: 0,        // timeDuration
		3: version,  // version
	}, []uint8{12, 10, 41, 1}, epoch, opts...)
	if err != nil {
		return nil, err
	}
	return g, nil
}
