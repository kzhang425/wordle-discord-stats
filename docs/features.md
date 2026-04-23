
## not finishing

starting and not finishing is a score of 7

## scoring features

all of these should map from player to value
and support query of either top k, or bottom k, or one player's value

### total elo

long term elo
treats single day as several 1v1 games
adjust elo, and move to next day

example:

alex: 3
bob: 4
charles: 5
alex v bob: alex W
bob v charles: bob W
alex v charles: alex W

sample starting value: 1500
sample k: 32

### min 20 simple

just take the average score of each player, over all time
then discard the players that have played under 20 games

### sliding week score

over the most recent 7 days
take the average score of each player

## misc features

### current streaks

count current streak of player
used to display in daemon report header

### all time streak

count best streak of player, and when
used as fun fact in daemon report header

### scores <= x

count number of times player got <= x
used as fun fact in daemon report header