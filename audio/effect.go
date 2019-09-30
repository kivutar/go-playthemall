package audio

import (
	"os"

	"github.com/libretro/ludo/settings"
	wav "github.com/youpy/go-wav"
	"golang.org/x/mobile/exp/audio/al"
)

// Effect is a static sound effect
type Effect struct {
	Format *wav.WavFormat
	source al.Source
}

// LoadEffect loads a wav into memory and prepare the buffer and source in OpenAL
func LoadEffect(filename string) (*Effect, error) {
	var e Effect
	e.source = al.GenSources(1)[0]
	buffer := al.GenBuffers(1)[0]
	e.source.SetGain(settings.Current.MenuAudioVolume)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := wav.NewReader(file)

	e.Format, err = reader.Format()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wav := []byte{}
	for {
		var data [4096]byte
		n, _ := reader.Read(data[:])
		if n == 0 {
			break
		}
		wav = append(wav, data[:]...)
	}

	buffer.BufferData(al.FormatStereo16, wav, int32(e.Format.SampleRate))
	e.source.QueueBuffers(buffer)
	al.DeleteBuffers(buffer)

	return &e, nil
}

// PlayEffect plays a sound effect
func PlayEffect(e *Effect) {
	al.PlaySources(e.source)
}

// SetEffectsVolume sets the audio volume of sound effects
func SetEffectsVolume(vol float32) {
	for _, e := range Effects {
		e.source.SetGain(vol)
	}
}
