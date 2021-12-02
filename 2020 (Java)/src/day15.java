import java.io.File;
import java.io.FileNotFoundException;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class day15 {

	public static void main(String[] args) throws FileNotFoundException {
		Scanner s = new Scanner(new File("day15.txt"));
		String[] nums = s.nextLine().split(",");
		int[] n = new int[nums.length];
		Map<Integer, Integer> m = new HashMap<>();
		int prev = 0;
		for (int i = 0; i < nums.length; i++) {
			m.put(Integer.valueOf(nums[i]), i);
			prev = Integer.valueOf(nums[i]);
		}
		int preI = m.size()-1;
		int current;
		for (int i = m.size(); i < 30000000; i++) {
			if (m.containsKey(prev) && preI-1 == m.get(prev)) {
				current = 1;
			}
			else if (!m.containsKey(prev)) {
				current = 0;
			}
			else {
				current = (i) - (m.get(prev) + 1);
			}
			m.put(prev, preI);
			prev = current;
			preI = i;
		}
		System.out.println(prev);
		
	}

}
