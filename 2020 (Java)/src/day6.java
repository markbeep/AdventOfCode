import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;

public class day6 {

	public static void main(String[] args) throws FileNotFoundException {
		File file = new File("day6.txt");
		Scanner scanner = new Scanner(file);
		int sum = 0;
		int l = 0;
		boolean[] visited = new boolean[26];
		while (scanner.hasNextLine()) {
			String line = scanner.nextLine();
			if (line.length() == 0) {
				for (int i = 0; i < visited.length; i++) {
					if (visited[i]) sum++;
				}
				visited = new boolean[26];
				l = 0;
				continue;
			}
			
			for (int i = 0; i < visited.length; i++) {
				if (l == 0 && line.contains(Character.toString(i+97)) || (visited[i] && line.contains(Character.toString(i+97)))) {
					visited[i] = true;
				}
				else {
					visited[i] = false;
				}
			}

			l++;
			
		}
		System.out.println(sum);

	}

}
