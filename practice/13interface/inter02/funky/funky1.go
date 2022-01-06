package funky
import "fmt"

type Funkytype struct {
	Phrase string
}

func (t Funkytype) GetAmero(i int) string {
	fmt.Printf("[%d] ", i)
	return "君は Funky Baby おどけてるよ " + t.Phrase
}

func (t *Funkytype) GetBmero(i int) string {
	fmt.Printf("[%d] ", i)
	return "君がいなけりゃ Baby I'm blue " + t.Phrase
}

