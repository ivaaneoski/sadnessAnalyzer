package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main(){
	fmt.Println("🧠 Welcome to your Sadness Report 🧠")

	fmt.Println("\033[35m")
	fmt.Println(`   
	
  
  ______                   __                                                 ______                       __                                         
 /      \                 |  \                                               /      \                     |  \                                        
|  $$$$$$\  ______    ____| $$ _______    ______    _______   _______       |  $$$$$$\ _______    ______  | $$ __    __  ________   ______    ______  
| $$___\$$ |      \  /      $$|       \  /      \  /       \ /       \      | $$__| $$|       \  |      \ | $$|  \  |  \|        \ /      \  /      \ 
 \$$    \   \$$$$$$\|  $$$$$$$| $$$$$$$\|  $$$$$$\|  $$$$$$$|  $$$$$$$      | $$    $$| $$$$$$$\  \$$$$$$\| $$| $$  | $$ \$$$$$$$$|  $$$$$$\|  $$$$$$\
 _\$$$$$$\ /      $$| $$  | $$| $$  | $$| $$    $$ \$$    \  \$$    \       | $$$$$$$$| $$  | $$ /      $$| $$| $$  | $$  /    $$ | $$    $$| $$   \$$
|  \__| $$|  $$$$$$$| $$__| $$| $$  | $$| $$$$$$$$ _\$$$$$$\ _\$$$$$$\      | $$  | $$| $$  | $$|  $$$$$$$| $$| $$__/ $$ /  $$$$_ | $$$$$$$$| $$      
 \$$    $$ \$$    $$ \$$    $$| $$  | $$ \$$     \|       $$|       $$      | $$  | $$| $$  | $$ \$$    $$| $$ \$$    $$|  $$    \ \$$     \| $$      
  \$$$$$$   \$$$$$$$  \$$$$$$$ \$$   \$$  \$$$$$$$ \$$$$$$$  \$$$$$$$        \$$   \$$ \$$   \$$  \$$$$$$$ \$$ _\$$$$$$$ \$$$$$$$$  \$$$$$$$ \$$      
                                                                                                              |  \__| $$                              
                                                                                                               \$$    $$                              
                                                                                                                \$$$$$$                               
	`)
	fmt.Println("\033[0m")

	fmt.Print("Tell me about your day : ")
	//Reading user input 
	reader := bufio.NewReader(os.Stdin)
	dayParagraph , _ := reader.ReadString('\n')

	// fmt.Print("\n You said : ")
	// fmt.Print(dayParagraph)

	//Cleaning the data input 
	dayParagraph = strings.TrimSpace(dayParagraph)
	dayParagraph = strings.ToLower(dayParagraph)


	//Applying the sadness calculator
	copiumData := analyzeCopium(dayParagraph)

	comment := getCopiumComment(copiumData)// retrives the if-else statements from copium_comments.go

	fmt.Println("\n🧾 Daily Copium Report:")
	fmt.Printf("📈\033[36m Copium Dose: %d%%\n", copiumData)
	fmt.Printf("\033[35m%s\033[0m\n",comment)


	fmt.Println("Thanks for trusting Daily Copium Report™. Stay strong, soldier! 🫡")
}

//Funtion to analyze the sadness levels 
func analyzeCopium(text string) int {
	sadWords := map[string]bool{
		"missed":       true,
		"failed":       true,
		"rejected":     true,
		"ghosted":      true,
		"forgot":       true,
		"broke":        true,
		"rained":       true,
		"ignored":      true,
		"late":         true,
		"alone":        true,
		"heartbroken":  true,
		"devastated":   true,
		"lost":         true,
		"grief":        true,
		"sorrow":       true,
		"unhappy":      true,
		"miserable":    true,
		"depressed":    true,
		"empty":        true,
		"numb":         true,
		"betrayed":     true,
		"abandoned":    true,
		"hurt":         true,
		"aching":       true,
		"bitter":       true,
		"disappointed": true,
		"regret":       true,
		"remorse":      true,
		"anguish":      true,
		"despair":      true,
		"suffering":    true,
		"torment":      true,
		"woeful":       true,
		"melancholy":   true,
		"downcast":     true,
		"gloomy":       true,
		"pessimistic":  true,
		"defeated":     true,
		"crushed":      true,
		"shattered":    true,
		"isolated":     true,
		"neglected":    true,
		"misunderstood": true,
		"unloved":      true,
		"forsaken":     true,
		"haunted":      true,
		"burdened":     true,
		"troubled":     true,
		"agonized":     true,
		"dismal":       true,
		"bleak":        true,
		"forlorn":      true,
		"weary":        true,
		"drained":      true,
		"cried":        true,
		"wept":         true,
		"mourned":      true,
		"lamented":     true,
		"struggled":    true,
		"battled":      true,
		"endured":      true,
		"survived":     true,
		"victim":       true,
		"casualty":     true,
		"tragedy":      true,
		"disaster":     true,
		"unlucky":      true,
		"cursed":       true,
		"hopeless":     true,
		"resigned":     true,
		"pained":       true,
		"wounded":      true,
		"scarred":      true,
		"fragile":      true,
		"vulnerable":   true,
		"helpless":     true,
		"powerless":    true,
		"yearned":      true,
		"longed":       true,
		"wished":       true,
		"if only":      true,
		"fuck":			true,
		"asshole":		true,
		"gay":			true,
		"cheated":		true,
		"ex":			true,
	}
	
	words := strings.Fields(text) //Splits the paragraph into words
	count := 0
  

	//Traversing through each word and checking after cleansing
	for _,word := range words { 
		word = cleanWord(word)
		if sadWords[word]{
			count++
		}
	}

	copium := count*10
	
	
	return copium
}

func cleanWord(word string) string {
	// Removing all non-alphabet characters at the start and end of the word
	word = strings.TrimFunc(word, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	// Returns the cleansed word
	return strings.ToLower(word)
}