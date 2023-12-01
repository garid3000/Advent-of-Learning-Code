#import tables
import std/strutils
import std/sequtils
import sugar

type
  Food = object # id: int
    ingredients: seq[string]
    allergens: seq[string]
var
  all_foods: seq[Food]
  ing2aller: seq[(string, string)]
  no_allr_ingr: seq[string]

proc `*`(s1s: seq[string], s2s: seq[string]): seq[string] = 
  result = s1s.filter(
    (s: string) -> bool => s in s2s
  )

proc `*`(f1: Food, f2: Food): Food = 
  result.ingredients = f1.ingredients * f2.ingredients
  result.allergens   = f1.allergens   * f2.allergens

#proc remove_from_all_food(ingr: string, allr:string) =
#  for i,f in all_foods:
#    all_foods[i].ingredients = all_foods[i].ingredients.filter(
#      (s: string) -> bool => (s != ingr)
#    )
#    all_foods[i].allergens = all_foods[i].allergens.filter(
#      (s: string) -> bool => (s != allr)
#    )

proc remove_ingr_from_all_food(ingr: string) =
  for i,f in all_foods:
    all_foods[i].ingredients = all_foods[i].ingredients.filter(
      (s: string) -> bool => (s != ingr)
    )

proc remove_allr_from_all_food(allr: string) =
  for i,f in all_foods:
    all_foods[i].allergens = all_foods[i].allergens.filter(
      (s: string) -> bool => (s != allr)
    )


proc populate_data() = 
  var instr: string
  while true:
    try:
      instr = stdin.readLine()
      echo instr
    except IOError:
      echo "STDIN finished"
      break

    all_foods.add(
      Food( ingredients: instr.split(" (contains ")[0].split(" "), 
      allergens: instr.split(" (contains ")[1].replace(")", "").split(", "))
    )
  echo all_foods

  while true:
    var changedFlag = true
    block oneReduction:
      for f1 in all_foods:
        for i2, f2 in all_foods:
          if f1 != f2:
            var f1f2 = f1 * f2
            echo f1, "\n", f2, "\n\t", f1f2
            if ( (f1f2.ingredients.len() == 1) and (f1f2.allergens.len() == 1) ):
              ing2aller.add( (f1f2.ingredients[0], f1f2.allergens[0]) )
              remove_ingr_from_all_food(f1f2.ingredients[0])
              remove_allr_from_all_food(f1f2.allergens[0])
              #echo "\t\tshiiiiiiiiiit\n\n"
              changedFlag = false
              break oneReduction
            elif ( (f2.ingredients.len() == 1) and (f2.allergens.len() == 1) ):
              ing2aller.add( (f2.ingredients[0], f2.allergens[0]) )
              remove_ingr_from_all_food(f2.ingredients[0])
              remove_allr_from_all_food(f2.allergens[0])
              #echo "\t\tshiiiiiiiiiit\n\n"
              changedFlag = false
              break oneReduction
            #elif (f2.allergens.len() == 0) and (f2.ingredients.len() != 0):
            #  no_allr_ingr.add(f2.ingredients[0])
            #  #all_foods.delete(i2)
            #  remove_ingr_from_all_food(f2.ingredients[0])
            #  changedFlag = false
            #  break oneReduction
            #elif (f2.allergens.len() == 0) and (f2.ingredients.len() == 0):
            #  all_foods.delete(i2)
            #  changedFlag = false
            #  break oneReduction
            elif (f2.allergens.len() == 0) and (f2.ingredients.len() == 0):
              all_foods.delete(i2)
              changedFlag = false
              break oneReduction



    if changedFlag:
      break

  echo "\n\n\n"
  echo ing2aller
  echo "\n\n\n"
  echo no_allr_ingr #all_foods
  echo "\n\n\n"
  echo all_foods




populate_data()
