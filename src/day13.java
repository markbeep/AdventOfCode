import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.Scanner;

public class day13 {

	public static void main(String[] args) throws FileNotFoundException{
		Scanner s = new Scanner(new File("day13.txt"));
		s.nextLine();
		String[] line = s.nextLine().split(",");

		// For each bus ID, there's an ArrayList with all values that need to be multiplied
		Map<Long, ArrayList<Long>> bus = new LinkedHashMap<>();
		
		// Used to store the index value that each bus ID needs to have
		Map<Long, Long> index = new LinkedHashMap<>();
		
		// Adds the bus ID and the index to the two arrays
		long j = 0;
		for (String e: line) {
			if (!e.equals("x")) {
				Long n = Long.valueOf(e);
				bus.put(n, new ArrayList<Long>());
				index.put(n, correct(n - j, n));
			}
			j++;
		}
		
		System.out.println(CRT(bus, index));
	}
	public static long CRT(Map<Long, ArrayList<Long>> bus, Map<Long, Long> index) {
		long sum = 0; // the sum of all values in "bus" multiplied
		long all = 1; // all keys multiplied
		
		// does CRT for each value in the map "bus"
		for (Long e: bus.keySet()) {
			all *= e;
			long n = index.get(e);
			
			// Adds all numbers that are not e to the "to multiply" ArrayList
			for (Long j: bus.keySet()) {
				if (e != j) {
					bus.get(e).add(j);
				}
			}
			// if the "multiplied value" MOD "the key" is not the bus ID, try multiplying by 2, 3, 4, ... until it works
			if (mult(bus.get(e)) % e != n % e) {
				for (long i = 2; i < 10000; i++) {
					if (i * mult(bus.get(e)) % e == n % e) {
						bus.get(e).add(i);
						break;
					}
				}
			}
			sum += mult(bus.get(e));
		}
		return sum % all;
	}
	/*
	 * Used to multiply all values in an arraylist
	 */
	public static long mult(ArrayList<Long> a) {
		long s = 1;
		for (long e: a) s *= e;
		return s;
	}
	/*
	 * Used to correct negative modulos (-2 mod 3 = 1 mod 3 for example)
	 */
	public static long correct(long a, long b) {
		while (a < 0) a += b;
		return a;
	}

}
