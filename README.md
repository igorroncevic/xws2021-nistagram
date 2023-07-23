# Asinhrono procesiranje podataka uz oslonac na Apache Kafka - diplomski rad

## Uvod
U diplomskom radu opisano je 
- asinhrono procesiranje u mikroservisnoj arhitekturi
- upotreba Apache Kafka platforme u distribuiranim sistemima za podršku asinhronog procesiranja podataka. 

Glavni cilj ovog rada je da približi razloge za korišćenje Kafke i sličnih platformi u rastućim mikroservisnim arhitekturama, kako aplikacije ne bi gubile na performansama u slučaju povećanja broja korisnika.

## Mikroservisi
Mikroservisi su mali, autonomni servisi koji sarađuju u cilju postizanja nekog cilja za korisnika usluga.

### Principi:
- Princip jedinstvene odgovornosti (eng. Single Responsibility Principle) - Definišemo granice nadležnosti servisa u odnosu na jedinične poslovne zahtjeve, čineći očiglednim namjenu datog servisa kroz čitav njegov životni vijek. 
  - Ovom separacijom izbjegavamo gomilanje funkcionalnosti unutar jednog servisa, spriječavajući ga da postane preveliki za održavanje.
- Mali - Nemoguće je dati konkretan broj linija koda ali ne treba pretjerivati. 
- Ukoliko je servis previše usitnjen i zahtjeva previše poziva ka drugim mikroservisima, može se osjetiti problem usporenja aplikacije. Što je servis manji, više maksimizujemo benefite po pitanju nezavisnosti i podjele nadležnosti, ali i mane po pitanju kompleksnosti.
- Nezavisni - neophodno je obezbjediti da se razvijaju, mijenjaju i isporučuju (eng. deploy) odvojeno od drugih, bez potrebe da se korisnici prilagođavaju
- Otporni na greške - otkazivanje jednog servisa ne smije da izazove kaskadno urušavanje ostalih servisa

### Mane:
- Implementacija transakcija i pretraga podataka su otežani i neophodno je koristiti dodatne šablone
- Lokacija servisa više nije unaprijed poznata, zato što se broj instanci servisa može mijenjati, pa samim tim i njihove lokacije. Neophodno je uključiti dodatne šablone za otkrivanje servisa, odnosno service discovery.
- Latentnost - servisi komuniciraju između sebe i svaki od dijelova te komunikacija oduzima vrijeme, kao i razmjena podataka između servisa, pa se može osjetiti određeno usporenje u radu.

## Asinhrono procesiranje
- Sinhrono procesiranje je uobičajeni način obrade podataka, gdje se zahtjevi obrađuju redom kojim stignu i u istom momentu kada su primljeni. 
- Aplikacija zauzima hardverske resurse dok ne završi sa obradom zahtjeva.
- Lakši je za praćenje jer odmah znamo da li je uspio ili nije
- Loš izbor za situacije kada su obrade zahtjevne, jer resursi ostaju predugo zauzeti
- U asinhronom procesiranju se obično koristi neki međusloj za usmjeravanje i upravljanje porukama, koji se ponaša kao posrednik između servisa.
- Servis će klijentu naknadno poslati odgovor pa klijent ne mora da zauzima hardverske resurse dok čeka odgovor, što povećava dostupnost servisa 

Postoje 2 tipa kanala preko kojih servisi asinhrono komuniciraju:
- Point-to-point, odnosno 1-1 kanali, koji dostavljaju poruku tačno jednom konzumentu koji čita iz kanala
- Publish/subscribe kanali koji dostavljaju poruku svakom od konzumenata koji se pretplatio na dati kanal

## Broker-based komunikacija
- Message broker je medijator kroz koji prolaze sve poruke. Pošiljalac šalje poruku u message brokera, a zatim je message broker dostavlja primaocu.
- Glavna prednost je što pošiljalac ne mora da zna adresu primaoca (što nije slučaj u brokerless komunikaciji)
- Druga prednost je zadržavanje, buffering, poruka dok servis ne bude mogao da ih obradi

