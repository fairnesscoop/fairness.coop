---
title: "Ceebios"
proposalTitle: "Démarche d'écoconception nouveau site Ceebios"
date: "2020-02-19"
author: Les coopérateurs Fairness
---

Sommaire

1. [Qui sommes nous ?](#1-qui-sommes-nous)
2. [Notre compréhension du besoin](#2-notre-compréhension-du-besoin)
3. [Méthodologie](#3-méthodologie)
4. [Ecoconception](#4-écoconception)
5. [Design UX / UI](#5-design-ux-ui)
6. [Préconisations technologiques](#6-préconisations-technologiques)
7. [Accessibilité](#7-accessibilité)
8. [Suivi d'audience](#8-suivi-d-audience)
9. [Hébergement](#9-hébergement)
10. [Tierce Maintenance Applicative (TMA)](#10-tierce-maintenance-applicative-tma)
11. [Résilience](#11-résilience)
12. [Contribution aux communs](#12-contribution-aux-communs)
13. [Livrables](#13-livrables)
14. [Budget et planning](#14-budget-et-planning)
15. [Conclusion](#conclusion)

## 1. Qui sommes nous ?

Fairness est une coopérative de conception responsable et agile des services numériques. Nous contribuons à l'économie
sociale et solidaire à la fois par notre statut mais aussi par nos actions : certification écoconception (formation greenit.fr),
membres ou contributeurs bénévoles auprès du collectif Conception Numérique Responsable, API Days sustainable initiative,
The Shift Project ([Carbonalyser](https://theshiftproject.org/carbonalyser-extension-navigateur/)), la Fing ou Latitudes.

Experts dans la réalisation de produits web et mobile, nous intégrons la qualité, l'accessibilité comme l'impact écologique
et sociétal dans les phases de conception puis d'accompagnement de nos clients.

[Mathieu, Nicolas, Thomas et Richard](/team) sont des artisans du logiciel : ils développent des services sur mesure, conseillent
leurs clients sur des choix d'architecture et donnent des formations autour des technologies Symfony, React ou Node.js et
l'écoconception.

[Hélène et Sébastien](/team) accompagnent et forment leurs clients autour des pratiques Lean et Agile. En tant que coach,
Scrum Master ou Product Owner, ils facilitent les discussions entre les parties prenantes de services numériques.

En savoir plus

* [Présentation de Fairness](/blog/2019/creation-cooperative-conception-responsable/)
* [Notre offre, compétences et démarche](/blog/2019/notre-offre/)
* [Nos références clients](/blog/2019/references/)

## 2. Notre compréhension du besoin

Ceebios souhaite refondre son site web en partant d'une feuille blanche en appliquant une démarche d'écoconception
c'est à dire réduire au minimum la quantité de ressources informatiques nécessaires pour apporter un service et une
information pertinentes aux utilisateurs du site. Il s'agit aussi d'une démarche de conception plus responsable
notamment en prenant en compte l'accessibilité dès le début du projet.

Le site actuel est trop gros : indice E de la page d'accueil sur ecoindex.fr, près de 6 Mo téléchargés et 114 requêtes
serveur effectuées.

Durant cette démarche d'écoconception, un guide de bonnes pratiques sera au fur et à mesure écrit à l'attention des
animateurs du site : rédacteurs, contributeurs, graphistes, développeurs...

La cible du site : les entreprises, les étudiants (une page dédiée) et le grand public (pour l'instant pas adressé
sur la version actuelle). Ceebios souhaite mettre mieux en avant ses activités. Il existe un espace adhérents avec
un seul mot de passe générique pour tous.
De plus, il est envisagé que le nouveau site soit en plusieurs langues (français, anglais).

L'optimisation du référencement naturel (SEO) est important : la sémantique HTML et les différentes entêtes seront
respectés. La conservation du référencement du site actuel est également nécessaire : garder les mêmes urls ou
rediriger les anciennes urls vers les nouvelles.

L'objectif est de publier cette nouvelle version avant l'évènement Biomim’expo qui se tien les 20 et 21 octobre 2020.

## 3. Méthodologie

Nous préconisons les approches agiles, Lean, d'innovation frugale, de circularité et d'éco-conception pour tous nos projets.
Elles nous permettront de faire les choix les plus pertinents en fonction des différentes problématiques et réflexions
rencontrées à chaque étape du projet.

Elles ont porté leurs fruits dans tous les projets pour lesquels nous l’avons préconisée et utilisée. Nous essayons toujours
de rester agile par rapport à cette méthodologie et donc de nous adapter au contexte projet.

### Vision produit

La vision d’un produit permet de donner une direction aux développements réalisés. Chaque membre de l’équipe projet est
plus impliqué lorsqu’il comprend cette vision produit.

Notre équipe de développement est au plus proche de vous et doit bien comprendre cette vision produit. Avec une équipe
entièrement dédiée au projet, nous savons par expérience que nous aurons des meilleurs résultats.

Cette vision produit doit toujours correspondre aux besoins utilisateurs. Notre équipe est sensibilisée aux problématiques
rencontrées par les utilisateurs de nos clients et sont capables de les aider à mieux cerner leurs besoins utilisateurs.

### S’adapter

Les difficultés à transcrire les besoins utilisateurs en outils informatiques ont poussées vers une approche plus
adaptative des méthodologies projet.

Les approches agiles se basent sur quelques principes fondamentaux :

* travailler en coopération pour faire émerger une solution qui correspond aux besoins utilisateur
* privilégier ce qui apporte le plus de valeur ajoutée aux utilisateurs
* montrer rapidement ce qu’on construit pour vérifier qu’on est sur la bonne voie
* savoir se remettre en cause : pour améliorer le produit et pour progresser en tant qu’équipe
* réaliser un produit de qualité pour pérenniser son utilisation

Sur la base de la méthodologie Scrum, nous avons réfléchi à une façon de travailler avec vous. C’est une base de travail
qui pourra être revue lors de notre réunion de cadrage et tout au long du projet.
L’agilité, c’est aussi s’adapter à la situation projet en partant de son expérience.

### L'équipe

* Chloé Lequette jouera le rôle de product owner et qui est disponible pour l'équipe
* Anne-Line et Chloé, graphistes Ceebios
* Richard Hanna, développeur senior Fairness, certifié écoconception web.
* Thomas Chatenet, développeur junior Fairness.

Pas de chef de projet, l'équipe s'auto-gère. Pas de relation client - prestataire, nous travaillons en équipe.

### Les différentes phases

* Co-création avec un interlocuteur unique porteur du projet (product owner) qui est disponible pour répondre
(aux heures de bureau) à nos questions tout au long du projet.
* Immersion métier : nous travaillons bien que lorsque nous avons compris votre métier et vos problématiques.
C'est la phase autrement appelée "Vis ma vie".
* Itérations graphiques : ateliers de travail pour la direction graphique du site. 3 rounds d'aller/retour pour 
finaliser le rendu général.
* Cadrage pour aligner l'équipe, sous forme d'ateliers pour écrire chaque User Story (la spécification fonctionnelle
la plus atomique et testable possible).
Réflexion sur les architectures technique, organisationnelle et de l'information : globalement le comment on fabrique le site.
* Itérations (de 2 semaines par exemple). Pendant chaque itération, au moins un atelier, pour définir le prochaine
itération, et défricher les itérations suivantes, et toujours avoir une vue macro sur l'objectif à atteindre.
* Rétrospective d'équipe afin de s'améliorer
* Mise en production finale
* TMA, maintenance évolutive et correction de bugs.

Les ateliers se font en physique à Paris ou en visioconférence.

Et, à chaque itération :

* Surveillance des impacts planning et budget
* Qualité logicielle pour pérenniser les développements
* Mise en préproduction et tests en continu
* Recueillir les feedbacks utilisateurs tout au long du projet pour valider que tout se passe bien
* Amélioration continue à l'aide de rétrospectives d'équipe
* Mise en production en continu (si possible)
* Capacité d'adapter le périmètre fonctionnel en cas de besoin

## 4. Écoconception

Les démarches d'écoconception et d'accessibilité sont en place dès le tout début du projet.

Nous formons une équipe : graphiste, porteur du projet, développeurs participent dès le départ à la co-création afin de
mettre en place ces bonnes pratiques.

La démarche d'écoconception est normée (ISO 14062) : c'est l'intégration des aspects environnementaux dans la conception
et le développement de produit. Cela passe par une définition d’une unité fonctionnelle à étudier
(« rechercher et consulter un article » par exemple). 

Cela passe par des optimisations comme le choix d'utiliser une police système, la taille et poids des images,
la génération de pages statiques, éviter les "embed" de carte ou de vidéo... différents leviers afin qu'une page web
reste une page web sobre, c'est à dire une information claire pour l'utilisateur, alléger de tout gras numérique.

Nous préconisons de concevoir avec une approche "mobile-first" afin d'aller à l'essentiel sur chaque écran. Informations
disponibles à la fois sur mobile et sur écran de bureau.

Pour mesurer l’empreinte technique de chaque page (poids, complexité, etc.) et l’empreinte environnementale associée
(gaz à effet de serre et eau), nous utiliserons ecoindex.fr. Nous pourrions ensemble déterminer un critère ou
des critères à respecter pour chacune des pages du site et un budget (Ko, empreinte environnemental) à respecter pour
les unités fonctionnelles étudiées (peut mobiliser une navigation sur plusieurs pages).

Le principe est d'être exemplaire. Nous vous soumettons aussi la proposition d'avoir une page sur votre site expliquant
la démarche d’écoconception et quels sont les impacts environnementaux et sociaux du numérique. Nous vous aiderons à
rédiger ce contenu.

Pour en savoir plus

* [Les arguments pour mettre en place une démarche de conception responsable / écoconception](https://www.greenit.fr/2019/10/08/6-arguments-pour-la-conception-responsable-des-services-numeriques/)
* [Présentation Retour d'expérience écoconception / conception responsable](https://speakerdeck.com/supertanuki/retour-dexperience-conception-numerique-responsable)
* [Présentaion Impact environnemental de mon site web](https://speakerdeck.com/supertanuki/impact-environnemental-de-mon-site-web)

## 5. Design UX / UI

Chloé Lequette a proposé de prendre le rôle de designer UX / UI pour ce projet.

Nous vous accompagnons à travers d'ateliers en amenant nos compétences en terme d'éco-conception. L'idée est de
travailler en amont du développement sur l'expérience utilisateur (UX) par exemple en établissant des maquettes des
principaux parcours utilisateurs. Ensuite, nous travaillerons le design graphique (UI), la page d'accueil et les autres
pages clé à déterminer ensemble.

## 6. Préconisations technologiques

### Approche orientée open source et artisanat du logiciel

Notre orientation Open Source nous pousse à utiliser un maximum de solutions issues de cette communauté.

Nous sommes des artisans du web et du logiciel. Nous adhérons à la pensée Software Craftwomanship et Craftmanship.

Notre approche s'inspire des conceptions en Domain Driven Design, en architecture hexagonale, ou en Tests Driven
Development.

### Gestion des contenus (CMS)

[Grav](https://getgrav.org/) est un CMS qui permet de concevoir des sites rapides basés sur la génération statique
de pages, c'est à dire sans utilisation de base de données à la consultation des pages.

Ce principe est le plus efficace car il sollicite le moins possible le serveur et pour l'utilisateur,
c'est une page qui se charge rapidement.

L'autre avantage de cette architecture, c'est la sécurité qui est renforcée du fait que le contenu est déjà généré et
qu'il y a donc moins de risque d'injection SQL par exemple.

Grav comporte des [fonctionnalités d'administration](https://learn.getgrav.org/16/admin-panel/page/editor) des contenus
du site accessibles même pour les personnes non techniques.

Grav gère nativement le [multi-langage](https://learn.getgrav.org/16/content/multi-language) des interfaces et des
contenus. Nous verrons ensemble quelle stratégie choisir pour la gestion des contenus : arborescence identique quelque
soit la langue ou bien chaque langue a sa propre arborescence de contenus.

Grav est un CMS open source écrit en PHP et qui s'appuie sur des composants [Symfony](https://symfony.com/)
et [Twig](https://twig.symfony.com/), ce qui est gage de qualité.
Le projet existe depuis plus de 6 ans et profite [d'une documentation de qualité et d'une importante communauté](https://getgrav.org/about).

## 7. Accessibilité

Il est important aujourd'hui que les produits numériques n'excluent pas. Nous recommandons une intégration HTML en
respectant les règles d'accessibilité des contenus Web (WCAG) 2.0.

## 8. Suivi d'audience

Pour suivre l'audience du site et en respect des données privées de l'utilisateur (RGPD), nous recommandons d'installer
une instance [Matomo](https://matomo.org/) et de l'auto-héberger.

Nous paramétrons Matomo afin que les données soient anonymisées permettant ainsi l'exemption de la demande de
consentement pour la pose d'un cookie (selon la CNIL).

## 9. Hébergement

Les hébergements "green" ne sont pas déterminants dans le fait d'écoconcevoir un service numérique.

Toutefois nous recommandons le suisse Infomaniak qui est un des seuls hébergeurs à avoir une
[charte écologique](https://www.infomaniak.com/fr/hebergeur-ecologique/charte-ecologique).

Le français OVH est une alternative intéressante en terme d'impact environnemental, et surtout vos données restent
en France.

Nous pouvons aussi envisager de travailler avec un hébergeur engagé chez les [Chatons](https://chatons.org/fr),
Collectif des Hébergeurs Alternatifs, Transparents, Ouverts, Neutres et Solidaires.

Les frais d'hébergement ou toute autre service sont à votre charge, vous permettant d'être indépendants de nous.

Par exemple, pour Infomaniak, le tarif est à partir de 5.75 €/mois et comprenant les prestations suivantes :

* 100 Go d'espace disque SSD
* Noms de domaines illimités
* Bases de données MySQL illimitées
* Sauvegarde dans un autre datacenter
* Certificats SSL EV et DV (Let's Encrypt, Sectigo) 

## 10. Tierce Maintenance Applicative (TMA)

Nous accompagnons nos clients dans le cadre de la Tierce Maintenance Applicative (TMA) par le biais d’un contrat sur
mesure. Il couvre les bugs et les petites évolutions.

Nous utilisons des outils de bug-tracking (Jira, Trello, sur Github ou autre) permettant de suivre les incidents et
leur résolution. L’outil permet de jalonner les évolutions, de programmer les maintenances et de prioriser les tickets.
Il permet de garder une trace des incidents et surtout de suivre leur résolution.
L’outil de bug-tracking sera fortement utile lors de la recette de l’application développée.

Pour monitorer la production et être alertés en cas de bugs, nous utilisons l'outil [Sentry](https://sentry.io/welcome/)
dont le coût est à partir de 26 $/mois.

## 11. Résilience

Nous cultivons votre indépendance vis à vis de nous. Vous pourrez facilement prendre un autre prestataire si besoin.
On fera tout pour que ça ne se produise pas à travers une coopération étroite.

Il est important par exemple d'avoir une documentation permettant pour un développeur de démarrer le projet rapidement.

Des tests unitaires et fonctionnels sont écrits pour garantir la fiabilité et la pérennité du code.
Le code métier est découplé du framework et du langage permettant une montée de version plus facilement
(Domain Driven Design, architecture hexagonale).

La documentation et les user stories (spécifications) sont stockées dans la base de code du projet afin de garantir
qu'elles sont toujours à portée de main.

## 12. Contribution aux communs

Nous préconisons - ce n'est pas une obligation - de mettre votre code en open-source : l'intérêt pour vous est la
transparence, le code ouvert à amélioration ou la correction de bugs, contribution aux communs et donc code réutilisable
pour d'autres usages.

Les contenus publiés peuvent être libérés sous licence Creative Commons.

## 13. Livrables

* Le code source hébergé sur un repository Github ou Gitlab qui vous appartient.
* Documentation démarrage du projet à destination des développeurs et intégrée au code ("Read me").
* Spécifications (ou users stories) conservées dans un dossier dédié dans la base de code (format à déterminer,
par exemple en markdown).
* Guide de bonnes pratiques co-écrit et à l'attention des animateurs du site.

## 14. Budget et planning

Comme présenté dans notre démarche, cette proposition n'est pas exhaustive et n'a pas lieu d'engagement.
Nous nous engageons sur des moyens avec un budget et un planning fixes, avec pour objectif un service sobre,
pertinent et convivial pour les utilisateurs.

Compte tenu des informations à notre connaissance, voici notre estimation pour ce projet.

* 3 ou 4 ateliers de co-conception
* Budget : 36 jours développeurs à 650 €HT / jour soit un total de 23400 €HT
* Planning : démarrage au plus tard en avril pour mise en ligne en septembre 2020

### Conditions de facturation

Nos conditions de règlement sont :

* 30% à la commande,
* 40% à mi-projet, après validation d'une version démo intermédiaire,
* 30% à la livraison,
* Paiement à 30 jours après établissement de la facture.

### Tierce Maintenance Applicative (TMA)

Provision de 6 jours développeur par an à 650 €HT / jour soit un total de 3900 €HT par année.
Cette provision pourra être ajustée à la baisse ou à la hausse selon vos besoins.

## Conclusion

Nous accompagnons nos clients dans la conception responsable de services numériques.
Nous privilégions les clients dont leurs activités sont responsables ou contributrices à l'intérêt général.

Autant que possible, nous ne souhaitons pas avoir une simple relation client / prestataire mais une relation de
partenaires, de transparence et de confiance.

Enfin, nous sommes enthousiastes à l’idée de travailler sur ce projet qui s'aligne avec nos valeurs et notre
sensibilité écologique.
