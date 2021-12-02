import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;

public class quickScanner {

		public static ArrayList<Integer> allInts(String fn) throws FileNotFoundException{
			File f = new File(fn);
			Scanner s = new Scanner(f);
			ArrayList<Integer> a = new ArrayList<Integer>();
			while (s.hasNextInt()) {
				a.add(s.nextInt());
			}
			s.close();
			return a;
		}
		
		public static ArrayList<String> allStrings(String fn) throws FileNotFoundException{
			File f = new File(fn);
			Scanner s = new Scanner(f);
			ArrayList<String> a = new ArrayList<String>();
			while (s.hasNext()) {
				a.add(s.next());
			}
			s.close();
			return a;
		}
		
		public static ArrayList<String> allLines(String fn) throws FileNotFoundException{
			File f = new File(fn);
			Scanner s = new Scanner(f);
			ArrayList<String> a = new ArrayList<String>();
			while (s.hasNextLine()) {
				a.add(s.nextLine());
			}
			s.close();
			return a;
		}
}
