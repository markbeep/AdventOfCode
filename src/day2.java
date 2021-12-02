import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;

public class day2 {

	public static void main(String[] args) throws FileNotFoundException {
		ArrayList<String> lines = quickScanner.allLines("day2.txt");
		char letter = '#';
		int min = 0;
		int lineIndex = 0;
		int max = 0;
		int count;
		int total = 0;
		String pw;
		for (int i = 0; i < lines.size(); i++) {
			System.out.println(lines.get(i));
			letter = '#';
			pw = "";
			for (int c = 0; c < lines.get(i).length(); c++) {
				if (lines.get(i).charAt(c) == ':') {
					letter = lines.get(i).charAt(c-1);
					max = Integer.parseInt(lines.get(i).substring(lineIndex + 1, c-2));
					pw = lines.get(i).substring(c+1);
					break;
				}
				if (lines.get(i).charAt(c) == '-') {
					min = Integer.parseInt(lines.get(i).substring(0, c));
					lineIndex = c;
				}
			}
			if ((pw.charAt(min) == letter && pw.charAt(max) != letter) || (pw.charAt(max) == letter && pw.charAt(min) != letter)) {
				total++;
			}
		}
		System.out.println(total);

	}

}
