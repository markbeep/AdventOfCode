import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Map;
import java.util.Scanner;
import java.util.TreeMap;

public class day7 {

	public static void main(String[] args) throws FileNotFoundException {
		File file = new File("day7.txt");
		Scanner scanner = new Scanner(file);
		Map<String, String[]> bags = new TreeMap<>();
		ArrayList<String> inside = new ArrayList<String>();
		while (scanner.hasNextLine()) {
			String line = scanner.nextLine();
			String[] parts = line.split("contain");
			String primary = parts[0].replace(" bags", "");
			primary = primary.substring(0, primary.length()-1);
			String[] otherBags = parts[1].split(",");
			for (int i = 0; i < otherBags.length; i++) {
				otherBags[i] = otherBags[i].substring(1).replace(".", "");
				if (otherBags[i].contains("bags")) {
					otherBags[i] = otherBags[i].replace("bags", "");
				}
				else {
					otherBags[i] = otherBags[i].replace("bag", "");
				}
				otherBags[i] = otherBags[i].substring(0, otherBags[i].length()-1);
			}
			if (otherBags[0] == "no other") otherBags = new String[0];
			
			bags.put(primary, otherBags);
		}
		int counter = 0;
		System.out.println(canIt(bags, "shiny gold"));
	}
	
	public static int canIt(Map<String, String[]> map, String bag) {
		int total = 1;
		if (!map.keySet().contains(bag)) return 1;
		for (String s: map.get(bag)) {
			if (s.contains("other")) {
				continue;
			}
			int amt = Integer.valueOf(s.substring(0, 1));
			s = s.substring(2);
			total += amt * canIt(map, s);
		}
		
		return total;
	}

}
