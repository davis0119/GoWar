# War

War is a card game which involves 2 players. It is typically a long simple card game where players draw a card (each starting off with half the deck) and beat the other player's card if their's outranks the other. If card ranks are the same, several cards are drawn and the duel continues by facing one card against another. This recursively continues until there is a winner. Winners of each round take the loser's played cards. The winner of the game is the last player standing with cards on their side. 

This program is a CLI app. Everything is playable on the command line. 

# Tweaks

Players are allowed to forfeit wars (this may be exercised if a player is running low on cards and would rather play more safe. When number of cards are low, the player usually has a sense of where their "good" cards are in their deck and would avoid risking losing such cards.)

Players are given 3 opportunities to deploy more troops -- in the event of a War event, a player may exercise a chance to add more troops to the war to increase the stakes of winning/losing.

# How to Play

1. go run war.go 
   - Prompt 'y' to enter the game. 'help' will always be a keyword to provide more information. 
   - Each round, you will be asked to 'Wage battle'. This represents a round in War (both players draw a card).
   - In the event where both players draw the same card, you will be given an option to continue with the War Event or to forfeit the war. Forfeiting may be handy if you remember you have very valuable cards near the top of your deck and do not wish to risk them. There will also be an option to deploy more troops (which would force your opponent to deploy more troops) to up the stakes. This is a limited action per game. 
2. You may exit the program by declining ('n') to continue the battle and not re-declaring War. 

# License 
MIT License

# Acknowledgements

Hanbang Wang for incredibly helpful resources! Everything detailed was also very appreciated for navigation through this project. 
