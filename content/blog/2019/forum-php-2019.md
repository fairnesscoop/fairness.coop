---
title: "Retour sur le ForumPHP 2019"
publishDate: "2019-11-08"
description: Nous étions au ForumPHP 2019 qui se tenait à guichet fermé encore une fois au Paris Marriott Rive Gauche.
readtime: 15
thumbnail: lightningtalk.jpg
author: Les coopérateurs Fairness
---
Des conférences, des *lightning-talks*, un "bar Open Source" mais aussi pour la première fois, un service de [vélotypie](https://afup.org/news/1060-la-velotypie-sinvite-au-forum-php-2019) était disponible pour retranscrire les conférences en direct, bien pratique pour les malentendants et même pour les quelques conférences en anglais.

![Echanges avec Kévin Dunglas pour podcast echo](/blog/images/podcast.jpg)

photo : [Julien Pauli](http://pic.jpauli.tech/#15721175616118)

## "Writing effective PHP" par Nuno Maduro

Faire du *Defensive programming* c'est comme conduire une voiture en toute sécurité selon Nuno. Pour éviter les bugs, un principe assez simple à implémenter : l'immutabilité des objets. Le seul point d'injection d'une classe devrait être à l'instanciation de l'objet via son contructeur.

De plus pour s'assurer de l'intégrité de cette instance, Nuno recommande de vérifier les entrées dans le constructeur en usant et abusant d'assertions et bien sûr en générant des exceptions si les paramètres d'entrée ne sont pas valides (InvalidArgumentException par exemple). Nuno insiste aussi sur le besoin de documenter son code. Une classe est en quelque sorte une API et la documenter permet d'avoir un contrat clair.

Passons nos classes en final. Une classe finale ne peut pas être héritée. Car étendre une classe qui étend une classe qui étend une autre... n'est pas une bonne pratique de code. L'héritage de classe est souvent signe d'une mauvaise architecture du code. Au lieu de l'héritage de classe, utilisons plutôt la [composition](https://speakerdeck.com/nunomaduro/writing-effective-php?slide=20).

Une bonne conférence de Nuno qui revient aux bases et invite à la minutie dans la création de nos classes.

## "Une application résiliente, dans un monde partiellement dégradé" par Pascal Martin

Quoi ? En lisant le titre, nous avons cru à une conférence sur la collapsologie. En fait, l'effondrement c'est le quotidien de nos applications. Les exemples sont courants. Pour la déclaration de revenu, et face à une saturation des serveurs, le ministère des finances a par exemple donné 48h supplémentaire pour faire la procédure.

Préférez-vous un service non disponible ou un service qui perde vos données ? Pascal nous invite à privilégier la satisfaction de l'utilisateur plutôt que de se satisfaire d'une vision technique de nos services.

Exemples : Quel pourcentage d'ajouts au panier réussis ? Combien de pages qui s'affichent en moins de 1,5 secondes ?

L'exemple #internetOfShit : des gens ne pouvaient pas rentrer chez eux à cause d'une panne des services de Google bloquant leur verrou connecté. C'est l'enfer des systèmes distribués.

Habituons-nous à ce qu'un service échoue. À l'échelle, les choses échouent beaucoup selon Pascal. Un système distribué n'est jamais tout le temps opérationnel. Il vit dans un écosystème partiellement dégradé.

Au lieu de la rendre complètement indisponible, une application pourrait être disponible en mode dégradé, donc pas forcément avec tous les contenus, encore une fois pour satisfaire un utilisateur.
Les pistes techniques :

* dimensionner en fonction de la durée des requêtes

* échouer rapidement avec des timeouts courts

* réessayer en cas d'échec avec un nombre d'essais limité

* *feature flipping* : désactiver les fonctionnalités coûteuses manuellement ou automatiquement

* déploiement continu

Selon Pascal, le métier du devops ce n'est pas que coder, configurer, déployer, mais c'est surtout fournir un service. Choisissons des technologies connues et éprouvées (boring technology). Chez Fairness on valide !

Et en cas d'incendie ? Communiquons en interne auprès des collègues, auprès des autres services, en externe (les clients) et régulièrement. Ayons des solutions toutes prêtes (ce qui signifie qu'elles ont été préparées), suivons des procédures, relayons-nous sur l'incident et sans oublier de se reposer !

Et nous n'avons évidemment pas tout retranscrit de cette conférence dense et très intéressante. Nous avons invité Pascal au micro de [podcast echo](https://podcastecho.github.io/) qu'on aura le plaisir d'entendre prochainement sur ce sujet.

## "Neuroatypie et IT : quelques conseils" par Alex Rock

Nous apprécions beaucoup les conférences non techniques. Nous ne pouvions pas rater celle d'Alex avec qui nous avions discuté lors d'un [épisode du podcast echo sur le recrutement des autistes](https://soundcloud.com/podcastecho/e24-recrutement-des-autistes-dans-l-it-avec-alex-rock).

Les autistes connus : Alan Turing, Greta Thunberg, Daniel Tammet qui parle 12 langues dont une qu'il a inventé, l'esperanto, et a appris l'islandais en une semaine.

Les causes potentielles de l'autisme : variations génétiques, facteurs environnementaux, dysfonctionnements cérébraux, déficits cognitifs, troubles psychologiques.
Les conséquences possibles : stress accru, inadaptabilité sociale, troubles de l'empathie, problèmes neuro-moteurs, incompréhension du second degré, angoisse, dépression, anxiété...

Microsoft a un programme de recrutement dédié pour les autistes : très majoritairement des hommes blancs hétéros. C'est un programme de diversité mais qui bizarrement n'inclut pas tout le monde. Les recommandations d'Alex si vous souhaitez recruter une personne neuroatypique :

* se renseigner sur les particularités de la personne, la laisser parler d'elle pour en savoir plus

* être clair et précis sur les process

* donner des conditions de travail adaptées. Les bureaux en *open-space* c'est 60% dans le secteur de la tech. Selon des études, l'open space est stressant pour n'importe qui. Pour la personne neuro-atypique c'est pire.

* communiquer : pratiquer l'écoute active, poser des questions

En fait, ces recommandations sont valables pour tous, pas que pour les neuro-atypiques.

En savoir plus : [épisode podcast echo Recrutement des autistes dans l'IT avec Alex Rock](https://soundcloud.com/podcastecho/e24-recrutement-des-autistes-dans-l-it-avec-alex-rock).

## "Si Darwin avait raison, l'agilité fonctionne par hasard" par François Zaninotto

Ou devrait-on dire, par "Professeur Eugène Fournier". François interprète ce professeur qui tente d'avoir un regard extérieur sur l'agilité en entreprise, en la comparant à la sélection naturelle selon Darwin. Conférence très décalée donc, peut-être un peu longue, mais rafraîchissante.

Comme la nature et les insectes, l'agilité en entreprise permet de s'adapter plus rapidement. Imposer les méthodes agiles c'est souvent intégrer un code génétique extérieur (ce qui a marché ailleurs) à une entreprise cible et forcément, on a une réaction de défense de l'hôte.

L'agilité augmente les variations et accélère l'adaptation. Les mutations qui sont générées au hasard sont parfois des succès, mais très souvent des échecs. Et ce n'est pas grave. Cela ajoute un peu de désorganisation mais il faut accepter les échecs pour avancer.

François (ou le professeur ?) propose ce néologisme en remplacement d'agiliste : "Agilitateur". La méthode Lean Startup permet de générer des erreurs et un jour par hasard, une fonctionnalité apparait. Pas faux ! D'où le titre ? L'agilité fonctionne par hasard.

Une startup, c'est une fourmi. Les banques et assurances sont des tortues. Expérimentons et échouons rapidement ("*Fail fast*"), comme une startup. Comme la fourmi, pas comme la tortue.

Quel est l'essentiel dans l'agilité ? François liste les 4 principes de la méthode "[Modern Agile](http://modernagile.org/)" :

* Faire de la sécurité un prérequis

* Apporter de la valeur en continu

* Expérimenter et apprendre rapidement

* Rendre les gens fantastiques

L'agilité n'est finalement pas révolutionnaire. Tout écosystème naturel peut évoluer rapidement sans dommage. Et dans ce que l'on conçoit, au delà de l'utile il y a le "beau".

François finit son "discours" avec un mot sur la catastrophe climatique et la disparition accélérée de la biodiversité.

## "Symfony HttpClient vs Guzzle vs HTTPlug" par Nicolas Grekas

Nicolas Grekas a illustré les différences entre Guzzle, HTTPlug et HttpClient afin de présenter le nouveau composant Symfony : HttpClient

En illustrant ses propos de démonstrations, il nous invite à tester Symfony HttpClient pour nos cas d'utilisations où l'on doit interroger des services externes.

Nicolas s'est beaucoup inspiré du meilleur des différentes bibliothèques et outils (curl, Guzzle, fopen, HttpPlug) et a pioché les bonnes idées des autres langages (axios, fetch, javascript, Golang) afin de proposer une interface la plus intuitive possible ([github.com/symfony/http-client-contracts](https://github.com/symfony/http-client-contracts)) de son composant [HttpClient](https://github.com/symfony/http-client). Et bonus pour le développeur, c'est qu'il n'a pas à devoir gérer manuellement des requêtes PSR-7 / PSR-18.

## "Tout pour se préparer à PHP 7.4" par Damien SEGUY

Damien ne s'est pas arrêté à la seule nouveauté du typage des propriétés mais a passé en revue l'ensemble des ajouts et suppressions à venir dans PHP. À base d'exemples, il a mis en avant ce qu'il y aura à changer de notre code pour pouvoir monter de version en toute transparence.

Bien évidemment, il est conseillé de se faire aider d'un outil d'analyse de code statique qui accompagnera les développeurs dans leur montée de version en leur remontant rapidement ce qui va changer.

## "Symfony Checker is coming" par Valentine Boineau

Nous n'avons pas pu assister à la conférence de Valentine mais nous l'avions rencontrée et nous avions échangé il y a quelques semaines sur le même sujet lors d'un épisode du podcast echo "[Symfony Checker, analyse de code et femmes dans la tech](https://soundcloud.com/podcastecho/e36-symfony-checker-analyse-de-code-et-femmes-dans-la-tech-avec-valentine-boineau)".

## "De CRUD à DDD, comment Meetic a sauvé son legacy" par Jean-Marie Lamodière

En 2015, les personnes chez Meetic ont commencé à réfléchir sur l'architecture technique qui était faite à base de CRUD et de fort couplage à la base de données. Le passage à une nouvelle architecture a été réalisé de manière progressive. Chaque nouvelle fonctionnalité devait être faite avec la nouvelle architecture, et si une nouvelle fonctionnalité avait besoin d'éléments codés avec l'ancienne architecture, c'était le moment de modifier l'ancienne. Ceci afin de ne pas inclure les anciens patterns dans la nouvelle façon de faire.

Très bonne mise en perspective de l'architecture hexagonale et du DDD (*Domain Driven Design*) par les équipes de Meetic : aide au maintien du code suite aux changements d'équipes, aux montées de versions et à l'évolution des technologies utilisées.

## "L'architecture progressive" par Matthieu Napoli

Mettre en place du DDD partout ? Non, ça dépend, c'est surtout pour des projets complexes. Car le CRUD, le SQL natif ou écrire des choses en dur ne sont pas des gros mots. Matthieu insiste sur le fait qu'un code legacy... finalement c'est un projet qui a réussi.

Et pensez à l'humain pour la pérennité du code : par exemple si un prestataire intervient sur une nouvelle base de code, il faut impliquer l'équipe entière afin qu'elle maîtrise la base de code.

## "Concevoir pour des futurs souhaitables" par Marie-Cécile Godwin et Thomas Di Luccio

À l'ère de l'anthropocène et des rapports du GIEC alarmants, Marie-Cécile et Thomas nous présentent différents scénarios de design-fiction, d'un futur géré par les GAFAM, un futur vert où on aurait tellement compensé nos émissions carbone en replantant des forêts qu'il n'y a plus de terres cultivables... Sont questionnés les rôles et les impacts de la tech et du design. A-t-on vraiement besoin d'une "dent connecté" ?

Enfin, Marie-Cécile et Thomas proposent des heuristiques pour agir dès maintenant :

* Prendre un nouveau départ en prenant conscience des limites planétaires et en redevant des terriens, démanteler ce qui nuit et repolitiser toutes nos activités, que l'on soit designer ou développeur.

* Adopter de nouvelles postures, s'armer émotionnellement, cultiver l'humilité et pratiquer une forme de radicalité.

* Expérimenter de nouvelles modalités d'action : préserver les savoirs et cultiver les communs, créer des outils ouverts et libres, adopter les low-tech, promouvoir la trans-disciplinarité, soyons exemplaires...

## Mais aussi

Nous avons adoré la keynote de fin "Pratiquons la physique avec Star Wars !" de Roland Lehoucq. Avec des exemples extraits de Star Wars, il a expliqué le vrai coût de l'énergie. Un sabre laser par exemple, c'est 1 gigawatt de puissance.

Par ailleurs, nous avons pris la parole lors d'un [lightning-talk de 5 minutes sur comment mesurer l'impact environnemental d'un site web](https://twitter.com/afup/status/1187397963385053189) en dressant les impacts environnementaux et sociaux notamment à la production de nos terminaux.

{{<rawhtml>}}
<p style="text-align: center"><img alt="Lightning Talk ForumPHP" src="/blog/images/lightningtalk.jpg" style="width: 50%;"></p>
{{< /rawhtml>}}

Nous avons traîné notre micro et avons enregistré [l'ambiance lors d'un épisode spécial ForumPHP](https://soundcloud.com/podcastecho/hors-serie-au-coeur-du-forumphp-2019) du podcast Echo. Écoutez aussi [notre entretien avec Kévin Dunglas pour un web éthique et décentralisé](https://soundcloud.com/podcastecho/38-favoriser-le-web-decentralise-et-ethique-avec-kevin-dunglas).

Nous tenions à remercier l'AFUP et tous les bénévoles pour l'organisation et à l'année prochaine les 20 ans !

