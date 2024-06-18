package main

import "fmt"

// Subsystem classes
type Projector struct{}

func (p Projector) On() {
	fmt.Println("Projector is on.")
}

func (p Projector) Off() {
	fmt.Println("Projector is off.")
}

type SoundSystem struct{}

func (s SoundSystem) On() {
	fmt.Println("Sound system is on.")
}

func (s SoundSystem) Off() {
	fmt.Println("Sound system is off.")
}

type MediaPlayer struct{}

func (m MediaPlayer) On() {
	fmt.Println("Media player is on.")
}

func (m MediaPlayer) Play() {
	fmt.Println("Media player is playing.")
}

func (m MediaPlayer) Off() {
	fmt.Println("Media player is off.")
}

// Facade
type HomeTheaterFacade struct {
	projector   Projector
	soundSystem SoundSystem
	mediaPlayer MediaPlayer
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		projector:   Projector{},
		soundSystem: SoundSystem{},
		mediaPlayer: MediaPlayer{},
	}
}

func (h *HomeTheaterFacade) WatchMovie() {
	h.projector.On()
	h.soundSystem.On()
	h.mediaPlayer.On()
	h.mediaPlayer.Play()
}

func (h *HomeTheaterFacade) EndMovie() {
	h.mediaPlayer.Off()
	h.soundSystem.Off()
	h.projector.Off()
}

// Client code
func main() {
	homeTheater := NewHomeTheaterFacade()
	homeTheater.WatchMovie()
	// Output:
	// Projector is on.
	// Sound system is on.
	// Media player is on.
	// Media player is playing.

	homeTheater.EndMovie()
	// Output:
	// Media player is off.
	// Sound system is off.
	// Projector is off.
}
