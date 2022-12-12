---
title: "Forum PHP 2022 : Papa et Maman nettoient l'Internet"
date: "2022-12-07"
description: |
  En tant que PO ou développeuse ou développeur backend, comment mettre en oeuvre les bonnes pratiques d'écoconception de services numérique ? Hélène et Mathieu ont donné une conférence au ForumPHP 2022 sur ce thème.
readtime: 8
thumbnail: fairness.png
author: Les coopérateurs Fairness
authorId: florimond-manca
---

Le 13 octobre 2022, Hélène Maître-Marchois et Mathieu Marchois, co-fondateurs de la coopérative, ont donné une conférence au titre qui a plus en interroger plus d'un : "Papa et Maman nettoient l'Internet".

[Visionner la rediffusion sur YouTube](https://www.youtube.com/watch?v=Pi8VcEAgEMY) (40 min)

Dans ce talk, Hélène et Mathieu ont proposé au public venu nombreux un condensé didactique des raisons et des bonnes pratiques pour _nettoyer Internet_.

## Rappel des constats

Il n'y a plus à tergiverser : il _faut_ réduire l'impact environnemental du numérique.

En 2020, le numérique représentait environ 3 à 4% des émissions de gaz à effet de serre dans le monde. Cet impact augmente de 8% par an en moyenne. (GreenIT, 2020 ; The Shift Project, 2019)

Internet, c'est physique : au-delà des datacenters, qui se comptent en millions, et des réseaux, les **terminaux utilisateurs** se comptent en dizaines de milliards et concentrent la majorité des impacts. Or les terminaux ont un très faible taux de recyclage (~5%) et nécessitent des métaux et matériaux critiques. Ainsi de l'indium, nécessaire pour la fabrication des écrans LCD. Hélène renvoie aux travaux d'Aurore Stéphant et du collectif SystExt. Nos métiers sont donc intimement liés aux **ressources minières**.

## Faire durer, éliminer, questionner, concevoir au plus juste

La piste principale ? **Faire durer les terminaux** le plus longtemps possible. Aujourd'hui en Europe, un smartphone est changé tous les 2 ans en moyenne. Au-delà du marketing, les développeurs ont aussi une responsabilité : ce renouvellement est aussi poussé par l'obsolescence logicielle, induite par des besoins de calcul toujours croissants. Il faut donc **agir dès la conception**.

Première étape : **éliminer le gras fonctionnel** qui alourdit les applications avec une très faible utilité. On estime ainsi que 70% à 80% des pages Web ne sont pas ou très peu ouvertes. En tant que PO, Hélène demande souvent "_pourquoi ?_", et obtient parfois des résultats spectaculaires. Un client a ainsi pu économiser du temps, de l'argent, et réduire les impacts côté utilisateur en remplaçant une fonctionnalité complexe de demande d'information complémentaire par une ligne téléphonique, car ces demandes étaient rares (quelques demandes à l'année).

Deuxième étape : **questionner l'existant** pour définir ce qui doit être gardé. Que faire des trackers et des centaines de requêtes déclenchés à chaque chargement de page ? On estime que le nombre de requêtes a été multiplié par 4 en 5 ans, et que le poids moyen d'une page Web est passé de 500 ko à plus de 2 Mo. En comparaison, la mission Apollo a envoyé des humains sur la Lune avec quelques kilo-octets.

La bonne nouvelle, c'est que ce nettoyage se révèle très satisfaisant : on obtient très vite de gros gains. C'est le cas de Bing a ainsi pu réduire sa charge serveur de 80% en réduisant de seulement 20% le nombre d'éléments affichés par page de recherche. <!-- Source nécessaire -->

Quand vient le besoin de développer de **nouvelles fonctionnalités**, il est impératif de faire travailler ensemble **PO, designers UX, et équipes de développement**. Cela permet de savoir de quoi l'utilisateur a _vraiment_ besoin. Le produit n'en sera que meilleur.

## Côté technique...

Une fois le fonctionnel traité, viennent les améliorations techniques. Le nerf de la guerre sera **les chiffres**, autrement dit la **mesure**. Hélas, les données d'impact environnemental manquent. Mais des outils permettent de faire un état des lieux :

* Outils généralistes : PageSpeed, web.dev, Blackfire...
* Spécialisés éco-conception : [EcoIndex](https://www.ecoindex.fr), Greenframe.io, Scaphandre, Carbonalyzer, Greenspector...

Attention à la CI : les runners consomment de la RAM et du CPU. Optimiser les tests, mettre les dépendances en cache, d'autant plus si la CI est lancée très régulièrement.

Il s'agit également de questionner notre **environnement de développement** : nombre d'écrans externes, matériel professionnel (téléphone, ordinateurs portables). Pour Hélène, développer sur du matériel qui ne soit pas dernier permet aussi de concevoir au plus proche de la situation des utilisateurs.

Rappel : la performance ne recoupe pas seulement la rapidité, mais également la disponibilité et un usage sobre des ressources (ISO 20510).

Au quotidien, comment alors rendre son code plus "green" ?

## Extrait de 10 bonnes pratiques

* Favoriser le sur-mesre à l'usage d'un CMS ;
* Choisir les formats de données adaptés : type de données dans la base de données, formats d'échange... Un mauvais type va gaspiller du stockage et de la mémoire ;
* Eviter les requêtes SQL dans les boucles ;
* Optimiser les requêtes aux bases de données : ouvrir le capot de l'ORM, ne récupérer que les colonnes absolument nécessaires, mettre en place de la pagination... ;
* Mettre les caches en RAM : c'est plus efficace qu'un cache sur disque dur ;
* Utiliser un cache HTTP (_reverse proxy_) : Nginx, Varnish... Cela réduit drastiquement le besoin de transfert de données et la charge serveur ;
* Réduire au nécessaire les logs serveur : régler les niveaux au plus juste, logger que ce qui est utile ;
* Favoriser le _Request Collapsing_ : regrouper les appels aux services distants en un seul ;
* Mettre en place une politique d'expiration et de suppression des données ;
* Utiliser la version la plus récente du langage : améliorations de fonctionnalités, de performance, de sécurité, très visible entre PHP 7 et PHP 5 par exemple, ou PHP 8 avec le compilateur JIT ;

Basé sur le [Guide des 115 bonnes pratiques du collectif Conception Numérique responsable (CNumR)](https://collectif.greenit.fr/outils.html), mis à jour en mai 2022.

## En conclusion

L'écoconception permet, dès sa prise en main, de rapidement obtenir de gros gains sur les impacts environnementaux de nos applicatifs. C'est un levier concret pour avoir un impact sur la façon dont le numérique pèse sur l'environnement et les utilisateurs.

À bientôt pour de nouvelles interventions !
