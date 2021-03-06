package feeder

import (
	"github.com/emmvou/wints/schema"
	"io"
	"strings"
)

var buf = `
"timestamp";"Filiere";"EstEtranger";"EstLabo";"GenreEtu";"PrenomEtu";"NomEtu";"NumSS";"NumEtu";"Adresse1Etu";"Adresse2Etu";"Adresse3Etu";"Adresse4Etu";"CPEtu";"VilleEtu";"PaysEtu";"DomEtu";"EmailEtu";"TelEtu";"DateNaissEtu";"Assurance";"AssuranceNo";"NomEntreprise";"SiteWebEnt";"Adresse1SiegeSocial";"Adresse2SiegeSocial";"Adresse3SiegeSocial";"Adresse4SiegeSocial";"CPSiegeSocial";"VilleSiegeSocial";"PaysSiegeSocial";"AdrSiegeSocial";"GenreDirEnt";"PrenomDirEnt";"NomDirEnt";"EmailDirEnt";"TelDirEnt";"QualiteDirEnt";"DateDebut";"DateFin";"DureeHebdo";"DureeStageJour";"DureeStageSemaine";"HoraireSpecif";"FerieEventuel";"MontantGrat";"Devise";"MontantGratEuros";"ModaliteGrat";"AvantagesGrat";"TitreStage";"DescrStage";"Adresse1Stage";"Adresse2Stage";"Adresse3Stage";"Adresse4Stage";"CPStage";"VilleStage";"PaysStage";"AdrStage";"GenreEncadreur";"PrenomEncadreur";"NomEncadreur";"EmailEncadreur";"TelEncadreur";"FctEncadreur";"GenreEnsResp";"PrenomEnsResp";"NomEnsResp";"EmailEnsResp";"TelEnsResp";"InfoComp";"convsignee";"Parcours";
"2015-03-12 10:45";"SI 5";"oui";"oui";"M.";"Romain";"ALEXANDRE";"192040602706694";"21001838";"235 boulevard Andre Breton";"les jonquilles batiment 9";"";"";"06600";"ANTIBES (France)";"FRANCE";"235 boulevard Andre Breton--les jonquilles batiment 9--06600 ANTIBES (France)";"alexandre.romain06@gmail.com";"+33668138450";"27/04/1992";"Matmut";"980000269722Y";"McGill University";"";"845 Sherbrooke Street West";"";"";"";"H3A 0G4";"MONTREAL";"CANADA";"845 Sherbrooke Street West--H3A 0G4 MONTREAL--CANADA";"M.";"Gregory";"DUDEK";"dudek@cs.mcgill.ca";"+1 (514) 398-7071";"Head of Department";"23/03/2015";"21/09/2015";"35";"127";"26";"9:00 - 17:00 each day, with 1 hour lunch break";"";"1500";"CAD";"1115,5";"Direct deposit";"";"Concern-Driven Software Development with TouchCORE";"TouchCORE is a multi-touch enabled software design modelling tool aimed at developing scalable and reusable software design models following the concern-driven software development paradigm. To prepare TouchCORE for the tool demonstration session of the 18th International Conference on Model-Driven Engineering Languages and Systems (MODELS 2015) in October 2015, several areas need to be improved, including model verification and model checking, model tracing, support for instantiation cardinalities and support for state diagram modelling. Upon arrival, the student will be assigned to one of the areas depending on her/his expertise and preference. In all cases, the student will work on changing the representation of models within the tool (metamodelling), adapting the manipulation of models by the tool (model transformations and weaving), as well as extending the graphical user interface (OpenGL, multi-touch gestures).";"School of Computer Science, 3480 University";"x";"x";"x";"H3A 0G4";"MONTREAL";"Canada";"School of Computer Science, 3480 University--H3A 0G4 MONTREAL--Canada";"M.";"J??rg";"KIENZLE";"Joerg.Kienzle@mcgill.ca";"+1 (514) 398-2049";"Associate Professor, Head of the Software Engineering Laboratory";"Mme";"Mireille";"BLAY-FORNARINO";"blay@unice.fr";"(33)4 92 96 51 61 /(33) 4 97 25 82 15";"";"";"";
"2015-03-18 20:21";"SI 5";"non";"non";"M.";"Martin";"ALFONSI";"1 90 04 67 482 350 34";"20808334";"R??sidence THESA";"210 Avenue ROUMANILLE";"";"";"06410";"BIOT";"FRANCE";"R??sidence THESA--210 Avenue ROUMANILLE--06410 BIOT";"alfonsi@polytech.unice.fr";"06 33 76 37 29";"16/04/1990";"GMF";"88.705343.65E";"INNO TSD";"http://www.inno-group.com/";"Place Joseph Bermond";"Ophira 1 ??? BP63";"";"";"06902";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"Place Joseph Bermond--Ophira 1 ??? BP63--06902 SOPHIA ANTIPOLIS CEDEX";"M.";"Marc ";"PATTINSON";"m.pattinson@inno-group.com";"04 92 38 84 10";"G??rant ";"23/03/2015";"22/09/2015";"35";"128";"26";"";"";"600";"Euros";"600";"";"";"D??veloppement d???applications web";"Inno labs est la cellule TIC de Inno Group. Dans le cadre de ses activit??s techniques, un stagiaire ing??nieur en d??veloppement logiciel est recherch??. Il aura pour principales activit??s :
    ??? D???assister le chef de projet dans la gestion quotidienne des projets ;
    ??? Participer ?? toutes les phases des projets (conception, r??alisation, exploitation, support) ;
    ??? Participer aux ??tudes de conseil (analyse de besoin client, sp??cifications, benchmarks technologiques) ;
    ??? Participer ?? la r??daction des propositions (r??ponses ?? appel d???offres / appels ?? projets).
";"Place Joseph Bermond";"Ophira 1 ??? BP63";"";"";"06902";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"Place Joseph Bermond--Ophira 1 ??? BP63--06902 SOPHIA ANTIPOLIS CEDEX";"M.";"Fabrice ";"CLARI";"f.clari@inno-group.com";"+33 4 92 38 84 19";"Senior ICT consultant";"M.";"Christian";"BREL";"brel@polytech.unice.fr";"06 52 48 53 67";"";"";"";
"2015-02-19 17:41";"SI 5";"non";"non";"M.";"Quentin";"BITSCHEN??";"191058305010171";"20901433";"391 chemin des bruy??res";"";"";"";"83550";"Vidauban";"FRANCE";"391 chemin des bruy??res--83550 Vidauban";"bitschene.quentin@yahoo.fr";"+33672131657";"31/05/1991";"MAIF";"1300020R";"Atos";"atos.net";"80 quai Voltaire";"";"";"";"95877";"Bezons";"FRANCE";"80 quai Voltaire--95877 Bezons";"M.";"C??dric";"Couget";"cedric.couget@atos.net";"+33(0)786284830";"Responsable d???agence";"30/03/2015";"30/09/2015";"35";"129";"26";"";"";"1000";"Euros";"1000";"";"";"Immersion op??rationnelle dans un projet ATOS pour un client d???envergure internationale. ";"Immersion op??rationnelle dans un projet ATOS pour un client d???envergure internationale. Vous int??grez notre division Technology Services d???Atos Int??gration, R??gion Midi-Pyr??n??es, site de Toulouse St Martin, sous la responsabilit?? de Monsieur C??dric Couget, Responsable d???agence.
 
Objectifs p??dagogiques :
 
Sous le tutorat de Monsieur Jean-Bernard Pugnet, Directeur de projet, vous int??grez l?????quipe de projet PARAMETRAGE ORANGE dans les locaux Atos ?? Toulouse (6, impasse Alice Guy) sur des activit??s de conception et d??veloppement d???un module (ORPEO) d???une application WEB.
 
Vous ferez l???apprentissage des m??thodes de travail op??rationnelles dans un environnement structur??. En fin de stage, vous aurez d??velopp?? les comp??tences techniques requises dans le cadre d???un projet de conception/d??veloppement en informatique de gestion.
 
Objectifs du stage :
 
-       D??couverte des projets, du client, du contexte, du sujet, de l???organisation et des process.
-       Prise de connaissance des projets PARAMETRAGE, COCPIT et du module ORPEO.
-       Mont??e en comp??tence sur les technologies de conception et d??veloppement utilis??es :
o    D??veloppement : PHP5, HTML, CSS, JAVASCRIPT
o    Base de donn??es : MYSQL, PostgreSQL
o    M??thodes : UML, AGILE (SCRUM)
-       L???objectif du stage est la refonte du module ORPEO :
o    conception d???un nouveau mod??le de base de donn??es relationnelle,
o    migration des donn??es,
o    adaptation de l???application ?? ce nouveau r??f??rentiel
o    d??veloppement de nouvelles fonctionnalit??s (dans la mesure du possible).
-       Vous serez amen??e ?? r??aliser ?? l???ensemble des activit??s : ??tude, conception, maquettage, d??veloppement, tests, int??gration en fonction du niveau de votre niveau de mont??e en comp??tence.
 
Objectifs p??dagogiques :
 
-       Int??gration dans une ??quipe de projet,
-       Mise en pratique vos comp??tences techniques,
-       Assimilation de l???organisation, le fonctionnement et le travail sur un projet en ESN,
-       Travail en ??quipe et esprit de service.";"6, impasse Alice Guy";"";"";"";"31300";"Toulouse";"FRANCE";"6, impasse Alice Guy--31300 Toulouse";"M.";"Jean-Bernard";"Pugnet";"jean-bernard.pugnet@atos.net";"+33 6 84 68 73 88";"Directeur de projet";"M.";"Philippe";"Renevier";"Philippe.RENEVIER@unice.fr";"+33 4 9296 5167";"Vous b??n??ficierez du RIE sur place pour prendre vos repas et d???un remboursement des transports en commun ?? hauteur de 50% sur justificatif.";"";"";
"2015-03-06 10:25";"SI 5";"non";"non";"M.";"Pierre";"BOUILLET";"191067425622144";"21210034";"48 chemin du Battieu";"";"";"";"74190";"PASSY";"FRANCE";"48 chemin du Battieu--74190 PASSY";"pierre_bouillet@hotmail.fr";"0659008593";"06/06/1991";"mae";"C003345329";"SOPRA BANKING";"";"3 RUE DU PRE FAUCON";"PAE DES GLAISINS";"";"";"74940";"ANNECY LE VIEUX";"FRANCE";"3 RUE DU PRE FAUCON--PAE DES GLAISINS--74940 ANNECY LE VIEUX";"Mme";"H??l??ne";"RIPPERT";"helene.rippert@soprabanking.com";"04.50.33.31.49";"Directeur d???agence";"30/03/2015";"30/09/2015";"35";"129";"26";"";"";"1097,74";"Euros";"1097,74";"";"Tickets restaurant, et ??ventuellement, indemnisation 50% abonnement transport en commun sur justificatif en note de frais.";"D??veloppement d???un progiciel de Reporting Bancaire";"Int??gr?? ?? une ??quipe de d??veloppement, dans le cadre du projet de refonte de l???outil de d??veloppement de mod??les de reporting, au sein de l???entit?? Sopra Banking Compliance R&D.

L???outil est architectur?? en couches :

-          Une IHM low cost base sur Ms Excel, rempla??able ?? terme par une future IHM
-          Une couche M??tier, service & batch en java JEE
-          Une couche data en XML/SGBD

Il s???agit donc de d??veloppement dans un environnement java, JEE, XML, Ms Excel, Eclipse, maven, JUnit, CI (Jenkins/Sonar).

Vous aurez donc d???abord ?? appr??hender l???architecture technique et logique du produit, et contribuer ?? sa mise en ??uvre/am??lioration. ";"3 RUE DU PRE FAUCON";"PAE DES GLAISINS";"";"";"74940";"ANNECY LE VIEUX";"FRANCE";"3 RUE DU PRE FAUCON--PAE DES GLAISINS--74940 ANNECY LE VIEUX";"Mme";"St??phane";"POLICET";" stephane.policet@soprabanking.com";"04.50.33.32.61 ";"Ing??nieur d?????tudes";"Mme";"Audrey ";"OCCELLO";"occello@polytech.unice.fr";"04 92 96 51 02 ";"";"";"";
"2015-03-10 11:00";"SI 5";"oui";"oui";"Mlle";"C??cile";"CAMILLIERI";"292086822461427";"21004323";"4 rue de Roppe";"";"";"";"68200";"MULHOUSE";"FRANCE";"4 rue de Roppe--68200 MULHOUSE";"cecile.camillieri@gmail.com";"+33680900327";"13/08/1992";"GMF";"88.341406.65K";"McGill University";"";"845 Sherbrooke Street West";"";"";"";"H3A 0G4";"MONTREAL";"CANADA";"845 Sherbrooke Street West--H3A 0G4 MONTREAL--CANADA";"M.";"Gregory";"DUDEK";"dudek@cs.mcgill.ca";"+1 (514) 398-7071";"Head of Department";"23/03/2015";"21/09/2015";"35";"127";"26";"9:00 - 17:00 each day, with 1 hour lunch break";"";"1500";"CAD";"1099,197";"Direct deposit";"";"Concern-Driven Software Development with TouchCORE";"TouchCORE is a multi-touch enabled software design modelling tool aimed at developing scalable and reusable software design models following the concern-driven software development paradigm. To prepare TouchCORE for the tool demonstration session of the 18th International Conference on Model-Driven Engineering Languages and Systems (MODELS 2015) in October 2015, several areas need to be improved, including model verification and model checking, model tracing, support for instantiation cardinalities and support for state diagram modelling. Upon arrival, the student will be assigned to one of the areas depending on her/his expertise and preference. 
In all cases, the student will work on changing the representation of models within the tool (metamodelling), adapting the manipulation of models by the tool (model transformations and weaving), as well as extending the graphical user interface (OpenGL, multi-touch gestures).";"School of Computer Science";"3480 University";"";"";" H3A 0E9";"MONTREAL";"CANADA";"School of Computer Science--3480 University--H3A 0E9 MONTREAL--CANADA";"M.";"J??rg";"Kienzle  KIENZLE";"Joerg.Kienzle@mcgill.ca";"+1 (514) 398-2049";"Associate Professor, Head of the Software Engineering Laboratory";"Mme";"Mireille";"BLAY-FORNARINO";"blay@unice.fr";"(+33)4 92 96 51 61";"";"";"";
"2014-12-16 22:00";"SI 5";"oui";"non";"M.";"Adrien";"CASANOVA";"192050608870764";"21000600";"4 rue barla";"";"";"";"06300";"NICE";"FRANCE";"4 rue barla--06300 NICE";"acasanov@polytech.unice.fr";"0661755892";"27/05/1992";"MAPA";"2225574B";"Garagesocial, Inc.";"https://www.garagesocial.com/#!home";"20 Park Plaza 4th Floor";"";"";"";"02116";"BOSTON, MA";"ETATS-UNIS";"20 Park Plaza 4th Floor--02116 BOSTON, MA--ETATS-UNIS";"M.";"Maxime";"RASSI";"maxime@garagesocial.com";"617-948-2530";"President";"23/03/2015";"18/09/2015";"40";"124";"26";"";"";"2700";"Dollars";"2 370";"";"";"Front End Development Internship";"Development of user interfaces working with Javascript MVCs frameworks and mockup HTML templates.

- Company: Garagesocial, Inc. is a new online community and marketplace for the automotive industry. Garagesocial, Inc. develops tools that allow users and companies to showcase vehicles, parts and services online and allow them to engage in networking and commercial activities.

- The technologies may or will include:
 Languages: HTML, CSS, Javascript, PHP, Python, Ruby, Objective C, SQL
 Processors: Compass, Coffeescript, Less
 Frameworks: Laravel, ROR, EmberJS, BackboneJS, MarionetteJS
 Infrastructure: Amazon Web Services (S3, EC2, Route53, RDS, ElasticCache, ElasticBeanstalk), ElasticSearch, Hadoop, Memcache, Solr

- The type of work involved may or will include:
 Prototyping of new features including UX Mockups & Graphic Creation
 Developing, Testing and Deploying New Features
 Cross browser testing
 Performance Monitoring and Optimization
 Mobile Development
 Data Store Schema design and optimization
 API Development";"20 Park Plaza 4th Floor";"";"";"";"02116";"BOSTON, MA";"Etats-Unis";"20 Park Plaza 4th Floor--02116 BOSTON, MA--Etats-Unis";"M.";"Maxime";"RASSI";"maxime@garagesocial.com";"617-948-2530";"President";"M.";"Anne-Marie";"PINNA-DERY";"pinna@polytech.unice.fr";"(+33) 4 92 96 51 62";"";"";"";
"2014-12-08 10:24";"SI 5";"non";"non";"M.";"Guy";"CHAMPOLLION";"190047511317802";"20808172";"8 rue Saint-Antoine";"";"";"";"06600";"ANTIBES";"FRANCE";"8 rue Saint-Antoine--06600 ANTIBES";"champoll@polytech.unice.fr";"0625899748";"15/04/1990";"Macif";"13057068";"Reador.NET";"reador.net";"Business P??le, Bat.B - Entr??e A - 2??me ??tage";"1047 Route des Dolines";"";"";"06560";"VALBONNE";"FRANCE";"Business P??le, Bat.B - Entr??e A - 2??me ??tage--1047 Route des Dolines--06560 VALBONNE";"M.";"Christophe";"DESCLAUX";"christophe@reador.net";"0762596417";"Fondateur de la startup Reador.NET";"01/12/2014";"16/06/2015";"35";"138";"29";"";"";"436,05";"Euros";"436,05";"Ch??que ou virement bancaire";"";"Stage d??veloppement client web";"L???objectif de ce stage est de proposer une optimisation de l???IHM existante en l???adaptant aux lecteurs et/ou r??dacteurs d???informations. L???IHM refondue sera utilis??e comme client officiel de l???application web Reador. C???est une ??tape critique du projet car elle permettra de fournir une interface graphique en ad??quation avec les attentes des utilisateurs (lecteurs, r??dacteurs d???informations???) du service. Elle permettra de valoriser le travail d???annotation des news effectu?? en amont. Vous devrez donc proposer et concevoir au cours de votre stage des am??liorations de l???IHM actuellement mise en place ?? l???aide du framework RubyOnRails et de JQuery. Vous aurez une grande libert?? dans le choix des fonctionnalit??s ?? impl??menter et ferez preuve d???initiative.";"Business P??le, Bat.B - Entr??e A - 2??me ??tage";"1047 Route des Dolines";"";"";"06560";"VALBONNE";"FRANCE";"Business P??le, Bat.B - Entr??e A - 2??me ??tage--1047 Route des Dolines--06560 VALBONNE";"M.";"Christophe";"Desclaux";"christophe@reador.net";"0762596417";"Fondateur de la startup Reador.NET";"Mme";"Anne-Marie";"Dery-Pinna";"pinna@polytech.unice.fr";"0661029387";"";"";"";
"2015-01-14 13:40";"SI 5";"non";"non";"M.";"Zhang";"CHEN";"190109921600030";"21210445";"18 avenue docteur fabre";"";"";"";"06160";"JUAN LES PINS";"FRANCE";"18 avenue docteur fabre--06160 JUAN LES PINS";"zchen@polytech.unice.fr";"0750358934";"01/10/1990";"Soci??t?? Courtage d???Assurance";"07001878";"ATOS";"";"80 quai Voltaire";"";"";"";"95877";"BEZONS CEDEX";"FRANCE";"80 quai Voltaire--95877 BEZONS CEDEX";"Mme";"Guiselene";"CIEUTAT";"guiselene.cieutat@atos.net";"04 97 15 79 11";"Rersponsable des Ressources Humaines";"09/03/2015";"09/09/2015";"35";"129";"26";"";"";"1000";"Euros";"1000";"";"";"Stage d???analyse d???impact ?? partir de bases de donn??es NoSQL -orient??es graphes";"Pour des besoins d???int??gration dans un syst??me de releasing et de tests, nous vous confierons l?????tude, la conception et le d??veloppement d???un outil d???Analyse d???impact et de Root Cause.
Cet outil sera utilis?? ?? des fins de gestion, d???aide ?? la d??cision, de tra??abilit?? et de reporting,
 Pour mener ce projet une ??tude en deux phases est n??cessaire :
?Etablir l?????tat de l???art dans les domaines d???Analyse d???impact et de recherche de Root Cause.
?Concevoir et d??velopper un outil d???Analyse d???impact et de Root Cause ?? base de Base de Donn??es Orient??e Graphes.";"Le Millenium 150 all??e Pierre Ziller";"";"";"";"06560";"SOPHIA ANTIPOLIS";"FRANCE";"Le Millenium 150 all??e Pierre Ziller--06560 SOPHIA ANTIPOLIS";"M.";"Salim";"AINOUCHE";"salim.ainouche@atos.net";"06 20 43 55 86";"Senior Software Designer";"M.";"Lionel";"FILLATRE";"lionel.fillatre@i3s.unice.fr";"0492942785";"";"";"";
"2015-04-30 12:32";"SI 5";"non";"non";"M.";"Julien";"CHIARAMELLO";"1 94 02 06 088 271 63";"21209527";"4 rue de Dijon";"";"";"";"06000";"NICE";"FRANCE";"4 rue de Dijon--06000 NICE";"chiarame@polytech.unice.fr";"0659455961";"04/02/1994";"ACM-IARD SA";"4228532";"QuantifiCare SA";"www.quantificare.com";"1180 route des Dolines";"B??timent Athena B";"BP 40051";"";"06901";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"1180 route des Dolines--B??timent Athena B--BP 40051--06901 SOPHIA ANTIPOLIS CEDEX";"Mme";"Pascale";"BUISSON";"pbuisson@quantificare.com";"04 92 91 54 40";"Directrice";"11/05/2015";"30/09/2015";"35";"98";"21";"";"";"700";"Euros";"700";"Virement";"Acc??s tickets restau de 8???, prise en charge soci??t?? de 4.60???";"Administrateur Syst??mes et R??seaux";"QuantifiCare est une entreprise dynamique en croissance r??guli??re qui est parmi les
leader mondiaux dans le d??veloppement d???applications et de services visant le march??
de la recherche m??dicale. 

Nous travaillons en relation avec des clients dans le monde entier et nous poss??dons
des bureaux ?? Sophia Antipolis et ?? San Mateo (US). 

Une partie de nos activit??s nous oblige ?? mettre en place toutes les solutions de
s??curit?? ?? notre disposition pour r??cup??rer des donn??es sensibles envoy??es par nos
clients.

Nous recherchons actuellement un stagiaire qui, sous la responsabilit?? du
responsable informatique et r??seaux, participera ?? : 

- L???installation, l???administration et la maintenance des machines clientes de
l???entreprise (majorit?? de client Windows 7/8/8.1, quelques Linux)

- La gestion des utilisateurs, en particulier la cr??ation de leurs comptes, de leurs
droits d???acc??s s??curis??s ainsi que de leurs certificats ??lectroniques personnels
(SSL) et/ou clefs SSH

La maintenance et l???administration des services aux utilisateurs avec une attention
- particuli??re port??e sur les outils permettant leurs acc??s aux ressources de
l???entreprise de mani??re s??curis??e (OpenVPN, SSL/TLS, SSH)

- La surveillance des sondes de s??curit?? des serveurs et l???analyse des ??ventuelles
alertes

- La maintenance des serveurs pour assurer leur int??grit?? op??rationnelle (tests
op??rationnels, analyse et application des mises ?? jour de s??curit??)

- La maintenance et la mise ?? jour des actifs r??seaux, en particulier les solutions de routage et de s??curit?? install??es entre les diff??rents r??seaux

- La gestion des achats du mat??riel informatique 

- ??ventuellement des interventions de maintenance ?? distance sur des machines chez nos
clients

En dehors de ces t??ches de fond, le stagiaire participera ??galement, en relation
avec divers d??partements de l???entreprise, ?? l?????tude, l???am??lioration et l?????volution
de l???infrastructure ?? serveur ?? de l???entreprise fournissant les services s??curis??s
destin??s ?? nos employ??s et ?? nos clients (serveurs exclusivement sous syst??mes
OpenSource Linux/FreeBSD). 

Force de proposition, il participera de mani??re active ?? l?????volution de nos
diff??rentes plate-formes, en particulier en apportant son expertise en mati??re de
s??curisation.

Patience, ??coute, ouverture d???esprit, r??activit?? ?? r??fl??chie ?? sont des qualit??s
ind??niables pour cette mission.

De bonnes notions de ?? hardware ?? seraient un plus.

QuantifiCare ??tant une soci??t?? tourn??e vers l???international, nous recherchons une
personne poss??dant un bon niveau d???anglais capable de lire / r??diger des proc??dures
techniques dans un anglais correct mais aussi capable d?????ventuellement communiquer
oralement avec nos employ??s am??ricains ainsi qu???avec notre client??le.";"1180 route des Dolines";"B??timent Athena B";"BP 40051";"";"06901";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"1180 route des Dolines--B??timent Athena B--BP 40051--06901 SOPHIA ANTIPOLIS CEDEX";"M.";"Matthieu";"MEURILLON";"mmeurillon@quantificare.com";"04 92 91 54 41";"Responsable informatique & r??seaux";"Mme";"Tamara";"REZK";"tamara.rezk@inria.fr";"04 97 15 53 37";"";"";"";
"2015-02-06 13:58";"SI 5";"oui";"oui";"Mlle";"Genevi??ve";"CIRERA";"292023155564686";"21206834";"1 rue la fontaine";"";"";"";"31450";"Donneville";"FRANCE";"1 rue la fontaine--31450 Donneville";"genevieve.cirera@gmail.com";"0647968604";"19/02/1992";"AXA";"3792084404";"NII";"";"2-1-2 Hitotsubashi Chiyoda-ku";"";"";"";"101-8430";"Tokyo";"JAPAN";"2-1-2 Hitotsubashi Chiyoda-ku--101-8430 Tokyo--JAPAN";"M.";"Masaru";"KITSUREGAWA";"nii-internship@nii.ac.jp";"03-4212-2000";"Director General";"15/03/2015";"30/08/2015";"35";"118";"24";"";"";"171000";"Yen";"1273,28";"";"";"AI system that solve physics problems of entrance exam for university";"General intelligent robots in daily life environment must observe and model external world, understand users??? instructions and intentions. In such situations, robots always face to ambiguity of information. Conventional approaches to solve the ambiguity were using common sense database/ontology, or asking users; however developers???/users??? cost is too huge.
You will engage on the software development for connection between 
natural language processing and physics simulator.
The languages used are perl, Modelica, C++.";"2-1-2 Hitotsubashi Chiyoda-ku";"";"";"";"101-8430";"Tokyo";"JAPAN";"2-1-2 Hitotsubashi Chiyoda-ku--101-8430 Tokyo--JAPAN";"M.";"Tetsunari";"Inamura";"inamura@nii.ac.jp";"03-4212-1605";"Associate Prof";"M.";"Fr??d??ric";"Precioso";"frederic.precioso@polytech.unice.fr";"+33 (0)4 92 96 51 43";"";"";"";
"2015-03-24 10:07";"SI 5";"oui";"oui";"M.";"H??l??ne";"COLLAVIZZA";"111112";"11112";"tttt";"";"";"";"06130";"mouans";"FRANCE";"tttt--06130 mouans";"helene.collavizza@gmail.com";"0606060606";"03/03/2015";"MAIF";"1234";"Essai H??l??ne";"";"fdgdfg";"";"";"";"06370";"Mouans";"FRANCE";"fdgdfg--06370 Mouans";"M.";"ffff";"ffff";"tt@tt.fr";"060606";"sfsfsf";"03/06/2015";"30/07/2015";"20";"40";"9";"";"";"100";"Euros";"100";"";"";"le beau stage";"dsdgsdgdgsdgsdgg";"sdsdgd";"";"";"";"02136";"kolp";"FRANCE";"sdsdgd--02136 kolp";"M.";"ddfdd";"dfdfdf";"ddd@tt.fr";"03030";"patron";"M.";"tttt";"pppp";"jj@ff.fr";"02 03 03 03";"";"";"";
"2015-02-26 11:32";"SI 5";"non";"non";"M.";"Cl??ment";"CRISTIN";"192054209506311";"21210627";"140 rue Albert Einstein appt C05";"";"";"";"06560";"Valbonne";"FRANCE";"140 rue Albert Einstein appt C05--06560 Valbonne";"cristin.clement@gmail.com";"+336 72 34 33 18";"05/05/1992";"MAIF";"1367936P";"Hewlett-Packard";"";"Z.I. de Courtaboeuf 1, avenue du Canada";"";"";"";"91947";"Les Ulis";"FRANCE";"Z.I. de Courtaboeuf 1, avenue du Canada--91947 Les Ulis";"M.";"Virginie";"MARESCHAL";"Virginie.mareschal@hp.com";"+334 80 32 14 87";"Charg??e d???administration RH";"16/03/2015";"15/09/2015";"35";"128";"26";"";"";"1220";"Euros";"1220";"";"";"HP R&D ??? Participer au d??veloppement de notre GUI Framework Telecom HP CMS ???Unified OSS Console???";"Vous rejoignez une ??quipe R&D dynamique, comp??tente et exp??riment??e dans le domaine de la
T??l??communication et du d??veloppement logiciel. 
Cette ??quipe bas??e ?? Sophia Antipolis est en charge d???un Web Framework UI nouvelle g??n??ration destin?? ?? int??grer les diff??rentes applications du Portfolio HP (Fault Management, Service Quality Management, Service Level Agreement Management, Events Correlation, Customer Experience, OSS Analytics...) et potentiellement les applications de nos clients dans un environnement graphique unifi?? et s??curis??.

Ce Framework UI est totalement ouvert ?? travers un SDK et est impl??ment?? sur des technologies web modernes (JavaScript MVVM ?? base de NodeJS / AngularJS) qui permettent un haut niveau d???int??gration, de customisations et de s??curit?? pour les diff??rentes applications t??l??coms. 
Il propose une exp??rience utilisateur riche et efficace sur diff??rents hardware (desktop, tablette...) en int??grant les concepts de web design ?? responsive ??. 
Un outil ???designer??? WYSIWYG permettra aux op??rateurs de construire graphiquement des tableaux de bord et des ??crans op??rationnels ?? partir des briques pr??d??finies disponibles.

Le stagiaire sera immerg?? dans une ??quipe R&D et participera au d??veloppement du Framework
commun et ?? l???enrichissement de ses diff??rents modules m??tiers en rajoutant de nouvelles
fonctionnalit??s en relation avec les responsables produits, les architectes et les ??quipes de R&D.

Il devra notamment:

- Enrichir le GUI Framework existant avec de nouvelles fonctionnalit??s et des services communs.
- Sp??cifier, d??signer et impl??menter des modules fonctionnels en relation avec des besoins clients Telecom.
- Fournir aux int??grateurs et aux clients d???HP une infrastructure de d??veloppement et de
d??ploiement de leurs propres modules, mise en plages et leurs widgets graphiques. (SDK, API, RestAPI, g??n??rateurs Yeoman...).
- Participer au d??veloppement du designer graphique WYSIWYG et ?? son int??gration dans le
produit.
- Enrichir notre biblioth??que de widgets graphiques. Certains widgets sont bas??es sur Highcharts et HighMaps (www.highcharts.com ).
- Concevoir et proposer des ??crans HTML responsives pour une exp??rience utilisateur efficace.
- Participer aux activit??s li??es ?? la mise en production en respectant les r??gles de qualit?? de l?????quipe R&D.
- Participer ?? la mise en place de tests unitaires et fonctionnels
- Participer ?? la documentation des produits et des diff??rentes APIs (REST, JavaScript...)
- Assister ?? la gestion de projet Agile bas?? sur SCRUM
- Fixer les probl??mes remont??s par nos utilisateurs.
- Participer aux pr??sentations techniques et aux d??monstrations clientes";"Marco Polo ??? B??timent B ??? Entr??e B1 ZAC du Font de l???Orme 1 BP1220 790 Avenue du Docteur Donat";"";"";"";"06254";"Mougins";"FRANCE";"Marco Polo ??? B??timent B ??? Entr??e B1 ZAC du Font de l???Orme 1 BP1220 790 Avenue du Docteur Donat--06254 Mougins";"M.";"Jean-Charles";"Picard";"jean-charles.picard@hp.com";"+33 4 22390127";"Responsable Developpement Logiciel";"M.";"S??bastien ";"Mosser";"mosser@i3s.unice.fr";"+334 92 96 50 58";"";"";"";
"2014-12-08 10:38";"SI 5";"non";"non";"M.";"Anthony";"DA MOTA";"191078313740995";"20901891";"30 Lotissement le Bois de la Combe";"";"";"";"83720";"Trans en Provence";"FRANCE";"30 Lotissement le Bois de la Combe--83720 Trans en Provence";"anthony.damota06@gmail.com";"0686072630";"24/10/1991";"LMDE";"3283564704";"Atos Toulouse";"";"6 Impasse Alice Guy";"";"";"";"31024  ";"TOULOUSE";"FRANCE";"6 Impasse Alice Guy--31024 TOULOUSE";"M.";"Julien";"BIREBENT";"julien.birebent@atos.net";"+33 5 34 36 32 23";"Delivery Manager ??? BPS / Telco IS";"30/03/2015";"30/09/2015";"35";"129";"26";"";"";"1062,20";"Euros";"1062,20";"";"";"Stage de d??veloppement Firefox OS";"Mission :
Il s???agira de r??aliser un prototype d???application mobile pour FirefoxOS servant dans un second temps de support d?????tude aux fonctionnalit??s de ce type d???application. Le stagiaire mettra notamment en relief ses r??sultats avec ce qu???il est possible de faire sur le syst??me d???exploitation Android.
Cette ??tude pourra porter sur les fonctionnalit??s disponibles mais ??galement les caract??ristiques telles que la s??curit??, la performance, l???outillage disponible dans cet environnement ou encore la fiabilit?? des applications d??velopp??es sur FirefoxOS.
 
La r??alisation du prototype ainsi que l?????tude comparative sera r??alis??e en utilisant la m??thodologie Scrum. Le p??rim??tre du prototype et de l?????tude sera donc affin?? tout au long du stage.
 
Le stagiaire sera int??gr?? au sein d???une ??quipe projet de l???agence toulousaine participant au d??veloppement d???applications similaires en utilisant la m??thode Scrum ??galement. ";"6 Impasse Alice Guy";"";"";"";"31024 ";"TOULOUSE";"FRANCE";"6 Impasse Alice Guy--31024 TOULOUSE";"M.";"Julien";"BIREBENT";"julien.birebent@atos.net";"+33 5 34 36 32 23";"Delivery Manager ??? BPS / Telco IS - Toulouse";"M.";"Michel";"BUFFA";"Michel.Buffa@unice.fr";"(33)-92-07-66-60";"";"";"";
"2015-03-26 07:44";"SI 5";"oui";"non";"Mlle";"Huinan";"DONG";"289099921600013";"dh210451";"Cite U jean medecin 25rue robert latouche";"";"";"";"06200";"Nice";"FRANCE";"Cite U jean medecin 25rue robert latouche--06200 Nice";"dlutdong@gmail.com";"+33605573389";"28/09/1989";"BNP";"177900968";"Price waterhouseCoopers";"http://www.pwccn.com/home/eng/index.html";"26/F., Office Tower A, Beijing Fortune Plaza,  7 Dongsanhuan Zhong Road, Chaoyang District";"";"";"";"100020";"Beijing ";"CHINE";"26/F., Office Tower A, Beijing Fortune Plaza,  7 Dongsanhuan Zhong Road, Chaoyang District--100020 Beijing--CHINE";"M.";"Depei";"SHEN";"offeree.support@cn.pwc.com";" +86106533 8888";"Directeur ";"23/03/2015";"23/09/2015";"40";"129";"26";"";"";"3000";"Yuan";"440";"";"";"Asset Liability Management";"1. ALM(Asset Liability Management) advisory.In this project PwC do the advisory work for the interest rate risk of banking book and liquidity risk management for the client, including gap analysis, optimizing the management structure, risk measurement for interest rate risk of banking book and liquidity risk management, and the related reporting roles. 

2. ALM system validation.PwC help the client propose system implementation roadmap, design ALM database and ALM models, set up ALM data standards and conduct the ALM system validation.

The role of the intern:
1.Using SQL to prepare the data for system validation;
2.Design the system validation program;
3.Verify report developed by sub-contractor;
4.Data requirement analysis, data gap analysis;
?";"26/F., Office Tower A, Beijing Fortune Plaza,  7 Dongsanhuan Zhong Road, Chaoyang District";"";"";"";"100020";"Beijing";"CHINE";"26/F., Office Tower A, Beijing Fortune Plaza,  7 Dongsanhuan Zhong Road, Chaoyang District--100020 Beijing--CHINE";"M.";"Yong";"LU";"pwc_recruit@foxmail.com";"+8613811630156";"directeur du projet";"M.";"Ioan";"BOND";"bond@polytech.unice.fr ";"+33678110112";"";"";"";
"2014-12-09 15:03";"SI 5";"non";"non";"M.";"Cl??ment";"DUFFAU";"192102403710940";"21002723";"13 rue des Petits Ponts";"";"";"";"06250";"MOUGINS LE HAUT";"FRANCE";"13 rue des Petits Ponts--06250 MOUGINS LE HAUT";"duffau@polytech.unice.fr";"0633796793";"02/10/1992";"PACIFICA";"2649317906";"Axonic";"http://www.axonic.fr/";"2720 Chemin Saint-Bernard ";"";"";"";"06224";"VALLAURIS";"FRANCE";"2720 Chemin Saint-Bernard--06224 VALLAURIS";"M.";"Marc";"ROUGE";"mrouge@axonic.fr";"0497213040";"PDG";"16/03/2015";"13/09/2015";"35";"126";"26";"";"";"1200";"Euros";"1200";"";"";"Stage R&D en d??veloppement logiciel";"Axonic, filiale du groupe MXM, est une startup bas??e ?? Sophia-Antipolis, 
sp??cialis??e dans le d??veloppement d???appareils m??dicaux actifs implantables 
d??di??s ?? la neuro-stimulation ayant pour objectif d???am??liorer les 
conditions de vie de patients atteints de maladies chroniques s??v??res ou 
d??g??n??ratives. 

Axonic recherche un stagiaire (Ing??nieur / Master 2) en d??veloppement logiciel pour int??grer 
son ??quipe et travailler sur la conception et le d??veloppement du 
""Framework Logiciel Axonic"". 

Ce Framework permet de g??rer, de piloter, et de contr??ler les produits 
Axonic, et notamment les param??tres de la stimulation. Il s???adresse ?? des 
utilisateurs qui ont pour m??tier la recherche clinique, la m??decine, la 
chirurgie mais aussi ?? terme directement aux patients. 

L?????quipe de d??veloppement logiciel suit une m??thode de d??veloppement Agile 
bas??e sur Scrum et l???int??gration continue. 

Dans un contexte de recherche et d??veloppement et d???innovation, le stagiaire 
sera en interface permanente avec tous les m??tiers de notre entreprise 
(micro-??lectronique, ??lectronique, m??canique, logiciel, clinique) et il 
d??couvrira les sp??cificit??s et les exigences li??es au domaine m??dical 
(respect de normes, s??curit??, couverture de tests, risques patients, 
??thique). 

Dans le cadre de ses attributions le stagiaire prendra en charge le 
d??veloppement complet des ""User Stories"" dont les priorit??s sont fix??es par 
nos ""Product Owners"" (??quipe de recherche clinique). 

Il pourra intervenir sur l???ensemble de l???application, les couches basses 
(ex: drivers permettant la communication avec les stimulateurs), le code 
m??tier (support et gestion de la stimulation), et/ou sur le code de 
l???interface homme-machine (IHM) en fonction des priorit??s et de ses 
attentes. 

Sur chacun des domaines, le stagiaire participera ?? la d??finition de 
l???architecture, ?? la conception, ?? l???impl??mentation, ?? ses tests, et ?? la 
r??daction de la documentation. 

Il pourra ??galement ??tre amen?? ?? ??tudier, proposer et mettre en place des 
solutions d???outillage permettant d???am??liorer nos pratiques de 
d??veloppement. ";"2720 Chemin Saint-Bernard ";"";"";"";"06224";"VALLAURIS";"FRANCE";"2720 Chemin Saint-Bernard--06224 VALLAURIS";"M.";"Pierrick";"Perret";"pperret@axonic.fr";"0497213040";"Head of Project Management Office";"M.";"Sebastien";"Mosser";"mosser@i3s.unice.fr";"0492965058";"J???ai un doute sur le tuteur enseignant entre le responsable sp??cialit?? AL et le responsable stage SI5";"";"";
"2015-01-08 18:05";"SI 5";"non";"non";"M.";"Thibaut";"DUFOUR";"193047815823714";"21002820";"37 Avenue Balzac";"";"";"";"92410";"VILLE d???AVRAY";"FRANCE";"37 Avenue Balzac--92410 VILLE d???AVRAY";"dufour@polytech.unice.fr";"0677910408";"09/04/1993";"Macif";"5176896";"Rivage Investment SAS";"www.rivageinvestment.com";"26 Rue du Quatre Septembre";"";"";"";"75002";"PARIS";"FRANCE";"26 Rue du Quatre Septembre--75002 PARIS";"M.";"Thierry";"SENG";"thierry.seng@rivageinvestment.com";"01 70 91 25 94";"Directeur Syst??mes Informatiques";"16/03/2015";"16/09/2015";"35";"129";"26";"";"";"1200";"Euros";"1200";"Virement bancaire";"Remboursement de frais professionnels sur pr??sentation des factures acquitt??es";"Ing??nieur D??veloppeur";"Le collaborateur fera partie int??grante de l?????quipe de d??veloppement de la plateforme propri??taire de gestion
et ?? ce titre aura les responsabilit??s suivantes :
- Responsabilit??s principales :
o d??velopper et maintenir la plateforme de gestion et ses outils (C++)
o d??velopper et maintenir la base de donn??es et ses outils (SQL Server)
o garantir la qualit??, la robustesse et l???adaptabilit?? de la plateforme de gestion
o ??tre force de proposition dans l???am??lioration des briques logicielles existantes
- Autres missions :
o soutenir la gestion courante des fonds (assistant de gestion front office : Excel)
o maintenir le site web de la soci??t?? (mises ?? jour, corrections, am??liorations)
o assurer une veille technologique sur les architectures concurrentes 
";"26 Rue du Quatre Septembre";"";"";"";"75002";"PARIS";"FRANCE";"26 Rue du Quatre Septembre--75002 PARIS";"M.";"Thierry";"SENG";"thierry.seng@rivageinvestment.com";"01 70 91 25 94";"Directeur Syst??mes Informatiques";"Mme";"Anne-Marie";"HUGUES";"hugues@unice.fr";"06 84 04 59 30";"";"";"";
"2015-03-24 21:47";"SI 5";"non";"non";"M.";"Fabien";"FOERSTER";"192068311802118";"21003654";"22 place des arcades";"";"";"";"06250";"MOUGINS";"FRANCE";"22 place des arcades--06250 MOUGINS";"fabienfoerster@gmail.com";"+33664513101";"05/06/1992";"Cr??dit Agricole";"931932908";"ATOS";"http://atos.net/";"River Ouest";"80 quai Voltaire";"";"";"95877";"BEZONS CEDEX";"FRANCE";"River Ouest--80 quai Voltaire--95877 BEZONS CEDEX";"Mme";"Guiselenne";"CIEUTAT";"anne.aime@atos.net";"04 97 15 79 11";" Responsable des Ressources Humaines";"07/04/2015";"20/09/2015";"35";"116";"24";"";"";"1000";"Euros";"1000";"virement bancaire";"";"Cartographie et tra??abilit??";"Le stage a pour objectif de d??finir une m??thode de calcul d???impact dans une cartographie multi-vue. 
En d??tail l?????tudiant doit :
??tudier les offres d???outils de mod??lisation multi paradigmes (UML, BPMN, MCD, autres)
impl??menter, dans un outil choisit suite ?? l?????tude, une cartographie multi-vue
d??finir une m??thode (voir une automatisation) du calcul d???impact dans un contexte de cartographie multi-vue
(si on a le temps) attacher ?? la m??thode une pond??ration en points de fonction des p??rim??tres d???impact d??tect??s";"Le Mill??nium";"150, all??e Pierre Ziller";"B.P. 279";"";"06905";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"Le Mill??nium--150, all??e Pierre Ziller--B.P. 279--06905 SOPHIA ANTIPOLIS CEDEX";"Mme";"Cl??mentine";"NEMO";"clementine.nemo@atos.net";"04 93 95 46 44";"Consultante Solution";"Mme";"Mireille";"BLAY-FORNARINO";"blay@i3s.unice.fr";"O4 92 96 51 61";"Alors premi??rement la position de mon encadrant au sein de l???entreprise peut etre sujette ?? modification suivant la r??ponse de mon encadrant.
Deuxi??mement mon sujet de stage peut ??tre un peu plus ??toff?? au besoin.";"";"";
"2015-02-25 22:08";"SI 5";"oui";"non";"Mlle";"Nancy";"FONG";"292069941713060";"21002581";"475 rue Evariste Galois, R??sidence Les Calades, n??3206";"";"";"";"06410";"Biot";"FRANCE";"475 rue Evariste Galois, R??sidence Les Calades, n??3206--06410 Biot";"nancy.fong06@gmail.com";"0610703448";"26/06/1992";"Matmut";"769 9040 72338 P 50";"mycs GmbH";"";"Friedrichstr. 123";"";"";"";"10117";"Berlin";"ALLEMAGNE";"Friedrichstr. 123--10117 Berlin--ALLEMAGNE";"M.";"Ka Chun";"To";"kachun@mycs.com";" +49 176 6118 1645";"Managing Director";"01/04/2015";"31/08/2015";"40";"107";"22";"";"";"1000";"Euros";"1000";"";"";"Web Development Internship";"- Learn how to code using the mycs technology stack
  - Javascript / CoffeeScript
  - AngularJS
  - React & Flux (still to be implemented)
  - node.js / hapijs
  - Postgres
  - Redis, Knex, Bootstrap, Grunt, Gulp, Bower, Jasmine, Karma

- Work using agile methodologies on a daily basis
  - Scrum
  - Continous Integration & Deployment
  - Self-organizing teams
  - Pair programming
  - MVP approach

- Get to know how a fully automated Microservice Architecture works on a production environment
  - RESTful
  - Amazon AWS
  - Docker
  - Ansible
  - Varnish

- Work on own Frontend projects

- Work on own Backend projects";"Friedrichstr. 123";"";"";"";"10117";"Berlin";"Allemagne";"Friedrichstr. 123--10117 Berlin--Allemagne";"M.";"Claudio";"Bredfeldt";"claudio@mycs.com";"+49 (0) 30/24333736";"Technology Director";"M.";"Michel";"Buffa";"buffa@unice.fr";"0662659345";"";"";"";
"2015-04-13 14:54";"SI 5";"non";"oui";"M.";"Luis";"GIOANNI";"191097511877859";"20905615";"45 rue Barberis";"";"";"";"06300";"Nice";"FRANCE";"45 rue Barberis--06300 Nice";"lgioanni@gmail.com";"0679306991";"26/09/1991";"12345";"12345";"Laboratoire I3S";"http://www.i3s.unice.fr";"2000 Route des Lucioles";"Les Algorithmes";"B??timent Euclide B";"";"06900";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"2000 Route des Lucioles--Les Algorithmes--B??timent Euclide B--06900 SOPHIA ANTIPOLIS CEDEX";"M.";"Michel";"RIVEILL";"sabine.barrere@i3s.unice.fr";"0492942705";"Directeur du Laboratoire I3S";"01/05/2015";"30/09/2015";"35";"106";"22";"";"";"500,51";"Euros";"500,51";"";"";" S??lection et composition opportuniste de capteurs pour la reconnaissance d???activit??s";"Contexte du stage :

Christel Dartigues Pallez et Fr??d??ric Precioso : Equipe Mind, Laboratoire I3S
St??phane Lavirotte et Jean-Yves Tigli : Equipe Rainbow, Laboratoire I3S

Nous disposons d???un ensemble de donn??es issues de diff??rents capteurs port??s (acc??l??rom??tres, magn??tom??tres, gyroscopes, centrales inertielles, ???) et ??quipant un environnement physique (d??tecteur de pr??sence, contacteurs, ???). Notre objectif est de reconna??tre les activit??s d???un utilisateur ?? partir de ces donn??es.
La probl??matique de l?????tude propos??e est que nous ne pouvons pas faire de pr??suppositions sur la disponibilit?? des capteurs ; on peut m??me avoir de nouveaux capteurs non connus a priori qui vont concourir ?? la reconnaissance de l???activit??. L???approche souhait??e est donc de r??aliser cette reconnaissance de mani??re totalement opportuniste, en fonction des capteurs disponibles ?? un instant donn??.
Nous avons d??j?? mis en place diff??rents algorithmes d???apprentissage permettant de reconna??tre l???activit?? d???un utilisateur. L???apprentissage de l???activit?? a ??t?? r??alis?? avec l???ensemble des donn??es de tous capteurs, puis pour chaque capteur ind??pendamment puis avec des configurations de capteurs en nombre variables. Les r??sultats obtenus sont coh??rents avec les autres travaux du domaine et il a ??t?? possible d???am??liorer ces r??sultats gr??ce ?? l???utilisation d???algorithmes d???apprentissage de type Random Forest.
Objectifs du stage
Le premier objectif de ce stage est de repartir des r??sultats obtenus pr??c??demment et de confirmer et justifier les am??liorations apport??es par l???apprentissage ?? base de Random Forest. Le but de cette premi??re partie de stage est de finaliser l???ensemble des informations n??cessaires pour publier les r??sultats obtenus dans une conf??rence scientifique.
Dans un second temps, il faudra ??tudier les approches possibles pour le m??canisme de s??lection des capteurs disponibles pour la reconnaissance d???une activit??. Ce deuxi??me objectif est donc d???am??liorer les r??sultats de reconnaissance d???activit?? tout en garantissant une approche opportuniste (?? savoir une s??lection parmi les capteurs disponibles ?? un instant donn??). Le m??canisme de s??lection pourra ??tre influenc?? par diff??rents crit??res de recherche (consommation ??nerg??tique minimis??e, optimisation de la reconnaissance, ???).
Le troisi??me objectif de ce stage sera d?????tudier la composition opportuniste des capteurs s??lectionn??s pour r??aliser une application auto-adaptative de reconnaissance d???un ensemble d???activit??s. Cette troisi??me partie devra d??buter par une ??tude bibliographique et proposer une approche pour la composition dynamique des donn??es des capteurs ou des r??sultats des apprentissages dans le but de reconna??tre un ensemble d???activit??s.
Comp??tences requises
Les comp??tences attendues pour traiter ce sujet sont :
-	Des connaissances sur les techniques d???apprentissage et la ma??trise d???un outil comme Matlab
-	Des connaissances sur le domaine de l???Intelligent Ambiante
-	Des comp??tences personnelles comme : initiative et force de proposition, autonomie, ???

R??f??rences
D. Roggen, A. Calatroni, K. Fr??ster, G. Tr??ster, P. Lukowicz, D. Bannach, A. Ferscha, M. Kurz, G. H??lzl, H. Sagha, H. Bayati, J. Mill??n and R. Chavarriaga. ?? Activity recognition in opportunistic sensor environments ??, Procedia Computer Science, vol 7, pp 173-174, 2011
G. H??lzl, M. Kurz and A. Ferscha. ?? Goal oriented opportunistic recognition of high-level composed activities using dynamically configured hidden markov models ??. In The 3rd International Conference on Ambient Systems, Networks and Technologies (ANT2012), 2012
M. Kurz, G. H??lzl, and Alois Ferscha. ?? Dynamic adaptation of opportunistic sensor configurations for continuous and accurate activity recognition ??. In Fourth International Conference on Adaptive and Self-Adaptive Systems and Applications (ADAPTIVE2012), July 22-27, Nice, France, July 2012
D. Roggen, A. Calatroni, K. F??rster, G. Tr??ster, P. Lukowicz et al. ?? Activity Recognition in Opportunistic Sensor Environments ??. The European Future Technologies Conference and Exhibition 2011, 2011.
";"Equipe Sparks";"";"";"";"06900";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"Equipe Sparks--06900 SOPHIA ANTIPOLIS CEDEX";"M.";"St??phane";"LAVIROTTE";"stephane.lavirotte@unice.fr";"0679672827";"Enseignant chercheur";"M.";"Jean-Yves";"TIGLI";"tigli@unice.fr";"0684245567";"";"";"";
"2015-05-15 10:13";"SI 5";"non";"non";"M.";"David";"GUERRERO";"191046444522936";"20906154";"210 Avenue Roumanille";"R??sidence Th??sa";"Appartement 222 B1";"";"06410";"BIOT";"FRANCE";"210 Avenue Roumanille--R??sidence Th??sa--Appartement 222 B1--06410 BIOT";"heldroe@gmail.com";"0629089011";"22/04/1991";"GAN";"A01385/091552079";"Ignilife France SAS";"www.ignilife.com";"27 rue Professeur Delvalle";"";"";"";"06300";"NICE";"FRANCE";"27 rue Professeur Delvalle--06300 NICE";"M.";"Fabrice";"Pakin";"fabrice@ignilife.com";"0673464911";"CEO";"11/05/2015";"28/08/2015";"35";"76";"16";"";"";"1000";"Euros";"1000";"";"20??? par mois de frais de transport";"Plate-forme objets connect??s";"La plate-forme des objets connect??s est une application ?? part enti??re ind??pendante de l???application Ignilife. L???id??e est d???externaliser les connecteurs actuels vers les API de Fitbit, Runkeeper et Withings (et d???autres ?? venir) vers cette application tout en repensant le mod??le d???int??gration. 
L???objectif est d???optimiser le temps et la facilit?? d???int??gration de nouveaux objets connect??s en assouplissant et en rendant g??n??rique le mod??le actuel. De plus, l???application doit permettre de d??livrer une donn??e formalis??e et ce quelque soit le provider source.  

Aspects techniques : l???application doit proposer un mod??le g??n??rique de provider et doit permettre de plugger de la mani??re la plus simple possible tous nouveaux providers en limitant au maximum le code sp??cifique. Les donn??es utilisateurs sont stock??es sur la plateforme selon un mod??le unique et doivent ??tre accessibles via une API s??curis??e. ";"27 rue Professeur Delvalle";"";"";"";"06300";"NICE";"FRANCE";"27 rue Professeur Delvalle--06300 NICE";"M.";"David";"BESSOUDO";"david@ignilife.com";"0689248157";"CTO";"M.";"Fr??d??ric";"PRECIOSO";"frederic.precioso@polytech.unice.fr";"0492965143";"";"";"";
"2015-03-18 15:27";"SI 5";"non";"non";"M.";"Jamal";"HENNANI";"192013417256369";"21209240";"2255 route des Dolines, Residence les Dolines APT 74";"";"";"";"06560";"VALBONNE";"FRANCE";"2255 route des Dolines, Residence les Dolines APT 74--06560 VALBONNE";"jamal.hennani@gmail.com";"0663795666";"27/01/1992";"ADH - SEGIA";"AC482864";"AMADEUS SAS";"";"485 Route du Pin Montard- Les Bouillides BP 69-";"";"";"";" 06902";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"485 Route du Pin Montard- Les Bouillides BP 69---06902 SOPHIA ANTIPOLIS CEDEX";"M.";"Pierre ";"PUIG";"stephanie.gasperini@amadeus.com";"04 97 15 45 81";"Talent Acquisition Manager";"01/04/2015";"30/09/2015";"37";"128";"26";"";"";"1100";"Euros";"1100";"Virement bancaire";"Le stagiaire b??n??ficiera de l???acc??s au restaurant d???entreprise et d???une prime de transport prenant en compte le trajet quotidien entre le domicile du stagiaire pendant l???ex??cution de son stage et l???entreprise (versement de cette prime selon les proc??dures d??finies par l???entreprise).";"Monitor and better control unexpected overbookings";"Our functional areas include the management of the bookings in airline inventory system (counters, bookings data, etc), the waitlist clearance process and the data reconciliation between reservation and inventory.
Main Responsibilities:
Our teams are responsible for sustaining critical traffic, reaching more than 2500 transactions per seconds at peak time.
This traffic is made of various transactions. The diversity of the flow, the throughput and the continuous evolution of our applications can be a source of problems impacting the accuracy of our booking counters.
In this complex environment, we implemented a system to automatically recover most of the problems. Indeed, inaccurate counters is a source of dissatisfaction for our partners. This can lead to overbookings, with potential impact on flight departure, or to empty seats which is a loss of revenue for the airlines.
We need to:
 -improve our visibility on the counter corrections done every day in our system. How many bookings were updated? Is there a serious impact for the flight? Is the flight close to departure? What are the tools to put in place to notify relevant people? What???s the cause of the problem?
 -have an efficient way to schedule some quality checks of our counters
The internship will participate in the implementation of a new application providing these functionalities to the teams.
The trainee will participate in following activities:
 -Get knowledge about airline inventory through existing documentation and presentations done by the team
 -Participate in defining the requirements with development and product definition teams
 -Help the team to design and implement the tools
 -Participate in the presentation to the teams";"2 rue du Vallon, Amadeus";"";"";"";"06560";"VALBONNE";"FRANCE";"2 rue du Vallon, Amadeus--06560 VALBONNE";"M.";"Maxime ";"ARMAND";"marmand@amadeus.com";"+33 4 9715 4971";"Manageur";"M.";"Francoise";"BAUDE";"baude@unice.fr";"+33 4 92 38 76 71";"";"";"";
"2015-03-19 12:14";"SI 5";"non";"non";"M.";"Swan";"JUMELLE-DUPUY";"189110608807665";"20904568";"3 rue guiglia";"";"";"";"06000";"NICE";"FRANCE";"3 rue guiglia--06000 NICE";"swan.jumelle@gmail.com";"+33618811975";"04/11/1989";"FILIA-MAIF";"5075611 N";"Ojingo Labs";"www.ojingolabs.com";"2101 23rd St";"";"";"";"94107";"San Francisco CA";"ETATS-UNIS";"2101 23rd St--94107 San Francisco CA--ETATS-UNIS";"M.";"Thomas";"MARBOIS";"tj@ojingolabs.com";"0667317118";"Directeur";"23/03/2015";"22/09/2015";"35";"128";"26";"";"";"3600";"Euros";"3600";"Virement";"";"D??voloppement d???interfaces admin";"Les d??veloppeurs serveur se trouvant en Russie, et vu la complexit?? des technologies mises en place, il est actuellement difficile de faire le lien entres les bugs client/serveur.
Le but du stage va consister ?? donner plus de visibilit?? aux d??veloppeurs Client, testeurs et administrateurs.
Il faudra mettre en place plusieurs interfaces web permettant:
- De connaitre l?????tat de sant?? du cluster de serveur (Monitoring OS, DB, APP)
- De visualiser les logs de l???application, ??ventuellement avoir des statistiques
- De visualiser la base de donn??es, et ??ventuellement interagir avec celle-ci

D???autres solutions pourront ??tre explor?? afin d???aider aux mieux les non-d??veloppeurs serveur ?? comprendre ce qu???il s???y passe.

L???impl??mentation de certaines fonctionnalit??s et r??solution de bugs cot?? serveur feront aussi parti des t??ches assign??es pendant le stage.";"1 place Mass??na";"";"";"";"06000";"NICE";"FRANCE";"1 place Mass??na--06000 NICE";"M.";"Thomas";"MARBOIS";"tj@ojingolabs.com";"0667317118";"Directeur";"M.";"Christian";"BREL";"brel@i3s.unice.fr";"+33 (0)4 92 96 50 74";"";"";"";
"2015-03-06 12:24";"SI 5";"non";"non";"M.";"Antoine";"LAVAIL";"191106938237790";"20900385";"247 impasse des Gen??vriers";"";"";"";"83100";"TOULON";"FRANCE";"247 impasse des Gen??vriers--83100 TOULON";"antoine.lavail@icloud.com";"0650558350";"23/10/1991";"MMA";"116363311";"Ojingo Labs";"http://www.ojingolabs.com";"2101 23rd St";"";"";"";"94107";"San Francisco, California";"USA";"2101 23rd St--94107 San Francisco, California--USA";"M.";"TJ";"MARBOIS";"tj@ojingolabs.com";"0667317118";"Fondateur";"23/03/2015";"22/09/2015";"35";"128";"26";"";"";"4000";"Euros";"4000";"";"";"D??veloppeur d???applications iOS";"Vous int??grerez l?????quipe compos??e de quatre d??veloppeurs iOS, de deux d??veloppeurs serveurs Vert.X, de trois graphistes et de deux project manager. Durant votre stage, vous participerez ?? la phase de reflexion et de cr??ation des projets en apportant votre vision technique sur la r??alisation d???applications aussi bien en terme d???ossature, d???ergonomie et de design de ces derni??res. Vous pourrez aussi conseiller le graphiste lors de la r??alisation des wireframes, et/ou maquettes en mati??re de format, de taille des fichiers, de compression et de portabilit?? mais aussi des chefs de projet lors de la r??daction du cahier des charges.";"1, place Mass??na";"";"";"";"06000";"NICE";"FRANCE";"1, place Mass??na--06000 NICE";"M.";"TJ";"MARBOIS";"tj@ojingolabs.com";"0667317118";"Fondateur";"M.";"Anne Marie";"DERY-PINNA";"pinna@polytech.unice.fr";"06 61 02 93 87";"";"";"";
"2015-02-18 11:40";"SI 5";"non";"non";"M.";"Paul";"LAVOINE";"19126748240459";"lp206655";"2255 route des Dolines";"";"";"";"06560";"Valbonne";"FRANCE";"2255 route des Dolines--06560 Valbonne";"paul.lavoine@hotmail.fr";"+33650789978";"03/12/1991";"maaf";"167148625K003";"Big Boss Studio";"http://www.bigbossstudio.com/";"4 rue de la libert??";"";"";"";"06000";"Nice";"FRANCE";"4 rue de la libert??--06000 Nice";"M.";"Eric";"Di Filippo";"edf@bigbossstudio.com";"+33 480805570";"CEO, Partner";"16/03/2015";"16/09/2015";"35";"129";"26";"10h - 12h30
14h-18h30";"";"508,20";"Euros";"508,20";"";"Tickets restaurant (20 X 8,50)";"Ing??nieur d??veloppeur d???applications mobiles";"??? participer ?? l?????tude et au d??veloppement de nouvelles applications mobiles en
prenant soin de respecter les ??tapes et les bonnes pratiques relatives au
d??veloppement logiciel, ce qui consistera en :

- participer aux sp??cifications en apportant l???expertise technique aux clients

- concevoir et mod??liser l???application

- d??velopper l???application et mettre en place des tests automatis??s

- participer ?? la recette de l???application

- mettre en production sur les stores d???applications

??? participer ?? la maintenance et ?? l???am??lioration des applications d??velopp??es par
Big Boss Studio

??? rechercher de nouveaux axes de d??veloppement et de nouvelles m??thodes par la
veille technologique et la r??alisation de prototypes ou par la participation aux
r??ponses aux appels d???offres

En parall??le et dans une d??marche d???am??lioration continue, participer au
d??veloppement des frameworks et outils internes permettant la mutualisation de code
et l???am??lioration de la qualit?? des applications d??velopp??es par Big Boss Studio.
";"4 rue de la libert??";"";"";"";"06000";"Nice";"FRANCE";"4 rue de la libert??--06000 Nice";"M.";"Cyril";"Chandelier";"cchandelier@big-boss-studio.com";"+33 683996977";"iOS Developer";"M.";"Philippe";"Renevier";"Philippe.RENEVIER@unice.fr";"+ 33 6 1844 9350";"";"";"";
"2015-03-10 10:19";"SI 5";"oui";"oui";"M.";"Hadrien";"LUTTIAU";"191047919128814";"21209982";"28 impasse des Grillons";"";"";"";"34540";"Balaruc les Bains";"FRANCE";"28 impasse des Grillons--34540 Balaruc les Bains";"hadrien.luttiau@gmail.com";"0782266556";"09/04/1991";"MAIF";"1112916 M";"National Institut of Informatics";"www.nii.ac.jp/en/";"2-1-2 Hitotsubashi, Chiyoda-ku";"";"";"";"101-8430";"TOKYO";"JAPAN";"2-1-2 Hitotsubashi, Chiyoda-ku--101-8430 TOKYO--JAPAN";"M.";"Henri";"ANGELINO";"nii-internship@nii.ac.jp";"+81-3-4212-2165";"Acting Director Global Liaison Office";"16/03/2015";"10/09/2015";"35";"123";"26";"The NII Internship program is not an employment, so it???s not appropriate to write as ""working hours"".
Instead, you could put like ""on a full-time basis""";"";"171000";"Yen";"1304,15";"";"";"Smart Service Compositions/Mashups in the City and the Web";"Intern students are expected to learn further knowledge and skills through collaborative activities for promoting our research activities and/or exploring further research topic. Depending on the status of the intern students (e.g., already have active master/phd topics or not), actual work in the internship can be determined flexibly through discussion. The outputs are typically joint papers and/or implemented software, but depend on the topic and the duration of the internship.

Subject:
Smart Service Compositions/Mashups in the City and the Web
Adaptation and Evolution Techniques in Service-based Systems (with Paradigms of ""Models at runtime"" / ""Requirements at runtime"")
Extension of Service Composition Techniques for Internet-of-Things (IoT) Services";"2-1-2 Hitotsubashi, Chiyoda-ku";"";"";"";"101-8430";"TOKYO";"JAPAN";"2-1-2 Hitotsubashi, Chiyoda-ku--101-8430 TOKYO--JAPAN";"M.";"Fuyuki";"ISHIKAWA";"f-ishikawa@nii.ac.jp ";"+81-3-4212-2675";"Associate Professor";"M.";"Ga??tan ";"REY";"gaetan.rey@unice.fr";" +33 4 92 96 51 44";"";"";"";
"2015-02-12 20:56";"SI 5";"oui";"non";"M.";"Nicolas";"MARQUEZ";"191079913816274";"20901829";"36 A val des castagnins";"les jardins de Ste Agn??s";"";"";"06500";"Menton";"FRANCE";"36 A val des castagnins--les jardins de Ste Agn??s--06500 Menton";"nmarquez@polytech.unice.fr";"0666664070";"10/07/1991";"PACIFICA";"1544166904";"Monaco Asset Management";"";"27 Boulevard Princesse Charlotte";"";"";"";"98000";"Monaco";"MONACO";"27 Boulevard Princesse Charlotte--98000 Monaco--MONACO";"M.";"Anthony";"Torriani";"atorriani@monacoasset.com";"+377 97 97 64 00";"CEO";"16/03/2015";"16/09/2015";"39";"129";"26";"8h30-12h30
14h-17h48";"";"1000";"Euros";"1000";"";"";"D??veloppement d???un outil de gestion de portefeuille";"Notre soci??t?? d??veloppe actuellement en interne un logiciel de gestion de portefeuille.
Le stagiaire sera en charge de consolider l?????quipe de d??veloppement et sera missionn?? sur deux axes importants :

- Am??liorer l???interface (c??t?? client ??? Visual C++)

- D??velopper de nouvelles fonctionnalit??s pour l???outil (Dont d??veloppement de proc??dures 

stock??es / fonctions sous PostgreSQL)";"27 Boulevard Princesse Charlotte";"";"";"";"98000";"Monaco";"Monaco";"27 Boulevard Princesse Charlotte--98000 Monaco--Monaco";"M.";"Vincent";"Froidefond";"VFroidefond@monacoasset.com";"(+377) 97 97 64 14";"IT Manager & Administration";"M.";"Anne-Marie";"HUGUES";"amhc@wanadoo.fr";"0684045930";"";"";"";
"2015-02-25 12:41";"SI 5";"non";"non";"M.";"Hugo";"MARTINEZ";"192045951258782";"20905088";"10 domaine de l???oratoire";"";"";"";"83300";"Draguignan";"FRANCE";"10 domaine de l???oratoire--83300 Draguignan";"hugo.83300@gmail.com";"0629698652";"08/04/1992";"LMDE";"5625166";"Worldline";"http://worldline.com";"80 Quai Voltaire ";"";"";"";"95870 ";"BEZONS";"FRANCE";"80 Quai Voltaire--95870 BEZONS";"M.";"Alice ";"WAUCQUIER";"alice.waucquier@atos.net";"03.20.60.79.42";"Recruteuse";"23/03/2015";"18/09/2015";"37";"124";"26";"";"";"1300";"Euros";"1300";"Virement banquaires";"Acc??s au restaurant d???entreprise

prise en charge des transports en commun ?? hauteur de 50%";"ADAPTIVE WEB DESIGN ET WEBSOCKET";"Le stage d??marre par l?????tude des solutions d???identification des 

informations de l???environnement de l???utilisateur et de son contexte 

d???utilisation (mobile derni??re g??n??ration, r??seau edge/3g/4g/wifi, ...), 

ainsi que l?????tude de la mise en place d???une solution de communication 

serveur vers client via les technologies WebSocket, Server-Sen events.

Le stagiaire proposera ensuite le cadre d???utilisation de ces donn??es dans 

nos services existants.

Puis il proc??dera ?? l???impl??mentation des propositions retenues en 

veillant toujours ?? optimiser les performances et ?? am??liorer l???exp??rience 

utilisateur.

Les technologies abord??es et produits utilis??s seront :

??? HTML5 : Websocket / eventSource

??? jQuery+Backbone.js

??? JSON

??? Java / SQL

??? Apache / Tomcat

??? Eclipse / Maven / Mercurial / Jenkins";"ZIA Rue de la Pointe";"";"";"";"59113 ";"SECLIN";"FRANCE";"ZIA Rue de la Pointe--59113 SECLIN";"M.";"Sylvain ";"POLLET-VILLARD";"sylvain.polletvillard@worldline.com";"+33 3 20 60 68 84";"Ing??nieur Etudes & Developpement";"M.";"Michel";"Buffa";"buffa@unice.fr";"(33)-92-07-66-60";"";"";"";
"2015-04-09 14:22";"SI 5";"non";"non";"M.";"Eric";"MASOERO";"191030608844255";"20905415";"171, chemin Barella";"";"";"";"06390";"Contes";"FRANCE";"171, chemin Barella--06390 Contes";"eric.masoero@gmail.com";"0621423027";"15/03/1991";"Aviva";"74486834";"Orange";"";"78 rue Olivier de Serres";"";"";"";"75505";"Paris Cedex";"FRANCE";"78 rue Olivier de Serres--75505 Paris Cedex";"Mme";"Fabienne";"Patinet";"fabienne.patinet@orange.com";"02 38 42 91 71";"DRH Site";"13/04/2015";"25/09/2015";"35";"115";"24";"";"";"1311,77";"Euros";"1311,77";"";"";"Nouveaux enablers pour le Mail et les Contacts Orange";"Mission :Orange propose ?? ses clients une suite d???outils de communication (Mail, SMS, MMS, Instant Messaging, Carnet
d???adresses, Calendrier) accessible depuis le Portail Orange mais ??galement depuis diff??rents supports : mobiles, tablettes, PC,
TV. Ces produits ??voluent rapidement, en fonction des avanc??es technologiques et marketing (exemple : Cloud, r??seau sociaux,
webIM...) mais ??galement en fonction des supports et des OS.
Au sein d???une ??quipe pluridisciplinaire et d???un projet de grande envergure, l?????tudiant devra contribuer ?? l???am??lioration des
services Mail et Carnet d???adresses d???Orange, permettant ?? l???op??rateur de rester leader ou ?? la pointe du march?? sur ces
produits strat??giques. L?????tudiant sera donc amen?? ?? ??tudier, int??grer, d??velopper et produire de nouvelles fonctionnalit??s
pertinentes pour l???utilisateur final.
En particulier, il s???int??ressera dans le cadre de ce stage aux Enablers, API et Backends Mail et Carnet d???adresses PIM. Il devra
??tudier, am??liorer ou d??velopper, puis int??grer de nouvelles API (REST), Enablers ou Backends. Il participera ??galement ?? la
mise en production de ces nouvelles API pour l???am??lioration des services Mail et Contacts Orange.
";"DSI DIGITAL FACTORY";"790 AV DR MAURICE DONAT";"BATIMENT MARCO POLO C2";"";"06250";"Mougins";"FRANCE";"DSI DIGITAL FACTORY--790 AV DR MAURICE DONAT--BATIMENT MARCO POLO C2--06250 Mougins";"M.";"Benoit";"Mercier";"benoit1.mercier@orange.com";"04 97 46 28 61";"Chef de projet SI";"M.";"Fabien";"Hermenier";"fabien.hermenier@unice.fr";"04 92 38 76 36";"";"";"";
"2015-03-29 21:08";"SI 5";"non";"non";"M.";"Marouan";"MESDOURI";"1880999000000093";"21410536";"25 , rue Robert Latouche R??sidence Jean Medecin";"";"25 , rue Robert Latouche R??sidence Jean Medecin";"";"06200";"NICE";"FRANCE";"25 , rue Robert Latouche R??sidence Jean Medecin--25 , rue Robert Latouche R??sidence Jean Medecin--06200 NICE";"marouanmesdouri@gmail.com";"0667302769";"02/09/1988";"lcl";"2006767539200";"APS int";"www.onduleurs.fr";"15 bis corniche Andr?? de Joly";"";"";"";"??06300";"NICE";"FRANCE";"15 bis corniche Andr?? de Joly--06300 NICE";"M.";"Jean Claude ";"??MATHIEU";"jcmathieu@onduleurs.fr";"06 11 35 53 73";"??G??rant";"09/04/2015";"09/09/2015";"35";"107";"22";"";"";"508.20";"Euros";"508.20";"";"";"DEVELOPPEMENT SITES WEB MARCHANDS";"SITE ONDULEURS.FR
D??veloppement et am??lioration du site de vente en ligne onduleurs.fr
D??veloppement d???un syst??me de veille concurrentielle sur les sites concurrents
Synchronisation des deux syst??mes entre eux en fonction de la politique tarifaire de la soci??t??.
Syst??me utilis?? PHP MYSQL 
ITE EVENTBOUTIK.COM Conception d???un Back office / CRM pour g??rer le site ";"15 bis corniche Andr?? de Joly";"";"";"";"06300";"NICE";"FRANCE";"15 bis corniche Andr?? de Joly--06300 NICE";"M.";"??Philippe";"??MATHIEU";"pmathieu@onduleurs.fr";"04 93 89 19 29";"Gestionnaire informatique";"M.";"Igor";"LITOVSKY";"lito@polytech.unice.fr";"04 92 96 51 24";"";"";"";
"2015-02-17 11:53";"SI 5";"oui";"non";"M.";"Thomas";"MONTANA";"1 93 10 75 112 632 12";"mt003150";"37c, Boulevard Gorbella";"Palais Johnny";"";"";"06100";"Nice";"FRANCE";"37c, Boulevard Gorbella--Palais Johnny--06100 Nice";"thmsmontana@gmail.com";"0674067141";"11/10/1993";"BPCE Assurances";"006108944";"GoEuro Corp.";"http://www.goeuro.com/";"Sonnenburger Str. 73";"";"";"";"10437";"Berlin";"GERMANY";"Sonnenburger Str. 73--10437 Berlin--GERMANY";"M.";"Naren";"Shaam";"naren.shaam@goeuro.com";"+4915234071141";"Chief Executive Officer";"01/04/2015";"31/08/2015";"40";"107";"22";"";"";"800";"Euros";"800";"";"???";"Web Frontend - Software Engineering Internship";"Due to the nature of our business and small length of release cycles, there will not be any long lasting projects nor projections on upcoming tasks. Oppositely, the internship will consist of a large number of small/medium length tasks. The whole picture of high velocity web and mobile web frontends will approached during this internship, as well as best practices of software development. 

We set up aggressive goals in terms of velocity, scalability, testability and maintainability. Thomas will be involved, from an abstract manner as of now, in all this.

From a technological standpoint, JavaScript will be the main language, in addition to HTML5, CSS3 and many others.";"Sonnenburger Str. 73";"";"";"";"10437";"Berlin";"Germany";"Sonnenburger Str. 73--10437 Berlin--Germany";"M.";"Louis";"Hache";"louis.have@goeuro.com";"+49 162 770 69 37";"Software Engineer / Web Frontend";"M.";"Michel";"Buffa";"buffa@polytech.unice.fr";"0662659345";"";"";"";
"2015-01-23 18:12";"SI 5";"oui";"oui";"M.";"J??r??me";"RANCATI";"192040602701944";"rj001117";"Traverse du vieux four cidex 434";"06330 Roquefort les Pins";"";"";"06330";"Roquefort les Pins";"FRANCE";"Traverse du vieux four cidex 434--06330 Roquefort les Pins--06330 Roquefort les Pins";"jerome.rancati@gmail.com";"+33628518057";"02/04/1992";"PROTEC BTP";"106170540 Y 003";"SnT Research Center";"http://wwwfr.uni.lu/snt";"4 Rue Alphonse Weicker, L-2721 Luxembourg";"";"";"";"L-2721";"LUXEMBOURG";"LUXEMBOURG";"4 Rue Alphonse Weicker, L-2721 Luxembourg--L-2721 LUXEMBOURG--LUXEMBOURG";"M.";"Bj??rn";"OTTERSTEN";"bjorn.ottersten@uni.lu";"(+352) 46 66 44 5665";"Directeur du Centre Interdisciplinaire ""Security, Reliability and Trust""";"15/03/2015";"15/09/2015";"40";"129";"26";"";"";"850";"Euros";"850";"";"";"Internship at the SnT Research Center working on a SmartHomes project";"In SmartHomes and SmartCities in general, various sensors and actuators are installed in the environmentto sense users activites. This not only allows to collect information about activities, but also to remotelyactivate devices to automate some scenario.
The ""Smart"" aspect of such environment is, nowadays, limited to reactive actions often encoded as sets ofrules interpreted by a specific engine to perform the automation. But the complexity to develop andconfigure such environment, specialized for each users, is today a blocking aspect. We are studying theuse of ""live machine learning"" to create a novel SmartHome engine, which should learn from users habits toalign the home settings to users??? everyday activity (i.e., by tuning the luminosity automatically based on the
regular manual setting of user).
This internship will offer two main aspects in collaboration with our industrial partner.
Firstly, the internship student will collaborate with the research team in order to learn and develop amachine learning algorithm to correlate sensors data in live. The outcome of this step will be integratedinto the Kevoree Modelling Framework.
Secondly, the student will have to develop an innovative web based dashboard (technology : GooglePolymer + Kevoree Modeling Framework), leveraging this machine learning algorithm. This dashboard should help users to understand and anticipate actions taken autonomously by the home with thissmart non-rules based engine (i.e., by displaying prediction and potential actions in future).
Because this project is included in a long run research activities, this subject has a great potential to becontinued as a PhD.
The requirements for this internship are:
High programming skills(Java/JS mainly).
Some knowledge about IoT/Sensors environments to understand how to select and process data.
Modeling knowledge to represent complex graph of dataThe machine learning aspect can be learnt during the internship, a first knowledge is a plus, but is not
mandatory.
Useful Links:
Google Polymer Framework: https://www.polymer-project.org/
Kevoree Modeling Framework: http://kevoree.org/kmf/";"6 rue Richard Coudenhove-Kalergi L-1359 Luxembourg-Kirchberg";"";"";"";"L-1359 ";"Luxembourg-Kirchberg";"LUXEMBOURG";"6 rue Richard Coudenhove-Kalergi L-1359 Luxembourg-Kirchberg--L-1359 Luxembourg-Kirchberg--LUXEMBOURG";"M.";"Fran??ois";"FOUQUET";"francois.fouquet@uni.lu";"+352 466644 5387 ";"Research Associate";"M.";"S??bastien";"MOSSER";"mosser@i3s.unice.fr";"0493925058";"";"";"";
"2015-03-19 13:29";"SI 5";"non";"non";"Mlle";"Roberta";"ROBERT";"000000000000";"21406870";"2255 Route des Dolines";"Chambre 203";"";"";"06560";"VALBONNE";"FRANCE";"2255 Route des Dolines--Chambre 203--06560 VALBONNE";"fleur.robert@gmail.com";"0695226334";"22/01/1992";"LCL";"1514341904";"SAP";"www.sap.com";"Tour SAP";"35 rue d???Alsace";"";"";"92309";"LEVALLOIS-PERRET";"FRANCE";"Tour SAP--35 rue d???Alsace--92309 LEVALLOIS-PERRET";"Mme";"Agn??s";"DESPLECHIN";"sylvine.eusebi@sap.com";"04 92 28 62 00";"Responsable du Service Recrutement";"01/04/2015";"31/07/2015";"35";"85";"18";"";"";"1 174";"Euros";"1 174";"";"";"Proxified Dynamic Security Testing";"This internship is based in the SAP Labs France Research Lab, in Sophia-Antipolis. The work will be
performed in the context of the Research Program ???Security & Trust???, and deals with dynamic security testing
and secure coding.
SAP has released a powerful platform called SAP HANA, which comprises two main components: the HANA
DB, an extremely efficient in-memory database, and the HANA XS engine, a server-side javascript based
application server. HANA application development is being done through an Integrated Development Editor
called WebIDE.
The development of an SAP application requires going through an internal security development lifecycle,
aiming at reducing the likeliness of introducing vulnerability at each development stage. When it comes to
the implementation phase, detecting code mistakes as they are being typed is the best way to ensure secure
code. For HANA development, a security guide was developed with requirements and best practices for
ensuring good code quality. The next step is to automate the verification of what is recommended in the
guide by the WebIDE itself.
A first prototype was already developed to this end, with a focus on static code analysis and on some
dynamic analysis. The next step is to extend on the dynamic analysis aspect, by keeping in mind the strong
requirements which come with it on performance and security. Dynamic security is on a thin edge between
defensive and offensive security. Offensive tools are a great asset for verifying certain security properties of
the application but can have detrimental effect if misused.
The goal of this internship is thus to assess different dynamic security testing scenarios and to implement a
proof of concept for those which comply to the performance and security requirements, in an effort to
enhance the WebIDE with a full set of offensive tools which can be used in a controlled fashion. One such
scenario will be to assess the possibility to develop and run an extension of or a tool similar to sqlmap
against SAP HANA applications through the WebIDE.
The candidate should be an expert in SQL, JavaScript and Web technologies. He should have a strong
background in (defensive) IT security and in penetration testing. He should be familiar with penetration
testing tools such as sqlmap and metasploit.
We expect that 20% of time will be dedicated to research activities, and 80% to development.";"805 Avenue Maurice Donat  Le Font de l???Orme";"";"";"";"06259";"MOUGINS";"FRANCE";"805 Avenue Maurice Donat  Le Font de l???Orme--06259 MOUGINS";"M.";"C??dric";"HERBERT";"cedric.hebert@sap.com";"04 92 28 62 00";"Team Leader";"M.";"Karima";"BOUDAOUD";"karima.boudaoud@polytech.unice.fr";"04 92 96 51 72";"";"";"";
"2015-03-20 14:23";"SI 5";"non";"non";"M.";"Victor";"SALL??";"193010502304469";"21206825";"2400 route des Dolines";"Res. Newton, Apt. A219";"";"";"06560";"VALBONNE";"FRANCE";"2400 route des Dolines--Res. Newton, Apt. A219--06560 VALBONNE";"victorsalle@outlook.com";"07.89.60.41.85";"28/01/1993";"Maaf";"159196140J002";"AVISTO";"www.avisto.com";"Space Antipolis 1";"2323, Porte 15, Chemin Saint-Bernard";"";"";"06220";"VALLAURIS";"FRANCE";"Space Antipolis 1--2323, Porte 15, Chemin Saint-Bernard--06220 VALLAURIS";"M.";"Christophe";"PARMENTIER";"christophe.parmentier@avisto.com";"04.92.38.74.71";"Responsable R??gional";"25/03/2015";"25/09/2015";"35";"129";"26";"9h-12h / 14h-18h";"";"915";"Euros";"915";"Mensuellement par virement avec fiche de paie";"- 1 ticket restaurant (8,60???) par jour
- 16,67??? par mois de frais de d??placement";"D??veloppement Android - MIDI DockStation & Android App";"Dans le cadre de d??veloppement d???une station d???accueil (dock station) MIDI et audio pour smartphone Android, vous participerez au d??veloppement (conception), au prototypage et ?? une ??tude de co??t pr??liminaire d???industrialisation de cette station d???accueil.

En suppl??ment de la fonction de chargement d???un smartphone, cette station d???accueil, destin??e avant tout aux audiophiles et compositeurs/musiciens, permet :
  - Le mixage et le traitement de sources sonores externes et du flux audio d???un smartphone Android,
  - La g??n??ration sonore gr??ce ?? une puce sonore propri??taire Elsys Design,
  - L???enregistrement du flux audio provenant de la station d???accueil sur carte SD du smartphone via l???application Android d??di??e,
  - Le contr??le d???instruments MIDI via toute application Android d??di??e,
  - La transmission d???information MIDI via WiFi (MIDI over WiFi)


Pour cela, vous devrez:

1. Valider les fonctionnalit??s MIDI et audio USB de la station d???accueil :
  - Mise en place et prise en main des outils de d??veloppement (SDK) pour Android,
  - Etude de la norme USB et des diff??rentes classes associ??es,
  - Conception et d??veloppement d???une application Android permettant d???envoyer/recevoir un flux audio vers/de la station - accueil via le port USB,
  - Conception et codage d???une application Android permettant d???envoyer/recevoir un flux MIDI vers/de la station d???accueil via le port USB,
  - Test et validation des fonctionnalit??s audio et MIDI USB de la station d???accueil,
  - Documents de conception et de validation

2. Cr??er une application audio media player/recorder permettant la lecture et l???enregistrement d???un fichier audio .WAV de/vers une carte SD ?? partir de la station d???accueil :
  - D??finir le cahier des charges de l???application,
  - Conception et codage de l???application en JAVA,
  - Test et validation de l???application avec la station d???accueil,
  - Documents de conception et de validation

3. Cr??er une application sp??cifique permettant le s??quen??age d?????v??nements MIDI, la lecture et l???enregistrement audio, le contr??le de la station d???accueil :
  - D??finir le cahier des charges de l???application,
  - D??finir des diff??rents formats de messages exclusifs MIDI entre l???application et la station d???accueil,
  - Conception et d??veloppement de l???application en JAVA,
  - Test et validation de l???application avec la station d???accueil,
  - Documents de conception et de validation";"Space Antipolis 1";"2323, Porte 15, Chemin Saint-Bernard";"";"";"06220";"VALLAURIS";"FRANCE";"Space Antipolis 1--2323, Porte 15, Chemin Saint-Bernard--06220 VALLAURIS";"M.";"Pierre";"PACCHIONI";"pierre.pacchioni@avisto.com";"04.92.38.74.72";"Directeur Technique";"M.";"Jean-Yves";"TIGLI";"jean-yves.tigli@unice.fr";"04.92.96.51.81";"";"";"";
"2015-03-12 01:19";"SI 5";"non";"non";"M.";"Rodrigo Augusto";"SCHELLER BOOS";"01B0JP00M31";"21407590";"45 Boulevard Pape Jean XXIII";"logement 43";"";"";"06300";"NICE";"FRANCE";"45 Boulevard Pape Jean XXIII--logement 43--06300 NICE";"rsboos@hotmail.com";"0782744908";"02/03/1995";"LCL";"1234556789";"Cadence Design Systems";"http://cadence.com/";"2655 Seely Avenue";"";"";"";"95134";"SAN JOSE";"USA";"2655 Seely Avenue--95134 SAN JOSE--USA";"M.";"Carolle";"PLANAS";"opicot@cadence.com";"04.89.87.30.00";"Human Resources Manager";"20/03/2015";"21/08/2015";"35";"108";"22";"";"";"1200";"Euros";"1200";"";"Droit au conges payes
Droit au titres restaurant
Droit au remboursement transport";"D??veloppement du Virtuoso Dashboard - environment graphique pour les outils Virtuoso";"Les outils Virtuoso sont des logiciels graphiques tr??s complexes (centaines de commandes, menus, langages d???extensions) manipulant des quantit??s importantes de donn??es. Cadence souhaite d??velopper un environnement graphique pour ses clients (Virtuoso Dashboard) permettant aux responsables des ??quipes de conception de mesurer les performances du logiciel (temps d???affichage, temps d???acc??s ?? la base de donn??es, etc.) et les taux d???utilisation des fonctionnalit??s les plus avanc??es. 
Ce logiciel acc??dera entre autres aux fichiers logs de Virtuoso, les organisera pour un acc??s rapide (data mining) et affichera des indicateurs de performance et d???utilisation.";"1080 route des dolines";"";"";"";"06560";"VALBONNE";"FRANCE";"1080 route des dolines--06560 VALBONNE";"M.";"Olivier";"PICOT";"opicot@cadence.com";"04.89.87.30.00";"Project Manager";"M.";"Michel";"RUHER";"ruher@i3s.unice.fr";"06";"";"";"";
"2015-03-17 19:27";"SI 5";"oui";"non";"M.";"Qihao";"TANG";"193029921600046";"21210465";"14 Rue de le Republique";"";"";"";"06600";"Antibes";"FRANCE";"14 Rue de le Republique--06600 Antibes";"qihao.tang@aliyun.com";"+33760456025";"21/02/1993";"LMDE";"06291980";"Everbright Securities";"http://www.ebscn.com/";"25 Taipingqiao St. Xicheng";"";"";"";"100034";"Beijing";"CHINA";"25 Taipingqiao St. Xicheng--100034 Beijing--CHINA";"M.";"Xinyuan";"Ge";"gexy@ebscn.com";"+8613528809518";"directeur g??n??ral";"30/03/2015";"30/09/2015";"40";"129";"26";"";"";"24000";"RMB";"3500";"";"";"Les produits d??riv??s financiers";"Le design et le pricing sur les produits d??riv??s financiers.
Am??lioration sur l???algorithme existant et d??v??loppement avec les programme.";"580 Nanjing West Rd. Jing???an";"";"";"";"200041";"Shanghai";"China";"580 Nanjing West Rd. Jing???an--200041 Shanghai--China";"M.";"Xinyuan";"Ge";"gexy@ebscn.com";"+8613528809518";"president g??n??ral";"M.";"Ioan";"Bond";"bond@polytech.unice.fr";"+33(0)492965141";"";"";"";
"2015-03-02 11:04";"SI 5";"non";"non";"M.";"William";"TASSOUX";"191103417258051";"21206992";"210 Avenue roumanille";"";"";"";"06410";"BIOT";"FRANCE";"210 Avenue roumanille--06410 BIOT";"william.tassoux@gmail.com";"0651894242";"30/10/1991";"Cabinet Gril Assurances";"025899952";"GFI Informatique";"http://www.gfi.fr/";"145 boulevard Victor Hugo";"";"";"";"93400";"SAINT-OUEN";"FRANCE";"145 boulevard Victor Hugo--93400 SAINT-OUEN";"M.";"St??phane";"Jourjon";"stephane.jourjon@gfi.fr";"0497155523";"Directeur de division";"09/03/2015";"09/09/2015";"35";"129";"26";"";"";"1200";"Euros";"1200";"Virement bancaire";"Carte avantage tickets restaurants
Prise en charge ?? hauteur de 50% des transports en commun sur justificatifs";"Evaluation, design and deployment of a distributed access control system";"The final purpose of the stage is to deploy a SSO (single-sign on) mechanism in the server-side application of a Graphical User Interface, in order to distribute authentication credentials to a plenty of possible systems the GUI is supposed to interact with.

The candidate will develop his project interacting with a team of 5/10 people, in GFI buildings located in Les Emeralds in Sophia Antipolis. 

The stage covers the requirements of a portion of a bigger project, meant to build-up a GUI able to communicate with multiple applications, to give the final users the possibility to interact with them through a single entry point.";"Emerald Square, B??timent B, Avenue Evariste Galois";"BP 199";"";"";"06904";"Sophia Antipolis";"FRANCE";"Emerald Square, B??timent B, Avenue Evariste Galois--BP 199--06904 Sophia Antipolis";"M.";"Emmanuel";"Jovenin";"emmanuel.jovenin@gfi.fr";"0489611288 ";"Delivery manager";"M.";"Jean-Paul";"Rigault";"jpr@polytech.unice.fr";"0492965133";"";"";"";
"2015-01-08 22:52";"SI 5";"non";"non";"Mlle";"Marie-Catherine";"TURCHINI";"291012B03319789";"21105993";"31 Boulevard PAOLI";"";"";"";"20200";"BASTIA";"FRANCE";"31 Boulevard PAOLI--20200 BASTIA";"turchini@polytech.unice.fr";"0615255288";"23/01/1991";"MAIF";"1530301D";"GoodBarber";"http://fr.goodbarber.com/";"12 Rue G??n??ral Fiorella";"";"";"";"20000";"AJACCIO";"FRANCE";"12 Rue G??n??ral Fiorella--20000 AJACCIO";"M.";"Dominique";"SIACCI";"siacci@duoapps.com";"0972134065";"Directeur g??n??ral";"30/03/2015";"27/09/2015";"35";"126";"26";"";"";"1500 NET";"Euros";"1500";"";"";"D??veloppement iOS (Objective-C) - D??veloppement d???un ensemble de fonctionnalit??s ?? destination des commerces de proximit??";"La plate-forme GoodBarber consiste en un ensemble de fonctionnalit??s permettant la cr??ation d???applications mobiles. Elle est le fruit de plusieurs mois de R&D ayant conduit ?? la cr??ation des diff??rents moteurs d???applications mobiles (pour iOS, Android, HTML5), qui constituent aujourd???hui une base technique tr??s solide pour l???ajout de nouvelles fonctionnalit??s.
Pour l???ann??e 2015, les d??veloppements au sein de GoodBarber vont s???orienter vers les interactions avec les utilisateurs finaux, et l???ajout de fonctionnalit??s du produit ?? destination des commerces de proximit??. Les objets connect??s et leur technologies associ??es (iBeacon, NFC, ???) seront une composante majeur de ces interactions. 

Ces fonctionnalit??s seront ax??es autour de la fid??lisation, et du concept drive-to-store. Dans ce cadre, nous imaginons une adaptation des moteurs techniques pour permettre le d??clenchement ??v??nementiel d???actions en fonction des profils utilisateurs, ou d???informations transmises par des objets connect??s, pour permettre la mise en place de scenarii marketing notamment ?? travers le geofencing.

Cet ensemble de fonctionnalit??s doit ??tre pens?? pour s???int??grer et s???inter-connecter au moteur GoodBarber, tant au niveau font-end natif (objet du stage) qu???au niveau back-end.";"12 Rue G??n??ral Fiorella";"";"";"";"20000";"AJACCIO";"FRANCE";"12 Rue G??n??ral Fiorella--20000 AJACCIO";"M.";"Dominique";"SIACCI";"siacci@duoapps.com";"0972134065";"Directeur g??n??ral";"M.";"Ga??tan";"REY";"Gaetan.Rey@unice.fr";"0492965144";"";"";"";
"2015-02-19 21:37";"SI 5";"non";"oui";"M.";"Michel";"VEDRINE";"1 91 06 78 646 509 73";"21209258";"41 rue Cl??ment Roassal";"";"";"";"06000";"Nice";"FRANCE";"41 rue Cl??ment Roassal--06000 Nice";"mvedrine@gmail.com";"0683158650";"23/06/1991";"MEP";"850445733";"CNRS D??l??gation C??te d???Azur";"";"Les Lucioles 1 ??? Campus Azur - 250 rue Albert Einstein CS 10 269";"";"";"";"06905";"Sophia Antipolis cedex";"FRANCE";"Les Lucioles 1 ??? Campus Azur - 250 rue Albert Einstein CS 10 269--06905 Sophia Antipolis cedex";"Mme";"B??atrice";"SAINT-CRICQ";"beatrice.saint-cricq@dr20.cnrs.fr";"0493954256";"d??l??gu??e R??gionale du CNRS C??te d???Azur";"16/03/2015";"16/09/2015";"35";"129";"26";"";"";"508,20";"Euros";"508,20";"Virements bancaires mensuels";"";"S??curisation des donn??es et de l?????change de donn??es dans un environnement mobile dans le cadre du projet PadDoc";"L???objectif de ce stage d???une dur??e de 6 mois sera de poursuivre un travail initi?? lors de 2 PFEs de derni??re ann??e d???ing??nieur concernant la s??curit?? de l?????change de donn??es en utilisant les protocoles de communication NFC et Bluetooth et la s??curit?? du stockage des donn??es sur un smartphone Android.

L?????tudiant devra donc :

- Finaliser l???impl??mentation des composants logiciels de s??curit?? n??cessaires pour s??curiser l?????change de donn??es en NFC et Bluetooth.
- Finaliser l???impl??mentation des composants logiciels n??cessaires ?? la s??curit?? du stockage des donn??es sur le mobile
- Faire une ??valuation de performances.
";"Laboratoire I3S UMR 7271 - GLC - Equipe RAINBOW ";"B??timent les Templiers 1 - Campus Sophi@Tech - 930 Route des Colles - BP 145";"";"";"06903";"Sophia Antipolis cedex";"FRANCE";"Laboratoire I3S UMR 7271 - GLC - Equipe RAINBOW--B??timent les Templiers 1 - Campus Sophi@Tech - 930 Route des Colles - BP 145--06903 Sophia Antipolis cedex";"Mme";"Karima";"Boudaoud";"karima@polytech.unice.fr";"0492965172 ";"Ma??tre de Conf??rences";"M.";"Jean-Yves";"Tigli";"tigli@polytech.unice.fr";"0684245567";"";"";"";
"2015-02-08 21:17";"SI 5";"non";"non";"M.";"Damien";"VIANO";"191090608832970";"21004752";"195";"chemin des Caillades";"";"";"06480";"La Colle sur Loup";"FRANCE";"195--chemin des Caillades--06480 La Colle sur Loup";"damien.viano06@gmail.com";"0667507301";"07/09/1991";"Maif";"2316612 M";"Sopra Steria";"http://www.soprasteria.com/";"3 RUE DU PRE FAUCON PAE DES GLAISINS 	 ";"";"";"";"74940";" ANNECY LE VIEUX";"FRANCE";"3 RUE DU PRE FAUCON PAE DES GLAISINS--74940 ANNECY LE VIEUX";"M.";"Fr??d??ric ";"Letot";"frederic.letot@soprasteria.com";" 0483150141";"Directeur d???agence";"09/03/2015";"30/09/2015";"35";"144";"29";"";"";"1120";"Euros";"1120";"";"Un acc??s ?? tous les avantages CE et une participation de Sopra aux frais de d??jeuner d???un montant de 5,28??? par jour travaill??.";"Mise en place d???une solution de statistiques et de supervision r??seaux";"Au sein de l?????quipe de d??veloppement de solutions de supervisions t??l??com, vous concevez
et ??laborez une solution de statistiques et de reporting r??seaux bas??e sur des produits
open source.
Vous ??tes en charge de??:
-> L?????tude des besoins et le cadrage du projet, pour les Network Operations Center
d???op??rateurs t??l??com,
-> Les sp??cifications fonctionnelles,
-> L?????tude de faisabilit?? et la proposition des solutions sur des briques Open Sources,
-> La r??daction du dossier d???architecture,
-> La mise en ??uvre des briques du projet sur un d??monstrateur,
-> L???int??gration des diff??rents composants sur un portail.";"Ath??na C, 1180 route des Dolines";"";"";"";"06904 ";"Sophia Antipolis";"FRANCE";"Ath??na C, 1180 route des Dolines--06904 Sophia Antipolis";"M.";"Bruno";"Cruanes";"bruno.cruanes@soprasteria.com";"04 83 15 00 00";"Directeur de projet";"M.";"Fr??d??ric";"Pr??cioso";"frederic.precioso@polytech.unice.fr";"0492965143";"";"";"";
"2015-03-17 16:57";"SI 5";"non";"non";"M.";"Robin";"VIVANT";"191113919828331";"21003782";"200 chemin du Beal";"";"";"";"06480";"La Colle sur Loup";"FRANCE";"200 chemin du Beal--06480 La Colle sur Loup";"robin.vivant@gmail.com";"0622699358";"23/11/1991";"MAIF";"1273527D";"OjingoLabs";"http://www.ojingolabs.com/";"2101 23rd Street ";"";"";"";"94107";"San Francisco";"USA - CALIFORNIE";"2101 23rd Street--94107 San Francisco--USA - CALIFORNIE";"M.";"TJ";"MARBOIS";"tj@ojingolabs.com";"0667317118";"Fondateur";"23/03/2015";"22/09/2015";"35";"128";"26";"";"";"3600";"Euros";"3600";"Virement banquaire";"";"Fullstack web engineer";"Vous int??grerez l?????quipe compos??e de quatre d??veloppeurs iOS, de deux d??veloppeurs serveurs Vert.X, de trois graphistes et de deux project manager. Durant votre stage, vous participerez ?? la phase de r??flexion et de cr??ation des projets en apportant votre vision technique sur la r??alisation d???applications aussi bien en terme d???ossature, d???ergonomie et de design de ces derni??res. Vous pourrez aussi conseiller le graphiste lors de la r??alisation des wireframes, et/ou maquettes.
Il y a besoin d???une solution monitoring pour les serveurs. Vous effectuerez une recherche sur ce qui se fait et participerez ?? la mise en place des clusters de monitoring et si besoin est, cr??erez des composants graphiques pour afficher des metrics personnalis??es.
La r??alisation de vues web sera ??galement demand?? pour certains besoins de l???application mobile et l???envoi de mails.";"1 place Massena";"";"";"";"06000";"NICE";"FRANCE";"1 place Massena--06000 NICE";"M.";"TJ";"MARBOIS";"tj@ojingolabs.com";"0667317118";"Fondateur";"M.";"Christian";"BREL";"brel@polytech.unice.fr";"04 92 96 51 62";"";"";"";`

type BufferReader struct {
	input string
}

type MockJournal struct{}

func (m *MockJournal) UserLog(u schema.User, msg string, err error)       {}
func (m *MockJournal) Log(em, msg string, err error)                      {}
func (m *MockJournal) Wipe()                                              {}
func (m *MockJournal) Access(method, url string, statusCode, latency int) {}
func (m *MockJournal) Logs() ([]string, error)                            { return []string{}, nil }
func (m *MockJournal) StreamLog(k string) (io.ReadCloser, error)          { return nil, nil }

func (b BufferReader) Reader(year int, promotion string) (io.Reader, error) {
	return strings.NewReader(b.input), nil
}

/*
func TestCSVParsing(t *testing.T) {
	r := BufferReader{input: buf}
	x := NewCsvConventions(r, []string{"si5"})
	conventions, errors := x.Import()
	assert.Nil(t, errors)
	assert.Equal(t, 0, len(conventions)) // TODO!
}
*/
