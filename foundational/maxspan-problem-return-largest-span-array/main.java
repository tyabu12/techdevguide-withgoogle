public class Problem {
  public int maxSpan(int[] nums) {
    if (nums.length == 0) return 0;
    int ans = 1;
    for (int i=0;ans<nums.length-i;i++) {
      for (int j=nums.length-1;ans<j-i+1;j--) {
        if (nums[j] == nums[i]) {
          ans = j - i + 1;
          break;
        }
      }
    }
    return  ans;
  }
}