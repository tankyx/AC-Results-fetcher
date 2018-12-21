# Comment lancer un server Assetto

* Ouvrir le dossier `ASSETTO SERVER` sur le bureau
* Ouvrir le dossier `acServer`
* Ouvrir le dossier `server`

## Lancer la récupération des résultats

* Lancer `AC-Results-fetcher.exe`
* Ouvrir Chrome
* Se connecter à `192.168.1.9:8080/#/results`

La page se met à jour automatiquement toutes les minutes.

## Créer la session

* Lancer `acServerManager.exe`
* cliquer sur `Paddock 4` sur la gauche

### Settings

Sur cette page, vous pouvez modifier les settings liés au serveur. Il vaut mieux ne toucher à rien vu que tout est correctement configuré.

Les seuls settings qui peuvent être changés sont ceux liés au assistances, `assists`.

* `Factory` signifie que les aides seront activées si la voiture le permet
* `Fixed` signifie que les aides seront toujours activées.
* `Auto clutch` est à cocher tout le temps.
* `Tyre Blankets` est à cocher tout le temps.

### Sessions / Weather

Ici, on choisit le type de session.

__Toujours laisser le booking coché, avec 1 en booking time__

Ensuite on choisit `Qualify` ou `Race`.

* `Qualify` -> On change `Qualify time` en fonction du temps de session.
* `Race` -> On change le `Race Minutes` en decochant `Extra`.

__Pas besoin de toucher au reste.__

### Tracks

Ici, on choisit le circuit __ET__ le nombre de personne dans la session (entre __2__ et __4__).

* `Clients Allowed` -> le nombre de personnes dans la session.

Ensuite on choisit dans la liste des circuits.

### Cars / Entry list

Sur la gauche, on a la liste de tous les modèles de voitures. Au centre, on a les peintures disponibles. A droite, on a la liste des voitures qui seront utilisés.

* Cocher les modèles de voitures qu'on souhaite utiliser.
* Au centre, on a toutes les voitures et peintures disponibles. On selectionne la ou les voitures souhaitées avec un __simple clic gauche__. La voiture devrait apparaître sur la droite.
* S'il y a une voiture en trop sur la droite, __double clic__ sur la voiture en trop pour l'enlever.
* Une fois la liste de voiture terminée, on clique sur la barre verticale `OPEN CAR / DRIVER VIEW`.
* Au centre on a la liste de voiture avec des champs à remplir. A droite on a la liste des simulateurs.
* Glisser-deposer les simulateurs sur les voitures dans la liste.
* Changer le champ `Name` par le nom du client. __NE RIEN CHANGER D'AUTRE__.

__NE PAS CLIQUER SUR CLEAR POUR ENLEVER LA VOITURE DE LA LISTE. JUSTE DOUBLE CLIC SUR LE NOM DE LA VOITURE.__
__NE PAS CLIQUER SUR STORE, CA VA ECRASER LA CONFIG DES SIMUS.__

### Advanced options / Manual

Pas besoin de toucher au `Advanced options`. Le `manual` recèle toutes les infos sur cette appli.

Ensuite :

* Cliquer sur `Save` sur la barre du bas, tout à gauche.

__NE PAS CLIQUER SUR LE RESTE__
__SURTOUT PAS SUR DELETE__

* Cliquer sur `START SERVER` en bas à droite.

### J'ai crée un nouveau preset parce que j'ai pas fait attention

Pas de panique. On selectionne `Paddock 4` et tout revient à la normale
