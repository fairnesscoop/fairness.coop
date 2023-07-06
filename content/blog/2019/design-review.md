---
title: "La design review : révolution du travail en équipe"
date: "2019-03-07"
description: "Arrivé dans ma nouvelle équipe depuis plus d'un an maintenant, je voudrais faire un point sur les bonnes pratiques que j'ai découvertes avec mes collègues et l'impact que ça a eu sur ma façon de travailler. Je vais commencer par une pratique un peu particulière qui a révolutionné ma façon de travailler en équipe : la revue de design ou 'design review'."
readtime: 5
thumbnail: design-review.png
author: Les coopérateurs Fairness
authorId: mathieu-marchois
gCO2e: 1.3
EcoindexLink: http://www.ecoindex.fr/resultats/?id=63632
aliases:
    - /blog/2019/design-review.html
---

L'idée de base est de faire une revue en équipe de l'architecture d'une User Story à prendre sur le sprint en cours. C'est le cas à chaque fois qu'un nouveau ticket va être pris avant le démarrage du développement. On fait en sorte que toute l'équipe participe à l'architecture et au nommage (classes, variables, namespaces...) avant qu'un ticket ne sorte de la colonne "to do" pour aller vers la "doing". Le but est que celui qui code soit bien au clair sur ce qu'il a à faire et que ceux qui reliront le code ensuite puissent avoir une idée précise de ce qu'ils doivent vérifier.

On y passe le temps qu'il faut, entre 5 et 20 minutes, avec tous les devs de l'équipe, un stylo et un tableau blanc.

![Design review](/blog/images/design-review.png)

Concrètement, celui qui va prendre le ticket conduit la revue de design. On commence par la lecture de la User Story qui a été rédigée pendant les ateliers avec le Product Owner. On discute ensuite en équipe de la façon dont on veut construire le code. La personne en charge de l'US dessine l'architecture sur le tableau sous forme de schéma ou de dessin d'interface avec les différentes interactions. Elle note aussi le ou les nommages choisis. Parfois, on regarde également le code existant pour voir comment notre User Story s'intègre à l'existant.

Le fait d'être plusieurs cerveaux à réfléchir à la fonctionnalité qu'il va falloir développer nous permet d'affiner nos compréhensions respectives, de les confronter et de lever les ambiguïtés s'il en reste. L'objectif étant de dresser une liste exhaustive et ordonnée de tâches techniques à faire pour considérer cette story comme terminée. Ce temps ensemble permet de simplifier le travail du développeur qui aura les grandes lignes déjà tracées.

Quand nous intégrons de nouvelles personnes dans l'équipe, la revue de design devient un temps de **partage de connaissances** sur les projets. Tout ce qui est expliqué en détails est du coup bien mieux intégré par l'ensemble de l'équipe. Nous profitons également de ce temps pour faire remonter les questions qui existeraient sur la façon de mener le développement. Il ne doit plus y avoir de zones d'ombre à la fin de la revue.

A l'issue de ces échanges, qui permettent donc d'optimiser au maximum le travail du développeur qui prend la story, il récapitule ce qu'il a l'intention de faire en s'appuyant sur ce qu'il y a d'écrit sur le tableau. Si tout le monde est d'accord avec sa vision de la tâche à accomplir, il peut partir sur le développement, seul ou en Pair Programming.

Nous gardons à l'esprit qu'il sera parfois nécessaire de revenir sur certains points de blocage qui seraient décelés en cours de développement. Si c'est le cas: mini design review concernant la question soulevée pour aider le développeur à avancer et à prendre la décision adéquate.

Nous prenons toujours ce temps pour nous poser ensemble avant de foncer tête baissée sur une US que l'on a parfois mal comprise. Cela peut nous faire gagner énormément de temps, notamment pour la code review : plus de débat sur le nommage ou sur l'architecture en aval du développement car cela a été fait en amont !

Ce matin, par exemple, lors d'une design review, certains voulaient ajouter une nouvelle route, un nouveau controller et un nouveau service dans l'application. D'autres devs de l'équipe leur ont démontré qu'ils n'étaient pas nécessaires car un service existait déjà pour cela. On estime un gain d'une demie-journée de boulot minimum pour 5 minutes de revue de design !

Je suis maintenant persuadé que passer un peu de temps au départ peut nous en faire gagner à l'arrivée tant en vitesse de développement qu'en **qualité de code** et en progression d'équipe sur le projet.

C'est **une vraie révolution** dans la perception de mon travail et de ma place au sein d'une équipe de développement. Par ces courts moments pris à réfléchir autour de notre code, nous créons la possibilité de développer **un vrai code d'équipe** et non plus une accumulation de nos tickets respectifs.

[Voir la vidéo de ScrumLife sur la design Review](https://www.youtube.com/watch?v=2Bre5j3SNjU) avec notre coopérateur Matthieu.
