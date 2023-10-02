import std/strutils

type
  Player = ref object
    deck: seq[int]
var
  p1 = Player()
  p2 = Player()

proc duplicateDeck(p: Player, first_n_card: int): Player =
  result = Player()
  result.deck = p.deck[0..<first_n_card]

proc drawTopCard(p: Player): int =
  result = p.deck[0]
  p.deck.delete(0)
  
proc `$`(p: Player): string =
  result = $p.deck

proc addCardAtBottom(p: Player, newCard: int) = 
  p.deck.add(newCard)

proc populateDecks() =
  var tmp: string
  while true:
    tmp = stdin.readLine()
    if tmp == "":
      break
    elif "Player 1:" in tmp:
      continue
    else:
      p1.addCardAtBottom(parseInt(tmp))

  while true:
    try:
      tmp = stdin.readLine()
    except IOError:
      break

    if tmp == "":
      break
    elif "Player 2:" in tmp:
      continue
    else:
      p2.addCardAtBottom(parseInt(tmp))


proc calculate_deck_score(s: seq[int]): int = 
  result = 0
  for i,element in s:
    result += (s.len() - i) * element


proc recursive_rounds(p1, p2: Player, depth:int): bool = # true means p1 won that round,false means p2 won
  echo "\n\n === Game ===, depht = ", depth, "\n\n"
  var round_inside_subgame = 0

  var permutation_logger: seq[(seq[int], seq[int])]

  while true:
    round_inside_subgame += 1
    echo  "\n-- Round ", round_inside_subgame, " (Game ", depth, ") -- "
    echo "Player 1's deck: ", p1
    echo "Player 2's deck: ", p2

    if (p1.deck, p2.deck) in permutation_logger:
      echo "p1 wins by deja vu this"
      return true
    else:
      permutation_logger.add((p1.deck, p2.deck))

    if p1.deck.len()==0:
      echo "p2 wins"
      return false        # means p2 wont this subgame
    elif p2.deck.len()==0:
      echo "p1 wins"
      return true         # means  p1 wont this subgame

    var p1_drawed_card = p1.drawTopCard()
    var p2_drawed_card = p2.drawTopCard()
    echo "Player 1 plays: ", p1_drawed_card
    echo "Player 2 plays: ", p2_drawed_card

    if (p1_drawed_card <= p1.deck.len()) and (p2_drawed_card <= p2.deck.len()):
      var new_p1 = p1.duplicateDeck(p1_drawed_card)
      var new_p2 = p2.duplicateDeck(p2_drawed_card)
      if recursive_rounds(new_p1, new_p2, depth + 1):
        echo "Player 1 wins the subgame"
        p1.addCardAtBottom(p1_drawed_card)
        p1.addCardAtBottom(p2_drawed_card)
      else:
        echo "Player 2 wins the round"
        p2.addCardAtBottom(p2_drawed_card)
        p2.addCardAtBottom(p1_drawed_card)

    else: # means chose without subgame
      if p1_drawed_card > p2_drawed_card:
        echo "Player 1 wins the round"
        p1.addCardAtBottom(p1_drawed_card)
        p1.addCardAtBottom(p2_drawed_card)
      else:
        echo "Player 2 wins the round"
        p2.addCardAtBottom(p2_drawed_card)
        p2.addCardAtBottom(p1_drawed_card)



proc main() = 
  populateDecks()
  echo p1
  echo p2

  if recursive_rounds(p1, p2, 1): 
    echo "\n\n\n\nfinished, P1 wont whole game"
    echo p1
    echo p2
    echo p1.deck.calculate_deck_score()
  else:
    echo "\n\n\n\nfinished, P2 wont whole game"
    echo p1
    echo p2
    echo p2.deck.calculate_deck_score()
main()
