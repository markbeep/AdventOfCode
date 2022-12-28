import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Scanner;

public class day5 {

	public static void main(String[] args) throws FileNotFoundException {
		File file = new File("day5.txt");
		Scanner scanner = new Scanner(file);
		int max = 0;
		ArrayList<Integer> seats = new ArrayList<Integer>();
		while (scanner.hasNextLine()) {
			String rowString = scanner.nextLine();
			
			int rowStart = 0;
			int rowEnd = 127;
			int colStart = 0;
			int colEnd = 7;
			int diff = 0;
			int row = 0;
			int col = 0;
			for (int i = 0; i < 10; i++) {
				if (i < 7) {
					// search row
					if (rowString.charAt(i) == 'F') {
						diff = rowEnd - rowStart;
						rowEnd -= diff / 2 + 1;
					}
					else {
						diff = rowEnd - rowStart;
						rowStart += diff / 2 + 1;
					}
				}
				else {
					// search col
					if (rowString.charAt(i) == 'L') {
						diff = colEnd - colStart;
						colEnd -= diff / 2 + 1;
					}
					else {
						diff = colEnd - colStart;
						colStart += diff / 2 + 1;
					}
				}
			}
			int id = (rowStart) * 8 + (colStart);
			seats.add(id);
			if (id > max) {
				max = id;
			}
		}
		ArrayList<Integer> test = new ArrayList<Integer>();
		for (int i = 48; i <= 874; i++) {
			test.add(i);
		}
		System.out.println(max);
		System.out.println(test.size());
		System.out.println(seats.size());
		System.out.println();
		Collections.sort(seats);
		int c = 0;
		for (int i: test) {
			if (!seats.contains(i)) {
				System.out.println(i);
			}
		}

	}

}
