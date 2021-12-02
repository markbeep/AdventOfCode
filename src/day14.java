import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;

public class day14 {

	public static void main(String[] args) throws FileNotFoundException {
		Scanner s = new Scanner(new File("day14.txt"));
		String mask = "";
		String[] line;
		long mem;
		long value;
		String[] bitA;
		Map<String, Long> bits = new HashMap<>();
		int l = 0;
		while (s.hasNextLine()) {
			line = s.nextLine().split(" = ");
			if (line[0].contains("mask")) {
				mask = line[1];
			}
			else {
				mem = Long.valueOf(line[0].replace("mem[", "").replace("]", ""));
				value = Long.valueOf(line[1]);
				bitA = masker(mask, mem, value);
				//System.out.println(String.join("", bitA));
				adder(bitA, bits, value);
			}
		}
		long sum = 0;
		for (String k: bits.keySet()) {
			sum += bits.get(k);
		}
		System.out.println(sum);
	}
	public static void adder(String[] mem, Map<String, Long> bits, long value) {
		if (!Arrays.asList(mem).contains("X")) {
			bits.put(String.join("", mem), value);
		}
		else {
			String[] copy;
			for (int c = 0; c < mem.length; c++) {
				if (mem[c].equals("X")) {
					copy = mem.clone();
					copy[c] = "0";
					adder(copy, bits, value);
					copy[c] = "1";
					adder(copy, bits, value);
					break;
				}
			}
		}
	}
	public static String[] masker(String mask, long mem, long value) {
		String bit = Long.toBinaryString(mem);  // task 1: value | task 2: mem
		bit = "0".repeat(36 - bit.length()) + bit;
		String[] bitA = bit.split("");
		String[] masA = mask.split("");
		for (int c = 0; c < masA.length; c++) {
			if (masA[c].equals("X")) {
				bitA[c] = "X";
			}
			else if (masA[c].equals("1")) {
				bitA[c] = "1";
			}
		}
		return bitA;
	}

}
