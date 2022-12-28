import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Scanner;

public class day10 {

	public static void main(String[] args) throws FileNotFoundException {
		File file = new File("day10.txt");
        Scanner s = new Scanner(file);
        ArrayList<Integer> a = new ArrayList<>();
        int d;
        int max = 0;
        while (s.hasNextLine()) {
            d = s.nextInt();
            a.add(d);
            if (max < d) max = d;
        }
        a.add(0);
        a.add(max+3);
		/*
		int prev = 0;
		int ones = 0;
		int thre = 1;
		for (int e: a.keySet()) {
			d = e - prev;
			if (d == 1) ones++;
			if (d == 3) thre++;
			prev = e;
		}
		System.out.println(ones);
		System.out.println(thre);
		System.out.println(ones*thre);
		*/
        Collections.sort(a);
        System.out.println(d(max+4, a));
        s.close();
    }
    public static long d(int m, List<Integer> a) {
        long[] p = new long[m];
        p[0] = 1;
        for (int k: a)
            if (k != 0)
                for (int y = 1; y <= 3; y++)
                    if (k-y >= 0) p[k] += p[k-y];
        return(p[p.length-1]);
    }

}
