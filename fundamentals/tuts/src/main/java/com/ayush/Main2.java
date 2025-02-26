package com.ayush;

import com.akstechies.Greet;

public class Main2  extends Greet{
    static Greet greet = new Greet();

    // this is a default modifier function hence not avalaible in another package
    // greet.greet_them()

    public static void main(String[] args) {
        Main2 main2 = new Main2();
        main2.student_subject();

        greet.college();
    }

    // You're encountering a syntax error because you're trying to execute a method call (main2.student_subject();) directly inside the class body, which is not allowed in Java. Method calls must be inside a method or a constructor.

    
}


