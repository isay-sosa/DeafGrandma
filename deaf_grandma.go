package DeafGrandma

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func speakWithGrandma(speak string) string {
	speakUpcase := strings.ToUpper(speak)
	if speak == speakUpcase {
		return "No, not since " + randomYear() + "!"
	}
	return "Huh?! Speak up, sonny!"
}

func randomYear() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	yearStr := "19"
	yearConcat := 30 + r1.Intn(20)
	return yearStr + strconv.Itoa(yearConcat)
}
