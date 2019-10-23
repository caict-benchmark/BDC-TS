package common

import (
	"io"
)

type serializerBceTSDBBulk struct {
}

func NewSerializerBceTSDBBulk() *serializerBceTSDBBulk {
	return &serializerBceTSDBBulk{}
}

func (s *serializerBceTSDBBulk) SerializePoint(w io.Writer, p *Point) (err error) {
	buf := scratchBufPool.Get().([]byte)

	for i := 0; i < len(p.TagKeys); i++ {
		buf = append(buf, p.TagValues[i]...)
	}

	if len(p.FieldKeys) > 0 {
		buf = append(buf, ' ')
	}

	for i := 0; i < len(p.FieldKeys); i++ {
		v := p.FieldValues[i]
		buf = fastFormatAppend(v, buf, false)

		if i+1 < len(p.FieldKeys) {
			buf = append(buf, ',')
		}
	}

	buf = append(buf, ' ')
	buf = fastFormatAppend(p.Timestamp.UTC().UnixNano(), buf, true)
	buf = append(buf, '\n')
	_, err = w.Write(buf)

	buf = buf[:0]
	scratchBufPool.Put(buf)

	return err
}

func (s *serializerBceTSDBBulk) SerializeSize(w io.Writer, points int64, values int64) error {
	return serializeSizeInText(w, points, values)
}
