import std/strutils

type
  Player = ref object
    deck: seq[int]
var
  p1 = Player()
  p2 = Player()

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

proc main() = 
  populateDecks()
  echo p1
  echo p2

  var round = 0
  while true:
    round += 1
    if (p1.deck.len() == 0) or (p2.deck.len()==0):
      break

    echo "-- Round ", round, "--"
    echo "Player 1's deck: ", p1
    echo "Player 2's deck: ", p2
    var p1_drawed_card = p1.drawTopCard()
    var p2_drawed_card = p2.drawTopCard()
    echo "Player 1 plays: ", p1_drawed_card
    echo "Player 2 plays: ", p2_drawed_card
    if p1_drawed_card > p2_drawed_card:
      echo "Player 1 wins the round"
      p1.addCardAtBottom(p1_drawed_card)
      p1.addCardAtBottom(p2_drawed_card)
    else:
      echo "Player 2 wins the round"
      p2.addCardAtBottom(p2_drawed_card)
      p2.addCardAtBottom(p1_drawed_card)
    echo ""
  
  echo "finsih"
  echo p1
  echo p2

  echo p1.deck.calculate_deck_score()
  echo p2.deck.calculate_deck_score()
main()
