import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;

public class day11 {

	public static void main(String[] args) throws FileNotFoundException {
		ArrayList<String> a = quickScanner.allLines("day11.txt");
		char[][] life = new char[a.size()][a.get(0).length()];
		for (int i = 0 ; i < a.size(); i ++) {
			for (int c = 0; c < a.get(i).length(); c++) {
				life[i][c] = a.get(i).charAt(c);
			}
		}
		printArray(life);
		char[][] newLife = copyArray(life);
		while (true) {
			for (int i = 0 ; i < a.size(); i ++) {
				for (int c = 0; c < a.get(i).length(); c++) {
					if (life[i][c] != '.') {
						int occ = D(life, i, c);
						//System.out.println("X: " + c + " | Y: " + i + " | occ: " + occ);
						if (occ >= 5 && life[i][c] == '#') newLife[i][c] = 'L';
						else if (occ == 0 && life[i][c] == 'L') newLife[i][c] = '#';
					}
				}
			}
			printArray(newLife);
			if (checkSame(life, newLife)) break;
			life = copyArray(newLife);
		}
		
		System.out.println("SUM: " + count(life));

	}
	
	public static int D(char[][] life, int i, int c) {
		int occ = 0;
		boolean[] ch = new boolean[8]; // 0 R, 1 DR, 2 D, 3 DL, 4 L, 5 UL, 6 U, 7 UR
		int[] n = new int[] {i, c, c, i, c, i};
		int[][] ass = new int[][] {{0, 2},{3, 2},{3, 1},{3, 4},{0, 4},{5, 4},{5, 1},{5, 2}};
		while (!ifBreak(ch)) {
			if (++n[2] > life[0].length-1) {
				ch[0] = true;
				ch[1] = true;
				ch[7] = true;
			}
			if (++n[3] > life.length-1) {
				for (int j = 1; j <= 3; j++) ch[j] = true;
			}
			if (--n[4] < 0) {
				for (int j = 3; j <= 5; j++) ch[j] = true;
			}
			if (--n[5] < 0) {
				for (int j = 5; j <= 7; j++) ch[j] = true;
			}
			for (int o = 0; o < 8; o++) {
				if (!ch[o]) {
					int y = n[ass[o][0]];
					int x = n[ass[o][1]];
					if (life[y][x] == '#') {
						occ++;
						ch[o] = true;
					}
					else if (life[y][x] == 'L') ch[o] = true;
				}
			}
		}
		return occ;
	}
	
	public static int countDiagonal(char[][] life, int i, int c) {
		int occ = 0;
		int right = c+1;
		int left = c-1;
		int up = i-1;
		int down = i+1;
		boolean[] checked = new boolean[8]; // 0 R, 1 DR, 2 D, 3 DL, 4 L, 5 UL, 6 U, 7 UR
		int[] nums = new int[] {i, c, c+1, i+1, c-1, i-1};
		int[][] ass = new int[][] {
			{0, 2}, // right
			{3, 2}, // down right
			{3, 1}, // down
			{3, 4}, // down left
			{0, 4}, // left
			{5, 4}, // up left
			{5, 1}, // up
			{5, 2}, // up right
		};
		/*
		 * y axis = 0
		 * x axis = 1
		 * right = 2
		 * down = 3
		 * left = 4
		 * up = 5
		 */
		while (true) {
			if (nums[2] > life[0].length-1) {
				checked[0] = true;
				checked[1] = true;
				checked[7] = true;
			}
			if (nums[3] > life.length-1) {
				for (int j = 1; j <= 3; j++) checked[j] = true;
			}
			if (nums[4] < 0) {
				for (int j = 3; j <= 5; j++) checked[j] = true;
			}
			if (nums[5] < 0) {
				for (int j = 5; j <= 7; j++) checked[j] = true;
			}
			
			for (int open = 0; open < 8; open++) {
				if (!checked[open]) {
					int y = nums[ass[open][0]];
					int x = nums[ass[open][1]];
					if (life[y][x] == '#') {
						occ++;
						checked[open] = true;
					}
					else if (life[y][x] == 'L') checked[open] = true;
				}
			}
			/*
			if (!checked[0]) {
				if (life[i][right] == '#') {
					occ++;
					checked[0] = true;
				}
				else if (life[i][right] == 'L') checked[0] = true;
			}
			if (!checked[1]) {
				if (life[down][right] == '#') {
					occ++;
					checked[1] = true;
				}
				else if (life[down][right] == 'L') checked[1] = true;
			}
			if (!checked[2]) {
				if (life[down][c] == '#') {
					occ++;
					checked[2] = true;
				}
				else if (life[down][c] == 'L') checked[2] = true;
			}
			if (!checked[3]) {
				if (life[down][left] == '#') {
					occ++;
					checked[3] = true;
				}
				else if (life[down][left] == 'L') checked[3] = true;
			}
			if (!checked[4]) {
				if (life[i][left] == '#') {
					occ++;
					checked[4] = true;
				}
				else if (life[i][left] == 'L') checked[4] = true;
			}
			if (!checked[5]) {
				if (life[up][left] == '#') {
					occ++;
					checked[5] = true;
				}
				else if (life[up][left] == 'L') checked[5] = true;
			}
			if (!checked[6]) {
				if (life[up][c] == '#') {
					occ++;
					checked[6] = true;
				}
				else if (life[up][c] == 'L') checked[6] = true;
			}
			if (!checked[7]) {
				if (life[up][right] == '#') {
					occ++;
					checked[7] = true;
				}
				else if (life[up][right] == 'L') checked[7] = true;
			}
			*/
			nums[5]--; // up
			nums[2]++; // right
			nums[3]++; // down
			nums[4]--; // left
			if (ifBreak(checked)) break;
		}
		return occ;
	}
	
	public static boolean ifBreak(boolean[] b) {
		for (boolean a: b)  if (!a) return false;
		return true;
	}
	
	public static int countNeighbors(char[][] life, int i, int c) {
		int occ = 0;
		for (int y = -1; y < 2; y++) {
			for (int x = -1; x < 2; x++) {
				if (c+x < 0 || c+x > life[i].length-1 || i+y < 0 || i+y > life.length-1 || (y == 0 && x == 0)) {
					continue;
				}
				else {
					if (life[i+y][c+x] == '#') {
						occ++;
					}
				}
			}
		}
		return occ;
	}
	
	public static int count(char[][] a) {
		int sum = 0;
		for (int y = 0 ; y < a.length; y++) {
			for (int x = 0; x < a[y].length; x++) {
				if (a[y][x] == '#') sum++;
			}
		}
		return sum;
	}
	
	public static boolean checkSame(char[][] a, char[][] b) {
		for (int y = 0 ; y < a.length; y++) {
			for (int x = 0; x < a[y].length; x++) {
				if (a[y][x] != b[y][x]) return false;
			}
		}
		return true;
	}
	
	public static void printArray(char[][] a) {
		String msg;
		for (int y = 0; y < a.length; y++) {
			msg = "";
			for (int x = 0; x < a[y].length; x++) {
				msg += a[y][x];
			}
			System.out.println(msg);
			
		}
		System.out.println();
	}
	
	public static char[][] copyArray(char[][] from) {
		char[][] to = new char[from.length][from[0].length];
		for (int y = 0 ; y < from.length; y++) {
			for (int x = 0; x < from[y].length; x++) {
				to[y][x] = from[y][x];
			}
		}
		return to;
	}

}
