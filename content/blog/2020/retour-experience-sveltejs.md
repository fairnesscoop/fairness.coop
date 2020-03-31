---
title: 'SvelteJS, coup de coeur 2020'
publishDate: '2020-03-02'
description: Le numérique d'aujourd'hui suppose un monde stable et en croissance alors que le monde de demain est incertain.
readtime: 10
thumbnail: numerique-et-crises.jpg
author: mathieu-marchois
---

Dans cet article je voudrais vous faire un retour d'expérience sur une techno que j'ai découverte il y a peu : SvelteJS !
Et pour spoiler un peu, je dois dire que cette techno front est absolument incroyable ! (et c'est un dev back qui le dit ...)

## Contexte

Actuellement en mission chez [RadioFrance](https://www.radiofrance.fr/) depuis plusieurs mois, je suis en charge de refondre complètement l'intranet permettant, entre autre, de visualiser / gérer les différents flux reçus de la Maison de la radio.

La refonte s'est portée sur les mêmes choix technologiques que l'existant, à savoir un backend en [Node.js](https://nodejs.org) (<3) et un front en [React](https://fr.reactjs.org/). Cela m'allait parfaitement étant à l'aise sur les deux technos.

Vous allez peut être vous demander pourquoi faire une refonte en réutilisant les mêmes technos ?

C'est un outil qui est passé d'équipes en équipes et devenait de plus en plus difficile à maintenir. Le code n'avait aucune cohérence, était très compliqué à installer (même localement), le choix d'architecture était trop complexe, code incompréhensible etc.
Ma mission était (et est toujours) de reprendre cet outil de zéro pour tout remettre au propre.

## Qu'est ce que SvelteJS ?

Je ne vous ferai pas de présentation détaillée, d'autres l'ont déjà fait et ce ne serait pas forcément le plus pertinent. Je tenais juste à vous faire mon retour d'expérience et comment j'ai réussi à convaincre RadioFrance de partir sur une technologie toute récente.

Quelques ressources utile si vous voulez découvrez ce language:

-   [Svelte.js : introduction au “compilateur en guise de framework”](https://medium.com/@nilmanduil/svelte-js-le-compilateur-en-guise-de-framework-5473f1d727f8)
-   [Svelte.js : le compilateur qui s'attaque à React et Vue](http://www.meanjs.fr/svelte-la-librairie-le-compilateur-plutot-qui-sattaque-a-react-et-vue/)

## Prise en main

J'ai découvert SvelteJS lors d'une présentation durant de la DotJS 2019. Cette présentation m'a vraiment emballé et je voulais absolument tester cette énième technologie front. Encore une, me direz vous. Oui mais pas n'importe laquelle !

Ayant une bonne connaissance de React, la prise en main a été assez rapide. J'ai suivi le tutoriel sur le [site officiel](https://svelte.dev/tutorial/basics) et je dois dire que j'ai été vraiment agréablement surpris.
C'est la première fois que je trouve une documentation aussi intuitive / interactive. Cela m'a permis de monter rapidement en compétence.

Très entousiaste à l'idée de me faire la main dessus, j'ai pris un projet interne chez Fairness pour le migrer en SvelteJS (à la base le front était en React).
La migration a été assez simple et rapide à faire, le projet étant encore en cours de développement il m'a fallait une petite semaine pour reprendre tout ce qui avait était fait.

Si vous voulez voir à quoi ca ressemble (dans la vrai vie, et non une todo-list), le projet [Permacoop](https://github.com/fairness/permacoop) est open-source.

## Intégration chez RadioFrance

### Présentation

Après m'être fait la main sur le projet interne, je me suis dit que ce serait bien d'en parler à mon équipe chez RadioFrance.
Petite réunion technique au cours de laquelle j'ai pu faire une présentation et expliquer quels seraient les avantages d'une migration vers SvelteJS :

-   Simplicité de prise en main du language.
-   3 fois moins de code à écrire (et ce n'est pas une blague, voir plus bas pour les chiffres) et donc maintenance plus simple.
-   Légèreté du code : Svelte est un compilateur qui compilera en VanillaJS uniquement le code dont vous avez besoin. Pas besoin donc de charger l'ensemble du framework. Svelte porte bien son nom !
-   Rapidité : 30 à 40 fois plus performant que React, Vue ou Angular. La raison majeure ? Il n'y a pas de DOM virtuel.

Cette réunion, qui a convaincu l'équipe, a débouché sur la réalisation d'un POC d'une semaine de l'intranet sur lequel je travaillais depuis le début de ma mission (4 mois de developpement dessus).

### Réalisation du POC

Pour réaliser ce POC, en plus de [Svelte](https://svelte.dev/), je me suis appuyé sur [Sapper](https://sapper.svelte.dev/) qui permet de builder des applications très rapidement, faire du SSR etc.

L'utilisation combinée de ses deux technos est d'une facilité déconcertante. Un exemple simple : **le routing**.

Qui n'a jamais perdu du temps à mettre en place tout un système de routing sur un projet React ou Vue ? Installation de librairies, configuration, mapping d'un composant à sa route etc.

Avec **Sapper** rien de plus simple. Toutes les vues accessibles se trouvent dans le dossier `routes`. Sapper se base sur le nom d'un fichier pour construire son URL.

Ex: `./routes/blog/posts.svelte` sera accessible via `blog/posts`.

Par dessus tout ça, il est également possible de rajouter des règles d'accès, grâce aux regex notamment, directement sur le nom des fichiers.

Ex: `./routes/blog/posts/[id([0-9]+)].svelte` sera accessible via `blog/posts/12`. Si la regex ne match pas notre URL, une 404 sera alors retournée.

A la fin de la semaine j'avais terminé la partie visualisation des données. La rapidité de la migration, la simplicité du code a convaincu le client de faire une bascule complète de l'intranet.

En plus de la réalisation de ce POC, il faut dire que le fait que le créateur soit [Rich Harris](https://github.com/Rich-Harris) développeur au [New York Times](https://www.nytimes.com/), a un peu facilité les choses, les corps de métiers étant assez similaire.

### Migration complète

## Conclusion