### Faktori pri izboru message brokera:
- Podržani programski jezici i komunikacioni standardi
- Redoslijed poruka
- Garantovana isporuka
- Perzistencija, izdržljivost
- Skalabilnost, latencija
- Tradeoff kod svake platforme, zavisi šta nam je bitno

### Prednosti broker-based komunikacije:
- Loose-coupling - klijent šalje zahtjev na kanal i nema nikakvu svijest o tome koja instanca servisa će da primi tu poruku
- Buffering (zadržavanje) poruka dok se ne procesiraju, što pomaže u otpornosti na padove, jer poruke neće biti izgubljene kao i u situacijama kada je povećan saobraćaj
- Fleksibilna komunikacija - podržan je svaki tip komunikacije → req/res, async req/res, one-way notifications, pub/sub itd.

### Mane broker-based komunikacije, iako su značajne, većina modernih riješenja se pobrinula oko njih
- Usko grlo performansi - jer ipak sve poruke prolaze tuda
- Potencijalna tačka urušavanja
- Dodatna komponenta koja se treba odžavati

## Apache Kafka
Publish/subscribe sistem za razmjenu poruka dizajniran da podnese ogromne količine poruka sa niskom latencijom za real-time obradu podataka. 
Podaci se čuvaju u redoslijedu i mogu biti distribuirani u sklopu sistema (sharded particije) da pruže dodatnu zaštitu od naglih padova i mogućnost skaliranja

### Bitni pojmovi
#### Poruka i batch
Jedinica prenosa podataka je poruka, koja je, što se tiče Kafke, obični niz bajtova čija semantika ili format nije bitan. 
Zbog perfomansi, često se upisuju u gomilama (batches) koje predstavljaju kolekcije poruka koje se upisuju u isti topic ili particiju.

#### Topic i particija
Poruke u Kafki su kategorizovane po topic-ima, kao folder na fajl sistemu. Poruke se dodaju u append-only režimu, dakle na kraj niza, i čitaju se od početka do kraja
Svaki topic biva razbijen u više particija. Particije su jedan od načina da se Kafka skalira, jer se mogu čuvati odvojeno jedna od druge na različitim serverima, što omogućava horizontalno skaliranje.

#### Proizvođači i konzumenti
Proizvođači kreiraju nove poruke i upisuju ih u određeni topic i obično ih raspoređuju ravnomjerno po svim particijama tog topica.
Konzumenti čitaju poruke iz jednog ili više topic-a na koje su pretplaćeni i to u redoslijedu u kojem su upisane u topic. Konzumenti vode računa o tome koje su poruke pročitali uz pomoć offset-a, što je samo jedan cijeli broj koji se inkrementuje sa svakom pročitanom porukom.

#### Brokeri i klasteri
Jedan Kafka server zove se broker. Broker prima poruka od proizvođača, dodjeljuje im ofset i čuva poruku na disku.
Kafka brokeri dizajnirani su da rade u okviru klastera. U svakom klasteru, jedan broker je glavni i zadužen za “administrativne” poslove kao što je nadgledanje rada drugih brokera i praćenje srušenih brokera.

## Zašto Kafka?
- Kafka bez problema upravlja mnoštvom klijenata (proizvođača i konzumenata), nezavisno od broja topica. Idealna je za agregaciju podataka koji dolaze iz više servisa, kao i konkurentno čitanje poruka od strane više konzumenata, bez preklapanja.
- Buffering, odnosno čuvanje poruka, što znači da konzumenti ne moraju da rade u realnom vremenu.
- Skalabilnost sa particijama jednog topica i klasterima brokera, što znači da je pogodna i za male i za enterprise aplikacije.
- Niska latencija i visoka dostupnost klijentima

## Kada koristiti Kafku?
- Praćenje aktivnosti korisnika je originalna upotreba zbog koje je LinkedIn dizajnirao Kafku. Korisnici interaguju sa aplikacijom i generišu poruke koje se upisuju u Kafku, koje se kasnije mogu obraditi i pružiti bolje iskustvo korisnicima na osnovu njihovih akcije.
- Metrika i logovanje - aplikacije mogu proizvoditi različite tipove poruka vezane za metrike o radu, koje sistem za monitoring može konzumirati.
- Procesuiranje stream-ova podataka koje se radi u realnom vremenu, kako poruke stižu. Korisnici mogu da pišu male aplikacije za obradu podataka po svojim zahtjevima.

