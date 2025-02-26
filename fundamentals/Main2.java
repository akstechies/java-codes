public class Main2 {

    // instance variables
    public String name;
    public String subject;
    public int intI;
    public Integer integerI;
    public char ch;

    final int total_marks = 100;
    int ayush_marks = 22;

    public static int z = 55;

    public Main2() {
        this.name = "Ayush";
    }


    public static void main(String[] args) {
        
        /* Types of Java Variables
         *- Local   
         *- Instance
         *- Static
        */

        // Local variable
        int x = 10;
        {
            x = 20;
        } 
        System.err.println(x);

        if (x < 40) {
            String result = "Less than 40";
            System.err.println(result);
        }
        // now result varialve above is in block scope - so below it will give error
        // System.err.println(result);


        // 2. Instance Variables
        /* Instance variables are non-static variables and are declared in a class outside of any method, constructor, or block. 
        - Initialization of an instance variable is not mandatory. Its default value is dependent on the data type of variable. For String it is null, for float it is 0.0f, for int it is 0, for Wrapper classes like Integer it is null, etc.
        - Scope of instance variables are throughout the class except the static contexts.
        - As instance variables are declared in a class, these variables are created when an object of the class is created and destroyed when the object is destroyed.
        */
        // create object of the class
        Main2 main2 = new Main2();
        System.out.println(main2.name);
        System.out.println(main2.intI);
        System.out.println(main2.integerI);
        System.out.println(main2.subject);
        System.out.println(main2.ch);   // empty 

        // The problem is that the static modifier is not allowed for local variables inside a method. To fix this, you should move the static variable declaration outside the 
        // static int z1 = 5;
        // since it is a static variable we can directly call it or call by class name without creating instance
        System.out.println(z);
        System.out.println(Main2.z);

        int z = 5;
        System.out.println(z);
        System.out.println(Main2.z);

        // since it is a non static we need to reference it by class
        System.out.println(main2.ayush_marks + " / " + main2.total_marks );
    }


}




