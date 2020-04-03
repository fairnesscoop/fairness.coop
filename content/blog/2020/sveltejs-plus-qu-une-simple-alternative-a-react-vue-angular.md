---
title: "SvelteJS, plus qu'une simple alternative à React, Vue ou Angular"
publishDate: '2020-04-02'
description: "Dans cet article je vous fais un retour d'expérience sur l'utilisation de SvelteJS, une migration d'application anciennement React. Nous verrons quels sont ses avantages."
readtime: 9
thumbnail: svelte.png
author: mathieu-marchois
gCO2e: 1.35
EcoindexLink: http://www.ecoindex.fr/resultats/?id=70654
---

Dans cet article je voudrais vous faire un retour d'expérience sur une techno que j'ai découverte il y a peu : SvelteJS !
Et pour spoiler un peu, je dois dire que cette techno front est absolument incroyable ! Et c'est un dev back qui le dit ...

## Contexte

Actuellement en mission chez [Radio France](https://www.radiofrance.fr/) depuis plusieurs mois, je suis en charge de refondre complètement l'intranet permettant, entre autre, de visualiser et de gérer les différents flux de la Maison de la Radio.

La refonte s'est portée sur les mêmes choix technologiques que l'existant, à savoir un backend en [Node.js](https://nodejs.org) (<3) et un front en [React](https://fr.reactjs.org/). Cela m'allait parfaitement étant à l'aise sur les deux technos.

Vous allez peut-être vous demander pourquoi faire une refonte en réutilisant les mêmes technos ?

C'est un outil qui est passé d'équipes en équipes et devenait de plus en plus difficile à maintenir. Le code n'avait aucune cohérence, était très compliqué à installer (même localement), le choix d'architecture était trop complexe, code incompréhensible etc.
Ma mission était (et est toujours) de reprendre cet outil de zéro pour tout remettre au propre.

## Qu'est ce que SvelteJS ?

Je ne vous ferai pas de présentation détaillée, d'autres l'ont déjà fait et ce ne serait pas forcément le plus pertinent. Je tenais juste à vous faire mon retour d'expérience et comment j'ai réussi à convaincre Radio France de partir sur une technologie toute récente.

Quelques ressources utiles si vous voulez découvrir ce language:

- [Svelte.js : introduction au “compilateur en guise de framework”](https://medium.com/@nilmanduil/svelte-js-le-compilateur-en-guise-de-framework-5473f1d727f8)
- [Svelte.js : le compilateur qui s'attaque à React et Vue](http://www.meanjs.fr/svelte-la-librairie-le-compilateur-plutot-qui-sattaque-a-react-et-vue/)

## Prise en main

J'ai découvert SvelteJS lors d'une conférence donnée à la [dotJS 2019](https://www.dotjs.io/schedule/dotjs-2019). Cette présentation m'a vraiment emballé et je voulais absolument tester cette énième technologie front. Encore une, me direz vous. Oui mais pas n'importe laquelle !

Ayant une bonne connaissance de React, la prise en main a été assez rapide. J'ai suivi le tutoriel sur le [site officiel de Svelte](https://svelte.dev/tutorial/basics) et je dois dire que j'ai été vraiment agréablement surpris.
C'est la première fois que je trouve une documentation aussi intuitive et interactive. Cela m'a permis de monter rapidement en compétence.

Très enthousiaste à l'idée de me faire la main dessus, j'ai pris [un projet interne](https://github.com/fairnesscoop/permacoop) de notre coopérative codé initialement en React pour le migrer en SvelteJS.
La migration a été assez simple et rapide à faire, le projet étant encore en cours de développement. Il m'a fallu une petite semaine pour reprendre tout ce qui avait été fait.

## Intégration chez Radio France

### Présentation

Après m'être fait la main sur [le projet interne](https://github.com/fairnesscoop/permacoop), je me suis dit que ce serait bien d'en parler à mon équipe chez Radio France.
Petite réunion technique au cours de laquelle j'ai pu faire une présentation et expliquer quels seraient les avantages d'une migration vers SvelteJS :

- Simplicité de prise en main du language.
- 3 fois moins de code à écrire (et ce n'est pas une blague, voir plus bas pour les chiffres) et donc maintenance plus simple.
- Légèreté du code : Svelte est un compilateur qui génère en _vanilla Javascript_ uniquement le code dont vous avez besoin. Inutile donc de charger l'ensemble du framework. Svelte porte bien son nom !
- Rapidité : 30 à 40 fois plus performant que React, Vue ou Angular. La raison majeure ? Il n'y a pas de DOM virtuel.

Cette réunion, qui a convaincu l'équipe, a débouché sur la réalisation d'une preuve de concept d'une semaine de l'intranet sur lequel je travaillais depuis le début de ma mission (4 mois de développement dessus).

### Preuve de concept

Pour réaliser cette preuve de concept, en plus de [Svelte](https://svelte.dev/), je me suis appuyé sur [Sapper](https://sapper.svelte.dev/) qui permet de builder des applications très rapidement, faire du SSR (Server Side Rendering), etc.

L'utilisation combinée de ces deux technos est d'une facilité déconcertante. Un exemple simple : **le routing**.

Qui n'a jamais perdu du temps à mettre en place tout un système de routing sur un projet React ou Vue ? Installation de librairies, configuration, mapping d'un composant à sa route etc.

Avec **Sapper** rien de plus simple. Toutes les vues accessibles se trouvent dans le dossier `routes`. Sapper se base sur le nom d'un fichier pour construire son URL.

Ex: `./routes/blog/posts.svelte` sera accessible via `blog/posts`.

Par dessus tout ça, il est également possible de rajouter des règles d'accès, grâce aux regex notamment, directement sur le nom des fichiers.

Ex: `./routes/blog/posts/[id([0-9]+)].svelte` sera accessible via `blog/posts/12`. Si la regex ne matche pas notre URL, une 404 sera alors retournée.

A la fin de la semaine j'avais terminé la partie visualisation des données. La rapidité de la migration, la simplicité du code a convaincu le client de faire une bascule complète de l'intranet.

En plus de la réalisation de cette preuve de concept, il faut dire que le fait que le créateur soit [Rich Harris](https://github.com/Rich-Harris) développeur au [New York Times](https://www.nytimes.com/), a un peu facilité les choses, les corps de métiers étant assez similaires.

### Migration

Quelques chiffres forcément beaucoup plus parlants que de longs discours. À fonctionnalités équivalentes:

Nombre de lignes de codes:

- React: ~29000
- SvelteJS: ~8000

Temps passé:

- React : ~2 mois
- SvelteJS: 2 semaines (aussi dû au fait que je connaissais déjà le métier)

Qu'est ce qui peut justifier une aussi grande différence ? Il y a plusieurs raisons, mais celle que je retiens principalement est le fait de se passer de [Redux](https://redux.js.org/). SvelteJS permet d'avoir une gestion d'état simplifiée et nécéssite donc moins de code, moins de configuration etc.

Bien que jeune, cette technologie est assez mature pour pouvoir être utilisée en production. Nous n'avons rencontré aucun problème lors du déploiement de l'intranet.

### Exemple de code

Pour des raisons de confidentialité, je ne peux vous partager du code provenant de l'application du client. Si vous voulez un exemple complet et concret vous pouvez regarder et vous inspirer du projet [Permacoop](https://github.com/fairnesscoop/permacoop), un ERP open-source (encore en cours de développement) éco-conçu pour les coopératives.

Le [serveur](https://github.com/fairnesscoop/permacoop/tree/master/server) a été développé avec le framework [NestJS](https://nestjs.com/) et le [client](https://github.com/fairnesscoop/permacoop/tree/master/client) avec SvelteJS. Toute contribution est bien évidement la bienvenue :)

## La suite ?

Cette migration s'est avérée être une véritable réussite et est toujours en cours de développement pour être iso-fonctionnelle à l'ancienne plateforme.

L'ensemble de l'équipe est convaincu par cette techno et prend même du plaisir à travailler dessus. Certains devs back, qui ne voulaient pas entendre parler de front, s'y mettent également.

Ce projet va servir de pilote pour potentiellement migrer toutes les applications _frontend_ de Radio France (France Inter, France Culture, France Musique, France Bleu, Fip, Mouv), aujourd'hui en React vers du SvelteJS.
Un grand changement à venir ? En tout cas, cette technologie fait l'unanimité grâce à sa performance, son poids léger et sa facilité de prise en main.

Je vous recommande vraiment de suivre avec attention l'évolution de ce projet. Ce n'est pas pour rien qu'il fait partie des [technos les plus plébiscitées de 2019](https://2019.stateofjs.com).