## Apache Kafka u demo aplikaciji društvene mreže
- Demo aplikacija zamišljena je kao platforma za dijeljenje slika i video materijala, inspirisana popularnom mrežom Instagram. 
- Registrovanje, dijeljenje sadržaja u vidu postova i storija, dopisivanje, čuvanje u kolekcije i hajlajte, verifikacija da bi postali influenseri, reklamne kampanje od strane agenata u vidu stori i post kampanja

### Arhitektura aplikacije:
Aplikacija je organizovana u 5 mikroservisa u Go-u koji komuniciraju preko gRPC-a, pristupa se preko gateway-a: servis za rad sa sadržajem, sa korisnicima, za dopisivanje, preporučivanje, agentska aplikacija

### Praćenje aktivnosti korisnika
Trenutno se koriste Prometheus i Grafana za monitoring performansi, međutim nemamo mogućnost prilagođavanja.
Proširivanje postojeće funkcionalnosti i dodavanje novih omogućava Kafka. 
- U prethodnu arhitekturu, dodat je novi, monitoring servis koje je pretplaćen na 2 topica (user-events i performance) i čuva eventove u bazu. Tim eventovima se može pristupiti preko REST endpointa
- Za praćenje korisničkih akcija demonstrativno su izabrane akcije od velike važnosti za korisnika (promjena šifre, login, ažuriranje podataka, ažuriranje kampanje itd) u slučaju napada. Prati se tip događaja, id korisnika, poruka sa detaljnim objašnjenjem. Svi događaji se mogu pregledati na Recent Activity stranici. Prikaz je demonstrativan i podaci se ne obrađuju, već služe isključivo za logovanje. Način obrade prelazi u domen poslovne logike koja ovdje nije obrađena.
- Za praćenje performansi, prate se važni neuspješni zahtjevi nesigurnih HTTP metoda (POST, PUT, DELETE). Nije potrebno logovati svaki uspješan zahtjev ili komentar na objavi, već samo one neuspješne, da bi administrator mogao upratiti uzrok. Prate se servis u kome se desila greška, funkcija i HTTP status.

## Zaključak
- Zašto koristiti asinhronu komunikaciju? 
    - Zbog povećanja broja korisnika, što zahtjeva daleko veću procesnu moć za obradu istih zahtjeva.
- Kada koristiti asinhronu komunikaciju? 
    - Kada god postoje neke obrade zahtjeva koje traju iole duže, kako ne bismo zauzimali hardverske resurse
- Koji su benefiti asinhrone komunikacije? 
    - Hardverski resursi su češće slobodni i povećana je dostupnost sistema za nove zahtjeve
- Apache Kafka je jedno od najpopularnijih opensource riješenja koje se koristi u svrhu asinhronog procesiranja, zbog niske latentnosti i visoke skalabilnosti. Pričali smo o njenim osnovnim pojmovima i mogućnostima korišćenja.
- U demo aplikaciji objasnili smo kako možemo koristiti Apache Kafku za logovanje korisničkih akcija. Iako se trenutno ne obrađuju, servisi ne moraju da brinu o tome kako će sačuvati događaji koje su generisali, već offloaduju taj posao na odvojen servis i nastavljaju sa radom.

# Kafka Notes

## Quick Start guide

For Ništagram, you will need to create two topics: "user-events" and "performance" using commands below.

## Container docs

- Kafka: https://hub.docker.com/r/wurstmeister/kafka

## Basic commands

### Entering Kafka container

- ```docker exec -it kafka /bin/sh```

## Topics

### Creating a topic

- ```cd /opt/kafka_2.13-2.7.0```
- ```./bin/kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic topicname```

### Listing topics

- ```./bin/kafka-topics.sh --list --zookeeper zookeeper:2181```

### Test read/write to a topic

- Write: ```bin/kafka-console-producer.sh --topic quickstart-events --bootstrap-server localhost:9092```
- Read: ```bin/kafka-console-consumer.sh --topic quickstart-events --from-beginning --bootstrap-server localhost:9092```
