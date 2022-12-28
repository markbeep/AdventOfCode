import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.util.Scanner;
import java.util.TreeMap;

public class day4 {

	public static void main(String[] args) throws FileNotFoundException {
		File file = new File("day4.txt");
		Scanner scanner = new Scanner(file);
		ArrayList<String> info = new ArrayList<String>();
		String txt = "";
		int total = 0;
		int c = 0;
		boolean br;
		while (scanner.hasNextLine()) {
			txt = " ";
			String line = scanner.nextLine();
			if (line.length() == 0) {
				for (String e: info) {
					txt += " " + e;
				}
				txt += " ";
				String[] needed = {" byr:", " iyr:", " eyr:", " hgt:", " hcl:", " ecl:", " pid:"};
				Map<String, String> counter = new TreeMap<>();
				c = 0;
				br = false;
				for (String s: needed) {
					int index = txt.indexOf(s);
					if (index == -1) {
						br = true;
						break;
					}
					int col = txt.indexOf(":", index)+1;
					int space = txt.indexOf(" ", col);
					String val = txt.substring(col, space);
					System.out.print(s);
					System.out.println(val);
					counter.put(s, val);
				}
				if (!br) {
					for (String v: new String[] {" byr:", " iyr:", " eyr:", " pid:"}) {
						try {
							if (v.equals(" pid:") && !(counter.get(v).length() == 9)) continue;
							int year = Integer.valueOf(counter.get(v));
							if (v.equals(" pid:")) c++;
							if (v.equals(" byr:") && year >= 1920 && year <= 2002) c++;
							if (v.equals(" iyr:") && year >= 2010 && year <= 2020) c++;
							if (v.equals(" eyr:") && year >= 2020 && year <= 2030) c++;
						}
						catch (NumberFormatException e) {}
					}
					if (counter.get(" hgt:").contains("cm") || counter.get(" hgt:").contains("in")) {
						try {
						if (counter.get(" hgt:").contains("cm")) {
							int h = Integer.valueOf(counter.get(" hgt:").substring(0, 3));
							if (h >= 150 && h <= 193) c++;
						}
						else {
							int h = Integer.valueOf(counter.get(" hgt:").substring(0, 2));
							if (h >= 59 && h <= 76) c++;
						}
						}catch (NumberFormatException e) {}
					}
					
					if (counter.get(" hcl:").length() == 7) {
						String val = counter.get(" hcl:");
						if (val.charAt(0) == '#') {
							val = val.substring(1);
							char[] chars = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'};
							int o = 0;
							for (int i = 0; i < 6; i++) {
								if (new String(chars).contains(val.substring(i, i+1))) o++;
							}
							if (o == 6) c++;
						}
					}
					String[] t = new String[] {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"};
					if (Arrays.asList(t).contains(counter.get(" ecl:"))) c++;
					total += (c == 7) ? 1: 0;
				}
				info = new ArrayList<String>();
				System.out.println(total);
			}
			info.add(line);
		}
		System.out.println(total);

	}

}
