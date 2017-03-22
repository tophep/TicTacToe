# TicTacToe


## Try It Out

TicTacToe requires [Golang](https://golang.org/dl/) to run.

```sh
$ git clone https://github.com/tophep/TicTacToe.git
$ cd TicTacToe
$ go run *.go
```


## The Problem

I recently had an interview with the following coding challenge:

1. Write a command-line program that lets two people play Tic-Tac-Toe.

2. Add replay functionality. Allow the user to enter the ID of a game and have that game played back to them.

3. Add an AI to play against. You’re not allowed to embed any knowledge of the game into your AI beyond making legal moves and recognizing that it has won or lost. Your program is expected to “learn” from the games it plays, until it masters the game and can play flawlessly.



## My Solution

Ok, Part 1 tests a lot of practical programming proficiency. Build a basic system by implementing the game logic, print the board nicely, parse some user input and update the game state. 

Part 2 -- is the system flexible enough to add features quickly -- to do this I had to add an array to each Game's State which represented the sequence of moves that were made. Then, these sequences just needed to be added to a global array representing the replays. To play a replay, iterate over the sequence and print the game board at each step. 

The final part was much trickier. After thinking for a bit, I decided the most important thing was for the AI to simply avoid losing. If both players play perfectly, the game is always a tie. So they're not really winning, they're just not losing.

For my AI to "learn" it needed to analyze previous games. This analysis took the form of matching the current game state to those in the Replays array. If the current game has X moves played so far, find all previous games where the first X moves were the same. Then, from this list of previous games, identify the games where the AI did not win nor tie and look ahead one move. Remove these "losing" moves from the list of valid moves, then choose one randomly. Importantly, if the AI is left with no moves to choose from, just pick any valid one at random -- the AI could have previously lost after trying each spot at this stage, but still hasn't tried all branches further down the tree of possible moves. 

Initially the AI will lose, a lot, but eventually it will win or tie a game. It will slowly accumulate a history of these non-losing paths.

I played the AI twice. Of course I won the first time, but the second time the AI reacted differently to my very first move. I had demonstrated proof of concept and this is where the interview ended.



## Taking Things To The Next Level

I couldn't help but wonder if the AI would actually work. I knew at least superficially it would avoid moves where it lost previously, but would it become perfect given enough time? Eventually I got too curious and began tinkering. Since there are 9! or 362,880 possible sequences in Tic-Tac-Toe, I couldn't exactly train the thing manually.

First things first, my algorithm for searching through the replays was pretty slow. Iterating through the replays array is O(n) where n is the number of previously played games. This would make the AI slower and slower, as the corpus of previous games grew. 

To speed things up I wrote a custom tree structure where each node has nine children, or as I like to call it, a Nine-Ary Tree. Each node stores a move and each index in the child node array represents a potential next move. Now after each game, I add the sequence to the tree -- start at the root and index into the children using the spot of the first move, place a new node there, then do the same with the new node and the next move, and so on...

EDIT: A much more succinct way to describe the structure is as a prefix tree (trie), but instead of words, sequences of moves are stored

Now, when looking up a sequence mid-game the AI just walks the tree, which ends up being O(1) due to a depth limit of 9. Additionally, when inserting into the tree, I cache the end result of the game at every node along the way. This way the AI can immediately know all potential results of a move, without searching further down every branch.

Ok, time to train the AI. I added an option to the command-line for this, which triggers a loop of the AI playing itself for as many games as the user specifies. Sadly, even after millions of games of training, I was able to beat the AI consistently.

I tweaked the AI to prefer moves that had no child moves, i.e. a win or tie. If no winning move is available, it prefers moves it has never made before, to flesh out the move tree. Finally, it defaults back to making the same "non-losing" moves as a last resort.

This tweak improved the AI's performance noticeably but it was still far from perfect. I realized the problem was the AI's training partner. When playing itself it would win in situtations where it shouldn't because it's opponent was choosing random moves, instead of optimal moves. It was learning bad habits from itself.



## The Minimax Algorithm

So to train my AI correctly I needed to build it a perfect training partner. With some minor Googling I came across the [Minimax Algorithm](http://neverstopbuilding.com/minimax). The algorithm comes from Game Theory and the basic idea is to find the move that is optimal for you, assuming your opponent will make the optimal counter move, and then you will make an optimal counter-counter move and so on...

Translated to code, this takes the form of a depth-first search of the tree of all game states. When the search reaches a move that ends the game, a score is returned: 1 if the current player won, 0 if it's a tie and -1 if the opponent won. Then when returning up the stack, if the move being analyzed is being made by the current player, the max score of all possible moves is returned. If it is instead the opponents turn, the minimum score is returned. It's not the easiest thing to describe this in words so [here is a good visual representation.](https://www.youtube.com/watch?v=zDskcx8FStA)

Now my AI has a responsible training partner and if it's trained with 10,000+ games it plays really well. Of course it will occasionally run into situations it hasn't encountered before (it prints this to the console to inform the user) but they tend to be further down the move tree so it will learn pretty quickly.

Oddly enough, after completing the Minimax implementation I realized it also satisfied the problem statement: "You’re not allowed to embed any knowledge of the game into your AI beyond making legal moves and recognizing that it has won or lost."

Minimax has no knowledge of the game's rules or logic and only needs to know which moves are legal and if they end in a win, loss or tie. So the last bit of the problem about "learning" is unecessary, and so is my whole AI. Oh well!

