public class Problem {
  public String withoutString(String base, String remove) {
    String s = base;
    remove = remove.toLowerCase();
    for (;;) {
      int idx = s.toLowerCase().indexOf(remove);
      if (idx == -1) break;
      s = s.substring(0, idx) + s.substring(idx+remove.length());
    }
    return s;
  }
}