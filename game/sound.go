package game

// func sound(file_name string, done chan bool) {
// 	f, err := os.Open("resources/" + file_name + ".mp3")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	streamer, format, err := mp3.Decode(f)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer streamer.Close()

// 	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

// 	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
// 		done <- true
// 	})))

// }
