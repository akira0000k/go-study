package monkey
import "fmt"

type Monkeytype struct {
	Phrase string
}

func (t Monkeytype) GetAmero(i int) string {
	fmt.Printf("[%d] ", i)
	return "君は Monkey Baby いかれてるよ " + t.Phrase
}

func (t Monkeytype) GetBmero(i int) string {
	fmt.Printf("[%d] ", i)
	return "君がいなけりゃ Baby I'm so sad " + t.Phrase
}

