# Random NPC generator

This is a random NPC generator for a campaign in D&amp;D 5e I runned. Some content is homebrew (e.g. Nations).

## Sample output

```
Nome: Muscled Madalynn
Eta: giovane
Sesso: F
Razza: Gnomo
Nazione: Beriel
Voce: afflitta
Carattere: focoso
Aspetto: 
	Capelli: Radi, Giallo pastello
	Occhi: Grigio ardesia scuro
	Altezza: 99
	Pelle: Blu oltremare
```

## How does it work

It simply picks randomly a gender, name, character, aspect, nationality, race and voice from a given list. So you can adapt this to your campaing by modifing the lists files.

## How to customize to your need

The list files can be found inside the folder `./generators/lists/`. The files name should be enough to identify which one does what.

## How to use it

First of all you need [a working installation of Go](https://golang.org/doc/install).

Now you can run in your terminal

```bash
$ go get github.com/CanobbioE/npcgen
```

You can then build, and run the generator with 

```bash
$ go build
$ ./npcgen
```

## Disclaimer
The text outputted is in Italian, to change it just change the lists files and the `String()` functions.
This is in an early stage and a fun/fast project, therefore there may never be additional features.
