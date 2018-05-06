def string_splosion(str):
  ans = str
  for i in range(1, len(str)):
    ans = str[0:len(str)-i] + ans
  return ans
