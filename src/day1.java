import java.io.FileNotFoundException;
import java.util.ArrayList;

public class day1 {

	public static void main(String[] args) throws FileNotFoundException {
		ArrayList<Integer> numbers = quickScanner.allInts("day1.txt");
		for (int i = 0; i < numbers.size(); i++) {
			for (int j = 0; j < numbers.size(); j++) {
				for (int k = 0; k < numbers.size(); k++) {
					if (numbers.get(i) + numbers.get(j) + numbers.get(k) == 2020) System.out.println(numbers.get(i) * numbers.get(j) * numbers.get(k));
				}
			}
		}

	}

}
