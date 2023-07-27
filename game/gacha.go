package game

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func Gacha() {
	fmt.Println("ガチャを回す。Enterを押してください。")
	fmt.Scanln()
	fmt.Println("ガチャン")
	// sound("door_lock")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("ガラガラ...")
	// sound("gatyagatya")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("ポン！")
	// sound("level_up")
	time.Sleep(1500 * time.Millisecond)

	charactor := generateCharactor()

	fmt.Println(charactor.Image)
	fmt.Println("なまえ : ", charactor.Name)
	fmt.Println("レアリティ : ", charactor.RaritytoStar())
	sound(charactor.SoundFile)

}

func sound(file_name string) {
	f, err := os.Open("resources/" + file_name + ".mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
