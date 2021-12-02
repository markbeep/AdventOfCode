import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;

public class day8 {

	public static void main(String[] args) throws FileNotFoundException {
		ArrayList<String> lines = quickScanner.allLines("day8.txt");
		boolean[] visited = new boolean[lines.size()];
		boolean[] wrong = new boolean[lines.size()];
		int amtC = 0;
		
		int i = 0;
		int acc = 0;
		int changed = 0;
		int c = 0;
		while (true) {
			if (i > lines.size()) break;
			i = i % lines.size();
			if (i == 0 && visited[i]) break;
			if (visited[i]) {
				wrong[changed] = true;
				i = 0;
				acc = 0;
				amtC = 0;
				visited = new boolean[lines.size()];
			}
			visited[i] = true;
			Scanner scanner = new Scanner(lines.get(i));
			String keyword = scanner.next();
			System.out.println("Index: " + i);
			System.out.println(c);
			int amt = Integer.valueOf(scanner.next());
			if (keyword.equals("nop")) {
				if (amt == 0) {
					wrong[i] = true;
					i++;
				}
				else if (amtC == 0 && !wrong[i]) {
					c++;
					changed = i;
					i += amt;
					amtC++;
				}
				else {
					i++;
				}
			}
			else if (keyword.equals("jmp")) {
				if (amtC == 0 && !wrong[i]) {
					c++;
					changed = i;
					i++;
					amtC++;
				}
				else i += amt;
			}
			else {
				acc += amt;
				i++;
			}
		}
		System.out.println(acc);

	}

}
