import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;

public class day12 {

	public static void main(String[] args) throws FileNotFoundException {
		Scanner s = new Scanner(new File("day12.txt"));
		ArrayList<String> cmd = new ArrayList<>();
		int sY = 0;
		int sX = 0;
		int[] w = {10, 1};
		char dir = 'E';  // subtask 1 (current facing direction)
		while (s.hasNextLine()) {
			String line = s.nextLine();
			cmd.add(line);
			char curDir = line.charAt(0);
			if (curDir == 'F') {
				for (int i = 0; i < Integer.valueOf(line.substring(1)); i++) {
					sX += w[0];
					sY += w[1];
				}		
			}
			switch (curDir) {
			case 'N':
				w[1] += Integer.valueOf(line.substring(1));
				break;
			case 'E':
				w[0] += Integer.valueOf(line.substring(1));
				break;
			case 'S':
				w[1] -= Integer.valueOf(line.substring(1));
				break;
			case 'W':
				w[0] -= Integer.valueOf(line.substring(1));
				break;
			case 'L':
				w = rotate(w[0], w[1], -Integer.valueOf(line.substring(1)));
				break;
			case 'R':
				w = rotate(w[0], w[1], Integer.valueOf(line.substring(1)));
				break;
			}
		}
		System.out.println("TOTAL: " + (Math.abs(sX) + Math.abs(sY)));
		
	}
	public static int[] rotate(int x, int y, int deg) {  // subtask 2
		int[] c = {x, y};
		int temp;
		int a = deg / 90;
		int d = (a < 0) ? -1: 1;
		for (int i = 0; i < Math.abs(a); i++) {
			temp=c[0]*-d;
			c[0]=c[1]*d;
			c[1]=temp;
		}
		return new int[] {c[0], c[1]};
	}
	public static char turn(char start, int degrees) {  // subtask 1
		char[] c = {'N', 'E', 'S', 'W'};
		int amt = degrees / 90;
		int ind = 0;
		for (int d = 0; d < 4; d++) if (start == c[d]) ind = d;
		while (amt < 0) amt += 4;
		return c[(ind + amt) % 4];
	}

}
