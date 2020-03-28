# Go_world_server

Go_world_server est une application web écrite en Go (aussi appelé Golang), langage créé par Google en 2009.

Cette application web renvoie l'heure et la date courante pour toutes les timezones passées par l'URL.

## Prérequis

Pour installer Go_world_server, les prérequis suivants doivent être satisfaits pour la machine:

* Linux dans sa version la plus récente,

* Go dans sa version la plus récente.

## Installer Go_world_server en développement ou de test sur une machine quelconque

Pour installer Go_world_server en version de développement ou de test, vérifiez bien que les prérequis listés ci-dessus sont remplis.

Ensuite, clonez le projet Go_world_server à l'aide de cette commande:

```bash
git clone https://gitlab.imerir.com/eric.ghoubiguian/go_world_server
```

Ensuite, placez-vous dans le répertoire créé par la commande précédente à l'aide ce celle-ci::

```bash
cd go_world_server
```

Exécutez après cette commande:

```bash
export GOPATH=${PWD}
```

__petite précision__: Cette commande permet de modifier la valeur de la variable d'environnement `GOPATH` qui spécifie la source du projet go.

Félicitations une fois cela fait, vous pouvez maintenant lancer le projet à l'aide de la commande suivante:

```bash
go run main.go
```
ça y est, le serveur est fonctionnel. Connectez-vous à l'aide du lien URL <a href="http://localhost:8080/">suivant</a>.

## Installer Go_world_server en production avec Docker

Il existe 2 procédés pour installer Go_world_server avec Docker: **à l'aide du fichier Dockerfile** et **à l'aide de l'image docker pullé**.

# Installer Go_world_server en production à l'aide du fichier Dockerfile

Un fichier nommé `Dockerfile` existe à la racine du projet.

Ce fichier permet à Docker de créer l'image de l'application web et de faire fonctionner ce dernier dans un container Docker.

Tout d'abord, placez-vous dans le projet à l'aide de la commande suivante:

```bash
cd go_world_server
```
Le succès de cette commande dépend du fait que le projet Go_world_server existe sur le serveur d'installation. Si ce n'est pas le cas, tout d'abord clonez-le à l'aide de la commande suivante:

```bash
git clone https://gitlab.imerir.com/eric.ghoubiguian/go_world_server
```
Ensuite, créez l'image Docker de Go_world_server à l'aide de la commande suivante:

```bash
docker image build -t go_world_server:latest .
```

__petite précision__: le caractère `.` précisé à la fin de la commande se réfère au répertoire courant.

Une fois cette commande terminée, l'image `go_world_server` est créée avec `go_world_server` comme nom et `latest` comme étiquette.

Pour vérifier, exécutez la commande suivante pour vérifier que l'image `go_world_server` a bien été créée avec succès:

```bash
docker image ls
```
Vous pouvez de plus consulter la liste de toutes les instructions qui ont été exécutées pour créer l'image Docker de Go_world_server:

```bash
docker image history go_world_server:latest
```
Il est maintenant temps de créer le container Docker et de le faire marcher, pour cela exécutez la commande suivante:

```bash
docker container run -d --name go_world_server -p 80:8080 go_world_server:latest
```
__petite précision__: 

Toutes mes félicitations, le container Docker de Go_world_server est créé et marche maintenant.
Pour s'en rendre compte, listez tous les containers Docker à l'aide de la commande suivante:

```bash
docker image ls
```

ou de celle-ci

```bash
docker image ls -a
```
Maintenant qu'il est installé, l'application web fonctionne maintenant.

Pour s'en rendre compte, il vous suffit d'ouvrir votre navigateur web favori puis d'entrer l'URL suivante:

```bash
<adresse_ip_du_container>:8080
```
Pour connaître l'adresse IP du container Docker de l'application web Go_world_server, faites la commande suivante:

```bash
docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' go_world_server
```
L'adresse IP du container s'affiche ensuite en dessous.

# Installer Go_world_server en production à l'aide de l'image docker pullé
