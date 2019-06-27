richiesta = preparaRichiesta();  
mandaRichiestaAsincrona(richiesta, function(risposta) {  
	visualizza(risposta);  
});  
visualizza("Siamo in attesa di risposta dal server ..."); 