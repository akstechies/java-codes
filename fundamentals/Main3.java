public class Main3 {

    public static void main(String[] args) {
        int i = 20;


        if (i < 10) {
            System.out.println("Less than 10");
        } else if (i < 20) {
            System.out.println("Less than 20");
        } else {
            System.out.println("Hah! More!!");
        }

        // multiple ways to write
        if (i == 20); System.out.println("i is " + i);

        if (i < 30) System.out.println("2- i is " + i);

        if (i < 30) 
            System.out.println("3- i is " + i);

        i = 54;
        if (i < 20) 
            System.out.println("3- i is " + i);
        else if (i < 30)
            System.out.println("I is:-- " + i);
        else
            System.out.println("else I is:-- " + i);


        // Switch case
        int num = 10;
        switch (num) {
            case 10:    // by this all below will be ececuted until break - keep this for same condition
            case 20:
                System.out.println("Num is 20");
                break;      // break is needed else all the cases below will also be executed
            case 30:
                System.out.println("Num is 30");
                break;
            default:
                System.err.println("Err! Default case");
        }


        // continue - if will skip the current iter and go to next iter
        for (int i1 = 0; i1 < 5; i1++) {
            if (i1 % 2 == 0) 
                continue;
            System.out.print(i1 + ", ");
        }
        System.out.println();
        // LOOPS
        for (int i2 = 0 ; i2 < 5; i2++) {
            System.out.print(i2 + ", ");
        }

        System.out.println();
        // for each loop for iteration over arrays or collection
        int[] arr = {1, 2, 3, 4, 5, 6};

        for (int a : arr) {
            System.out.print(a + ", ");
        }

        System.out.println();
        int x = 1;
        while (x < 5) {
            System.out.print(x + ", ");
            x++;
        }

        System.out.println();
        // do while - here what will happen then anything inside do will first exercute no matter what
        int z = 1;
        do {
            System.out.println("Do while");
            z++;
        } while (z < 1);


    }


    
}