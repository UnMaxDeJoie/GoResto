package managers

import (
	"github.com/jordan-wright/email"
	"log"
	"net/http"
	"net/smtp"
)

func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	// Créer un e-mail
	e := email.NewEmail()
	e.From = "Votre nom <goresto833@gmail.com>"
	e.To = []string{"maxait93@gmail.com"}
	e.Subject = "Test d'envoi d'e-mail"
	e.Text = []byte("Ceci est un test d'envoi d'e-mail en utilisant GoMail.")

	// Authentification SMTP
	auth := smtp.PlainAuth("testmaxgo", "goresto833@gmail.com", "Goresto23!!", "smtp.gmail.com")

	// Envoi de l'e-mail
	err := e.Send("smtp.gmail.com:587", auth)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("E-mail envoyé avec succès !")
}
