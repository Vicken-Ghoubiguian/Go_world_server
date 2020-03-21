# Go_world_server

Go_world_server est une application web écrite en Go (aussi appelé Golang), langage créé par Google en 2009.

Cette application web renvoie l'heure et la date courante pour toutes les timezones passées par l'URL.

## Prérequis

Pour installer Go_world_server, les prérequis suivants doivent être satisfaits pour la machine:

* Linux dans sa version la plus récente,

* Go dans sa version la plus récente.

## Installer Go_world_server en version de développement ou de test sur une machine quelconque

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

## Installer Go_world_server en version de production dans un container Docker
