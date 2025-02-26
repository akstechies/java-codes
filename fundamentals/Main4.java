public class Main4 {

    private int num;
    
    public void getMessage() {
        System.out.println("Hello World");
    }

    public int addTwoNos(int a, int b) {
        return a + b;
    }

    public static void method_1() {
        System.out.println("Method 1");
    }

    public void method_2() {
        this.method_1();    // this calls static but should not be inside static func
        System.out.println("Method 2");
    }

    public void setNum(int num) {
        this.num = num;
    }

    public int getNum() {
        return this.num;
    }

    public static void main(String[] args) {
        Main4 main4 = new Main4();
        main4.getMessage();

        int sum = main4.addTwoNos(77, 41);
        System.out.println(sum);

        // we can call static method without creating an instance
        Main4.method_1();

        // these are instance methods as this requires object of the class
        main4.setNum(24);
        System.out.println(main4.getNum());

        // we can call static method directly also
        method_1();


        // In Java, Access modifiers helps to restrict the scope of a class, constructor, variable, method, or data member.
        /* 
        Default – No keyword required
        Private
        Protected
        Public
        */
    }

}
