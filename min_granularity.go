package gid

func NewMinGranularityIdGenerator(epoch int64, deviceId, mode, version uint64, opts ...OptionFunc) (IdServer, error) {
	g, err := NewIdGenerator(2, 0, map[int]uint64{
		0: 0,        // sequence
		1: deviceId, // deviceId
		2: 0,        // timeDuration
		3: mode,     // mode: embed|rpc|http
		4: version,  // version
	}, []uint8{10, 10, 41, 2, 1}, epoch, opts...)
	if err != nil {
		return nil, err
	}
	return g, nil
}
