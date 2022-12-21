import sympy
from sympy import solve

strexpr = input()
# //print(strexpr)
humn = sympy.Symbol("humn")
expr = eval(strexpr)
print(expr)
print("humn =", solve(expr, humn))
