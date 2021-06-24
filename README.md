# exampleREST

S použitím knihovny gofiber.io navrhněte REST službu pro správu klientů.

Nemusíte řešit autorizaci (oAuth/session/…), všechny endpointy budou veřejně dostupné.

Služba bude pracovat se dvěma typama klientů:
	
osoba - obsahuje jméno, příjmení a datum narození
	
firma - obsahuje název společnosti, ičo a zastupující osobu (jméno a příjmení)

Pro práci s databázou doporučujeme použít knihovnu gorm.io.

Služba musí umět obsloužit následující typy požadavků:
	
výpis všech klientů (GET /clients)

vytvoření nového klienta (POST /clients)

aktualizace klienta (PUT /clients/{id})

odstranění klienta (DELETE /clients/{id})
