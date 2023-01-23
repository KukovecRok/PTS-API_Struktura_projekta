# Osnovna struktura Go projekta

 - V [main.go](main.go) inicializiramo in zaženemo dele aplikacije (inicializacija baze, zagon HTTP strežnika)
 - V [Router.go](Router.go) registriramo svoje REST poti in middleware funkcije 
 - V [API](API) sprejmemo HTTP podatke, jih deserializiramo v ustrezne strukture in nato kličemo dalje funkcijo iz Logic.go
 - V [Logic](Logic/) Imamo definirano glavno logiko programa. Obdelava podatkov in pošiljanje teh v bazo, računanja, preverjanje pravilnosti podatkov
 
 ---
 
 ### Mapa DB
 
 V to mapo damo kodo za komunikazijo z bazo
  - V [DB.go](DB/DB.go) imamo podatkovne strukture in interface za bazo
  - V [MariaDB](DB/MariaDB) ali kateri drugi datoteki glede na bazo pa imamo implementiran interface iz [DB.go](DB/DB.go) 
