package main

import (
	"embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/michaelzhao21/shapes/data"
	"github.com/michaelzhao21/shapes/game"
	"github.com/michaelzhao21/shapes/graphics"
)

var (
	//go:embed assets/*
	assetsFS embed.FS

	songs map[string]data.Song
)

type Game struct {
	State game.GameState

	audioContext *audio.Context

	songAudioPlayer *audio.Player

	hitsoundPlayer *audio.Player
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyJ) || inpututil.IsKeyJustPressed(ebiten.KeyF) {
		// As audioPlayer has one stream and remembers the playing position,
		// rewinding is needed before playing when reusing audioPlayer.
		if err := g.hitsoundPlayer.Rewind(); err != nil {
			return err
		}

		g.hitsoundPlayer.Play()
	}

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.songAudioPlayer.Play()
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.songAudioPlayer.Pause()
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.State {
	case game.LOADING:
		graphics.DrawLoading(screen)
	case game.MAIN_MENU:
		graphics.DrawMainMenu(screen)
	case game.PLAYING:
		graphics.DrawPlaying(screen)
	}

	// Draw the image
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(50, 50)
	// op.GeoM.Scale(1.0/8, 1.0/8)
	// screen.DrawImage(img, op)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 450
}

func main() {
	// Initialize loggers
	data.CreateLoggers()

	// Share the assets filesystem
	data.AssetsFS = assetsFS

	// Load all songs
	songs = data.LoadSongs()

	// Print out all song names
	for k := range songs {
		data.GetInfoLogger().Println("Loaded song:", k)
	}

	// Setup game window
	ebiten.SetWindowSize(1600, 900)
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Phantom Paws")

	// AUDIO STUFF
	audioContext := audio.NewContext(44100)
	audioFile, err := data.AssetsFS.Open("assets/sounds/hitsound.wav")
	if err != nil {
		data.GetErrorLogger().Fatal(err)
	}
	wavFile, err := wav.DecodeWithSampleRate(44100, audioFile)
	if err != nil {
		data.GetErrorLogger().Fatal(err)
	}
	audioPlayer, err := audioContext.NewPlayer(wavFile)
	if err != nil {
		data.GetErrorLogger().Fatal(err)
	}

	songFile := data.OpenSongFile(data.Song{
		Name: "00000000 Lights Camera Action",
	})
	songWavFile, err := mp3.DecodeWithSampleRate(44100, songFile)
	if err != nil {
		data.GetErrorLogger().Fatal(err)
	}
	songAudioPlayer, err := audioContext.NewPlayer(songWavFile)
	if err != nil {
		data.GetErrorLogger().Fatal(err)
	}

	// Create game object
	game := Game{
		State:           game.LOADING,
		audioContext:    audioContext,
		hitsoundPlayer:  audioPlayer,
		songAudioPlayer: songAudioPlayer,
	}

	// Start the game
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
