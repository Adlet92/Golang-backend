# Groupie-tracker

Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.

It will be given an [API](https://groupietrackers.herokuapp.com/api), that consists in four parts:

    - The first one, artists, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

    - The second one, locations, consists in their last and/or upcoming concert locations.

    - The third one, dates, consists in their last and/or upcoming concert dates.

    - And the last one, relation, does the link between all the other parts, artists, dates and locations.

Given all this you should build a user friendly website where you can display the bands info through several data visualizations.

## Usage

```bash
go run ./cmd
```
Copy the link which has appeared
```bash
Server launched at http://127.0.0.1:8080
```
Push explore button to discover more information about bands

