public class Main1 {
    public static void main(String[] args) {
        System.out.print("Hello World");
        System.out.println("- AYush");      // this prints without new line

        // printing numbers
        System.out.println(30);
        System.out.println(2 * 3);
        System.out.println(44 + 6);

        // error
        System.err.println("Error bro");


        /*** Java Variables */
        /*
         => String - stores text, such as "Hello". String values are surrounded by double quotes
         => int - stores integers (whole numbers), without decimals, such as 123 or -123
         => float - stores floating point numbers, with decimals, such as 19.99 or -19.99
         => char - stores single characters, such as 'a' or 'B'. Char values are surrounded by single quotes
         => boolean - stores values with two states: true or false
        */

        String name = "Ayush";
        int num = 10;
        float grade = 2.8f;     // by default the points are treated as double so better ad f to tell the java it is to be treated as float
        double points = 77.95;
        char ch = 'a';
        boolean is_passed = true;

        System.out.println(name + ", " + num + ", " + 
            grade + ", " + points + ", " + ch + ", " + is_passed
        );

        // If you don't want others (or yourself) to overwrite existing values, use the final keyword (this will declare the variable as "final" or "constant", which means unchangeable and read-only):
        final int TOTAL_MARKS = 100;
        System.out.println("Total marks is: " + TOTAL_MARKS);


        // declare multiple variables
        int x = 10, y = 20, z = 15;
        System.out.println(x + y + z);

        int a, b, c;
        a = b = c = 10;
        System.out.println(a + b + c);

        // Primitive data types - includes byte, short, int, long, float, double, boolean and char
        byte position = 4;
        System.out.println(position);
        
        short height = 5; 
        System.out.print(height);
        System.out.println(" ft.");

        long distance = 985;
        System.out.print(distance);
        System.out.println(" km.");

        // double can also be given `d`
        double dbl = 296.45d;
        System.out.println(dbl);


        /** Java Type Casting */
        /*
        Type casting is when you assign a value of one primitive data type to another type.

         - In Java, there are two types of casting:

            => Widening Casting (automatically) - converting a smaller type to a larger type size
             -- byte -> short -> char -> int -> long -> float -> double

            => Narrowing Casting (manually) - converting a larger type to a smaller size type
             -- double -> float -> long -> int -> char -> short -> byte
        */
        // Widening Casting
        int myInt = 10;
        double myDbl = myInt;       // Automatic casting: int to double
        System.out.println(myDbl);

        // Narrowing casting
        // ~ Narrowing casting must be done manually by placing the type in parentheses () in front of the value:

        short myShort = (short) myDbl;      // if not use typecast then error will be cannot convert double to short
        System.out.println(myShort);
    }
}