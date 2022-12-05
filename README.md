# War

War is a card game which involves 2 players. It is typically a long simple card game where players draw a card (each starting off with half the deck) and beat the other player's card if their's outranks the other. If card ranks are the same, several cards are drawn and the duel continues by facing one card against another. This recursively continues until there is a winner. Winners of each round take the loser's played cards. The winner of the game is the last player standing with cards on their side. 

# Tweaks
Additional rules are added for increased game complexity. Some may be substituted.

Players are allowed to forfeit wars (this may be exercised if a player is running low on cards and would rather play more safe. When number of cards are low, the player usually has a sense of where their "good" cards are in their deck and would avoid risking losing such cards.)

Players are given 3 opportunities to "re-strike" -- in the event of a War event, a player may exercise one "opportunity" to select one of the 3 face down cards in a War to attempt to beat the other player's currently played card. 

The player will play against a Bot with different difficulty settings, different behaviors (some randomness / informed decisions), and automated dialogue.

Other rules may be added if further complexity is desired. 

# Example Code
```
game := war.NewGame() 
game.InitPlayer(user) 
game.InitBot(difficulty) 
game.SplitCards() 

for !game.Over() {
    game.CommenceRound() // if card ranks are the same 
    /* within CommenceRound(), we check for ties and grant the option to forfeit
        for both players like below 
        '''
        if isWar {
            playerForfeit := grantForfeitOption() 
        }
        '''
    */
    game.PrintBattleResults() 
}
game.AnnounceWinner() 
```
# License 
MIT License
