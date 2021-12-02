import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Collections;

public class day9 {

	public static void main(String[] args) throws FileNotFoundException {
		ArrayList<Integer> numbers = quickScanner.allInts("day9.txt");
		boolean g;
		/*
		for (int k = 0; k < numbers.size(); k++) {
			g = false;
			if (k < 25) g = true;
			System.out.println("We at: " + numbers.get(k));
			for (int i = k - 25; i < k; i++) {
				for (int j = k - 25; j < k; j++) {
					if (i < 0) i = 0;
					if (j < 0) j = 0;
					System.out.println(i + " + " + j);
					System.out.println(numbers.get(j) + numbers.get(i));
					if (numbers.get(j) + numbers.get(i) == numbers.get(k)) g = true;
				}
			}
			System.out.println(g);
			System.out.println(numbers.get(k));
			if (!g) break;	
		}
		*/
		// 69316178
		int l = 11691646;
		ArrayList<Integer> dings = new ArrayList<Integer>();
		for (int k = 0; k < numbers.size(); k++) {
			if (numbers.get(k) == l) {
				System.out.println("FOUND");
				for (int i = 0; i < k; i++) {
					if (i < 0) continue;
					dings.add(numbers.get(i));
				}
			}
		}
		int sum = 0;
		boolean found = false;
		int min = Integer.MAX_VALUE;
		System.out.println(dings.toString());
		for (int d = 0; d < dings.size(); d++) {
			for (int i = 2; i < 26; i++) {
				sum = 0;
				ArrayList<Integer> correct = new ArrayList<Integer>();
				for (int j = d; j < i + d; j++) {
					if (j >= dings.size() || j < 0) continue;
					System.out.print(j + " ");
					correct.add(dings.get(j));
					sum += dings.get(j);
					if (sum == l) {
						System.out.println(correct.toString());
						System.out.println(Collections.min(correct)+Collections.max(correct));
						found = true;
						break;
					}
				}

				System.out.println(sum);
				if (sum <= l) System.out.println(sum);
				//if (sum > l) break;
				//System.out.println(correct.toString());
				if (found) break;
			}
			if (found) break;
		}
		
	}

}
