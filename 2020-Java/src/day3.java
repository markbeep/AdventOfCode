import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class day3 {

	public static void main(String[] args) throws FileNotFoundException {
		File file = new File("day3.txt");
		int[] down = {1, 1, 1, 1, 2};
		int[] right = {1, 3, 5, 7, 1};
		int total = 1;
		for (int i = 0; i < down.length; i++) {
			Scanner scanner = new Scanner(file);
			int place = 0;
			int trees = 0;
			String line = scanner.nextLine();
			int height = 0;
			while (true) {
				//System.out.println(height);
				if (place >= line.length()) place = place - line.length();
				if (line.charAt(place) == '#') trees++;
				place += right[i];
				if (!scanner.hasNextLine()) break;
				for (int j = 0; j < down[i]; j++) {
					height++;
					line = scanner.nextLine();
				}
			}
			System.out.println(trees);
			total *= trees;
		}
		System.out.println(total);
	}
}
