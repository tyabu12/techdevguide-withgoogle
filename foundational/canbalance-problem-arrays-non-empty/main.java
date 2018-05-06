public class Problem {
  public boolean canBalance(int[] nums) {
    int lsum = 0, rsum = 0;
    
    for (int n : nums) rsum += n;

    for (int i=0;i<nums.length-1;i++) {
      lsum += nums[i];
      rsum -= nums[i];
      if (lsum == rsum) return true;
    }
    
    return false;
  }
}