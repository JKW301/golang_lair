package main

import (
	"fmt"
	"time"
)

// -----------------------
// Structures et Méthodes
// -----------------------
type Person struct {
	Name string
	Age  int
}

// -----------------------
// Méthode attachée à la structure Person
// -----------------------
/* Greet retourne une chaîne de caractères contenant un message de salutation
 * personnalisé avec le nom et l'âge de la personne. Cette méthode utilise
 * fmt.Sprintf pour formater la chaîne de manière conviviale.
*/
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I'm %d years old.", p.Name, p.Age)
}

// -----------------------
// Interfaces et Polymorphisme
// -----------------------
// Greeter est une interface qui définit une méthode Greet.
// Toute structure qui implémente cette méthode peut être considérée comme un Greeter.
// Cela permet d'utiliser le polymorphisme, où différentes structures peuvent être traitées de manière uniforme
// tant qu'elles implémentent cette interface.
type Greeter interface {
	Greet() string
}

// -----------------------
// Goroutines et Channels
// -----------------------

// Fonction qui simule un travail concurrent
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(500 * time.Millisecond) // Simuler un traitement
		results <- job * 2 // Retourner le double du job comme résultat
	}
}

// -----------------------
// Gestion des erreurs
// -----------------------
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// -----------------------
// Fonction principale
// -----------------------
func main() {
	// 1. Typage explicite et implicite
	var explicitVar int = 42
	implicitVar := "I am a string"
	fmt.Println("Explicit Variable:", explicitVar)
	fmt.Println("Implicit Variable:", implicitVar)

	// 2. Utilisation des Structures et Méthodes
	person := Person{Name: "Alice", Age: 30}
	fmt.Println(person.Greet())

	// 3. Interfaces et Polymorphisme
	var greeter Greeter = person
	fmt.Println("Interface says:", greeter.Greet())

	// 4. Goroutines et Channels
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Lancement des workers en goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Envoi des jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Récupération des résultats
	for r := 1; r <= numJobs; r++ {
		fmt.Printf("Result: %d\n", <-results)
	}

	// 5. Gestion des erreurs
	a, b := 10, 0
	result, err := divide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	// 6. Gestion de la concurrence avec les goroutines
	fmt.Println("Starting a concurrent task...")
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Concurrent task completed!")
	}()
	time.Sleep(2 * time.Second)

	// 7. Utilisation de la bibliothèque standard
	fmt.Println("Current time:", time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println("Go program completed successfully!")
}
