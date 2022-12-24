package tcp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"time"
)

type tcpTransmitMessageHeader struct {
	Header   [4]byte
	Length   uint32
	Checksum uint32
	Reserved [8]byte
}

type tcpReceiveMessageHeader struct {
	Length   uint32
	Checksum uint32
	Reserved [8]byte
}

var startSequence = [4]byte{0x18, 0x96, 0x18, 0x20}

func buildHeader(payload []byte) tcpTransmitMessageHeader {
	header := tcpTransmitMessageHeader{
		Header:   startSequence,
		Length:   uint32(len(payload)),
		Checksum: 0,
		Reserved: [8]byte{0, 0, 0, 0, 0, 0, 0, 0},
	}
	for _, b := range payload {
		header.Checksum += uint32(b)
	}

	return header
}

func (t *Transport) writeMessage(payload string) error {
	if t.conn == nil {
		return errors.New("t not connected")
	}

	workingBuf := new(bytes.Buffer)
	castPayload := []byte(payload)
	header := buildHeader(castPayload)
	headerBufErr := binary.Write(workingBuf, binary.LittleEndian, header)
	if headerBufErr != nil {
		return headerBufErr
	}
	_, payloadBufErr := workingBuf.Write(castPayload)
	if payloadBufErr != nil {
		return payloadBufErr
	}

	_, writeErr := workingBuf.WriteTo(t.conn)
	return writeErr
}

func (t *Transport) readMessage(timeout time.Duration) ([]byte, error) {
	if t.conn == nil {
		return nil, errors.New("t not connected")
	}

	t.conn.SetReadDeadline(time.Now().Add(timeout))

	// Read in until we match the whole fixed start of message sequence
	currentStartSequenceIndex := 0

	nextByte := make([]byte, 1)
	var startSeqErr error
	for currentStartSequenceIndex < len(startSequence) {
		_, startSeqErr = t.conn.Read(nextByte)
		if startSeqErr != nil {
			return nil, startSeqErr
		}
		if nextByte[0] == startSequence[currentStartSequenceIndex] {
			currentStartSequenceIndex++
		} else {
			currentStartSequenceIndex = 0
		}
	}

	// get the rest of the header
	header := tcpReceiveMessageHeader{}
	err := binary.Read(t.conn, binary.LittleEndian, &header)
	if err != nil {
		panic(err)
	}

	// and the payload
	messageBuf := make([]byte, header.Length)
	_, payloadErr := t.conn.Read(messageBuf)
	return messageBuf, payloadErr
}
