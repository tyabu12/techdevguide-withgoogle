public class Problem {
  public int sumNumbers(String str) {
    int ans = 0, mem = 0;
    for (char c : str.toCharArray()) {
      if ('1' <= c && c <= '9') {
        mem = mem * 10 + c - '0';
      } else if (mem > 0) {
        ans += mem;
        mem = 0;
      }
    }
    return ans + mem;
  }
}