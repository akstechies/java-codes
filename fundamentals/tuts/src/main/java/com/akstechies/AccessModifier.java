package com.akstechies;

public class AccessModifier {
    static void get_codes() {
        // System.out.println("Hello world!");

        Greet greet = new Greet();
        greet.greet_them();

        // this is a private access modifier hence not available in another class
        // greet.secret_info();

        Greet.get_secret_info();
        // greet.get_secret_info();     // this works but since it is static call is using above

        greet.student_subject();


        // in java method should be called inside a method or a constructor, calling inside class body is not allowed

        greet.college();
    }
}
