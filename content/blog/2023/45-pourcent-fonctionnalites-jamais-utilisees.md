---
title: |
    "45% des fonctionnalités jamais utilisées" : qu'en est-il vraiment ?
date: "2023-03-31"
description: |
    C'est un chiffre que nous utilisons parfois en conférence et dans nos communications. Mais d'où provient-il ? Après un peu d'archéologie, il s'avère ne pas être si solide. Ce qui ne remet pas en cause l'enjeu qu'il sert souvent à introduire : éliminer les fonctionnalités inutiles.
readtime: 7
thumbnail: loupe.jpg
author: Les coopérateurs Fairness
authorId: florimond-manca
---

Dans le cadre de nos actions de sensibilisation à la conception numérique responsable, il est nécessaire de s'appuyer sur des faits vérifiables.

Tandis que je préparais une intervention, j'ai buté sur la vérification d'un chiffre que nous utilisons de temps en temps en conférence. 70% des fonctionnalités seraient "non-essentielles", et 45% d'entre elles ne seraient tout bonnement jamais utilisées. D'où cela provient-il ? 

## Menons l'enquête

Nous tirons à la base cette affirmation de la [bonne pratique BP_001](https://github.com/cnumr/best-practices/blob/main/chapters/BP_001_fr.md) de la [checklist des 115 bonnes pratiques](https://collectif.greenit.fr/ecoconception-web/115-bonnes-pratiques-eco-conception_web.html) éditée par le collectif Conception Numérique Responsable (CNumR). La description de cette bonne pratique débute ainsi :

> Plusieurs études (Cast Software et Standish Group, notamment) démontrent que 70 % des fonctionnalités demandées par les utilisateurs ne sont pas essentielles et que 45 % ne sont jamais utilisées. [...]

Nous pouvons donc aller voir ce que "Cast Software" et "Standish Group" ont à dire sur le sujet.

Du côté de Cast Software, je n'ai pas trouvé de référence traitant d'un sujet similaire.

J'ai en revanche trouvé un fil à remonter du côté de Standish Group.

### Les études de Standish Group

Standish Group a publié diverses études dans le domaine de la qualité des projets logiciels.

L'étude plus connue semble être le "CHAOS Report". Depuis la [première édition en 1994](https://www.standishgroup.com/sample_research_files/chaos_report_1994.pdf), ce rapport s'intéresse à la réussite ou non des projets logiciels (_software projects_). On apprend ainsi que sur les cas étudiés, "31.1% des projets sont annulés avoir même d'avoir été terminés". Il n'y est cependant pas question de l'utilisation effective des fonctionnalités.

Au global, l'objet des CHAOS Reports semble avoir été (et de toujours être) de montrer l'avantage des méthodes agiles notamment par rapport aux méthodes en cascade (_Waterfall_) dans le domaine du logiciel. Cela a probablement occassionné des biais. Notons par exemple les critiques méthodologiques formulées dans le papier [The Rise and Fall of the Chaos Report Figures](https://dl.acm.org/doi/abs/10.1109/MS.2009.154), publié dans la revue IEEE en 2010.

En tout état de cause cependant, le CHAOS Report n'est pas la source que l'on cherche.

C'est pourtant bien chez Standish Group que la source primaire se trouve, comme la suite de l'enquête nous l'apprendra.

### Reprises et critiques du chiffre

En tapant "Standish Group features never used" dans un moteur de recherche, je suis tombé sur deux billets de blog.

Le premier billet s'intitule [Why 45% of all software features in production are NEVER used](https://www.linkedin.com/pulse/why-45-all-software-features-production-never-used-david-rice/) et fut publié par David Rice en 2016. Il comporte un graphique en camembert appuyant le titre du billet. La somme des portions "Jamais utilisées" et "Rarement utilisées" est de 64%. Le chiffre de 70% serait donc déjà une approximation. Le graphique cite la source suivante : "Standish Group, Study of 2000 projects at 1000 companies". Dommage, on ne va pas pouvoir retrouver un rapport ou une étude précise avec cela.

Dans un commentaire, un internaute cite cet autre billet : [Are 64% of Features Really Rarely or Never Used?](https://www.mountaingoatsoftware.com/blog/are-64-of-features-really-rarely-or-never-used). Dans ce billet critique datant de 2015, l'auteur Mike Cohn signale que la métrique selon laquelle "64% des fonctionnalités sont rarement voire jamais utilisées" proviendrait d'une présentation introductive (_keynote_) donnée par Jim Johnson, à l'époque président de Standish Group, lors de la conférence XP 2002. Selon l'auteur, qui indique avoir contacté Standish Group directement, cette affirmation se basait sur l'étude de "4 applications internes". Il ne s'agirait donc pas d'une affirmation basée sur un grand jeu de données, ni même sur des applications commerciales. Le billet résume les données du Standish Group remises dans leur contexte par ce graphique :

![Graphique en camembert intitulé "Feature Use in Four Internal-Use Products. Valeurs : never = 45%, rarely = 19%, sometimes = 16%, often = 13%, always = 7%. Source : Jim Johnson, Chairman of the Standish Group, Keynote "ROI, It's Your Job", Third International Conference on Extreme Programming, Alghero, Italy, May, 26-29, 2002.](https://www.mountaingoatsoftware.com/uploads/blog/Chart001.jpg)

La conférence XP 2002 est également citée dans un papier d'un chercheur de l'Université du Québec à Montréal publié en mai 2009 et intitulé [Agile development: Issues and avenues requiring a substantial enhancement of the business perspective in large projects](https://www.researchgate.net/publication/221186023_Agile_development_Issues_and_avenues_requiring_a_substantial_enhancement_of_the_business_perspective_in_large_projects). En page 62, l'auteur y fait référence pour illustrer le "_software bloat_", défini comme "un phénomène en vertu duquel les applications incluent des fonctionnalités à la valeur métier douteuse". La figure, reproduite ci-dessous, est sous-titrée : "Features use presented by Standish Group at the XP 2002 conference".

![Graphique en camembert éclaté et en perspective, intitulé "Actual Use of Requested Features". Valeurs : identiques au graphique précédent. Pas de source intutilée sur la figure.](https://www.researchgate.net/profile/Ghislain-Levesque/publication/221186023/figure/fig2/AS:643929196204042@1530536092127/Features-use-presented-by-Standish-Group-at-the-XP-2002-conference.png)

On progresse. Il ne nous reste plus qu'à vérifier : est-ce bien ce qui a été dit lors de la conférence XP 2002 ? Sur quelles sources s'appuyait alors le président de Standish Group pour prononcer cette affirmation ?

### Retour en 2002

En tapant "XP 2002 Jim Johnson" dans un moteur de recherche, je tombe dans les premiers résultats sur le billet de blog [The XP 2002 Conference](https://www.martinfowler.com/articles/xp2002.html), rédigé en juillet 2002 par Martin Fowler, célèbre pour ses billets de blog sur l'architecture et la qualité logicielle.

Selon Martin Fowler, Jim Johnson avait à l'époque cité deux sources :

* "Une étude de DuPont" (s'agit-il bien de l'entreprise de chimie américaine du même nom ?) indiquant que "seulement 25% des fonctionnalités d'un système sont réellement nécessaires" ;
* "Une étude du Standish Group" où l'on retrouve le chiffre de 45% de fonctionnalités jamais utilisées et 20% souvent ou toujours utilisées, ce qui concorde avec les données rapportées plus haut.

Point crucial, Martin Fowler a mis dans son billet le lien vers les slides originales de Jim Johnson. Un coup de [Wayback Machine](https://web.archive.org), plus tard, il s'avère qu'elles sont toujours disponibles [ici](https://web.archive.org/web/20060814204423/http://ciclamino.dibe.unige.it/xp2002/talksinfo/johnson.pdf).

Après inspection, la slide qui nous intéresse est "Features & Function Usage", en page 5. Il n'y a pas d'autre précision que le camembert que nous avons déjà vu plus haut. Aucune étude, rapport ou élément de méthodologie n'est cité. Zut.

Avant de perdre espoir, lançons-nous dans de dernières recherches web. Je suis tombé sur ce [rapport client](https://www.standishgroup.com/sample_research_files/Modernization.pdf) (PDF, 1.6 Mo) de Standish Group et daté de 2010. En page 15, Standish Group cite une "étude de 1996" (toujours sans référence précise, je comprends donc qu'elle n'a pas été publiée ?), réalisée par leurs soins. Elle s'appuierait sur "_100 custom development applications_" (que peut bien vouloir dire ce terme ?) et aboutit aux chiffres que l'on connaît désormais si bien. L'étude a consisté à faire l'inventaire des fonctionnalités et à comparer avec des déclarations issues d'ateliers avec des utilisateurs. Alors... 4 applications internes d'après le contact de Mike Cohn ou 100 applications d'après ce document ?

J'ai bien l'impression que l'enquête pourrait encore continuer. Mais il faut bien ressortir la tête de l'eau au bout d'un moment. Alors je vais me risquer à une conclusion.

## Retour vers le présent

À ce stade, il semble qu'il faille constater que l'idée qu'une quasi-majorité des fonctionnalités dans les applications informatiques soient peu ou pas utilisées repose sur une littérature plutôt maigre. D'autant plus que les données datent du tournant des années 2000, alors que depuis les méthodes dites agiles pourraient très bien avoir complètement renversé la situation, comme

N'y aurait-il pas des travaux plus fiables et plus à jour permettant de le vérifier ?

Par une recherche rapide, j'ai trouvé un [rapport](https://go.pendo.io/rs/185-LQW-370/images/2019%20Feature%20Adoption%20Report%20Digital.pdf) publié en 2019 par la société Pendo, qui édite une plateforme d'analytics pour la recherche utilisateur. La société conclut, à partir de données anonymisées d'utilisation de sa plateforme que, 56% des fonctionnalités seraient "rarement utilisées", et 24% "jamais utilisées", soit un total de 80% de fonctionnalités pas ou peu utilisées. Ces proportions empiriques rejoignent les chiffres de Standish Group.

Mais là aussi, la critique s'impose.

Si des fonctionnalités sont peu utilisées, on peut supposer qu'elles sont peu _utiles_. Il faudrait donc les reconsidérer, comme l'y invite une démarche écoconception.

À l'inverse, la société Pendo semble supposer que ces fonctionnalités ne nécessitent pas d'être reconsidérées, et que ce sont les utilisateurs qui ne les découvrent pas assez, ou ne comprennent pas suffisamment leur intérêt. Le résumé du rapport indique ainsi qu'il est possible d'augmenter les ratios d'utilisation par diverses stratégies pour "pousser à adopter les fonctionnalités".

Il faut probablement imputer cet angle de vue au positionnement commercial de cette société, qui a un intérêt à permettre à ses clients de justifiés des prix élevés pour une quantité de fonctionnalités large et effectivement utilisées, y compris par des techniques d'adoption quasi-forcée.

De la même façon, [ce billet](https://chartio.com/learn/product-analytics/what-is-feature-usage/), publié sur le blog d'une solution d'analytics ciblant visiblement des startups, ne cite que des actions visant à accroître le taux d'utilisation fonctionnelle une fois celui-ci mesuré. L'exemple de la pop-up suggérant d'écouter de la musique chaque fois que l'utilisateur active une application de marche à pied et se déplace n'est certainement pas sans rappeler certaines évolutions dans les OS mobiles majoritaires.

Cela participe d'un accroissement continu de l'emprise et de la place du numérique dans notre vie quotidienne. Au contraire, le Numérique responsable vise aussi à apaiser notre relation au numérique. Tout en étant lui aussi [source potentielle d'économies financières](/blog/2019/arguments-pour-la-conception-responsable-des-services-numeriques/).

## Qu'en conclure ?

Le plus important dans tout cela n'est peut-être pas l'exactitude du chiffre avancé. Remarquons que dans la bonne pratique BP001, le chiffre avancé servait surtout à illustrer l'enjeu suivant :

> En réduisant la couverture et la profondeur fonctionnelle de l’application, on abaisse son coût de développement initial, sa dette technique et les impacts environnementaux associés.

Autrement dit, du point de vue de l'éconception de services numériques, la limitation du périmètre fonctionnel qui permet de répondre efficacement aux besoins utilisateurs reste une priorité.

Le fait qu'une fraction des fonctionnalités ne soit pas utilisée peut servir de moteur. Mais cette fraction est difficile à estimer. Il se peut qu'elle doit déjà modérée dans les organisations ayant déjà adopté des approches qui tendent à minimiser les développements de fonctionnalités inutiles (développement agile, _lean startup_, recherche utilisateur, etc).

Pour nos propres présentations, il conviendrait néanmoins de prendre plus de précautions en citant ce chiffre, voire d'illustrer l'enjeu du questionnement des fonctionnalités différemment.

Nous pourrions par exemple plutôt énoncer ceci : "Quelques études empiriques et retours d'expérience suggèrent qu'une partie des fonctionnalités d'une application ne sont probablement peu ou pas utilisées. Leur remise en question et le cas échéant leur élimination est une source de réduction de coûts de développement et de maintenance, de dette technique, mais aussi d'impacts environnementaux".
