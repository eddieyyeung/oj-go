L(X[0:m-1], Y[0:n-1])

if X[m-1] == Y[n-1]
  L(X[0:m-1], Y[0:n-1]) = L(X[0:m-2], Y[0:n-2]) + 1
else if X[m-1] != Y[n-1]
  L(X[0:m-1], Y[0:n-1]) = math.Max(L(X[0:m-2], Y[0:n-1]), L(X[0:m-1], Y[0:n-2]))